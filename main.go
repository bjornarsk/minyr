package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bjornarsk/minyr/yr"
	// "strconv"
	// "strings"
	// "github.com/bjornarsk/funtemps/conv"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Venligst velg convert, average eller exit:")

		if !scanner.Scan() {
			break
		}
		input = scanner.Text()

		switch input {
		case "q", "exit":
			fmt.Println("exit")
			return

		case "convert":
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit...")
			// funksjon som gjør åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil
			yr.ConvertTemperature()

		case "average":
			fmt.Println("Gjennomsnitt-kalkulator")

			for {
				yr.AverageTemperature()

				var input2 string
				scanjn := bufio.NewScanner(os.Stdin)
				fmt.Println("Avslutte? (j/n)")
				for scanjn.Scan() {
					input2 = scanjn.Text()
					if input2 == "j" {
						break
					} else if input2 == "n" {
						break
					}
				}
				if input2 == "j" {
					break
				}
			}
		}
	}
	fmt.Println("Avslutter program.")
}
