package main 
import "fmt"

var arr = [6] int {2, 5, 1, 6, 3, 9}

func main() {
	var temp, y int
	pass := 0
	for pass < len(arr)-1 {
		imax := pass
		ctr := imax + 1
		for ctr < len(arr) {
			if arr[ctr]<arr[imax] {
				imax = ctr
			}
			ctr++
		}
		temp = arr[imax]
		arr[imax] = arr[pass]
		arr[pass] = temp
		pass++
	}
	for i:= 0; i<len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Scanln(&y)
	s(y)
}

func s(y int) {
	var tengah int
	found := false
	kiri := 0
	kanan := len(arr)
	for kiri<=kanan && !found {
		tengah = (kiri+kanan)/2
		if y == arr[tengah] {
			found = true
		} else if y > arr[tengah] {
			kiri = tengah+1
		} else if y < arr[tengah] {
			kanan = tengah
		}
	}
	
	if found {
		fmt.Print(arr[tengah], "*")
	}
}