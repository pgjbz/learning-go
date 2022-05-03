package main

func main() {
	sum := 0
	for i := 0; i < 10; i++ { //for in Go dont have paren
		sum += i
	}
	println(sum)

	sum = 1

	for sum < 1000 { //while cond
		sum += sum
	}

	println(sum)

	sum = 0
	for { //for is while

		sum += 1
		if sum >= 100 {
			break
		}

	}

}
