package db

type (
	Config struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		Driver   string
	}
)
