package super

func (ctx *supertest) Auth(username, password string) {
	if username != "" && password != "" {
    ctx.request.httpRequest.SetBasicAuth(username, password)
	}
}