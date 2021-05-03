## Go SuperTest

Go Supertest is HTTP Client Testing only for Gin Framework, inspired by Supertest package library HTTP Client Testing for
Express.js Framework.

### Example Usage

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func main() {
	router := SetupRouter()
	router.Run()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	router.GET("/", GetMethod)
	router.POST("/", PostMethod)

	return router
}
```

### Example Usage Test

```go
package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	super "github.com/restuwahyu13/go-supertest/supertest"
	util "github.com/restuwahyu13/go-supertest/utils"
)

var router = SetupRouter()

func TestGetMethod(t *testing.T) {
	supertest := super.NewSuperTest(router)

	_, err := supertest.GET("/")
	supertest.Set("Content-Type", "application/json")

	if err != nil {
		t.Error(err)
	}

	supertest.End(func(rr *httptest.ResponseRecorder) {

    response := util.Parse(rr.Body.Bytes())
		t.Log(response)

		assert.Equal(t, rr.Code, response.StatusCode)
    assert.Equal(t, http.MethodGet, response.Method)
    assert.Equal(t, "fetch request using get method", response.Message)
	})
}

func TestPostMethod(t *testing.T) {
	supertest := super.NewSuperTest(router)

	payload := gin.H{
		"email" : "restuwahyu13@gmail.com",
		"password" : "bukopin12",
	}

	_, err := supertest.POST("/")
	supertest.Set("Content-Type", "application/json")
  supertest.Send(payload)

	if err != nil {
		t.Error(err)
	}

	supertest.End(func(rr *httptest.ResponseRecorder) {

    response := util.Parse(rr.Body.Bytes())
		t.Log(response)

		assert.Equal(t, rr.Code, response.StatusCode)
    assert.Equal(t, http.MethodPost, response.Method)
    assert.Equal(t, "fetch request using post method", response.Message)
	})
}
```
