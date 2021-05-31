package http_test

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-rest-api-kata/config"
	memoryHttp "golang-rest-api-kata/internal/memory/delivery/http"
	"golang-rest-api-kata/internal/memory/entity"
	"golang-rest-api-kata/internal/memory/mock"
	"golang-rest-api-kata/pkg/logger"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMemoryHandlers_CreateInMemory(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)

	cfg := &config.Config{
		Logger: config.Logger{
			Development: true,
		},
	}

	apiLogger := logger.NewApiLogger(cfg)
	memoryHandlers := memoryHttp.NewMemoryHandlers(cfg, service, apiLogger)

	e := &entity.Memory{
		Key:   "test",
		Value: "test",
	}

	service.EXPECT().CreateInMemory("test", "test").Return(e, nil)

	ts := httptest.NewServer(memoryHandlers.CreateInMemory())
	defer ts.Close()

	payload := fmt.Sprintf(`{
	    "key": "test",
    	"value": "test"
	}`)

	resp, err := http.Post(ts.URL+"/in-memory", "application/json", strings.NewReader(payload))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	r := &entity.Memory{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "test", r.Key)
	assert.Equal(t, "test", r.Value)
}

func TestMemoryHandlers_GetInMemory(t *testing.T) {

}
