package yr_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/bjornarsk/minyr/yr"
)

// Tester fra oppgavebeskrivelsen

// antall linjer i filen er 16756

func TestCountLines(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "kjevik-temp-celsius-20220318-20230318.csv", want: 16756},
	}

	for _, tc := range tests {
		got := yr.CountLines(tc.input)
		if got != tc.want {
			t.Errorf("%v: want %v, got %v,", tc.input, tc.want, got)
		}
	}
}

// gitt "Kjevik;SN39040;18.03.2022 01:50;6" ønsker å få (want) "Kjevik;SN39040;18.03.2022 01:50;42.8"
// gitt "Kjevik;SN39040;18.03.2022 01:50;0" ønsker å få (want) "Kjevik;SN39040;18.03.2022 01:50;32.0"
// gitt "Kjevik;SN39040;18.03.2022 01:50;-11" ønsker å få (want) "Kjevik;SN39040;18.03.2022 01:50;12.2"

// gitt "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;" ønsker å få (want)
// "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av STUDENTENS_NAVN"

func TestConvertLines(t *testing.T) {

	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
			want: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Bjørnar"},
	}

	for _, tc := range tests {
		got := yr.ProcessLine(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestAverageTemperature(t *testing.T) {

	type test struct {
		input   string
		wantAvg string
		wantErr error
	}

	tests := []test{
		{
			input:   "kjevik-temp-celsius-20220318-20230318.csv",
			wantAvg: "8.56",
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		// Open the test csv file
		file, err := os.Open(tc.input)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		// Read the lines from the csv file
		scanner := bufio.NewScanner(file)

		var lines []string

		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}

		// Call the function to get the actual average temperature
		actualAvg, err := yr.GetAverageTemperature(lines, "celsius")
		if err != tc.wantErr {
			t.Fatalf("expected error: %v, but got error: %v", tc.wantErr, err)
		}

		// Compare the actual and expected average temperature
		if actualAvg != tc.wantAvg {
			t.Errorf("expected average temperature %v, but got %v", tc.wantAvg, actualAvg)
		}
	}
}
