package basic

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ErrorExample() {
	errorExample1()
	readEq("123 3")
	readEq("123 abc")

}

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return fmt.Sprintf("Password length is %d, but should be at least %d", err.Len, err.RequireLen)
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{Len: len(password), RequireLen: 8}
	}
	return nil
}

func errorExample1() {
	err := RegisterAccount("myID", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("Account registered successfully!")
	}
}

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, scanner.Err()
	}
	word := scanner.Text()
	num, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s err:%w", word, err)
	}
	return num, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}
