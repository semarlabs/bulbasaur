package resources

import (
	"bulbasaur/internal/config"
	"bulbasaur/internal/resources/httpclient"
	"bulbasaur/internal/shared/constants"
	"fmt"
	"net/http"
)

//implement in the new resource
type Resource interface {
	Check() error
	Reset() error
	Close() error
}

type ClientResource interface {
	Resource
	SetRequestHeader(key, value string) error
	Request(method, path string, body []byte) (err error)
	GetResponse() (int, http.Header, []byte, error)
}

func Create(cfg *config.Resource) (Resource, error) {
	switch cfg.Type {
	case constants.ResHttpClient:
		return httpclient.New(cfg)
	default:
		return nil, fmt.Errorf("resource type not found: %s", cfg.Type)
	}
}
