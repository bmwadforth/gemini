package mapper

import BookService "svc-template/protocol_buffers/book_service"

func MapCreateBookRequest(request *BookService.CreateBookRequest) BookService.Book {
	return BookService.Book{
		Isbn:   request.Isbn,
		Title:  request.Title,
		Author: request.Author,
	}
}
