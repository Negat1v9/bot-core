package bot

type Config struct {
	TgBotToken string `toml:"TgBotToken"`
	DBPath     string `toml:"db_path"`
}

func NewConfig() *Config {
	return &Config{
		TgBotToken: "",
	}
}
