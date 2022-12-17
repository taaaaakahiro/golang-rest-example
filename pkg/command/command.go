package command

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
)

const (
	exitOK  = 0
	exitErr = 1
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	logger := log.Default()
	logger.Print("Run App")

	cnf := config.LoadCofig()
	fmt.Print(cnf.Port)

	return exitOK
}
