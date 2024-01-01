package app

// here is what I do:
// First, I need to gather all necessary configs
// This will include:
// * which tempates to use
// * what database should be used
// * else...

type config struct {
	data1 string
}

func GetConfigs() *config {
	return &config{
		data1: "test_data",
	}
}
