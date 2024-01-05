# Template for gRPC services
This template allows you to quickly scaffold a go repo to create a gRPC service.

It comes with a variety of boilerplate code to connect to and consume commonly used services in a gRPC microservice.

## Pre-requisites
1. Protoc - https://grpc.io/docs/protoc-installation/
2. The application authenticates with Google cloud platform via the `GOOGLE_APPLICATION_CREDENTIALS` environment variable. When deployed, the cloud run instance automatically has access to the credentials. When developing locally you will need to set the variable. See https://cloud.google.com/docs/authentication/application-default-credentials
3. When changes are pushed to the main branch, [Cloud build](https://cloud.google.com/build?hl=en) will automatically build the code and deploy it into [Cloud run](https://cloud.google.com/run?hl=en).

## Documentation
1. Protocol Buffers - https://protobuf.dev/programming-guides/proto3
2. gRPC - https://grpc.io/docs/languages/go/quickstart/

## Guide
Create a folder under `src`. This will be a gRPC service. You can then create a `.proto` file within that directory (see book_service example). Once you have defined your `.proto` file, you can generate the necessary code with the following command (run from the root directory). For example, the command to generate the go code for the book_service:
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./protocol_buffers/book_service/book_service.proto    
```
You can then create a new folder under src and implement your gRPC service by using the code that was generated in the above command (see book_service example).
