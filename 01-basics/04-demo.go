package main

import "fmt"

type CalcOperation func(int, int) int

func main() {
	/*
		fmt.Println(add(100, 200))
		fmt.Println(subtract(100, 200))
	*/

	/*
		invokerWithLog(add, 100, 200)
		invokerWithLog(subtract, 100, 200)
	*/
	addWithLog := logInvoker(add)
	subtractWithLog := logInvoker(subtract)

	fmt.Println(addWithLog(100, 200))
	fmt.Println(subtractWithLog(100, 200))
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func invokerWithLog(oper func(int, int) int, x, y int) {
	fmt.Println("invocation started")
	result := oper(x, y)
	fmt.Println(result)
	fmt.Println("invocation completed")
}

/*
func logInvoker(oper func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		fmt.Println("invocation started")
		result := oper(x, y)
		fmt.Println("invocation completed")
		return result
	}
}
*/

func logInvoker(oper CalcOperation) CalcOperation {
	return func(x, y int) int {
		fmt.Println("invocation started")
		result := oper(x, y)
		fmt.Println("invocation completed")
		return result
	}
}

/*
func add(x, y int) int {
	fmt.Println("invocation started")
	result := x + y
	fmt.Println("invocation completed")
	return result
}

func subtract(x, y int) int {
	fmt.Println("invocation started")
	result := x - y
	fmt.Println("invocation completed")
	return result
}
*/
