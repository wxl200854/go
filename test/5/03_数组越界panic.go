package main

func test(x int) {
	var a [10]int
	a[x] = 100
}

func main() {
	test(20)
}
