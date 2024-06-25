package app

type config struct {
	serveStatic bool
	port        string
	address     string
}

// TODO: reads configs from ENV or file
func GetConfigs() *config {
	return &config{
		serveStatic: false,
		port:        "8080",
		address:     "localhost",
	}
}
