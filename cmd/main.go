package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
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
	apiKey := "lgtrEDSGBbFm3eEoQL6oCrCNXRaWzAGJL4jAdvZIrtqqApWaBMKfW7WDnObIhGAt"
	apiSecret := "qrw3iV9BynFEdwrMmmP4vkvusWeLCv1fC6O50veodZ5gBRZ4uKFBNoYeVn2v0Xpz"

	resp, err := apiClient.TradeApi.AccountInformationUSERDATA(ctx, timestamp, apiSecret, "", apiKey)
	if err != nil {
		return
	}
	localVarBody, err := ioutil.ReadAll(resp.Body)

	logrus.Println("res %s %v", string(localVarBody), err)
}
