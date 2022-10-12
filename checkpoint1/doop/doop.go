// Instructions
// Write a program that is called doop.

// The program has to be used with three arguments:

// A value
// An operator, one of : +, -, /, *, %
// Another value
// In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

// The program has to handle the modulo and division operations by 0 as shown on the output examples below.

// Usage
// $ go run .
// $ go run . 1 + 1 | cat -e
// 2$
// $ go run . hello + 1
// $ go run . 1 p 1
// $ go run . 1 / 0 | cat -e
// No division by 0$
// $ go run . 1 % 0 | cat -e
// No modulo by 0$
// $ go run . 9223372036854775807 + 1
// $ go run . -9223372036854775809 - 3
// $ go run . 9223372036854775807 "*" 3
// $ go run . 1 "*" 1
// 1
// $ go run . 1 "*" -1
// -1
// $

package cp1

import (
	"os"
)

func main() {
	f := Atoi(os.Args[1])
	op := os.Args[2]
	s := Atoi(os.Args[3])
	d := 0
	if len(os.Args) != 4 {
		os.Exit(0)
	}
	switch op {
	case "/":
		if s == 0 {
			os.Stdout.WriteString("No division by 0")
			os.Exit(0)
		} else {
			d = f / s
		}
	case "%":
		if s == 0 {
			os.Stdout.WriteString("No modulu by 0")
			os.Exit(0)
		} else {
			d = f % s
		}
	case "*":
		d = f * s
	case "-":
		d = f - s
	case "+":
		d = f + s
	default:
		os.Exit(0)
	}
	if d > 9223372036854775807 || d < -9223372036854775807 {
		os.Exit(0)
	}
	os.Stdout.WriteString(Itoa(d))

}
func Itoa(n int) string {
	num := 0
	str := ""
	var neg bool
	if n < 0 {
		neg = true
		n *= -1
	}
	for n > 0 {
		num = n % 10
		str = string(num+48) + str
		n /= 10
	}
	if neg {
		str = "-" + str
	}
	return str
}
func Atoi(s string) int {
	num := 0
	var neg bool = false
	if s[0] == '-' {
		s = s[1:]
		neg = true
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			os.Exit(0)
		}
		num = num*10 + int(ch-48)
	}
	if neg {
		num *= -1
	}
	return num
}
