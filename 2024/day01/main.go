package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
)

func main() {
    file, scanner, err := readInputFile("input1.txt")
    if err != nil {
        fmt.Errorf("Error reading input in main: %e", err)
        return
    }
    defer file.Close()

    left, right, numbs := leftAndRight(scanner)

    one := partOne(left, right)
    two := partTwo(left, numbs)

    fmt.Println(one)
    fmt.Println(two)
}

func partOne(left, right []int) int {
    sumDists := 0
    for i := range left {
        sumDists += Abs(right[i] - left[i])
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

func leftAndRight(scanner *bufio.Scanner) ([]int, []int, map[int]int){
    var left, right []int
    numbs := map[int]int{}
    for scanner.Scan() {
        var nl, nr int
        line := scanner.Text()
        fmt.Sscanf(line, "%d %d", &nl, &nr)
        left, right = append(left, nl), append(right, nr)
        numbs[nr]++
    }

    slices.Sort(left)
    slices.Sort(right)

    return left, right, numbs
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func readInputFile(fileName string) (*os.File, *bufio.Scanner, error) {
    file, err := os.Open(fileName)
    if err != nil {
        return nil, nil, fmt.Errorf("Error opening file: %e", err)
    }

    return file, bufio.NewScanner(file), nil
}
