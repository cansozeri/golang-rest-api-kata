package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"golang-rest-api-kata/pkg/validator"
	"io"
	"io/ioutil"
	"net/http"
)

// GetConfigPath Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

func ReadRequest(r *http.Request, to interface{}, v *validator.ApiValidator) (err error) {
	if err = r.ParseForm(); err != nil {
		return err
	}

	if err = schema.NewDecoder().Decode(to, r.Form); err != nil {
		return err
	}

	if v != nil {
		if err = v.Validator.Struct(to); err != nil {
			return err
		}
	}
	return err
}

func ReadRequestBody(body io.Reader, to interface{}, v *validator.ApiValidator) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, to)
	if err != nil {
		return fmt.Errorf(err.Error()+" Body is: %s", string(data))
	}

	if v != nil {
		if err = v.Validator.Struct(to); err != nil {
			return err
		}
	}
	return nil
}

func Render(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if data == nil {
		return nil
	}
	return json.NewEncoder(w).Encode(data)
}
