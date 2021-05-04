package supertest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	util "github.com/restuwahyu13/go-supertest/utils"
)

type User struct {
	Name string `json:"name"`
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

func DeleteMethod(ctx *gin.Context) {

	userId := ctx.Param("id")
	userData := make(map[string]string, 5)

	userData["name-1"] = "john doe"
	userData["name-2"] = "jane doe"
	userData["name-3"] = "james bond"
	userData["name-4"] = "curt cobain"
	userData["name-5"] = "rorona zoro"

	delete(userData, "name-" + userId)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"method": http.MethodPost,
		"message": "fetch request using delete method",
		"data": userData,
	})
}

func PutMethod(ctx *gin.Context) {

	var input User

	userId := ctx.Param("id")
	ctx.ShouldBindJSON(&input)

	userData := make(map[string]string, 5)

	userData["name-1"] = "john doe"
	userData["name-2"] = "jane doe"
	userData["name-3"] = "james bond"
	userData["name-4"] = "curt cobain"
	userData["name-5"] = "rorona zoro"

	userData["name-" + userId] = input.Name

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"method": http.MethodPost,
		"message": "fetch request using put method",
		"data": userData,
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	router.GET("/", GetMethod)
	router.POST("/", PostMethod)
	router.DELETE("/:id", DeleteMethod)
	router.PUT("/:id", PutMethod)

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
		"name" : "restu wahyu saputra",
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

func TestDeleteMethod(t *testing.T) {
	test := NewSuperTest(router, t)

	test.Delete("/" + fmt.Sprintf("%v", 5))
	test.Set("Content-Type", "application/json")
	test.End(func(rr *httptest.ResponseRecorder) {

		response := util.Parse(rr.Body.Bytes())

		assert.Equal(t, rr.Code, http.StatusOK)
    assert.Equal(t, http.MethodPost, response.Method)
    assert.Equal(t, "fetch request using delete method", response.Message)

		encoded , _:= json.Marshal(response.Data)

		var mapping map[string]interface{}
		json.Unmarshal(encoded, &mapping)

		assert.Equal(t, 4, len(mapping))
	})
}

func TestPutMethod(t *testing.T) {
	test := NewSuperTest(router, t)

	payload := gin.H{
		"name" : "restu wahyu saputra",
	}

	test.Put("/" + fmt.Sprintf("%v", 1))
	test.Send(payload)
	test.Set("Content-Type", "application/json")
	test.End(func(rr *httptest.ResponseRecorder) {

		response := util.Parse(rr.Body.Bytes())

		assert.Equal(t, rr.Code, http.StatusOK)
    assert.Equal(t, http.MethodPost, response.Method)
    assert.Equal(t, "fetch request using put method", response.Message)

		encoded , _:= json.Marshal(response.Data)

		var mapping map[string]interface{}
		json.Unmarshal(encoded, &mapping)

		assert.Equal(t, "restu wahyu saputra", mapping["name-1"])
	})
}