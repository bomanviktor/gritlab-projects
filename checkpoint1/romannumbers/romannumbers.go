// Instructions:
// Write a program called rn. The objective is to convert a number, given as an argument, into a roman number and print it with its roman number calculation.
//
// The program should have a limit of 4000. In case of an invalid number, for example "hello" or 0 the program should print ERROR: cannot convert to roman digit.

// allowed imports: fmt.Printf, os.*

package cp1

import (
	"fmt"
	"os"
)

const r1, r5, r10, r50, r100, r500, r1000 = "I", "V", "X", "L", "C", "D", "M"
const error string = "ERROR: cannot convert to roman digit."

func rn(s string) {
	if len(os.Args) != 2 {
		fmt.Printf(error)
		os.Exit(0)
	}
	output := ""
	calc := ""
	leng := len(os.Args[1])
	if leng == 0 || leng == 1 && os.Args[1] == "0" {
		fmt.Printf(error)
		os.Exit(0)
	}
	for _, ch := range os.Args[1] {
		if ch == '0' {
			leng--
			continue
		}
		switch leng {
		case 4:
			if ch >= '4' {
				fmt.Printf(error)
				os.Exit(0)
			} else {
				for i := '0'; i < '4'; i++ {
					output += r1000
					calc += r1000 + "+"
				}
			}
		case 3:
			if ch == '9' {
				output += r100 + r1000
				calc += "(" + r100 + "-" + r1000 + ")+"
			} else if ch == '4' {
				output += r100 + r500
				calc += "(" + r100 + "-" + r500 + ")+"
			} else if ch >= '5' && ch < '9' {
				output += r500
				calc += r500 + "+"
				for i := '5'; i < ch; i++ {
					output += r100
					calc += r100 + "+"
				}
			} else if ch >= '0' && ch < '4' {
				for i := '4'; i > ch; i-- {
					output += r100
					calc += r100 + "+"
				}
			}
		case 2:
			if ch == '9' {
				output += r10 + r100
				calc += "(" + r10 + "-" + r100 + ")+"
			} else if ch == '4' {
				output += r10 + r50
				calc += "(" + r10 + "-" + r50 + ")+"
			} else if ch >= '5' && ch < '9' {
				output += r50
				calc += r50 + "+"
				for i := '5'; i < ch; i++ {
					output += r10
					calc += r10 + "+"
				}
			} else if ch >= '0' && ch < '4' {
				for i := '4'; i > ch; i-- {
					output += r10
					calc += r10 + "+"
				}
			}
		case 1:
			if ch == '9' {
				output += r1 + r10
				calc += "(" + r1 + "-" + r10 + ")+"
			} else if ch == '4' {
				output += r1 + r5
				calc += "(" + r1 + "-" + r5 + ")+"
			} else if ch >= '5' && ch < '9' {
				output += r5
				calc += r5 + "+"
				for i := '5'; i < ch; i++ {
					output += r1
					calc += r1 + "+"
				}
			} else if ch >= '0' && ch < '4' {
				for i := '4'; i > ch; i-- {
					output += r1
					calc += r1 + "+"
				}
			}

		}
		leng--
	}
	fmt.Printf(output + "\n")
	fmt.Printf(calc[:len(calc)-1])
}
