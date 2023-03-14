package main

func main() {
	x := 0
	if x >= 0 {
		println("x")
	}
	println("-----------------")
	if n := "abc"; x > 0 {
		println(n[2])
		println(string(n[2]))
	} else if x < 0 {
		println(n[1])
		println(string(n[1]))
	} else {
		println(n[0])
		println(string(n[0]))
	}
}
