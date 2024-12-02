package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
    "log"
)

type InputData struct {
    left    []int
    right   []int
    numbs   map[int]int
}

func main() {
    file, scanner, err := readInputFile("input1.txt")
    if err != nil {
        log.Fatal("Error reading input in main:", err)
    }
    defer file.Close()

    data := parseInputData(scanner)

    one := partOne(data.left, data.right)
    two := partTwo(data.left, data.numbs)

    fmt.Println(one)
    fmt.Println(two)
}

func partOne(left, right []int) int {
    sumDists := 0
    for i := range left {
        sumDists += abs(right[i] - left[i])
    }
    return sumDists
}

func partTwo(left []int, numbs map[int]int) int {
    score := 0
    for i := range left {
        score += left[i] * numbs[left[i]]
    }
    return score
}

func parseInputData(scanner *bufio.Scanner) InputData {
    data := InputData{
        numbs: make(map[int]int),
    }

    for scanner.Scan() {
        var nl, nr int
        line := scanner.Text()
        fmt.Sscanf(line, "%d %d", &nl, &nr)
        data.left = append(data.left, nl)
        data.right = append(data.right, nr)
        data.numbs[nr]++
    }

    slices.Sort(data.left)
    slices.Sort(data.right)

    return data
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func readInputFile(fileName string) (*os.File, *bufio.Scanner, error) {
    file, err := os.Open(fileName)
    if err != nil {
        return nil, nil, fmt.Errorf("opening file: %w", err)
    }

    return file, bufio.NewScanner(file), nil
}
