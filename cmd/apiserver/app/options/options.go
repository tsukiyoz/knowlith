package options

import (
	"github.com/spf13/pflag"
	"github.com/tsukiyoz/knowlith/internal/apiserver"
	genericoptions "github.com/tsukiyoz/knowlith/pkg/options"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

type ServerOptions struct {
	HttpOptions *genericoptions.HttpOptions `json:"http" mapstructure:"http"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		HttpOptions: genericoptions.NewHttpOptions(),
	}
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	o.HttpOptions.AddFlags(fs)
}

func (o *ServerOptions) Validate() error {
	errs := []error{}
	errs = append(errs, o.HttpOptions.Validate()...)
	return utilerrors.NewAggregate(errs)
}

func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		HttpOptions: o.HttpOptions,
	}, nil
}
