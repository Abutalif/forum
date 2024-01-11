package app

type config struct {
	data1 string
}

func GetConfigs() *config {
	return &config{
		data1: "test_data",
	}
}
