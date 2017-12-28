package coincheck

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBalance(t *testing.T) {
	key := os.Getenv("COINCHECK_KEY")
	secret := os.Getenv("COINCHECK_SECRET")

	api := New(key, secret)
	ret, body, err := api.GetBalance()
	require.NoError(t, err, nil)
	log.Printf("err:%v", err)
	log.Printf("ret:%v", ret)
	log.Printf("body:%s", string(body))

	return
}
