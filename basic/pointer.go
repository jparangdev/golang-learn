package basic

import "fmt"

type Data struct {
	value int
	data  [200]int
}

type PUser struct {
	Name string
	Age  int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func PointerExample() {
	var data Data
	ChangeData(&data)

	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	//var data2 Data
	//var p *Data = &data2
	var p *Data = &Data{}

	p.value = 123
	p.data[50] = 456
	fmt.Printf("value = %d\n", p.value)
	fmt.Printf("data[100] = %d\n", p.data[100])

	u := NewPUser("John", 25)
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Age: %d\n", u.Age)
}

func NewPUser(name string, age int) *PUser {
	var u = PUser{name, age}
	return &u
}
