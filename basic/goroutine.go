package basic

import (
	"fmt"
	"sync"
	"time"
)

func GoroutineExample() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
	fmt.Println("")

	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go SumAtoB(i, 1, 1000000000)
	}

	wg.Wait()
	fmt.Println("all calculation is done")
}

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '사', '아', '자'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

var wg sync.WaitGroup

func SumAtoB(n, a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d => %d to %d sum is %d\n", n, a, b, sum)
	wg.Done()
}
