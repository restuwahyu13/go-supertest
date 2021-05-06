package supertest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

/**
* @description -> core funcionality
 */

type Options struct {
	Key string
	Value interface{}
}

type payload struct {
	path   string
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
	test   *testing.T
	payload
	response
	request
}

/**
* @description -> parent core funcionality
 */

func NewSuperTest(router *gin.Engine, test *testing.T) *supertest {
	return &supertest{router: router, test: test}
}

/**
* @description -> http client for get request
 */

func (ctx *supertest) Get(url string) {
	ctx.payload.path = url
	ctx.payload.method = http.MethodGet
}

/**
* @description -> http client for post request
 */

func (ctx *supertest) Post(url string) {
	ctx.payload.path = url
	ctx.payload.method = http.MethodPost
}

/**
* @description -> http client for delete request
 */

func (ctx *supertest) Delete(url string) {
	ctx.payload.path = url
	ctx.payload.method = http.MethodDelete
}

/**
* @description -> http client for put request
 */

func (ctx *supertest) Put(url string) {
	ctx.payload.path = url
	ctx.payload.method = http.MethodPut
}

/**
* @description -> http client for patch request
 */

func (ctx *supertest) Patch(url string) {
	ctx.payload.path = url
	ctx.payload.method = http.MethodPatch
}

/**
* @description -> http client for head request
 */

func (ctx *supertest) Head(url string) {
	ctx.payload.path = url
	ctx.payload.method = http.MethodHead
}

/**
* @description -> http client for options request
 */

func (ctx *supertest) Options(url string) {
	ctx.payload.path = url
	ctx.payload.method = http.MethodOptions
}
