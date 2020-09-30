package config

type Config struct {
	CurrencyRateApi string  `json:"currency_rate_api"`
	Radius          float64 `json:"radius"`
	CurrencyCode    string  `json:"currency_code"`
}
