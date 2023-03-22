package yr

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/bjornarsk/funtemps/conv"
)

func ConvertTemperature() {
	// Check if the file already exists
	if _, err := os.Stat("output-test.csv"); err == nil {
		fmt.Print("Filen eksisterer allerede. Vil du generere filen på nytt? (j/n): ")
		var overwriteInput string
		fmt.Scanln(&overwriteInput)
		fmt.Println("Genererer filen på nytt...")
		if strings.ToLower(overwriteInput) == "n" {
			fmt.Println("Går tilbake til hovedmeny")
			return
		}
	}

	// Open the input file
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create the output file
	outputFile, err := os.Create("output-test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Create a writer to write to the output file
	outputWriter := bufio.NewWriter(outputFile)

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(file)

	// Write the first line to the output file
	if scanner.Scan() {
		_, err := outputWriter.WriteString(scanner.Text() + "\n")
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
		_, err = outputWriter.WriteString(outputLine + "\n")
		if err != nil {
			panic(err)
		}
	}

	// Flush the writer to ensure all data is written to the file
	err = outputWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}

	fmt.Println("Ferdig!")
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

func AverageTemperature() {
	// Open the csv file
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the lines from the csv file
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Prompt the user for temperature unit
	fmt.Println("Velg temperaturenhet (celsius/fahr):")
	var unit string
	fmt.Scan(&unit)

	// Calculate the average temperature
	var sum float64
	count := 0
	for i, line := range lines {
		if i == 0 {
			continue // ignore header line
		}
		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			log.Fatalf("unexpected number of fields in line %d: %d", i, len(fields))
		}
		if fields[3] == "" {
			continue // ignore line with empty temperature field
		}
		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			log.Fatalf("could not parse temperature in line %d: %s", i, err)
		}

		if unit == "fahr" {
			// Convert back to Fahrenheit
			temperature = conv.CelsiusToFahrenheit(temperature)
		}
		sum += temperature
		count++
	}

	if unit == "fahr" {
		average := sum / float64(count)
		average = math.Round(average*100) / 100 // round to two decimal places
		fmt.Printf("Gjennomsnittlig temperatur: %.2f°F\n", average)
	} else {
		average := sum / float64(count)
		fmt.Printf("Gjennomsnittlig temperatur: %.2f°C\n", average)
	}
}
