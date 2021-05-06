package supertest

import (
	"net/http"
	"net/http/httptest"
)

func (ctx *supertest) End(handleFunc func(req *http.Request, rr *httptest.ResponseRecorder)) {
	ctx.router.ServeHTTP(ctx.response.httpResponse, ctx.request.httpRequest)
	handleFunc(ctx.request.httpRequest, ctx.response.httpResponse)
}
