package controllers

import (
	"fmt"
	"github.com/inhuman/geo-currency-rate/config"
	"github.com/inhuman/geo-currency-rate/geometry"
	"github.com/inhuman/geo-currency-rate/rate_client"
	"github.com/rs/zerolog"
	"time"
)

type RateRequestOpts struct {
	Conf       *config.Config
	Logger     *zerolog.Logger
	Point      *geometry.Point
	RateClient *rate_client.RateClient
}

const dateFormat = "02.01.2006"

func ProcessRateRequest(opts RateRequestOpts) (*rate_client.RateResponse, error) {

	if !geometry.IsInRadius(opts.Conf.Radius, *opts.Point) {
		return nil, fmt.Errorf("point {%f, %f} not in radius %f",
			opts.Point.X, opts.Point.Y, opts.Conf.Radius)
	}

	quadrant := geometry.QuadrantDetector(*opts.Point)

	currentTime := time.Now()

	switch quadrant {

	case 1:

		// date today
		d := currentTime.Format(dateFormat)
		return opts.RateClient.GetRate(d)

	case 2:

		// date yesterday
		d := currentTime.Add(-24 * time.Hour).Format(dateFormat)
		return opts.RateClient.GetRate(d)

	case 3:

		//  date day before yesterday
		d := currentTime.Add(-48 * time.Hour).Format(dateFormat)
		return opts.RateClient.GetRate(d)

	case 4:

		// date tomorrow
		d := currentTime.Add(24 * time.Hour).Format(dateFormat)
		return opts.RateClient.GetRate(d)

	default:
		return nil, fmt.Errorf("invalid quadrant %d for point {%f, %f}",
			quadrant, opts.Point.X, opts.Point.Y)

	}
}
