package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Position struct {
	y, x int           
	direction int       // 0 up 1 right 2 down 3 left            R = ++, L --
}


var (
	// lines = make([]string, 0, 300)
	grid [300][180] byte
	cmds = make([]string, 0, 5000)
	pos  = Position{}
)

func advance_1(){
	dir := pos.direction % 4
	if dir == 0 {                              //up
		if (grid[pos.y-1][pos.x] == '.') {pos.y = pos.y - 1}
	} else if dir == 1 || dir == -3 {          //righ
		if (grid[pos.y][pos.x+1] == '.') {pos.x = pos.x + 1}
	} else if dir == 2 || dir == -2 {          // down
		if (grid[pos.y+1][pos.x] == '.') {pos.y = pos.y + 1}
	} else if dir == 3 || dir == -1 {          //left
		if (grid[pos.y][pos.x-1] == '.') {pos.x = pos.x - 1}
	}
}

func initialize_grid(){
	for i:=0; i<300; i++ {
		for j:=0; j<180; j++ {
			grid[i][j] = '#'
		}
	}
}

func set_row(row int, line string) {
	for j, chr := range line {
		if chr == '\n' {
			panic(1)
		}
		if line[j] == ' '{ 
			grid[row][j+1] = '#'
		} else {
			grid[row][j+1] = line[j]
		}
	}
}

func print_grid(row_from, row_end int) {
	for i:=row_from; i<=row_end; i++ {
		for j:=0; j<180; j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}

func parse_command(line string) {
	var tmpbufffer string;
	lastisnum := false
	for i,L:=0,len(line); i<L; i++ {
		if line[i] == 'R' || line[i] == 'L' {
			cmds = append(cmds, tmpbufffer) 
			cmds = append(cmds, string(line[i])) 
			lastisnum = false
			tmpbufffer = ""
		} else {
			tmpbufffer += string(line[i])
			lastisnum = true
		}
	}
	if lastisnum {
		cmds = append(cmds, tmpbufffer) 
	}
}

func main() {
	initialize_grid()
	fmt.Printf("AoC2022 day22\n")

	fname := "/home/garid/Documents/advent/AoC-2022/day22/input.txt" ;
	// fname := "/home/garid/Documents/advent/AoC-2022/day22/test"

	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}
	islastline := false
	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		line, ret1 := reader.ReadString('\n')
		if ret1 == io.EOF {
			fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			break
		}
		// fmt.Printf("%v", line)
		if line == "\n" {
			fmt.Printf("=========================================\n")
			islastline = true
		}
		if !islastline {
			set_row(i+1, line[:len(line)-1])
		} else {
			parse_command(line[:len(line)-1])
		}
		// fmt.Sscanf(line, "%d\n", &val)
	}
	// print_grid(0, 299)
	fmt.Printf("%v\n", cmds)

}
