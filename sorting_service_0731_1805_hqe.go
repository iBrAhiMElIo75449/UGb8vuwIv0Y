// 代码生成时间: 2025-07-31 18:05:42
package main

import (
    "fmt"
    "sort"
)

// SortingService is a service to perform sorting operations.
type SortingService struct {}

// BubbleSort sorts a slice of integers using bubble sort algorithm.
func (s *SortingService) BubbleSort(numbers []int) ([]int, error) {
    if numbers == nil {
        return nil, fmt.Errorf("input slice is nil")
    }
    
    for i := 0; i < len(numbers); i++ {
        for j := 0; j < len(numbers)-i-1; j++ {
            if numbers[j] > numbers[j+1] {
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
            }
        }
    }
    return numbers, nil
}

// InsertionSort sorts a slice of integers using insertion sort algorithm.
func (s *SortingService) InsertionSort(numbers []int) ([]int, error) {
    if numbers == nil {
        return nil, fmt.Errorf("input slice is nil")
    }
    
    for i := 1; i < len(numbers); i++ {
        key := numbers[i]
        j := i - 1
        for j >= 0 && numbers[j] > key {
            numbers[j+1] = numbers[j]
            j = j - 1
        }
        numbers[j+1] = key
    }
    return numbers, nil
}

// QuickSort sorts a slice of integers using quick sort algorithm.
func (s *SortingService) QuickSort(numbers []int) ([]int, error) {
    if numbers == nil || len(numbers) < 2 {
        return numbers, nil // No need to sort single or nil element
    }
    return quickSort(numbers, 0, len(numbers)-1), nil
}

// quickSort is a helper function for the QuickSort algorithm.
func quickSort(numbers []int, low, high int) []int {
    if low < high {
        pi := partition(numbers, low, high)
        quickSort(numbers, low, pi-1) // Before pi
        quickSort(numbers, pi+1, high) // After pi
    }
    return numbers
}

// partition is a helper function for quickSort.
func partition(numbers []int, low, high int) int {
    i := low - 1
    pivot := numbers[high]
    for j := low; j < high; j++ {
        if numbers[j] < pivot {
            i++
            numbers[i], numbers[j] = numbers[j], numbers[i]
        }
    }
    numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
    return i + 1
}

func main() {
    service := SortingService{}
    numbers := []int{64, 34, 25, 12, 22, 11, 90}
    
    // Example usage of BubbleSort
    sortedNumbers, err := service.BubbleSort(numbers)
    if err != nil {
        fmt.Println("Error during BubbleSort: ", err)
    } else {
        fmt.Println("BubbleSort: ", sortedNumbers)
    }
    
    // Example usage of InsertionSort
    sortedNumbers, err = service.InsertionSort(numbers)
    if err != nil {
        fmt.Println("Error during InsertionSort: ", err)
    } else {
        fmt.Println("InsertionSort: ", sortedNumbers)
    }
    
    // Example usage of QuickSort
    sortedNumbers, err = service.QuickSort(numbers)
    if err != nil {
        fmt.Println("Error during QuickSort: ", err)
    } else {
        fmt.Println("QuickSort: ", sortedNumbers)
    }
}
