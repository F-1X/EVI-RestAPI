go-get-firebase:
	go get firebase.google.com/go/v4@latest

install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files




run:
	go run cmd/main.go 

swag:
	swag fmt
	swag init -g ./cmd/main.go -o cmd/docs

cover:
	go tool cover