package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

var once sync.Once
var logger zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {
		file, err := os.OpenFile(
			fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02-15-04-05")),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)

		if err != nil {
			panic(err)
		}

		writer := zerolog.MultiLevelWriter(
			zerolog.ConsoleWriter{Out: os.Stdout},
			file,
		)

		logger = zerolog.New(writer).With().Timestamp().Logger()
	})

	return logger
}
