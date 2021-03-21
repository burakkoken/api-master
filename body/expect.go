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

		switch chain.GetValue().(type) {
		case []byte:
			err := json.Unmarshal(chain.GetValue().([]byte), value)

			if err != nil {
				assert.NoError(t, err)
				return
			}

			chain.Next(string(chain.GetValue().([]byte)))
		case string:
			err := json.Unmarshal([]byte(chain.GetValue().(string)), value)

			if err != nil {
				assert.NoError(t, err)
				return
			}

			chain.Next(chain.GetValue())
		}

		chain.GetContext().Put(ContextKeyBoundInstance, value)
	}
}

func Xml(value interface{}) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {

		switch chain.GetValue().(type) {
		case []byte:
			err := xml.Unmarshal(chain.GetValue().([]byte), value)

			if err != nil {
				assert.NoError(t, err)
				return
			}

			chain.Next(string(chain.GetValue().([]byte)))
		case string:
			err := xml.Unmarshal([]byte(chain.GetValue().(string)), value)

			if err != nil {
				assert.NoError(t, err)
				return
			}

			chain.Next(chain.GetValue())
		}

		chain.GetContext().Put(ContextKeyBoundInstance, value)
	}
}

func Text(value *string) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {

		switch chain.GetValue().(type) {
		case []byte:
			*value = string(chain.GetValue().([]byte))
		case string:
			*value = chain.GetValue().(string)
		}

		chain.Next(*value)
		chain.GetContext().Put(ContextKeyBoundInstance, *value)
	}
}

func Nil() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.Nil(t, chain.GetValue())
	}
}

func NotNil() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {
		assert.NotNil(t, chain.GetValue())
	}
}

func Empty() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {

		switch chain.GetValue().(type) {
		case string:
			assert.Empty(t, chain.GetValue())
		case []byte:
			assert.Empty(t, string(chain.GetValue().([]byte)))
		}

	}
}

func NotEmpty() Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {

		switch chain.GetValue().(type) {
		case string:
			assert.NotEmpty(t, chain.GetValue())
		case []byte:
			assert.NotEmpty(t, string(chain.GetValue().([]byte)))
		}

	}
}

func Equal(value string) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {

		switch chain.GetValue().(type) {
		case string:
			assert.Equal(t, value, chain.GetValue())
		case []byte:
			assert.Equal(t, value, string(chain.GetValue().([]byte)))
		}

	}
}

func NotEqual(value string) Expect {
	return func(t *testing.T, chain *expect.Chain, response *http.Response) {

		switch chain.GetValue().(type) {
		case string:
			assert.NotEqual(t, value, chain.GetValue())
		case []byte:
			assert.NotEqual(t, value, string(chain.GetValue().([]byte)))
		}

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
