package supertest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func(ctx *supertest) Send(payload interface{}) {

		request  := make(chan *http.Request, 1)

		if ctx.payload.method == "GET" || ctx.payload.method == "DELETE" {
				req, err :=  http.NewRequest(ctx.payload.method, ctx.payload.path, bytes.NewBuffer([]byte(nil)))
				request <- req
				ctx.request.httpRequest = req

				if err != nil  {
					ctx.test.Fail()
					return
				}
		} else {
				response, err := json.Marshal(payload)

				if err != nil  {
					ctx.test.Fail()
					return
				}

				req, err := http.NewRequest(ctx.payload.method, ctx.payload.path, bytes.NewBuffer(response))
				request <- req
				ctx.request.httpRequest = req

				if err != nil {
					ctx.test.Error(err.Error())
					return
				}
		}

		rr := httptest.NewRecorder()
		ctx.response.httpResponse = rr

		ctx.router.ServeHTTP(rr, <- request)
}
