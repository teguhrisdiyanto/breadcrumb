SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
PACKAGE = $(shell Get-Content go.mod -head 1 | Foreach-Object { $$data = $$_ -split " "; "{0}" -f $$data[1]})
BIN = crud-rpc.exe

build: 	generate
	go build .

generate:
	protoc -Iproto --go_opt=module=${PACKAGE} --go-grpc_opt=module=${PACKAGE} --go_out=. proto/*.proto --go-grpc_out=. proto/*.proto

bump: generate
	go get -u ./...

clean:
	rm proto/*.pb.go
	rm crud-rpc

# protoc -Iproto --go_out=./proto --go_opt=path=source_relative\ --go-grpc_out=./proto --go-grpc_opt=paths=source_relative\ proto/*.proto
#protoc  --go_out=github.com/teguhrisdiyanto/crud-rpc/proto --go_opt=path=source_relative\ --go-grpc_out=github.com/teguhrisdiyanto/crud-rpc/proto --go-grpc_opt=paths=source_relative\ proto/*.proto


#berhasil
# protoc --go-grpc_out=github.com/teguhrisdiyanto/crud-rp --go_grpc_out=paths=source_relative\  proto/company-api.proto

#berhasil
# protoc --go_out=. --go_opt=paths=source_relative  proto/company.proto


# protoc --go_out=.  proto/address/v1/address.proto
# protoc --go-grpc_out=.  proto/address/v1/address_api.proto


# protoc --go_out=. --go-grpc_out=. proto/address/v1/*.proto
