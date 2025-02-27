/* Copyright © 2025 Milos Zivlak

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

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
