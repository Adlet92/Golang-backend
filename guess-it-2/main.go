package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var k []int
	counter := 1
	for reader.Scan() {
		t := reader.Text()
		n, err := strconv.Atoi(t)
		if err != nil {
			fmt.Print(err)
			return
		}
		k = append(k, n)
		counter++
		if len(k) < 3 {
			fmt.Println(100, 200)
		} else {
			b, a := Count(k)
			res := b*float64(counter) + a
			fmt.Println(res-45, res+45)
		}
	}
}

func Count(k []int) (float64, float64) {
	var (
		SumY  int
		SumX  int
		SumX2 int
		SumXY int
	)

	for _, i := range k {
		SumY += i
	}
	for h, _ := range k {
		SumX += h
		SumX2 += h * h
	}
	for b := 0; b < len(k); b++ {
		SumXY += b * k[b]
	}
	b := float64(len(k)*SumXY-SumX*SumY) / float64(len(k)*SumX2-SumX*SumX)
	a := float64(SumY*SumX2-SumX*SumXY)/float64(len(k)*SumX2-SumX*SumX) + b
	return b, a
}
