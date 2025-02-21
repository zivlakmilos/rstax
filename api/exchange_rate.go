package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const ExchangeRateBaseUrl = "https://kurs.resenje.org/api/v1/currencies/"

type ExchangeRate struct {
	Code           string  `json:"code"`
	Date           string  `json:"date"`
	DateFrom       string  `json:"date_from"`
	Number         int     `json:"number"`
	Parity         int     `json:"parity"`
	Cash_buy       float64 `json:"cash_buy"`
	Cash_sell      float64 `json:"cash_sell"`
	ExchangeBuy    float64 `json:"exchange_buy"`
	ExchangeMiddle float64 `json:"exchange_middle"`
	ExchangeSell   float64 `json:"exchange_sell"`
}

func GetExchangeRate(currency string, date time.Time) (*ExchangeRate, error) {
	url := fmt.Sprintf("%s%s/rates/%s", ExchangeRateBaseUrl, currency, date.Format("2006-01-02"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ExchangeRate
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
