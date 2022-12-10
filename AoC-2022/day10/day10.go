package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	X = 1
	cycle = 0
	sum = 0
)

func inc_cycle(){
	cycle++
	fmt.Printf("*")
	if (cycle % 40) == 20{
		fmt.Printf("----Cycle: %d, X:=%d-----", cycle, X)
		sum += cycle * X
	}
	
}

func noop(){
	//takes one cycle on effect
	inc_cycle()
}
func addx(val int){
	inc_cycle()
	inc_cycle()
	X += val
}

func readline(line string){
	if len(line) < 4{
		fmt.Printf("string is less than 4, len=%d, %s", len(line), line)
		panic(5)
	}
	if line[:4] == "noop"{
		noop()
	} else if line[:4] == "addx"{
		var v int
		fmt.Sscanf(line, "addx %d", &v)
		addx(v)
	} else {
		fmt.Printf("key word is not noop/addx, len=%d, %s", len(line), line)
		panic(6)
	}


}

func main(){
	//initial coordinates of chain
	fmt.Printf("AoC2022 day 10\n");
	fname := "/home/garid/Documents/advent/AoC-2022/day10/input.txt"
	//fname := "/home/garid/Documents/advent/AoC-2022/day10/test1"
	file, err:= os.Open(fname)
	if err != nil{
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}
	reader := bufio.NewReader(file)
	for i:=0;;i++{
		line, ret := reader.ReadString('\n')
		if ret == io.EOF{
			fmt.Printf("File has ended. Total %d lines.\n", i)
			break
		}
		line = line[:len(line)-1]
		fmt.Printf("%d\t%s", i, line)
		readline(line)
		fmt.Println()
	}
	//finished instructions
	fmt.Printf("finished: sum=%d\n", sum) // this is part1
}
