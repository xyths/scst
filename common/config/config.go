package config

type PairConfig struct {
	A        string
	B        string
	AmountIn float64
}

type CoralConfig struct {
	Url     string // like "http://127.0.0.1:8545"
	Address string
}

type MonitorGlobalConfig struct {
	Coral CoralConfig
	Pairs []PairConfig
}
