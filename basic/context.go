package basic

import (
	"context"
	"fmt"
)

func ContextExample() {
	wg.Add(1)
	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)
	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg.Done()
}
