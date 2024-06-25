package app

type config struct {
	apiOnly bool // TODO: must be a better way to put that
}

// reads configs from ENV or file
func GetConfigs() *config {
	return &config{
		apiOnly: false,
	}
}
