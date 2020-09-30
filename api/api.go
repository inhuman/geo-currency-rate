package api

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/inhuman/geo-currency-rate/config"
	"github.com/inhuman/geo-currency-rate/controllers"
	"github.com/inhuman/geo-currency-rate/geometry"
	"github.com/inhuman/geo-currency-rate/rate_client"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type Opts struct {
	Logger     *zerolog.Logger
	Config     *config.Config
	RateClient *rate_client.RateClient
}

type Container struct {
	conf       *config.Config
	logger     *zerolog.Logger
	rateClient *rate_client.RateClient
}

func Run(opts Opts) error {

	con := &Container{
		conf:       opts.Config,
		logger:     opts.Logger,
		rateClient: opts.RateClient,
	}

	r := gin.New()

	apiLog := opts.Logger.With().Str("source", "api").Logger()

	r.Use(logger.SetLogger(logger.Config{
		Logger: &apiLog,
		UTC:    true,
	}))

	r.GET("/rates", con.GetRate)

	return r.Run(":8080")
}

func (con *Container) GetRate(c *gin.Context) {

	xStr := c.Query("x")
	yStr := c.Query("y")

	x, err := strconv.ParseFloat(xStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	y, err := strconv.ParseFloat(yStr, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := controllers.ProcessRateRequest(controllers.RateRequestOpts{
		Conf: con.conf,
		Point: &geometry.Point{
			X: x,
			Y: y,
		},
		RateClient: con.rateClient,
		Logger:     con.logger,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, resp)
}
