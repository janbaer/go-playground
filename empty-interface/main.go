package main

import "fmt"

func main() {
	test("Hello world")
	test(1)
	test(2.2)
	testAny("Hello world")
}

func test(t interface{}) {
	// fmt.Printf("%#v %T\n", t, t)
	switch t := t.(type) {
	case string:
		fmt.Printf("t is a string: %s\n", t)
	case int:
		fmt.Printf("t is a int: %d\n", t)
	default:
		fmt.Printf("t is %T: %v\n", t, t)
	}
}

func testAny(t any) {
	switch t := t.(type) {
	case string:
		fmt.Printf("t is a string: %s\n", t)
	case int:
		fmt.Printf("t is a int: %d\n", t)
	default:
		fmt.Printf("t is %T: %v\n", t, t)
	}
}
