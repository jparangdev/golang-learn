package basic

import "fmt"

func StringsExample() {

	str1 := "Hello\t'World'\n"
	str2 := `Go is "awesome"!\nGo is simple and \t 'powerful'`

	fmt.Println(str1)
	fmt.Println(str2)

	//utf-8 is default
	poet1 := "죽는 날 까지 하늘을 우러러 \n 한 점 부끄럼이 없기를, \n잎새에 이는 바람에도\n나는 괴로워했다.\n"
	poet2 := `죽는 날 까지 하늘을 우러러
한 점 부끄럼이 없기를, 
잎새에 이는 바람에도
나는 괴로워했다.`

	fmt.Println(poet1)
	fmt.Println(poet2)

	var char rune = '한'
	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)

	str3 := "가나다라마바사"
	str4 := "abcdefg"
	fmt.Printf("len(str3) = %d\n", len(str3))
	fmt.Printf("len(str4) = %d\n", len(str4))

	runes := []rune(str3)
	fmt.Printf("len(runes) = %d\n", len(runes))

	str5 := "Hello 월드"
	for i := 0; i < len(str5); i++ {
		fmt.Printf("Type:%T value:%d char:%c\n", str5[i], str5[i], str5[i])
	}
	arr := []rune(str5)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Type:%T value:%d char:%c\n", arr[i], arr[i], arr[i])
	}

	for _, v := range str5 {
		fmt.Printf("Type:%T value:%d char:%c\n", v, v, v)
	}

	str6 := "Hello"
	str7 := "World"
	fmt.Println(str6 + " " + str7)
	fmt.Println(len(str6 + str7))

	str8 := "Hello"
	str9 := "Hell"
	str10 := "Hello"

	fmt.Printf("%s == %s : %v\n", str8, str9, str8 == str9)
	fmt.Printf("%s != %s : %v\n", str8, str9, str8 != str9)
	fmt.Printf("%s == %s : %v\n", str8, str10, str8 == str10)
	fmt.Printf("%s != %s : %v\n", str8, str10, str8 != str10)

	str11 := "BBB"
	str12 := "aaaaAAA"
	str13 := "BBAD"
	str14 := "ZZZ"
	fmt.Printf("%s > %s : %v\n", str11, str12, str11 > str12)
	fmt.Printf("%s < %s : %v\n", str11, str13, str11 < str13)
	fmt.Printf("%s <= %s : %v\n", str11, str14, str11 <= str14)

}
