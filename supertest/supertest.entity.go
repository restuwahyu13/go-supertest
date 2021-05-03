package super

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)


type body struct {
	Data interface{}
}

type response struct {
	Response *httptest.ResponseRecorder
}

type request struct {
	Request *http.Request
}

type supertest struct {
	router *gin.Engine
	test *testing.T
	body
	response
	request
}