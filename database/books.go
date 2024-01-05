package database

import (
	"context"
	"errors"
	BookService "svc-template/protocol_buffers/book_service"
	"svc-template/util"
)

func GetBookByISBN(isbn int64) util.DataResponse[*BookService.Book] {
	var book *BookService.Book
	dataResponse := util.NewDataResponse("success", book)
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	docs, err := client.Collection("books").Where("Isbn", "==", isbn).Documents(ctx).GetAll()
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	if len(docs) == 0 {
		dataResponse.SetError(errors.New("no books found"), util.DbresultNotFound)
		return dataResponse
	}

	if len(docs) > 1 {
		dataResponse.SetError(errors.New("error multiple books found"), util.DbresultError)
		return dataResponse
	}

	err = docs[0].DataTo(&book)
	if err != nil {
		dataResponse.SetError(errors.New("error unable to deserialize record"), util.DbresultError)
		return dataResponse
	}
	dataResponse.SetData(book)

	return dataResponse
}

func GetBooksByAuthor(author string) util.DataResponse[*[]BookService.Book] {
	books := make([]BookService.Book, 0, 5)
	var book BookService.Book
	dataResponse := util.NewDataResponse("success", &books)
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	docs, err := client.Collection("books").Where("Author", "==", author).Documents(ctx).GetAll()
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	if len(docs) == 0 {
		dataResponse.SetError(errors.New("no books found"), util.DbresultNotFound)
		return dataResponse
	}

	for _, doc := range docs {
		err := doc.DataTo(&book)
		if err != nil {
			dataResponse.SetError(errors.New("error unable to deserialize record"), util.DbresultError)
			return dataResponse
		}

		books = append(books, book)
	}

	dataResponse.SetData(&books)

	return dataResponse
}

func GetBooks() util.DataResponse[*[]BookService.Book] {
	books := make([]BookService.Book, 0, 5)
	var book BookService.Book
	dataResponse := util.NewDataResponse("success", &books)
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	docs, _ := client.Collection("books").Documents(ctx).GetAll()

	for _, doc := range docs {
		err := doc.DataTo(&book)
		if err != nil {
			dataResponse.SetError(errors.New("error unable to deserialize record"), util.DbresultError)
			return dataResponse
		}

		books = append(books, book)
	}

	dataResponse.SetData(&books)

	return dataResponse
}

func CreateBook(request *BookService.Book) util.DataResponse[string] {
	dataResponse := util.NewDataResponse("success", "")
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	docRef, _, err := client.Collection("books").Add(ctx, request)
	if err != nil {
		dataResponse.SetError(err, util.DbresultError)
		return dataResponse
	}

	dataResponse.SetData(docRef.ID)

	return dataResponse
}
