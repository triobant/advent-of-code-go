package main

import (
    "testing"
)

func TestPart1(t *testing.T) {
    tests := []struct {
        name        string
        input       string
        expected    int
    }{
        {
            name: "Day02 Sample Input Part 1",
            input: "input_base.txt",
            expected: 2,
        },
    }

    for i, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            file, scanner, err := readInputFile(tt.input)
            if err != nil {
                t.Fatalf(" Test %v - %s FAIL: error reading file: %v", i, tt.name, err)
            }
            defer file.Close()

            reports := getReports(scanner)
            actual := partOne(reports)

            if actual != tt.expected {
                t.Errorf("Test %v - %s FAIL: expected safe reports: %v, actual: %v", i, tt.name, tt.expected, actual)
            }
        })
    }
}
