package options

type HttpOptions struct {
	Addr string `json:"addr" mapstructure:"addr"`
}

func NewHttpOptions() *HttpOptions {
	return &HttpOptions{
		Addr: "0.0.0.0:8080",
	}
}
