package supertest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	util "github.com/restuwahyu13/go-supertest/utils"
)

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func GetMethod(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"method": http.MethodGet,
		"message": "fetch request using get method",
		"data": nil,
	})
}

func PostMethod(ctx *gin.Context) {
	var input User
	ctx.ShouldBindJSON(&input)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"method": http.MethodPost,
		"message": "fetch request using post method",
		"data": input,
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	router.GET("/", GetMethod)
	router.POST("/", PostMethod)

	router.Run()

	return router
}

var router = SetupRouter()

func TestGetMethod(t *testing.T) {
	test := NewSuperTest(router, t)

	test.Get("/")
	test.Set("Content-Type", "application/json")
	test.End(func(rr *httptest.ResponseRecorder) {

		response := util.Parse(rr.Body.Bytes())

		assert.Equal(t, rr.Code, http.StatusOK)
    assert.Equal(t, http.MethodGet, response.Method)
    assert.Equal(t, "fetch request using get method", response.Message)
	})
}

func TestPostMethod(t *testing.T) {
	test := NewSuperTest(router, t)

	payload := gin.H{
		"email" : "restuwahyu13@gmail.com",
		"password" : "bukopin12",
	}

	test.Post("/")
	test.Send(payload)
	test.Set("Content-Type", "application/json")
	test.End(func(rr *httptest.ResponseRecorder) {

		response := util.Parse(rr.Body.Bytes())

		assert.Equal(t, rr.Code, http.StatusOK)
    assert.Equal(t, http.MethodPost, response.Method)
    assert.Equal(t, "fetch request using post method", response.Message)
	})
}