package main

import (
	"fmt"
	conf "github.com/basharal/config"
	"github.com/kerwinan/go/kit/log"
	"github.com/kerwinan/kwina-demo/bootstrap"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.X.With("err", r, "stack", string(debug.Stack())).Errorf("panic")
		}
		log.X.Flush()
	}()

	if err := bootstrap.Run(); err != nil {
		log.X.With("err", err).Errorf("bootstrap error")
	}
}
