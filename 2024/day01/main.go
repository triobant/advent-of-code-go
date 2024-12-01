package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
)

func main() {
    file, err := os.Open("input1.txt")
    if err != nil {
        fmt.Errorf("Error opening file: %e", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    if err != nil {
        fmt.Errorf("Error opening file: %e", err)
    }

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

    sumDists, score := 0, 0
    for i := range left {
        sumDists += abs(right[i] - left[i])
        score += left[i] * numbs[left[i]]
    }

    fmt.Println(sumDists)
    fmt.Println(score)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
