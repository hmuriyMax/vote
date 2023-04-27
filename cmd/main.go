package main

import (
	"context"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	ctx.Err()
}
