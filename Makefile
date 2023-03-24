mocks:
	mockgen -source=./internal/service/service.go -destination=./internal/service/mockservice.go -package=service