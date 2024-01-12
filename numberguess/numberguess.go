package numberguess

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Game() {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	cnt := 1
	for {
		fmt.Printf("input number :")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("only input number")
		} else {
			if n > r {
				fmt.Println("Too high!")
			} else if n < r {
				fmt.Println("Too low!")
			} else {
				fmt.Println("Congratulations! You guessed the number correctly! count : ", cnt)
				break
			}
			cnt++
		}
	}
}

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
