// quickSort project main.go
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	n := 5000000
	arr := make([]int, n)

	for i, _ := range arr {
		arr[i] = rand.Intn(100000) + 1
	}

	sn := time.Now()
	quickSort(arr, 0, n-1)
	fmt.Println(time.Since(sn))

}

func quickSort(arr []int, start, stop int) {
	if start >= stop {
		return
	}

	i := start
	j := stop

	flag := arr[i]

	for j != i {
		for j != i && arr[j] >= flag {
			j--
		}
		arr[i] = arr[j]
		for j != i && arr[i] < flag {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = flag

	if i-1 > start {
		quickSort(arr, start, i-1)
	}
	if stop > i+1 {
		quickSort(arr, i+1, stop)
	}
}
