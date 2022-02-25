package string

import "fmt"

func complexNumberMultiply(num1 string, num2 string) string {
	a, b, c, d := 0, 0, 0, 0

	fmt.Sscanf(num1, "%d+%di", &a, &c)
	fmt.Sscanf(num2, "%d+%di", &b, &d)

	return fmt.Sprintf("%d+%di", a*b-c*d, a*d+b*c)
}
