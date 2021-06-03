package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-rest-api-kata/pkg/utils"
	"golang-rest-api-kata/pkg/validator"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	t.Parallel()

	configPath := "docker"
	assert.Equal(t, "./config/config-docker", utils.GetConfigPath(configPath))

	configPath = ""
	assert.Equal(t, "./config/config-local", utils.GetConfigPath(configPath))

}

func TestReadRequest(t *testing.T) {
	type Test struct {
		Test string `validate:"required"`
	}

	request := httptest.NewRequest("GET", "/test?test=1", nil)
	validate, _ := validator.NewValidate()
	err := utils.ReadRequest(request, &Test{}, validate)

	assert.Equal(t, nil, err)
}

func TestReadRequestBody(t *testing.T) {

	type Test struct {
		Test string `json:"test" validate:"required"`
	}

	payload := fmt.Sprintf(`{"test": "2016-01-01"}`)

	request := httptest.NewRequest("POST", "/test", strings.NewReader(payload))
	validate, _ := validator.NewValidate()
	err := utils.ReadRequestBody(request.Body, &Test{}, validate)

	assert.Equal(t, nil, err)
}

func TestRender(t *testing.T) {

	type Test struct {
		Test string `json:"test" validate:"required"`
	}

	payload := &Test{
		Test: "test",
	}

	w := httptest.NewRecorder()

	err := utils.Render(w, 200, payload)

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, w.Code)
}
