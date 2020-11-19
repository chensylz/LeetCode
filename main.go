package main

import (
	"GoLeetCode/SoredFunction"
	"GoLeetCode/Tools"
	"fmt"
)

func main() {
	fmt.Println(SoredFunction.QuickSort(Tools.GenerateRandArray(10)))
	fmt.Println(SoredFunction.MergeSort(Tools.GenerateRandArray(10)))
	fmt.Println(SoredFunction.BubbleSort(Tools.GenerateRandArray(10)))
	fmt.Println(SoredFunction.SelectionSort(Tools.GenerateRandArray(10)))
}
