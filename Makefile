run: |
	gofmt -w .
	go run main.go

mock-service:
	mockgen -source=domain/service/user.go -destination=domain/service/mock_user_engine_db.go -package=service

