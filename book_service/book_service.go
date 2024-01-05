package svc_template

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"svc-template/database"
	"svc-template/mapper"
	BookService "svc-template/protocol_buffers/book_service"
	"svc-template/util"
)

type Server struct {
	BookService.UnimplementedBookServiceServer
}

func (s *Server) GetBook(ctx context.Context, in *BookService.GetBookRequest) (*BookService.Book, error) {
	dataResponse := database.GetBookByISBN(in.Isbn)
	if dataResponse.GetError() != nil {
		err := status.Error(dataResponse.GetCodeFromDBResult(), dataResponse.GetErrorMessage())
		return nil, err
	}

	return dataResponse.Data, nil
}

func (s *Server) GetBooksViaAuthor(in *BookService.GetBookViaAuthor, srv BookService.BookService_GetBooksViaAuthorServer) error {
	dataResponse := database.GetBooksByAuthor(in.Author)
	if dataResponse.GetError() != nil {
		err := status.Error(dataResponse.GetCodeFromDBResult(), dataResponse.GetErrorMessage())
		return err
	}

	for _, book := range *dataResponse.Data {
		if err := srv.Send(&book); err != nil {
			util.SLogger.Errorf("Error occurred: %v", err)
			err := status.Error(codes.Internal, err.Error())
			return err
		}
	}

	return nil
}

func (s *Server) GetGreatestBook(srv BookService.BookService_GetGreatestBookServer) error {
	recv, err := srv.Recv()
	if err != nil {
		util.SLogger.Errorf("Error occurred: %v", err)
		return err
	}

	util.SLogger.Infof("GetGreatestBook with req :%v", recv)

	response := BookService.Book{
		Isbn:   recv.Isbn,
		Title:  "GetGreatestBook Method",
		Author: "Whiskey0",
	}

	if err := srv.SendAndClose(&response); err != nil {
		util.SLogger.Errorf("Error occurred: %v", err)
		return err
	}

	return nil
}

func (s *Server) GetBooks(srv BookService.BookService_GetBooksServer) error {
	_, err := srv.Recv()
	if err != nil {
		util.SLogger.Errorf("Error occurred: %v", err)
		err := status.Error(codes.Internal, err.Error())
		return err
	}

	dataResponse := database.GetBooks()
	if dataResponse.GetError() != nil {
		err := status.Error(dataResponse.GetCodeFromDBResult(), dataResponse.GetErrorMessage())
		return err
	}

	for _, book := range *dataResponse.Data {
		if err := srv.Send(&book); err != nil {
			util.SLogger.Errorf("Error occurred: %v", err)
			err := status.Error(codes.Internal, err.Error())
			return err
		}
	}

	return nil
}

func (s *Server) CreateBook(ctx context.Context, in *BookService.CreateBookRequest) (*BookService.CreateBookResponse, error) {
	createBook := mapper.MapCreateBookRequest(in)
	dataResponse := database.CreateBook(&createBook)
	if dataResponse.GetError() != nil {
		err := status.Error(dataResponse.GetCodeFromDBResult(), dataResponse.GetErrorMessage())
		return nil, err
	}

	return &BookService.CreateBookResponse{Id: dataResponse.Data}, nil
}
