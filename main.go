package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/inhuman/geo-currency-rate/api"
	"github.com/inhuman/geo-currency-rate/config"
	"github.com/inhuman/geo-currency-rate/log"
	"github.com/inhuman/geo-currency-rate/rate_client"
	"github.com/rs/zerolog"
	"io/ioutil"
	"os"
)

func main() {

	l := log.NewLogger()

	err := realMain(l)
	if err != nil {
		l.Err(err).Str("source", "main").Send()
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func realMain(l *zerolog.Logger) error {

	fileName := flag.String("file", "config.json", "path to json config")
	flag.Parse()

	jsonFile, err := os.Open(*fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			l.Err(err).Str("source", "main").Send()
		}
	}()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	conf := config.Config{}

	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return err
	}

	l.Debug().Str("config", fmt.Sprintf("%+v", conf)).Send()

	rateClient, err := rate_client.NewRateClient(conf)
	if err != nil {
		return err
	}

	opts := api.Opts{
		Logger:     l,
		Config:     &conf,
		RateClient: rateClient,
	}

	if err := api.Run(opts); err != nil {
		return err
	}

	return nil
}
