package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	// "time"
)

// var (
// 	powers = make([]int, 0, 100)
// )
func snafu2dec(line string) int  {
	val := 0
	for i,L := 0, len(line); i<L; i++{
		val *= 5;
		switch(line[i]) {
		case '=':
			val-=2
		case '-':
			val-=1
		case '0':
			val+=0
		case '1':
			val+=1
		case '2':
			val+=2
		default:
			fmt.Printf("panic at: %s's %d-th char %c\n", line, i, line[i])
			panic(1)
		}
	}
	return val;
}


func dec2fivebased(val int, s string) string {
	// we want to know val = d * 5 + remainder
	d := val / 5;
	remainder := val - 5 * d;

	if d == 0{
		return strconv.Itoa(remainder)
	}
	return dec2fivebased(d, s) + strconv.Itoa(remainder)
}


func dec2snafu(val int, s string) string {
	// we want to know val = d * 5 + remainder
	d := val / 5;
	remainder := val - 5 * d;
	// fmt.Printf("\t\t%d %s %d %d\n", val, s, d, remainder)
	if d == 0{
		// return strconv.Itoa(remainder)
		if remainder == 0 || remainder == 1 || remainder == 2 {
			return strconv.Itoa(remainder)
		} else if remainder == 3 {
			d += 1
			return dec2snafu(d, s) + "=" 
		} else if remainder == 4 {
			d += 1
			return dec2snafu(d, s) + "-" 
		} else {
			panic(2)
		}
	}

	if remainder == 0 || remainder == 1 || remainder == 2 {
		return dec2snafu(d, s) + strconv.Itoa(remainder)
	} else if remainder == 3 {
		d += 1
		return dec2snafu(d, s) + "=" 
	} else if remainder == 4 {
		d += 1
		return dec2snafu(d, s) + "-" 
	} else {
		panic(2)
	}
}



func main() {
	fmt.Printf("AoC2022 day25\n")
	fname := "/home/garid/Documents/advent/AoC-2022/day25/input.txt" ;
	// fname := "/home/garid/Documents/advent/AoC-2022/day25/test"
	// fname := "/home/garid/Documents/advent/AoC-2022/day25/test_small"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}
	reader := bufio.NewReader(file)
	sum := 0
	for i := 0; ; i++ {
		line, ret1 := reader.ReadString('\n')
		if ret1 == io.EOF {
			fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			break
		}
		line = line[:len(line)-1]
		val := snafu2dec(line)
		fmt.Printf("%d\t%v\t%v\n", i, line , val)
		sum+=val
	}
	

	fmt.Printf("sum: %d\t%v\t%v\n", sum, dec2fivebased(sum, ""), dec2snafu(sum, "") )
}
