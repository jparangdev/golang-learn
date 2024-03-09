package basic

import (
	"fmt"
)

// House struct type, with the properties Address (string), Size (integer),
// Price (float64) and Type (string).
type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

// User struct type, with the properties Name (string), ID (string) and Age (integer).
type User struct {
	Name string
	ID   string
	Age  int
}

// VIPUser struct type. It embeds the User struct, and has additional properties
// VIPLevel (integer) and Price (integer).
type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func StructExample() {

	// Declaring a variable 'house' of type House
	var house House
	// Assigning values to the properties of 'house'
	house.Address = "seoul, south korea"
	house.Type = "apartment"
	house.Size = 28
	house.Price = 9.8

	// Print the properties of 'house'
	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입:", house.Type)

	// Initializing a variable of type User
	user := User{"송하나", "hana", 23}
	// Initializing a variable of type VIPUser
	vip := VIPUser{
		User{"화랑", "hawrang", 40},
		3,
		250,
	}

	// Print the properties of 'user'
	fmt.Printf("User: %s ID: %s Age: %d\n", user.Name, user.ID, user.Age)
	// Print the properties of 'vip'
	fmt.Printf("VIP User: %s ID: %s Age: %d VIP Lv: %d VIP price : %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
