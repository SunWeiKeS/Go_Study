package main

func main() {
	for i := 0; i < 5; i++ {
		println("i=", i)
	}
	for j := 4; j >= 0; j-- {
		println("j=", j)
	}

	x := 0
	for x < 5 { //相当于while(x<5){...}
		println(x)
		x++
	}

	x1 := 4
	for { //相当于 while(true){...}
		println(x1)
		x1--
		if x < 0 {
			break
		}
	}
}
