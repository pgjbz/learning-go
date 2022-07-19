package iteration

func Repeat(value string, qtt int) string {
	finalValue := value
	for i := 0; i < qtt-1; i++ {
		finalValue += value
	}
	return finalValue
}
