package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type service struct {
	input  int64
	result string
}

type Service interface {
	thousands(input int64)
	hundreds(input int64)
	tens(input int64)
	units(input int64)
}

func main() {

	s := &service{}

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
		convInput, err := strconv.ParseInt(dodgeInput, 10, 32)

		s.input = convInput
		if err != nil {
			log.Println("A number... Whats wrong with you???")
		} else if s.input > 3999 {
			log.Println("Number must be smaller that 3999")
		} else {

			if s.input/1000 != 0 {

				s.thousands(s.input)
			}

			if s.input/100 != 0 {

				s.hundreds(s.input)
			}

			if s.input/10 != 0 {

				s.tens(s.input)
			}

			s.units(s.input)

			log.Println("Your result. Take it and go!!: " + s.result)

			s.result = ""

		}
	}
}

func (s *service) thousands(input int64) {
	th := input / 1000
	for x := th; x != 0; x-- {
		s.result += "M"
	}

	s.input = input % 1000
}

func (s *service) hundreds(input int64) {
	hu := input / 100
	if hu == 9 {
		s.result += "CM"
	} else if hu > 5 && hu < 9 {
		s.result += "D"
		hu -= 5
		for x := hu; x != 0; x-- {
			s.result += "C"
		}
	} else if hu == 5 {
		s.result += "D"
	} else if hu == 4 {
		s.result += "CD"
	} else {
		for x := hu; x != 0; x-- {
			s.result += "C"
		}
	}

	s.input = input % 100
}

func (s *service) tens(input int64) {
	te := input / 10
	if te == 9 {
		s.result += "XC"
	} else if te > 5 && te < 9 {
		s.result += "L"
		te -= 5
		for x := te; x != 0; x-- {
			s.result += "X"
		}
	} else if te == 5 {
		s.result += "L"
	} else if te == 4 {
		s.result += "XL"
	} else {
		for x := te; x != 0; x-- {
			s.result += "X"
		}
	}
	s.input = input % 10
}

func (s *service) units(input int64) {
	un := input / 1
	if un == 9 {
		s.result += "IX"
	} else if un > 5 && un < 9 {
		s.result += "V"
		un -= 5
		for x := un; x != 0; x-- {
			s.result += "I"
		}
	} else if un == 5 {
		s.result += "V"
	} else if un == 4 {
		s.result += "IV"
	} else {
		for x := un; x != 0; x-- {
			s.result += "I"
		}
	}
}
