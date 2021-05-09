package bootstrap

import (
	"github.com/kerwinan/go/kit/log"
	"os"
	"os/signal"
	"syscall"
)

func Run() error {
	p, err := setupProvider()
	if err != nil {
		return err
	}
	defer p.Close()
	pch := p.Run()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-ch:
		log.X.With("signal", sig).Infof("receive exit signal")
		return nil
	case err := <-pch:
		return err
	}
}
