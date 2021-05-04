package supertest

func (ctx *supertest) Auth(username, password string) {
	ctx.request.httpRequest.SetBasicAuth(username, password)
}
