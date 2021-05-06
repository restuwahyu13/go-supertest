package supertest

func (ctx *supertest) Set(key, value string) {
	ctx.request.httpRequest.Header.Add(key, value)
}
