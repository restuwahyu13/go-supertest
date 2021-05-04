## Go SuperTest

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/restuwahyu13/go-supertest?style=flat-square) [![Go Report Card](https://goreportcard.com/badge/github.com/restuwahyu13/go-supertest)](https://goreportcard.com/report/github.com/restuwahyu13/go-supertest) ![GitHub issues](https://img.shields.io/github/issues/restuwahyu13/go-supertest?style=flat-square) ![GitHub closed issues](https://img.shields.io/github/issues-closed/restuwahyu13/go-supertest?style=flat-square) ![GitHub contributors](https://img.shields.io/github/contributors/restuwahyu13/go-supertest?style=flat-square)

Go Supertest is minimalize HTTP Client Testing only for Gin Framework, inspired by [Supertest](https://www.npmjs.com/package/supertest) package library HTTP Client Testing for Express.js Framework.

- [API Documentation](#go-supertest)
  * [Installation](#installation)
  * [API Reference](#api-reference)
    + [NewSuperTest](#newsupertest)
    + [Get](#get-url-string-)
    + [Post](#post-url-string-)
    + [Delete](#delete-url-string-)
    + [Put](#put-url-string-)
    + [Patch](#patch-url-string-)
    + [Head](#head-url-string-)
    + [Options](#options-url-string-)
    + [Send](#send-payload-interface-)
    + [End](#end-handle-funcrr-httptestresponserecorder-)
    + [Set](#set-key-value-string-)
    + [Auth](#auth-key-value-string-)
    + [Timeout](#timeout-timetype-string-value-timeduration-)
  * [Example Usage](#example-usage)
    + [Main Setup](#main-setup)
    + [Test Setup](#test-setup)
  * [API Status Reference](#api-status-reference)
  * [Bugs](#bugs)
  * [Contributing](#contributing)
  * [License](#license)
 
## Installation

```sh
$ go get -u https://github.com/restuwahyu13/go-supertest
```

### API Reference

**Important** if you use http request using `Get` or `Delete` you must be use `Send` with nil value and `Send` must be added before `Set`, please check example usage about this package is working.

  - #### NewSuperTest( router *gin.Engine, test *testing.T )
  
    | Method       | Description                             |
    | ------------ | --------------------------------------- |
    | NewSuperTest | instance method for create unit testing | 

  - #### Get( url string )
  
    | Method | Description                                       |
    | ------ | ------------------------------------------------- |
    | Get    | for handling request http client using GET method | 

  - #### Post( url string )
    
    | Method | Description                                        | 
    | ------ | -------------------------------------------------- |
    | Post   | for handling request http client using POST method |

  - #### Delete( url string )
  
    | Method | Description                                          | 
    | ------ | ---------------------------------------------------- |
    | Delete | for handling request http client using Delete method |

  - #### Put( url string )
  
    | Method | Description                                       |
    | ------ | ------------------------------------------------- |
    | Put    | for handling request http client using Put method | 

  - #### Patch( url string )
  
    | Method | Description                                         |
    | ------ | --------------------------------------------------- |
    | Patch  | for handling request http client using Patch method | 

  - #### Head( url string )
  
    | Method | Description                                        |
    | ------ | -------------------------------------------------- |
    | Head   | for handling request http client using Head method | 

  - #### Options( url string )
  
    | Method  | Description                                           |
    | ------- | ----------------------------------------------------- |
    | Options | for handling request http client using Options method | 

  - #### Send( payload interface{} )
  
    | Method | Description                        |
    | ------ | ---------------------------------- |
    | Send   | send request data needed by client | 

  - #### End( handle func(rr *httptest.ResponseRecorder) )
  
    | Method | Description                          |
    | ------ | ------------------------------------ |
    | End    | receive responses from http requests | 

  - #### Set( key, value string )
  
    | Method | Description                                              |
    | ------ | -------------------------------------------------------- |
    | Set    | set your headers before sending http request into client | 

  - #### Auth( key, value string )
  
    | Method | Description                                                              |
    | ------ | ------------------------------------------------------------------------ |
    | Auth   | pass the username or password if you are using HTTP Basic authentication | 

  - #### Timeout( timeType string, value time.Duration )
  
    | Method  | Description                                  |
    | ------- | -------------------------------------------- |
    | Timeout | set delay before sending request into client | 


### Example Usage

- #### Main Setup

  ```go
  package main
  
  type User struct {
    Name string `json:"name"`
  }

  func GetMethod(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
      "statusCode": 200,
      "method":     http.MethodGet,
      "message":    "fetch request using get method",
      "data":       nil,
    })
  }

  func PostMethod(ctx *gin.Context) {
    var input User
    ctx.ShouldBindJSON(&input)

    ctx.JSON(http.StatusOK, gin.H{
      "statusCode": 200,
      "method":     http.MethodPost,
      "message":    "fetch request using post method",
      "data":       input,
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

    delete(userData, "name-"+userId)

    ctx.JSON(http.StatusOK, gin.H{
      "statusCode": 200,
      "method":     http.MethodPost,
      "message":    "fetch request using delete method",
      "data":       userData,
    })
  }
  
  func main() {
    router := SetupRouter
    router.Run()
  }
  
  func SetupRouter() *gin.Engine {
    router := gin.Default()
    
    gin.SetMode(gin.TestMode)

    router.GET("/", GetMethod)
    router.POST("/", PostMethod)
    router.DELETE("/:id", DeleteMethod)
    router.PUT("/:id", PutMethod)

    return router
  }
  ```


- #### Test Setup

  ```go
  package main

  var router = SetupRouter()

  func TestGetMethod(t *testing.T) {
    test := NewSuperTest(router, t)

    test.Get("/")
    test.Send(nil)
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
      "name": "restu wahyu saputra",
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
    test.Send(nil)
    test.Set("Content-Type", "application/json")
    test.End(func(rr *httptest.ResponseRecorder) {

      response := util.Parse(rr.Body.Bytes())

      assert.Equal(t, rr.Code, http.StatusOK)
      assert.Equal(t, http.MethodPost, response.Method)
      assert.Equal(t, "fetch request using delete method", response.Message)

      encoded, _ := json.Marshal(response.Data)

      var mapping map[string]interface{}
      json.Unmarshal(encoded, &mapping)

      assert.Equal(t, 4, len(mapping))
    })
  }

  func TestPutMethod(t *testing.T) {
    test := NewSuperTest(router, t)

    payload := gin.H{
      "name": "restu wahyu saputra",
    }

    test.Put("/" + fmt.Sprintf("%v", 1))
    test.Send(payload)
    test.Set("Content-Type", "application/json")
    test.End(func(rr *httptest.ResponseRecorder) {

      response := util.Parse(rr.Body.Bytes())

      assert.Equal(t, rr.Code, http.StatusOK)
      assert.Equal(t, http.MethodPost, response.Method)
      assert.Equal(t, "fetch request using put method", response.Message)

      encoded, _ := json.Marshal(response.Data)

      var mapping map[string]interface{}
      json.Unmarshal(encoded, &mapping)

      assert.Equal(t, "restu wahyu saputra", mapping["name-1"])
    })
  }
  ```

### API Status Reference

| Name    | Ready | Description                                                                      |
| ------- | ----- | -------------------------------------------------------------------------------- |
| Get     | *yes* | for handling request http client using GET method                                |
| Post    | *yes* | for handling request http client using POST method                               |
| Delete  | *yes* | for handling request http client using Delete method                             |
| Put     | *yes* | for handling request http client using Put method                                |
| Patch   | *yes* | for handling request http client using Patch method                              |
| Head    | *yes* | for handling request http client using Head method                               |
| Options | *yes* | for handling request http client using Options method                            |
| Send    | *yes* | send request data needed by client                                               |
| End     | *yes* | receive responses from http requests                                             |
| Set     | *yes* | set your headers before sending http request into client                         |
| Auth    | *yes* | pass the username or password if you are using HTTP Basic authentication         |
| Timeout | *yes* | set delay before sending request into client                                     |
| Expect  | *yes* | expect the given data to match                                                   |
| Attach  | *no*  | handle requests from files or image uploads if you are using multipart/form-data |
| Field   | *no*  | handle data submitted from form/field if you are using multipart/form-data       | 

### Bugs

For information on bugs related to package libraries, please visit [here](https://github.com/restuwahyu13/go-supertest/issues)

### Contributing

Want to make **Go Supertest** more perfect ? Let's contribute and follow the [contribution guide.](https://github.com/restuwahyu13/go-supertest/blob/main/CONTRIBUTING.md)

### License

- [MIT License](https://github.com/restuwahyu13/go-supertest/blob/main/LICENSE.md)

<p align="right" style="padding: 5px; border-radius: 100%; background-color: red; font-size: 2rem;">
  <b><a href="#go-supertest">BACK TO TOP</a></b>
</p>