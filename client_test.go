package apimaster

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func (suite *ExampleTestSuite) TestExample() {
	testServer := NewTestServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
	}))

	responseValue := &HttpBinGetResponse{}

	response := suite.client.GET("http://httpbin.org/get").Expect()
	response.Status(Equal(200))
	response.ElapsedTime(LessOrEqual(time.Millisecond * 500))
	response.Header(
		Get("Content-Type"), NotEmpty(), Equal("application/json"),
	)
	response.Headers(Contains("Content-Type"))
	response.ContentType(
		NotEmpty(), Equal("application/json"),
	)
	response.ContentLength(NotEqual(200))
	response.Body(
		NotNil(),
		Bind(responseValue),
		IsValid(),
	)

	testServer.Close()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
