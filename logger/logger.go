package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var file *os.File

func Load() error {
	if _, err := os.ReadDir("logs"); err != nil {
		return err
	}

	name := "logs/" + time.Now().Format(time.DateTime) + ".log"
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(nil)
	}

	log.SetOutput(io.MultiWriter(os.Stdout, file))

	return nil
}

func Unload() error {
	return file.Close()
}
