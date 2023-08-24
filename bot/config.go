package bot

type Config struct {
	TgBotToken string `toml:"TgBotToken"`
}

func NewConfig() *Config {
	return &Config{
		TgBotToken: "",
	}
}
