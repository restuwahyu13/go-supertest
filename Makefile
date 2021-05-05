goinstall:
	@go mod download

gotest:
	@richgo test -v ./supertest/...

gofix:
	@go fix ./supertest/...