package options

import (
	"time"

	"github.com/spf13/pflag"
)

var _ IOptions = (*HttpOptions)(nil)

type HttpOptions struct {
	Network string        `json:"network" mapstructure:"network"`
	Addr    string        `json:"addr" mapstructure:"addr"`
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
}

func NewHttpOptions() *HttpOptions {
	return &HttpOptions{
		Network: "tcp",
		Addr:    "0.0.0.0:8080",
		Timeout: 30 * time.Second,
	}
}

func (o *HttpOptions) Validate() []error {
	if o == nil {
		return nil
	}

	errors := []error{}
	if err := ValidateAddress(o.Addr); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func (o *HttpOptions) AddFlags(fs *pflag.FlagSet, prefixs ...string) {
	fs.StringVar(&o.Network, "http.network", o.Network, "Specify the network for the HTTP server.")
	fs.StringVar(&o.Addr, "http.addr", o.Addr, "Specify the HTTP server bind address and port.")
	fs.DurationVar(&o.Timeout, "http.timeout", o.Timeout, "Timeout for server connections.")
}

func (o *HttpOptions) Complete() error {
	return nil
}
