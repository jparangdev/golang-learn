package basic

import (
	"fmt"
)

type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcle\n", parcel)
}

type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다. \n", parcel)
}

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func InterfaceExample() {
	koreaPostSender := &PostSender{}
	SendBook("어린 왕자", koreaPostSender)

	fedexSender := &FedexSender{}
	SendBook("Harry Potter", fedexSender)
}
