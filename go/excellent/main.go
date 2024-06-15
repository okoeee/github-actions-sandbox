package main

func main() {
	EvenOrOdd(10)
	println("Hello, World!")
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
