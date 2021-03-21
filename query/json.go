package query

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)
import "github.com/tidwall/gjson"

type JsonQuery struct {
	json string
	t    *testing.T
}

func NewJsonQuery(t *testing.T, data []byte) *JsonQuery {
	var jsonMap map[string]interface{}
	err := json.Unmarshal(data, &jsonMap)

	if err != nil {
		panic("Invalid Json : " + err.Error())
	}

	sanitizedData, _ := json.Marshal(jsonMap)

	return &JsonQuery{
		string(sanitizedData),
		t,
	}
}

func (query *JsonQuery) Get(path string) *JsonQuery {
	result := gjson.Get(query.json, path)
	query.json = result.String()
	return query
}

func (query *JsonQuery) Empty() *JsonQuery {
	assert.Empty(query.t, query.json)
	return query
}

func (query *JsonQuery) NotEmpty() *JsonQuery {
	assert.NotEmpty(query.t, query.json)
	return query
}

func (query *JsonQuery) Equal(value interface{}) *JsonQuery {
	assert.Equal(query.t, query.toJsonString(value), query.json)
	return query
}

func (query *JsonQuery) NotEqual(value interface{}) *JsonQuery {
	assert.NotEqual(query.t, query.toJsonString(value), value)
	return query
}

func (query *JsonQuery) Len(length int) *JsonQuery {
	assert.NotEmpty(query.t, query.json, "Json object must not be empty")

	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(query.json), &jsonMap)

	if err != nil {
		panic("Invalid json : " + err.Error())
	}

	assert.Len(query.t, jsonMap, length)
	return query
}

func (query *JsonQuery) Bind(value interface{}) *JsonQuery {
	err := json.Unmarshal([]byte(query.json), value)

	if err != nil {
		assert.NoError(query.t, err)
	}

	return query
}

func (query *JsonQuery) String() string {
	return query.json
}

func (query *JsonQuery) toJsonString(value interface{}) string {
	jsonValue, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}

	return string(jsonValue)
}
