test:
	@go test ./...

doc:
	@godoc -http=:8080 -index