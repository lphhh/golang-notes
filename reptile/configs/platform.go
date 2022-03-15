package configs

const (
	Binance = iota + 1
	Ftx
)

const (
	// SucceededPush is log block
	Succeeded = "succeeded"
	// FailedPush is log block
	Failed = "failed"
)

func GetPlatforms() []string {
	var platforms = []string{"Binance", "Ftx"}
	return platforms
}

func GetPlatform(platform int) string {
	switch platform {
	case Binance:
		return "Binance"
	case Ftx:
		return "Ftx"
	default:
		return "unknown"
	}
}
