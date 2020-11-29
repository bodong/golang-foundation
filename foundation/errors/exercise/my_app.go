package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "err : %s\n", err)
		os.Exit(1)
	}
}

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)

	if err != nil {
		return errors.Wrap(err, "unable to open pid file (is server running?)")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warn: can't remove pid file - %s", err)
	}

	strPid := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPid)
	if err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}
