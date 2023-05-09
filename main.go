package main

import (
	"fmt"
	"syscall"

	"github.com/judwhite/go-svc"
)

type program struct {
}

func main() {
	if err := svc.Run(&program{}, syscall.SIGINT, syscall.SIGTERM); err != nil {
		fmt.Println()
	}
}

func (p *program) Init(env svc.Environment) error {
	return nil
}

func (p *program) Start() error {
	return nil
}

func (p *program) Stop() error {
	return nil
}
