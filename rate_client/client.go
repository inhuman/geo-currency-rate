package rate_client

import (
	"encoding/xml"
	"fmt"
	"github.com/inhuman/geo-currency-rate/config"
	"golang.org/x/text/encoding/charmap"
	"io"
	"net/http"
	"time"
)

type RateClient struct {
	ApiUrl       string
	CurrencyCode string
	http.Client
}

func NewRateClient(conf config.Config) (*RateClient, error) {
	httpClient := http.Client{
		Timeout: 1500 * time.Millisecond,
	}

	return &RateClient{
		ApiUrl:       conf.CurrencyRateApi,
		CurrencyCode: conf.CurrencyCode,
		Client:       httpClient,
	}, nil
}

type RateResponse struct {
	Date         string `json:"date"`
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

func (c *RateClient) GetRate(date string) (*RateResponse, error) {

	req, err := c.buildGetRateRequest(date)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rates := ValCurs{}

	err = c.decodeRateResponse(&rates, resp)
	if err != nil {
		return nil, err
	}

	for _, rate := range rates.Valute {
		if rate.CharCode == c.CurrencyCode {
			return &RateResponse{
				Date:         date,
				CurrencyCode: c.CurrencyCode,
				Value:        rate.Value,
			}, nil

		}
	}

	return nil, fmt.Errorf("currency code %s not available", c.CurrencyCode)
}

func (c *RateClient) buildGetRateRequest(date string) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.ApiUrl, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("date_req", date)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (c *RateClient) decodeRateResponse(rates *ValCurs, resp *http.Response) error {
	d := xml.NewDecoder(resp.Body)

	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}
	return d.Decode(&rates)
}

type ValCurs struct {
	Valute []Valute
}

type Valute struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}
