package response_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ywardhana/goapi/response"
)

func TestWrite(t *testing.T) {
	var (
		data = struct {
			Message string `json:"message"`
		}{
			Message: "test message",
		}

		jsonDecode struct {
			Data struct {
				Message string `json:"message"`
			} `json:"data"`
			Meta struct {
				HttpStatus int `json:"http_status"`
			} `json:"meta"`
		}
	)

	meta := response.MetaInfo{
		HttpStatus: http.StatusOK,
	}

	resp := response.Response{
		Meta: meta,
		Data: data,
	}
	writer := httptest.NewRecorder()
	response.Write(writer, resp)

	result := writer.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		assert.Nil(t, err)
	}

	err = json.Unmarshal(body, &jsonDecode)
	if err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, data.Message, jsonDecode.Data.Message)
}

func TestOKWithMeta(t *testing.T) {
	var (
		data = struct {
			Message string `json:"message"`
		}{
			Message: "test message",
		}
		meta = response.MetaInfo{
			Offset: 10,
			Limit:  20,
		}

		jsonDecode struct {
			Data struct {
				Message string `json:"message"`
			} `json:"data"`
			Meta struct {
				HttpStatus int `json:"http_status"`
				Offset     int `json:"offset"`
				Limit      int `json:"limit"`
			} `json:"meta"`
		}
	)

	writer := httptest.NewRecorder()
	response.OKWithMeta(writer, data, "", meta)

	result := writer.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		assert.Nil(t, err)
	}

	err = json.Unmarshal(body, &jsonDecode)
	if err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, meta.Limit, jsonDecode.Meta.Limit)
	assert.Equal(t, meta.Offset, jsonDecode.Meta.Offset)
	assert.Equal(t, data.Message, jsonDecode.Data.Message)
}

func TestOK(t *testing.T) {
	var (
		data = struct {
			Message string `json:"message"`
		}{
			Message: "test message",
		}

		jsonDecode struct {
			Data struct {
				Message string `json:"message"`
			} `json:"data"`
			Meta struct {
				HttpStatus int `json:"http_status"`
			} `json:"meta"`
		}
	)

	writer := httptest.NewRecorder()
	response.OK(writer, data, "")

	result := writer.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		assert.Nil(t, err)
	}

	err = json.Unmarshal(body, &jsonDecode)
	if err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, data.Message, jsonDecode.Data.Message)
}

func TestCreated(t *testing.T) {
	var (
		data = struct {
			Message string `json:"message"`
		}{
			Message: "test message",
		}

		jsonDecode struct {
			Data struct {
				Message string `json:"message"`
			} `json:"data"`
			Meta struct {
				HttpStatus int `json:"http_status"`
			} `json:"meta"`
		}
	)

	writer := httptest.NewRecorder()
	response.Created(writer, data, "")

	result := writer.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		assert.Nil(t, err)
	}

	err = json.Unmarshal(body, &jsonDecode)
	if err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, http.StatusCreated, result.StatusCode)
	assert.Equal(t, data.Message, jsonDecode.Data.Message)
}

func TestError(t *testing.T) {
	var (
		jsonDecode struct {
			Message string `json:"message"`
			Meta    struct {
				HttpStatus int `json:"http_status"`
			} `json:"meta"`
		}
	)

	testError := errors.New("test error")
	writer := httptest.NewRecorder()
	response.Error(writer, testError, http.StatusUnprocessableEntity)

	result := writer.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		assert.Nil(t, err)
	}

	err = json.Unmarshal(body, &jsonDecode)
	if err != nil {
		assert.Nil(t, err)
	}
	log.Println(string(body))
	assert.Equal(t, http.StatusUnprocessableEntity, result.StatusCode)
	assert.Equal(t, testError.Error(), jsonDecode.Message)
}
