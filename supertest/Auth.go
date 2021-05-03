package super

func (ctx *supertest) Auth(username, password string) {
	if username != "" && password != "" {
    ctx.request.Request.SetBasicAuth(username, password)
	}
}