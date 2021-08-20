package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"strconv"
	"time"

	binanceSpot "binance-trade-bot/pkg/binance_spot_api_v1"
)

func init() {
	logFile := false
	if logFile {
		fileWriter := lumberjack.Logger{}
		logrus.SetOutput(&fileWriter)
	}

	logrus.SetLevel(logrus.TraceLevel)
}

func main() {
	logrus.Debugf("Starting..")
	defer logrus.Debugf("Shutdown.")

	cfg := &binanceSpot.Configuration{
		BasePath: "https://testnet.binance.vision",
		// Scheme:,
		//DefaultHeader:,
		// UserAgent:,
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives:   false,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     time.Second * 10,
				TLSHandshakeTimeout: time.Second,
			},
			Timeout: time.Second * 10,
		},
	}
	apiClient := binanceSpot.NewAPIClient(cfg)

	ctx := context.Background()
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	apiSecret := "KTCu2dTCp66eMUCR45RIuUiQkftR7hrzmxYg13i0EBuo1DWhFyJeg1WjQVpTMdfG"
	apiKey := "WuKmhxvuuB70ALYjFoKZMEQVLcQKrqPb3hpzOk8HKJG41zBo7szhuGgeJsQcYXaM"

	signature := buildSignature(, []byte(apiSecret))
	resp, err := apiClient.TradeApi.AccountInformationUSERDATA(ctx, timestamp, signature, "", apiKey)

	logrus.Println("res %v %v", resp, err)
}

func buildSignature(data []byte, apiSecret []byte) string {
	h := hmac.New(sha256.New, apiSecret)
	h.Write(data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
	// return crypto.createHmac('sha256', config.API_SECRET).update(data).digest('hex')
}
