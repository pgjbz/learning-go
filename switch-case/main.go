package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux. FUCK YEEEH")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	i := 10

	switch i {
	case 4:
		fmt.Println("meeeh")
	case 4 + 6:
		fmt.Println("uau")
	default:
		fmt.Println("we tryed")
	}

	switch {
	case i < 3:
		fmt.Println("small")
	case i < 5:
		fmt.Println("very very good")
	case i > 6:
		fmt.Println("aberracao da natureza")
	}
}
