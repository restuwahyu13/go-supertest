package supertest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

// type core interface {
// 	Send(payload interface{})
// 	End(handle func(rr *httptest.ResponseRecorder))
// 	Set(key, value string)
// 	Auth(key, value string)
// 	Timeout(timeType string, value time.Duration)
// }

type SuperTest interface {
	Get(url string)
	Post(url string)
	Put(url string)
	Delete(url string)
	Send(payload interface{})
	End(handle func(rr *httptest.ResponseRecorder))
	Set(key, value string)
	Auth(key, value string)
	Timeout(timeType string, value time.Duration)
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

type supertest struct {
	router *gin.Engine
	test *testing.T
	payload
	response
	request
}

func NewSuperTest(router *gin.Engine, test *testing.T) *supertest {
	return &supertest{router: router, test: test}
}

/**
* @description -> http client for get request
*/

func(ctx *supertest) Get(url string)  {
	ctx.payload.path = url
	ctx.payload.method = http.MethodGet
	ctx.Send(nil)
}

/**
* @description -> http client for post request
*/

func(ctx *supertest) Post(url string)  {
	ctx.payload.path = url
	ctx.payload.method = http.MethodPost
}

/**
* @description -> http client for delete request
*/

func(ctx *supertest) Delete(url string)   {
	ctx.payload.path = url
	ctx.payload.method = http.MethodDelete
	ctx.Send(nil)
}

/**
* @description -> http client for put request
*/

func(ctx *supertest) Put(url string)  {
	ctx.payload.path = url
	ctx.payload.method = http.MethodPut
}