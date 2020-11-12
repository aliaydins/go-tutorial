package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Decleration
	var fname, lname string = "John", "Doe"
	var i int = 5

	z := 10.1

	fmt.Printf("%s %s", fname, lname)
	fmt.Println(reflect.TypeOf(i))  //int
	fmt.Printf("%T -- %v \n", i, i) // int -- 5
	fmt.Printf("%T --%f\n", z, z)   // float64--10.100000

	// Block variable decleration
	// var (
	// 	product  = "Mobile"
	// 	quantity = 50
	// )

	// %b	base 2
	// %c	the character represented by the corresponding Unicode code point
	// %d	base 10
	// %o	base 8
	// %s   string
	// %d   integer
	// %q	a single-quoted character literal safely escaped with Go syntax.
	// %x	base 16, with lower-case letters for a-f
	// %X	base 16, with upper-case letters for A-F
	// %U	Unicode format: U+1234; same as "U+%04
	// %T   Type of variable
	// \t   tab
	// \n   new line

	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	// 42       101010          0X2A

	// utf-8
	// for i := 60; i < 122; i++ {
	// 	fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	// }
	// 60 	 111100 	 3c 	 '<'
	// 61 	 111101 	 3d 	 '='
	// 62 	 111110 	 3e 	 '>'
	// 64 	 1000000 	 40 	 '@'
	// 65 	 1000001 	 41 	 'A'

	// a := 10
	// b := "golang"
	// c := 4.17
	// d := true
	// e := "Hello"
	// f := `Do you like my hat?`
	// g := 'A'

	// fmt.Printf("%T - %v\n", a, a) //int
	// fmt.Printf("%T - %v\n", b, b) //string
	// fmt.Printf("%T - %v\n", c, c) //int
	// fmt.Printf("%T - %v\n", d, d) //float64
	// fmt.Printf("%T - %v\n", e, e) //bool
	// fmt.Printf("%T - %v\n", f, f) //string
	// fmt.Printf("%T - %v\n", g, g) //int32 utf8

	// var a int     // 0
	// var s string  // " " space
	// var d bool // false
	// var c float64 //0

	// int		64 bit
	// int8		-128 , 127
	// int16	-32768 , 32767
	// int32	2147483648 , 2147483647
	// int64	-9223372036854775808 , 9223372036854775807
	// uint		64 bit
	// uint8	0 , 255
	// uint16	0 , 65535
	// uint32	0 , 4294967295
	// uint64	0 , 18446744073709551615
	// byte		0 , 255
	// rune		-2147483648 , 2147483647 	like int32

	// >= , <=, == ,!= , && (AND),||(OR),!(NOT)

	//keywords
	// break        default      func         interface    select
	// case         defer        go           map          struct
	// chan         else         goto         package      switch
	// const        fallthrough  if           range        type
	// continue     for          import       return       var

	// const a = 42
	// const b = 42.78
	// const c = "James Bond"

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n", c)

	// Type conversion

	// var test int = 5
	// var test2 string = string(5)

	// fmt.Printf("%T - %v\n", test, test)   // int - 5
	// fmt.Printf("%T - %v\n", test2, test2) // string -  number not using as a string " "

	// Also you can assign function
	// func1 := func() int {
	// 	return 5 + 4
	// }

	// fmt.Printf("%T- %v\n", func1, func1()) //Type is func() restult is 9

}
