package app

// TODO:
// First, gather all necessary configs.
// * which tempates to use
// * what databases should be used
// * ...

type config struct {
	data1 string
}

func GetConfigs() *config {
	return &config{
		data1: "test_data",
	}
}
