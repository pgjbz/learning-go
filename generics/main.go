package main

type Calculable interface {
    float64 | float32 | uint | int32 | uint32 | int
}

func sum[T Calculable](values []T) T {
    var sum T
    for _, value := range values {
        sum += value
    }
    return sum
}

func main() {
    var ints []int = []int{10, 20, 30, 40, 50}
    var floats []float64 = []float64{10.0, 20.0, 30.0, 40.0, 50.0}
    println(sum(ints))
    println(sum(floats))
}