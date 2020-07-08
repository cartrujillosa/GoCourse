run-server:
	go build server.go chat.go && ./server
run-client:
	go run client.go