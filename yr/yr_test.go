package yr

import (
	"fmt"
	"os"
	"testing"
)

// Tester fra oppgavebeskrivelsen

/* antall linjer i filen er 16756

gitt "Kjevik;SN39040;18.03.2022 01:50;6" ønsker å få (want) "Kjevik;SN39040;18.03.2022 01:50;42.8"

gitt "Kjevik;SN39040;07.03.2023 18:20;0" ønsker å få (want) "Kjevik;SN39040;07.03.2023 18:20;32"

gitt "Kjevik;SN39040;08.03.2023 02:20;-11" ønsker å få (want) "Kjevik;SN39040;08.03.2023 02:20;12.2"

gitt "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;" ønsker å få (want) "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av STUDENTENS_NAVN", hvor STUDENTENS_NAVN er navn på studenten som leverer besvarelsen

*/

func TestCountLines(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "kjevik-temp-celsius-20220318-20230318.csv", want: 16756},
	}

	for _, tc := range tests {
		got := countLines(tc.input)
		if got != tc.want {
			t.Errorf("%v: want %v, got %v,", tc.input, tc.want, got)
		}
	}
}

func TestProcessLine(t *testing.T) {
	// Create a temporary input file with a single line
	tmpfile, err := os.CreateTemp("", "test_input")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer tmpfile.Close()

	// Write a line to the input file
	fmt.Fprintf(tmpfile, "Kjevik;SN39040;18.03.2022 01:50;6\n")

	// Call processLine() on the input line
	got := processLine("Kjevik;SN39040;18.03.2022 01:50;6")

	// Check that the output is as expected
	want := "Kjevik;SN39040;18.03.2022 01:50;42.8"
	if got != want {
		t.Errorf("processLine() = %q, want %q", got, want)
		fmt.Println("Actual output: ", got)
	}
}

func TestProcessLine2(t *testing.T) {
	// Create a temporary input file with a single line
	tmpfile, err := os.CreateTemp("", "test_input")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer tmpfile.Close()

	// Write a line to the input file
	fmt.Fprintf(tmpfile, "Kjevik;SN39040;07.03.2023 18:20;0\n")

	// Call processLine() on the input line
	got := processLine("Kjevik;SN39040;07.03.2023 18:20;0")

	// Check that the output is as expected
	want := "Kjevik;SN39040;07.03.2023 18:20;32.0"
	if got != want {
		t.Errorf("processLine() = %q, want %q", got, want)
		fmt.Println("Actual output: ", got)
	}
}

func TestProcessLine3(t *testing.T) {
	// Create a temporary input file with a single line
	tmpfile, err := os.CreateTemp("", "test_input")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer tmpfile.Close()

	// Write a line to the input file
	fmt.Fprintf(tmpfile, "Kjevik;SN39040;08.03.2023 02:20;-11\n")

	// Call processLine() on the input line
	got := processLine("Kjevik;SN39040;08.03.2023 02:20;-11")

	// Check that the output is as expected
	want := "Kjevik;SN39040;08.03.2023 02:20;12.2"
	if got != want {
		t.Errorf("processLine() = %q, want %q", got, want)
		fmt.Println("Actual output: ", got)
	}
}

func TestProcessLineLast(t *testing.T) {
	// Create a temporary input file with a single line
	tmpfile, err := os.CreateTemp("", "test_input")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer tmpfile.Close()

	// Write a line to the input file
	fmt.Fprintf(tmpfile, "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;\n")

	// Call processLine() on the input line
	got := processLine("Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;")

	// Check that the output is as expected
	want := "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Bjørnar"
	if got != want {
		t.Errorf("processLine() = %q, want %q", got, want)
		fmt.Println("Actual output: ", got)
	}
}

// test 2
