package healthcheck

type Config struct {
	Addr string `yaml:"addr" env:"ADDR" env-default:":8080"`
}
