package body

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/burakkoken/api-master/expect"
	v "github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var validator = v.New()

const HeaderContentType = "Content-Type"

const (
	ContentTypeApplicationJson = "application/json"
	ContentTypeApplicationXml  = "application/xml"
)

const ContextKeyBody = "ContextKeyBody"
const ContextKeyBoundInstance = "ContextKeyBoundInstance"

type Expect func(t *testing.T, chain *expect.Chain, response *http.Response)

func Json(value interface{}) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		err := json.Unmarshal(chain.GetValue().([]byte), value)
		if err != nil {
			assert.NoError(t, err)
		}
		chain.GetContext().Put(ContextKeyBoundInstance, value)
	}
}

func Xml(value interface{}) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		err := xml.Unmarshal(chain.GetValue().([]byte), value)
		if err != nil {
			assert.NoError(t, err)
		}
		chain.GetContext().Put(ContextKeyBoundInstance, value)
	}
}

func Text(value *string) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		*value = string(chain.GetValue().([]byte))
		chain.GetContext().Put(ContextKeyBoundInstance, *value)
	}
}

func Nil() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Empty(t, chain.GetValue().([]byte))
	}
}

func NotNil() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotEmpty(t, chain.GetValue())
	}
}

func Empty() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Empty(t, string(chain.GetValue().([]byte)))
	}
}

func NotEmpty() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotEmpty(t, string(chain.GetValue().([]byte)))
	}
}

func Equal(value string) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Equal(t, value, string(chain.GetValue().([]byte)))
	}
}

func NotEqual(value string) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotEqual(t, value, string(chain.GetValue().([]byte)))
	}
}

func IsValid() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		boundInstance := chain.GetContext().Get(ContextKeyBoundInstance)

		switch boundInstance.(type) {
		case string:
			return
		default:
			validate(t, boundInstance)
		}

	}
}

func validate(t *testing.T, value interface{}) {
	err := validator.Struct(value)

	if err != nil {

		var errMessage string
		if _, ok := err.(*v.InvalidValidationError); ok {
			errMessage = err.Error()

		} else {
			for index, err := range err.(v.ValidationErrors) {

				errMessage += errMessage +
					fmt.Sprintf("%d. '%s' validation error on %s (%s), \nTag Value: %v, \nValue : %v\n",
						index+1,
						err.Tag(),
						err.StructNamespace(),
						err.Type(),
						err.Param(),
						err.Value(),
					)

			}
		}

		assert.NoError(t, err, errMessage)
	}
}
