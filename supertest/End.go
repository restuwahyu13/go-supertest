package supertest

import "net/http/httptest"

func (ctx *supertest) End(handleFunc func(rr *httptest.ResponseRecorder)) {
	handleFunc(ctx.response.httpResponse)
}
