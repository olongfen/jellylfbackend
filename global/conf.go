package global

type Configs struct {
	HTTPPort uint   `mapstructure:"HTTP_PORT" `
	DBDsn    string `mapstructure:"DB_DSN"`
}
