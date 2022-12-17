package command

import (
	"context"
	"fmt"
	"os"
)

const (
	exitOK  = 0
	exitErr = 1
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	fmt.Println(44444)

	return exitOK
}
