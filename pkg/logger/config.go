package logger

type Config struct {
	Level      string            `yaml:"level" env:"LEVEL" env-default:"info"`
	Format     string            `json:"format" yaml:"format" env:"FORMAT" env-default:"json"`
	DevMode    bool              `yaml:"devMode" env:"DEV_MODE" env-default:"false"`
	BaseFields map[string]string `yaml:"baseFields" env:"BASE_FIELDS"`
}
