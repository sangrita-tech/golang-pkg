package healthcheck

type Configs struct {
	Addr string `yaml:"addr" env:"ADDR" env-default:":8080"`
}
