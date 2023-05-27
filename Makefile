gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=./pb --go_out=./pb

clean:
	rm pb/*.go 

run:
	go run main.go 
