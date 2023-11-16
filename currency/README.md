make sure you run below command 
Install command:

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH=$PATH:$(go env GOPATH)/bin

Install package...
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc


use require_unimplemented_servers flag in protoc command to avoid unimplemented method error

Install  grpcurl using command `brew install grpcurl`.. this is same of curl command for rest
Below command will return us all the servie on grpc service
grpcurl --plaintext localhost:9092 list

similary  `grpcurl --plaintext localhost:9092 list Currency` command returns the method on Currency service

Use below command in terminal to test the GetRate function on grpc service Currency

**grpcurl --plaintext -d '{"Base":"USD", "Destination":"INR"}' localhost:9092 Currency.GetRate**
