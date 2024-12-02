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
            name: "Sample Input Part 1",
            input: "input_base.txt",
            expected: 11,
        },
    }

    for i, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            file, scanner, err := readInputFile(tt.input)
            if err != nil {
                t.Fatalf(" Test %v - %s FAIL: error reading file: %v", i, tt.name, err)
            }
            defer file.Close()

            left, right, _ := leftAndRight(scanner)
            actual := partOne(left, right)

            if actual != tt.expected {
                t.Errorf("Test %v - %s FAIL: expected dist: %v, actual: %v", i, tt.name, tt.expected, actual)
            }
        })
    }
}

func TestPart2(t *testing.T) {
    tests := []struct {
        name        string
        input       string
        expected    int
    }{
        {
            name: "Sample Input Part 2",
            input: "input_base.txt",
            expected: 31,
        },
    }

    for i, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            file, scanner, err := readInputFile(tt.input)
            if err != nil {
                t.Fatalf(" Test %v - %s FAIL: error reading file: %v", i, tt.name, err)
            }
            defer file.Close()

            left, _, numbs := leftAndRight(scanner)
            actual := partTwo(left, numbs)

            if actual != tt.expected {
                t.Errorf("Test %v - %s FAIL: expected dist: %v, actual: %v", i, tt.name, tt.expected, actual)
            }
        })
    }
}
