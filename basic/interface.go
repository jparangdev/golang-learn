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

	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{"john", 15})
	s := &Student{"Hu", 17}
	PrintNameAndAge(s)
}

type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Stringer interface {
	String() string
}

func PrintNameAndAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("Age : %d\n", s.Age)
}
