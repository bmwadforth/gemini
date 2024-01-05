package gemini

import (
	"context"
	"errors"
	GeminiService "gemini/protocol_buffers/gemini_service"
	"gemini/util"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Server struct {
	GeminiService.UnimplementedGeminiServer
}

func streamParts(srv GeminiService.Gemini_QueryGeminiServer, parts []genai.Part, option int, error error) error {
	var response string
	var responseError string
	if error != nil {
		responseError = error.Error()
	} else {
		responseError = ""
	}

	for _, part := range parts {
		switch part.(type) {
		case genai.Text:
			response = string(part.(genai.Text))
		case genai.Blob:
			util.SLogger.Errorf("Received invalid type of data from Gemini response")
			continue
		default:
			util.SLogger.Errorf("Received invalid type of data from Gemini response")
			continue
		}

		if err := srv.Send(&GeminiService.QueryResponse{Response: response, Option: int32(option), Error: responseError}); err != nil {
			util.SLogger.Errorf("Error occurred: %v", err)
			err := status.Error(codes.Internal, err.Error())
			return err
		}
	}

	return nil
}

func (s *Server) QueryGemini(in *GeminiService.QueryRequest, srv GeminiService.Gemini_QueryGeminiServer) error {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(util.Config.GeminiApiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-and-image input (multimodal), use the gemini-pro-vision model. gemini-pro for text based only
	model := client.GenerativeModel("gemini-pro")

	prompt := genai.Text(in.Query)
	iter := model.GenerateContentStream(ctx, prompt)
	var blockedError *genai.BlockedError

	for {
		blockedError = nil
		resp, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		switch {
		case errors.As(err, &blockedError):
			{
				// If non-nil, the model's response was blocked.
				// Consult the Candidate and SafetyRatings fields for details.
				if blockedError.Candidate != nil {
					streamParts(srv, blockedError.Candidate.Content.Parts, 0, blockedError)
				}

				// If non-nil, there was a problem with the prompt.
				if blockedError.PromptFeedback != nil {
					if err := srv.Send(&GeminiService.QueryResponse{Response: blockedError.PromptFeedback.BlockReason.String(), Option: 0, Error: blockedError.Error()}); err != nil {
						util.SLogger.Errorf("Error occurred: %v", err)
						err := status.Error(codes.Internal, err.Error())
						return err
					}
				}
				util.SLogger.Fatal(err)
			}
		default:
			if err != nil {
				util.SLogger.Fatal(err)
			}
		}

		if resp != nil {
			if len(resp.Candidates) > 1 {
				for i, candidate := range resp.Candidates {
					streamParts(srv, candidate.Content.Parts, i, nil)
				}
			} else {
				streamParts(srv, resp.Candidates[0].Content.Parts, 0, nil)
			}
		}
	}

	return nil
}
