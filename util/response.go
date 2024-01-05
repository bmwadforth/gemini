package util

import (
	"google.golang.org/grpc/codes"
)

type EDatabaseResult uint

var (
	DbresultNotFound   EDatabaseResult = 0
	DbresultError      EDatabaseResult = 1
	DbresultOk         EDatabaseResult = 2
	DbresultIncomplete EDatabaseResult = 3
)

type DataResponse[T any] struct {
	Message        string
	error          error
	databaseResult EDatabaseResult
	Data           T
}

func NewDataResponse[T any](message string, data T) DataResponse[T] {
	return DataResponse[T]{
		Message:        message,
		Data:           data,
		databaseResult: DbresultOk,
		error:          nil,
	}
}

func (de *DataResponse[T]) SetData(data T) {
	de.Data = data
}

func (de *DataResponse[T]) SetDatabaseResult(result EDatabaseResult) {
	de.databaseResult = result
}

func (de *DataResponse[T]) SetError(error error, result EDatabaseResult) {
	SLogger.Errorf("an error has occurred: %v", error.Error())
	de.error = error
	de.databaseResult = result
}

func (de *DataResponse[T]) GetError() error {
	return de.error
}

func (de *DataResponse[T]) GetDataResult() EDatabaseResult {
	return de.databaseResult
}

func (de *DataResponse[T]) GetCodeFromDBResult() codes.Code {
	var code codes.Code
	switch de.databaseResult {
	case DbresultOk:
		code = codes.OK
	case DbresultIncomplete:
		code = codes.Unknown
	case DbresultNotFound:
		code = codes.NotFound
	case DbresultError:
		code = codes.Internal
	default:
		code = codes.Unknown
	}

	return code
}

func (de *DataResponse[T]) GetErrorMessage() string {
	return de.error.Error()
}
