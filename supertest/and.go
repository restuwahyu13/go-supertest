package supertest

import "net/http/httptest"

func (ctx *supertest) End(handleFunc func(rr *httptest.ResponseRecorder)) {
	ctx.router.ServeHTTP(ctx.response.httpResponse, ctx.request.httpRequest)
	handleFunc(ctx.response.httpResponse)
}
