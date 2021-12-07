package utils

func SumInts(arr []int, manipulator func(n int) int) int {
	sum := 0
	for _, n := range arr {
		if manipulator != nil {
			sum += manipulator(n)
		} else {
			sum += n
		}
	}
	return sum
}

func ConvertStringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i] = ConvertStringToInt(str)
	}
	return ints
}
