package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	setupLog()
	cfg, err := readConfig("/never/found/path/ghost.yml")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", cfg)
}

//Config type
type Config struct {
}

//readConfig
func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open config file")
	}

	defer file.Close()

	cfg := &Config{}
	return cfg, nil

}

func setupLog() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	log.SetOutput(out)
}
