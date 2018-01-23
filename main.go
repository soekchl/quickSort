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

	n := 500000
	arr := make([]int32, n)
	arr1 := make([]int32, n)

	for i, _ := range arr {
		arr[i] = rand.Int31n(100000) + 1
	}
	copy(arr1, arr)
	//	fmt.Println(arr)

	sn := time.Now()
	quickSort(arr, 0, n-1)
	fmt.Println(time.Since(sn))

	ch := make(chan bool, 1)
	sn = time.Now()
	go quickSort1(arr1, 0, n-1, ch)
	<-ch
	fmt.Println(time.Since(sn))

	for i, v := range arr {
		if v != arr1[i] {
			fmt.Println("排序错误！")
			return
		}
	}
}

// 改成多协程
func quickSort1(arr []int32, start, stop int, ch chan bool) {
	if start >= stop {
		ch <- true
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

	c := make(chan bool, 3)
	n := 0

	if i-1 > start {
		n++
		go quickSort1(arr, start, i-1, c)
	}

	if stop > i+1 {
		n++
		go quickSort1(arr, i+1, stop, c)
	}

	for i = 0; i < n; i++ {
		<-c
	}
	ch <- true
}

func quickSort(arr []int32, start, stop int) {
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
