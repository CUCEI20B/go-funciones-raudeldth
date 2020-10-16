package main

import "fmt"

func Burbuja(s []int64)  {
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s) - 1; j++ {
			if s[j] > s[j+1] {
				temp := s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
}

func main()  {
	a := []int64{8,4, 6, 3, 7, 10, 1, 2, 5, 9}
	Burbuja(a)
}
