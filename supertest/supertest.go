package super

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

type SuperTest interface {
	Get(url string) (*httptest.ResponseRecorder, *http.Request, error)
	Post(url string) (*httptest.ResponseRecorder, *http.Request, error)
	Put(url string) (*httptest.ResponseRecorder, *http.Request, error)
	Delete(url string) (*httptest.ResponseRecorder, *http.Request, error)
}

type payload struct {
	path string
	method string
}

type response struct {
	httpResponse *httptest.ResponseRecorder
}

type request struct {
	httpRequest *http.Request
}

type core interface {
	Send(payload interface{})
	End(handle func(rr *httptest.ResponseRecorder))
	Set(key, value string)
	Auth(key, value string)
}

type supertest struct {
	router *gin.Engine
	test *testing.T
	payload
	response
	request
	core
}

func NewSuperTest(router *gin.Engine, test *testing.T) *supertest {
	return &supertest{router: router, test: test}
}

/**
* @description -> http client for get request
*/

func(ctx *supertest) Get(url string) (*httptest.ResponseRecorder)  {

	time.Sleep(time.Second * 1)

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte(nil)))

	if err != nil {
		ctx.test.Error(err.Error())
		return nil
	}

	rr := httptest.NewRecorder()
	ctx.router.ServeHTTP(rr, req)

	ctx.request.httpRequest = req
	ctx.response.httpResponse = rr

	return rr
}

/**
* @description -> http client for post request
*/

func(ctx *supertest) Post(url string) core {
	ctx.payload.path = url
	ctx.payload.method = http.MethodPost
	return ctx
}

/**
* @description -> http client for delete request
*/

func(ctx *supertest) Delete(url string) (*httptest.ResponseRecorder)  {

	time.Sleep(time.Second * 1)

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte(nil)))

	if err != nil {
		ctx.test.Error(err.Error())
		return nil
	}

	rr := httptest.NewRecorder()
	ctx.router.ServeHTTP(rr, req)

	ctx.request.httpRequest = req
	ctx.response.httpResponse = rr

	return rr
}

/**
* @description -> http client for put request
*/

func(ctx *supertest) Put(url string) core {
	ctx.payload.path = url
	ctx.payload.method = http.MethodPost
	return ctx
}