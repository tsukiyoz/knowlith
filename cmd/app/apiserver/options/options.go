package options

import (
	"github.com/tsukiyoz/knowlith/internal/apiserver"
	genericoptions "github.com/tsukiyoz/knowlith/pkg/options"
)

type ServerOptions struct {
	HttpOptions *genericoptions.HttpOptions `json:"http" mapstructure:"http"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		HttpOptions: genericoptions.NewHttpOptions(),
	}
}

func (o *ServerOptions) Validate() error {
	return nil
}

func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		HttpOptions: o.HttpOptions,
	}, nil
}
