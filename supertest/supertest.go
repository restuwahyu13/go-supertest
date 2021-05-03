package super

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

type SuperTest interface {
	GET(url string) (*httptest.ResponseRecorder, *http.Request, error)
	POST(url string) (*httptest.ResponseRecorder, *http.Request, error)
	PUT(url string) (*httptest.ResponseRecorder, *http.Request, error)
	DELETE(url string) (*httptest.ResponseRecorder, *http.Request, error)
	Send(payload interface{})
	Set(key, value string)
	End(handle func(rr *httptest.ResponseRecorder))
	Auth(key, value string)
	Field(key , value string)
}

func NewSuperTest(router *gin.Engine, test *testing.T) *supertest {
	return &supertest{router: router, test: test}
}

func(ctx *supertest) GET(url string) (*httptest.ResponseRecorder)  {

	time.Sleep(time.Second * 1)

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte(nil)))
	ctx.request.Request = req

	if err != nil {
		ctx.test.Error(err.Error())
		return nil
	}

	rr := httptest.NewRecorder()
	ctx.response.Response = rr

	ctx.router.ServeHTTP(rr, req)

	return rr
}

func(ctx *supertest) POST(url string) (*httptest.ResponseRecorder) {

	time.Sleep(time.Second * 1)

	response, err := json.Marshal(ctx.body.Data)

	if err != nil  {
		ctx.test.Error(err.Error())
		return nil
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(response))
	ctx.request.Request = req

	if err != nil {
		ctx.test.Error(err.Error())
		return nil
	}

	rr := httptest.NewRecorder()
	ctx.response.Response = rr

	ctx.router.ServeHTTP(rr, req)

	return rr
}

func(ctx *supertest) DELETE(url string) (*httptest.ResponseRecorder)  {

	time.Sleep(time.Second * 1)

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte(nil)))
	ctx.request.Request = req

	if err != nil {
		ctx.test.Error(err.Error())
		return nil
	}

	rr := httptest.NewRecorder()
	ctx.response.Response = rr

	ctx.router.ServeHTTP(rr, req)

	return rr
}

func(ctx *supertest) PUT(url string) (*httptest.ResponseRecorder) {

	time.Sleep(time.Second * 1)

	response, err := json.Marshal(ctx.body.Data)

	if err != nil  {
		ctx.test.Error(err.Error())
		return nil
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(response))
	ctx.request.Request = req

	if err != nil {
		ctx.test.Error(err.Error())
		return nil
	}

	rr := httptest.NewRecorder()
	ctx.response.Response = rr

	ctx.router.ServeHTTP(rr, req)

	return rr
}