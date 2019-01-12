package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var result string

	fmt.Println("Roman Numeral Converter by Jugdish Schatz")
	fmt.Println("-----------------------------------------")

	for {

		log.Println("Enter a number for conversion")
		reader := bufio.NewReader(os.Stdin)
		dodgeInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		dodgeInput = strings.TrimSpace(dodgeInput)
		input, err := strconv.ParseInt(dodgeInput, 10, 32)
		if err != nil {
			log.Println("A number... Whats wrong with you???")
		} else if input > 3999 {
			log.Println("Number must be smaller that 3999")
		} else {

			if input/1000 != 0 {

				th := input / 1000
				for x := th; x != 0; x-- {
					result += "M"
				}

				input = input % 1000
			}

			if input/100 != 0 {

				hu := input / 100
				if hu == 9 {
					result += "CM"
				} else if hu > 5 && hu < 9 {
					result += "D"
					hu -= 5
					for x := hu; x != 0; x-- {
						result += "C"
					}
				} else if hu == 5 {
					result += "D"
				} else if hu == 4 {
					result += "CD"
				} else {
					for x := hu; x != 0; x-- {
						result += "C"
					}
				}

				input = input % 100
			}

			if input/10 != 0 {

				te := input / 10
				if te == 9 {
					result += "XC"
				} else if te > 5 && te < 9 {
					result += "L"
					te -= 5
					for x := te; x != 0; x-- {
						result += "X"
					}
				} else if te == 5 {
					result += "L"
				} else if te == 4 {
					result += "XL"
				} else {
					for x := te; x != 0; x-- {
						result += "X"
					}
				}
				input = input % 10
			}

			un := input / 1
			if un == 9 {
				result += "IX"
			} else if un > 5 && un < 9 {
				result += "V"
				un -= 5
				for x := un; x != 0; x-- {
					result += "I"
				}
			} else if un == 5 {
				result += "V"
			} else if un == 4 {
				result += "IV"
			} else {
				for x := un; x != 0; x-- {
					result += "I"
				}
			}

			log.Println("Your result. Take it and go!!: " + result)

			result = ""

		}
	}
}
