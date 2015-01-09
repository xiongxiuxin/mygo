package main

import "fmt"

func main() {
	_, _, nickName := getName()
	fmt.Println(nickName)
	var value1 complex64 // 由2个float32构成的复数类型
	value1 = 3.2 + 12i
	value2 := 3.2 + 12i        // value2是complex128类型
	value3 := complex(3.2, 12) // value3结果同 value2
	fmt.Println(value1, value2, value3)
	fmt.Println(real(value3), imag(value3))

}

func getName() (firstName, lastName, nickName string) {
	return "xiong", "xiuxin", "xxx"
}
