package q3

import (
	"fmt"
	"sort"
)

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}

	sort.Ints(numbers)

	menor := numbers[0]
	maior := numbers[len(numbers)-1]
	sum := 0

	for i := 1; i < len(numbers)-1; i++ {
		sum += numbers[i]
	}

	return menor, maior, float64(sum / (len(numbers) - 2)), nil
}
