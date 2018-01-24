package coincheck

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const (
	// APIBaseURL coincheck API endpoint
	APIBaseURL = "https://coincheck.com/api"
)

func init() {
	log.SetFlags(log.Lshortfile)

}

// New returns an instantiated coincheck struct
func New(apiKey, apiSecret string) *Coincheck {
	client := NewClient(apiKey, apiSecret)
	return &Coincheck{client}
}

// NewWithCustomTimeout returns an instantiated coincheck struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Coincheck {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &Coincheck{client}
}

// Coincheck represent a coincheck client
type Coincheck struct {
	client *Client
}

// SetDebug represent a coincheck client
func (b *Coincheck) SetDebug(d bool) {
	b.client.debug = d
}

// GetBalance ..
func (b *Coincheck) GetBalance() (res BalanceResponse, r []byte, err error) {
	r, err = b.client.do("GET", fmt.Sprintf("accounts/balance"), nil, true)
	if err != nil {
		return
	}
	//log.Printf("r:%s", string(r))
	//log.Printf("body=%s", string(r))

	if err = json.Unmarshal(r, &res); err != nil {

		er := ErrorResponse{}
		if err = json.Unmarshal(r, &er); err == nil {
			err = fmt.Errorf("Error: success=%v, %s", er.Success, er.Error)
			return
		}
		log.Printf("body=%s", string(r))
		return
	}

	// ok
	return
}

// GetBalance ..
func (b *Coincheck) GetTicker() (res TickerResponse, r []byte, err error) {
	r, err = b.client.do("GET", fmt.Sprintf("ticker"), nil, false)
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &res); err != nil {
		er := ErrorResponse{}
		if err = json.Unmarshal(r, &er); err == nil {
			err = fmt.Errorf("Error: success=%v, %s", er.Success, er.Error)
			return
		}
		log.Printf("body=%s", string(r))
		return
	}

	// ok
	return
}
