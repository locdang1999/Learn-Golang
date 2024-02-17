package main

import "fmt"

/*
* Primitive Data Types:  Kiểu dữ liệu nguyên thủy
* 1. Boolean
* 2. Numberics: Integer, Float, Complex
* 3. Text: String, Byte, Rune
 */

func main() {
	//1> Boolean
	var a bool = true //mặc định là false
	//var b bool = 1 == 1
	//var b bool = 1 != 2
	//var b bool = 1 <= 2
	//var b bool = 1 >= 2
	var b bool = 1 > 2
	//var b bool = 1 < 2

	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)

	//1> Numberics:
	//Integer
	//-Signed integer: int8 int16 int32 int64
	//-Unsigned integer: uint8 uint16 uint32 uint64

	var num int8 = 32
	var num1 int8 = -32
	fmt.Printf("%v, %T \n", num, num)
	fmt.Printf("%v, %T \n", num1, num1)
	//Thao tác BIT
	fmt.Printf("%v, %T \n", num&num1, num&num1)
	fmt.Printf("%v, %T \n", num|num1, num|num1)
	fmt.Printf("%v, %T \n", num^num1, num^num1)
	fmt.Printf("%v, %T \n", num&^num1, num&^num1)
	//Dich bit
	fmt.Printf("%v, %T \n", num<<3, num<<3)
	fmt.Printf("%v, %T \n", num>>3, num>>3)
	//Float
	//Float32
	//Float64
}
