package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
)

func checkall4diff(str string) bool{
	//check 1st
	if (str[0] == str[1]) || (str[0] == str[2]) || (str[0] == str[3]){
		return false
	}

	//check 2nd
	if  (str[1] == str[2]) || (str[1] == str[3]){
		return false
	}

	//check 3nd
	if  str[2] == str[3] {
		return false
	}
	return true
}

func check14diff(str string) bool{
	for i:=0;i<14;i++{
		//i is the checking index
		for j:=i+1;j<14;j++{
			if str[i] == str[j]{
				return false
			}
		}
	}
	return true
}


func checkNdiff(str string, N int) bool{
	for i:=0;i<N;i++{
		//i is the checking index
		for j:=i+1;j<N;j++{
			if str[i] == str[j]{
				return false
			}
		}
	}
	return true
}

func subroutine(str string) int{
	//var x = 3
	for i,L:=0, len(str); i<L;i++{
		if checkall4diff(str[i:i+4]){
			return i + 4
		}
	}
	fmt.Printf("no 4 different letters")
	panic(2)
}


func subroutine_part2(str string) int{
	//var x = 3
	for i,L:=0, len(str); i<L;i++{
		if check14diff(str[i:i+14]){
			return i + 14
		}
	}
	fmt.Printf("no 14 different letters")
	panic(2)
}

func main(){
	fmt.Println("AoC2022-day6");
	//fname := "/home/garid/Documents/advent/AoC-2022/day6/input.txt"
	fname := "/home/garid/Documents/advent/AoC-2022/day6/test3"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Error %v, when file opening %s\n", err, fname)
		fmt.Printf("Please check file path")
		panic(err)
	}
	reader := bufio.NewReader(file)
	for i:=0 ;; i++{
		line, ret := reader.ReadString('\n')
		if ret == io.EOF{
			fmt.Println("This is the end of file: ")
			return 
		}
		fmt.Printf("%d %s\n", i, line)
		fmt.Printf("4diff: %d\n", subroutine(line))
		fmt.Printf("14diff %d\n", subroutine_part2(line))
	}
}
