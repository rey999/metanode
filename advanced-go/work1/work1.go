package main

func add(a *int) {
	*a += 10
}

func main() {
	a := 10
	add(&a)
	println(a)
}
