package main

import (
	"context"

	"go-kit-demo/gateway/internal/cmd"
)

func main() {
	cmd.Main.Run(context.Background())
}
