package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	//"sort"
)

type Grid struct {
	grid [][]byte;
	xlen int;
	ylen int;
	move_count [][]int;
}

var themap Grid;
var startx, starty int;
var endx, endy int;


func populate_grid(i int, line string){
	themap.xlen = len(line)
	if (strings.Contains(line, "S")){
		starty = themap.ylen
		startx = strings.Index(line, "S")
		line = strings.Replace(line, "S", "a", 1)
	}
	if (strings.Contains(line, "E")){
		endy = themap.ylen
		endx = strings.Index(line, "E")
		line = strings.Replace(line, "E", "z", 1)
	}

	themap.grid = append(themap.grid, []byte(line))
	themap.ylen++

	tmp_row_move_count := make([]int, themap.xlen)
	for i := range tmp_row_move_count {
		tmp_row_move_count[i] = 99999
	}
	themap.move_count = append(themap.move_count, tmp_row_move_count)
}

func fancy_print_map(){
	fmt.Printf("ylen:%v\nxlen:%v\n", themap.ylen, themap.xlen)
	for i:=0;i<themap.ylen;i++{
		fmt.Printf("%s\n", string(themap.grid[i]))
	}
}

func fancy_print_count(){
	fmt.Printf("ylen:%v\nxlen:%v\n", themap.ylen, themap.xlen)
	for i:=0;i<themap.ylen;i++{
		fmt.Printf("%v %T\n", themap.move_count[i], themap.move_count[i])
	}
}

func map_explorer(y int, x int, count int){
	// explores maps in 4 direction via the recursion
	count++;
	if themap.move_count[y][x] < count{
		//i.e. somewhere I already came here with less ammount
		return
	}
	themap.move_count[y][x] = count
	currentheight := themap.grid[y][x]

	// try up:    y-1
	if (y != 0) { // check the limit condition
        //.Printf("v")
		dest_height := themap.grid[y-1][x]
		if dest_height - currentheight <= 1 {
			map_explorer(y - 1, x, count)
		}
	}

	// try down:  y+1
	if (y != (themap.ylen-1)) { // check the limit condition
        //.Printf("^")
		dest_height := themap.grid[y+1][x]
		if dest_height - currentheight <= 1 {
			map_explorer(y + 1, x, count)
		}
	}

	// try left:    x-1
	if (x != 0) { // check the limit condition
        //.Printf("<")
		dest_height := themap.grid[y][x-1]
		if dest_height - currentheight <= 1 {
			map_explorer(y, x-1, count)
		}
	}

	// try down:  x+1
	if (x != (themap.xlen-1)) { // check the limit condition
        //.Printf(">")
		dest_height := themap.grid[y][x+1]
		if dest_height - currentheight <= 1 {
			map_explorer(y, x+1, count)
		}
	}
}

func main() {
	fmt.Printf("AoC2022 day 12\n")

	fname := "/home/garid/Documents/advent/AoC-2022/day12/test"
	//fname := "/home/garid/Documents/advent/AoC-2022/day12/input.txt"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		line, ret := reader.ReadString('\n')
		if ret == io.EOF {
			fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			break
		}
		line = line[:len(line)-1]
		populate_grid(i, line)
		fmt.Printf("%d\t%s %d\n", i, line, len(line))
	}


	fmt.Printf("Finished\n") // this is part1
	//fmt.Printf("%v\n", themap)
	fancy_print_map()
	fancy_print_count()
	fmt.Printf("startx: %d\n", startx)
	fmt.Printf("starty: %d\n", starty)
	fmt.Printf("endx: %d\n", endx)
	fmt.Printf("endy: %d\n", endy)

	fmt.Printf("Recursive exploration\n") // this is part1


	map_explorer(starty, startx, -1)

	fancy_print_map()
	fancy_print_count()

	fmt.Printf("Steps: %v\n", themap.move_count[endy][endx])
}
