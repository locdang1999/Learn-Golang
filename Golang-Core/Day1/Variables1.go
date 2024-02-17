package main

import (
	"fmt"
	"strconv"
)

/*
* Định nghĩa biến trong lập trình
* Cú pháp khai bao biến
* Global and block scope
* Shadow
* Declared and not used
* Export global scope
* Naming convention --> camel
* Convert Type
 */

var age4 int = 4

var (
	n int    = 10
	m int    = 20
	h string = "Lộc Đặng"
)

var Address string = "Tp.Hồ Chí Minh" // export var

func main() {
	//1. Biến là vùng nhớ được cấp phát bởi chương trình dùng để lưu trữ giá trị của 1 kiểu dữ liệu nào đó tại 1 thời điểm nhất định và được truy xuất bởi tên biến

	//2. Cú pháp
	var age int
	age = 0

	var age1 int = 1

	age2 := 2 //toán tử := có ý nghĩa vừa khai báo vừa gán value cho biến
	age3 := 2.5
	name := "Loc"
	fmt.Println(age, age1, age2)
	fmt.Printf("%v, %T \n", age1, age1) //%v: xuất ra giá trị, %V: xuất ra kiểu của biến
	fmt.Printf("%v, %T \n", age3, age3)
	fmt.Printf("%v, %T \n", name, name)
	fmt.Println(age4, n, m, h)

	//3.Global and block scope
	/* Block là các biến nằm trong hàm
	 * Global có thể khai báo bên ngoài hàm main() và chỉ khai báo theo kiểu thứ 2
	 * Có thể khai báo nhiều biến trong 1 Global scope
	 */

	//4.Shadow: Khai báo 1 biến trong block scope trùng với global scope thì biến trong block scope sẽ shadow lại
	//tức nghĩa biến khai báo lại đó chỉ có tác dụng trong nội khu block scope mà không ảnh hưởng đến global scope
	var n int = 50
	fmt.Println(n)

	//5: Declared and not used: khi khai báo biến thì phải dùng nếu ko goalng sẽ đưa lỗi vì để sẽ tốn bộ nhớ
	//name1 := "Hello"
	fmt.Println("Hello")

	//6: Export global scope: Muốn export 1 biến global ra bên ngoài cho package khác dùng thì viết hoa biến đó lên chỉ cần chữ đầu tiên

	//7: Naming convention --> camel

	//8: Convert Type

	var num1 int = 10

	var num2 float32 = float32(num1)

	var num3 float64 = 10.5

	var text13 string = "3.5"

	var num4 int = int(num3)

	//Int --> String
	var text string = strconv.Itoa(num1)

	//String --> Float
	//fl1, error := strconv.ParseFloat(text13, 64)
	fl1, _ := strconv.ParseFloat(text13, 64)
	//var fl1, _ = strconv.ParseFloat(text13, 64)

	//Float --> String
	var num5 float64 = 121.123456789
	//text1 := strconv.FormatFloat(num3, 'g', 6, 64)
	var text1 string = strconv.FormatFloat(num5, 'g', 12, 64)
	var text2 string = fmt.Sprintf("%v", num3)

	fmt.Printf("%v, %T \n", num1, num1)
	fmt.Printf("%v, %T \n", num2, num2)
	fmt.Printf("%v, %T \n", num3, num3)
	fmt.Printf("%v, %T \n", num4, num4)
	fmt.Printf("%v, %T \n", text, text)
	//fmt.Printf("%v, %T \n", text1, text1)
	//if error != nil {
	//	fmt.Println(error)
	//} else {
	//
	//}
	fmt.Printf("%v, %T \n", fl1, fl1)
	fmt.Printf("%v, %T \n", text1, text1)
	fmt.Printf("%v, %T \n", text2, text2)
}
