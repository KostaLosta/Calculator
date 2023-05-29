package printResult

import "fmt"

func PrintResult(sum int) {
	result := fmt.Sprintf("Ваш результат: %d", sum)

	fmt.Println(result)
}
