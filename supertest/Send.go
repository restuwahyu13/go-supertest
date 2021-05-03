package super

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"
)

func(ctx *supertest) Send(payload interface{}) {

		time.Sleep(time.Second * 1)

		response, err := json.Marshal(payload)

		if err != nil  {
			ctx.test.Error(err.Error())
		}

		req, err := http.NewRequest(ctx.payload.method, ctx.payload.path, bytes.NewBuffer(response))
		ctx.request.httpRequest = req

		if err != nil {
			ctx.test.Error(err.Error())
		}

		rr := httptest.NewRecorder()
		ctx.response.httpResponse = rr

		ctx.router.ServeHTTP(rr, req)
}
