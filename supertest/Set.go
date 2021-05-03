package super

func(ctx *supertest) Set(key, value string)  {
	ctx.request.httpRequest.Header.Set(key,value)
}