package apimaster

import (
	"github.com/burakkoken/api-master/body"
	"github.com/burakkoken/api-master/clength"
	"github.com/burakkoken/api-master/ctype"
	"github.com/burakkoken/api-master/header"
	"github.com/burakkoken/api-master/headers"
	"github.com/burakkoken/api-master/status"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestServer(handler http.Handler) *httptest.Server {
	return httptest.NewServer(handler)
}

type ExampleTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *ExampleTestSuite) SetupSuite() {
	suite.client = NewClient(suite.T())
}

type HttpBinGetResponse struct {
	Headers map[string]string `json:"headers" xml:"headers" validate:"required"`
}

type User struct {
	Name string `json:"name" xml:"name"`
}

func (suite *ExampleTestSuite) TestExample() {
	testServer := NewTestServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
	}))

	responseValue := &HttpBinGetResponse{}
	responseValue2 := &HttpBinGetResponse{}

	response := suite.client.POST("http://httpbin.org/post").WithJson(&User{Name: "TEST"}).Expect()
	response.Status(
		status.Equal(200),
	)

	response.Header(
		header.Get("Content-Type"),
		header.NotEmpty(),
		header.Equal("application/json"),
	)
	response.Headers(
		headers.Contains("Content-Type"),
	)
	response.ContentType(
		ctype.NotEmpty(), ctype.Equal("application/json"),
	)
	response.ContentLength(
		clength.NotEqual(200),
	)
	var str string
	response.Body(
		body.NotNil(),
		body.NotEmpty(),
		body.Equal("test"),
		body.Json(responseValue),
		body.Json(responseValue2),
		body.IsValid(),
		body.Text(&str),
	)

	testServer.Close()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
