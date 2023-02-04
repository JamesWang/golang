# golang

##gRPC
- Prepare
  - get proto support
    go get github.com/golang/protobuf/proto
    go get -u github.com/golang/protobuf/protoc-gen-go
  - get grpc support
    go get google.golang.org/grpc
- Create proto file(s)
- Generate stub/skeleton files
  - Generate proto/gRPC related go files
    protoc --proto_path=proto proto/*.proto --go_out=./datafiles --go-grpc_out=./datafiles
- When register a server
  - define:

            type server struct {
	            pb.UnimplementedMoneyTransactionServer  //add similar line(s) when you get compile errors
            }
  - implement required functions/methods,
  - register the server
  
            pb.RegisterMoneyTransactionServer(s, &server{})
##Two methods of deleting an item from a slice
![image](https://user-images.githubusercontent.com/324429/216785926-45f43881-8f89-4f6c-94d3-dc7b7d9a0aad.png)
