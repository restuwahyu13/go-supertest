goinstall:
	@go mod download

gotest:
	@go test -v ./supertest/...

gorich:
	@richgo test -v ./supertest/...

gofix:
	@go fix ./supertest/...