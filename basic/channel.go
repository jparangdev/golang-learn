package basic

import (
	"fmt"
	"sync"
	"time"
)

var wgCh sync.WaitGroup
var startTime = time.Now()

func ChannelExample() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)
	fmt.Println("Start Factory")

	wgCh.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)
	wgCh.Wait()
	fmt.Println("Close the factory")
}

type Car struct {
	Body  string
	Tire  string
	Color string
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(time.Second * 10)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Volvo XC90"
			tireCh <- car
		case <-after:
			close(tireCh)
			wgCh.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Michelin"
		paintCh <- car
	}
	wgCh.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wgCh.Done()
}
