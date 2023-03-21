package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bjornarsk/funtemps/conv"
)

func main() {
	// Open the file
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(file)

	// Create the output file
	outputFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Create a writer to write to the output file
	writer := bufio.NewWriter(outputFile)

	// Write the first line to the output file
	if scanner.Scan() {
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	var outputLine string

	// Loop through each remaining line in the input file
	for scanner.Scan() {
		line := scanner.Text()

		// Skip the line if it's empty
		if line == "" {
			continue
		}

		// Split the line into fields separated by ";"
		fields := strings.Split(line, ";")

		// Extract the last field from the line, if there is one
		var lastField string
		if len(fields) > 0 {
			lastField = fields[len(fields)-1]
		}

		// Convert the last field to Fahrenheit, if there is one
		var convertedField string
		if lastField != "" {
			var err error
			convertedField, err = convertLastField(lastField)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
		}

		// Replace the original last field with the converted field, if there is one
		if convertedField != "" {
			fields[len(fields)-1] = convertedField
		}

		if line[0:7] == "Data er" {
			outputLine = "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Bjørnar"
		} else {
			// Put the fields back together with semicolons in between
			outputLine = strings.Join(fields, ";")
		}

		// Write the output line to the output file
		_, err = writer.WriteString(outputLine + "\n")
		if err != nil {
			panic(err)
		}
	}

	// Flush the writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func convertLastField(lastField string) (string, error) {
	// Return an error if the last field is empty
	if lastField == "" {
		return "", fmt.Errorf("last field is empty")
	}

	// Convert the last field to a float
	celsius, err := strconv.ParseFloat(lastField, 64)
	if err != nil {
		return "", err
	}

	// Convert Celsius to Fahrenheit
	fahrenheit := conv.CelsiusToFahrenheit(celsius)

	// Convert the Fahrenheit temperature back to a string
	return fmt.Sprintf("%.1f", fahrenheit), nil
}
