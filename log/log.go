package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

func NewLogger() *zerolog.Logger {

	f, err := os.OpenFile("service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	l := zerolog.New(f).With().Timestamp().Logger()
	return &l
}
