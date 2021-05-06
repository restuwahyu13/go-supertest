package supertest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (ctx *supertest) Send(payload interface{}) {

	response, err := json.Marshal(payload)

	if err != nil {
		ctx.test.Error(err.Error())
		return
	}

	req, err := http.NewRequest(ctx.payload.method, ctx.payload.path, bytes.NewBuffer(response))
	req.Header.Add("Access-Control-Allow-Origin", "*")
	req.Header.Add("Access-Control-Allow-Headers", "*")
	req.Header.Add("Access-Control-Expose-Headers", "*")
	req.Header.Add("User-Agent", "go-supertest/0.0.1")

	ctx.request.httpRequest = req

	if err != nil {
		ctx.test.Error(err.Error())
		return
	}

	rr := httptest.NewRecorder()
	ctx.response.httpResponse = rr
}
