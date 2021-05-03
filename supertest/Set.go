package super

func(ctx *supertest) Set(key, value string)  {
	ctx.request.Request.Header.Set(key,value)
}