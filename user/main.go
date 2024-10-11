package main

import (
	"context"

	"go-kit-demo/user/internal/cmd"
)

func main() {
	cmd.Main.Run(context.Background())
}
