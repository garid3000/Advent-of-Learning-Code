package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	X = 1      // X means apperantly x-axis, sprite pos
	cycle = 0
	sum = 0
	tmprow = [40]byte{'y'}
	display = [6][40]byte {{'x'}}
	row_count = 0
)

func abs(v int) int{
	if v < 0 {
		return -v
	}
	return v
}

func init_arrays(){ //not needed, just for the debuggin purposes
	for i:=0;i<6;i++{
		for j:=0;j<40;j++{
			display[i][j] = 'x'
		}
	}
}

func display_message(){
	for i:=0;i<6;i++{
		for j:=0;j<40;j++{
			fmt.Printf("%c", display[i][j])
		}
		fmt.Printf("\n")
	}
}

func inc_cycle(){
	pos := cycle % 40
	if abs(pos - X) <= 1 {
		tmprow[pos] = '#'
	} else {
		tmprow[pos] = ' '
	}

	cycle++
	fmt.Printf("*")
	if (cycle % 40) == 20{
		fmt.Printf("----Cycle: %d, X:=%d-----", cycle, X)
		sum += cycle * X
	} else if (cycle % 40) == 0 {
		for j:=0; j<40; j++{
			display[cycle/40 - 1][j] = tmprow[j]
		}
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

	init_arrays()
	display_message()
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
	display_message()
}
