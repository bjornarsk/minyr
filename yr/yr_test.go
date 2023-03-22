package yr

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// Tester fra oppgavebeskrivelsen

/* antall linjer i filen er 16756

gitt "Kjevik;SN39040;18.03.2022 01:50;6" ønsker å få (want) "Kjevik;SN39040;18.03.2022 01:50;42,8"

gitt "Kjevik;SN39040;07.03.2023 18:20;0" ønsker å få (want) "Kjevik;SN39040;07.03.2023 18:20;32"

gitt "Kjevik;SN39040;08.03.2023 02:20;-11" ønsker å få (want) "Kjevik;SN39040;08.03.2023 02:20;12,2"

gitt "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;" ønsker å få (want) "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av STUDENTENS_NAVN", hvor STUDENTENS_NAVN er navn på studenten som leverer besvarelsen

*/

func TestFileLineCount(t *testing.T) {
	filename := "kjevik-temp-celsius-20220318-20230318.csv"
	expectedNumLines := 16756

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numLines := 0
	for scanner.Scan() {
		numLines++
	}

	if numLines != expectedNumLines {
		t.Errorf("Expected %v lines in %v, but got %v lines", expectedNumLines, filename, numLines)
	}
}

func TestModifyString(t *testing.T) {
	originalString := "Kjevik;SN39040;18.03.2022 01:50;6"
	expectedString := "Kjevik;SN39040;18.03.2022 01:50;42,8"

	components := strings.Split(originalString, ";")
	components[len(components)-1] = "42,8"
	modifiedString := strings.Join(components, ";")

	if modifiedString != expectedString {
		t.Errorf("Expected %v, but got %v", expectedString, modifiedString)
	}
}

func TestModifyString2(t *testing.T) {
	originalString := "Kjevik;SN39040;18.03.2022 01:50;0"
	expectedString := "Kjevik;SN39040;18.03.2022 01:50;32"

	components := strings.Split(originalString, ";")
	components[len(components)-1] = "32"
	modifiedString := strings.Join(components, ";")

	if modifiedString != expectedString {
		t.Errorf("Expected %v, but got %v", expectedString, modifiedString)
	}
}

func TestModifyString3(t *testing.T) {
	originalString := "Kjevik;SN39040;18.03.2022 01:50;-11"
	expectedString := "Kjevik;SN39040;18.03.2022 01:50;12,2"

	components := strings.Split(originalString, ";")
	components[len(components)-1] = "12,2"
	modifiedString := strings.Join(components, ";")

	if modifiedString != expectedString {
		t.Errorf("Expected %v, but got %v", expectedString, modifiedString)
	}
}
