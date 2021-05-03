package super

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

func NewSuperTest(router *gin.Engine) *supertest {
	return &supertest{router: router}
}

func(ctx *supertest) GET(url string) (*httptest.ResponseRecorder, error)  {

	time.Sleep(time.Second * 1)

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte(nil)))
	ctx.request.Request = req

	rr := httptest.NewRecorder()
	ctx.response.Response = rr

	ctx.router.ServeHTTP(rr, req)

	return rr,  err
}

func(ctx *supertest) POST(url string) (*httptest.ResponseRecorder, error) {

	time.Sleep(time.Second * 1)

	response, err := json.Marshal(ctx.body.Data)

	if err != nil  {
		logrus.Error(err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(response))
	ctx.request.Request = req

	rr := httptest.NewRecorder()
	ctx.response.Response = rr

	ctx.router.ServeHTTP(rr, req)

	return rr,  err
}

func(ctx *supertest) DELETE(url string) (*httptest.ResponseRecorder, error)  {

	time.Sleep(time.Second * 1)

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte(nil)))
	ctx.request.Request = req

	rr := httptest.NewRecorder()
	ctx.response.Response = rr

	ctx.router.ServeHTTP(rr, req)

	return rr,  err
}

func(ctx *supertest) PUT(url string) (*httptest.ResponseRecorder, error) {

	time.Sleep(time.Second * 1)

	response, err := json.Marshal(ctx.body.Data)

	if err != nil  {
		logrus.Error(err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(response))
	ctx.request.Request = req

	rr := httptest.NewRecorder()
	ctx.response.Response = rr

	ctx.router.ServeHTTP(rr, req)

	return rr,  err
}