BINARY_NAME=main.out
 
MAIN_FILE="cmd/log_handler/main.go"

build:
	@go build -o ${BINARY_NAME} ${MAIN_FILE}
 
test:
	@go test ./...

run:
	@go build -o ${BINARY_NAME} ${MAIN_FILE}
	@./${BINARY_NAME}
 
clean:
	@go clean
	@rm ${BINARY_NAME}