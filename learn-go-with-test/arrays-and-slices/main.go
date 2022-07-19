package main

func Sum(ns []int) int {
	i := 0
	for _, v := range ns {
		i += v
	}
	return i
}

func SumAll(args ...[]int) []int {
	var sums []int
	for _, arg := range args {
		sums = append(sums, Sum(arg))
	}
	return sums
}

func SumAllTail(args ...[]int) []int {
	var sums []int
	for _, arg := range args {
		if len(arg) == 0 {
			sums = append(sums, 0)
			continue
		}
		sums = append(sums, Sum(arg[1:]))
	}
	return sums
}

func main() {

}
