package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
	"log"
	//"encoding/csv" // test
//"sort"
)

type Grid struct {
	grid [][]byte; //byte is uint8
	xlen int;
	ylen int;
	move_count [][]int;
}

type yx struct {
	y, x int
}

var themap Grid;
var cellwith_a []yx;
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
	for j,L:=0,len(line);j<L;j++{
		if line[j] == 'a'{
			tmp := yx{}
			tmp.y = i
			tmp.x = j
			cellwith_a = append(cellwith_a, tmp)
		}
	}

	themap.grid = append(themap.grid, []byte(line))
	themap.ylen++

	tmp_row_move_count := make([]int, themap.xlen)
	for i := range tmp_row_move_count {
		tmp_row_move_count[i] = -1
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
	for i:=0;i<themap.ylen;i++{
		fmt.Printf("%v %T\n", themap.move_count[i], themap.move_count[i])
	}
}

func map_explorer_mark_dest(y int, x int, count int){
	currentheight := int(themap.grid[y][x])

	// try up:    y-1
	if (y != 0) { // check the limit condition
		dest_height := int(themap.grid[y-1][x])
		if dest_height - currentheight <= 1 {
			if themap.move_count[y-1][x] < 0 {
				themap.move_count[y-1][x] = count + 1
			}
		}
	}

	// try down:  y+1
	if (y != (themap.ylen-1)) { // check the limit condition
		dest_height := int(themap.grid[y+1][x])
		if dest_height - currentheight <= 1 {
			if themap.move_count[y+1][x] < 0{
				themap.move_count[y+1][x] = count + 1
			}
		}
	}

	// try left:    x-1
	if (x != 0) { // check the limit condition
		dest_height := int(themap.grid[y][x-1])
		if dest_height - currentheight <= 1 {
			if themap.move_count[y][x-1] < 0{
				themap.move_count[y][x-1] = count + 1
			}
		}
	}

	// try down:  x+1
	if (x != (themap.xlen-1)) { // check the limit condition
		dest_height := int(themap.grid[y][x+1])
		if dest_height - currentheight <= 1 {
			if themap.move_count[y][x+1] < 0 {
				themap.move_count[y][x+1] = count + 1
			}
		}
	}
}

func explore_marked_n_cells(val int){
	for y:=0; y<themap.ylen; y++{
		for x:=0; x<themap.xlen; x++{
			if (themap.move_count[y][x] == val){
				map_explorer_mark_dest(y,x, val)
			}
		}
	}
}


func writecsv(fname string){ // this was for the debugging
	s := ""
	for y:=0; y<themap.ylen;  y++{
		for x:=0; x<themap.xlen;  x++{
			s += strconv.Itoa(themap.move_count[y][x]) + "\t"
		}
		s += "\n"
	}

	outpath := fname + ".csv"
	err := os.WriteFile(outpath, []byte(s), 0644)
    if err != nil {
        log.Fatal(err)
    }
}


func main() {
	fmt.Printf("AoC2022 day 12\n")

	//fname := "/home/garid/Documents/advent/AoC-2022/day12/test"
	fname := "/home/garid/Documents/advent/AoC-2022/day12/input.txt"
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
	//fancy_print_map()
	//fancy_print_count()
	fmt.Printf("startx: %d\n", startx)
	fmt.Printf("starty: %d\n", starty)
	fmt.Printf("endx: %d\n", endx)
	fmt.Printf("endy: %d\n", endy)

	themap.move_count[starty][startx] = 0
	//fancy_print_count()
	for i:=0;i<1000;i++{
		explore_marked_n_cells(i)
		if themap.move_count[endy][endx] > 0 {
			break
		}
	}
	writecsv("last.csv")
	fancy_print_count()
	fmt.Printf("Steps: %v\n", themap.move_count[endy][endx])
	thestepsneeded := themap.move_count[endy][endx]
	var best_posx, best_posy int
	//max_a_that_lead_E = themap.move_count[endy][endx]
	//recursive_until_a(endy, endx)
	//
	//fmt.Printf("=========")
	//fmt.Printf("val: %v\n", max_a_that_lead_E)
	//fmt.Printf("y    %v\n", max_a_that_lead_E_y)
	//fmt.Printf("x    %v\n", max_a_that_lead_E_x)
	fmt.Printf("%v\t%v\n", cellwith_a, len(cellwith_a))
	for i,L:=0,len(cellwith_a); i<L; i++{
		//fmt.Printf("--%d--%v-----------------------------", i, cellwith_a[i])
		for ii:=0; ii<themap.ylen; ii++{
			for jj:=0; jj<themap.xlen; jj++{
				themap.move_count[ii][jj] = -1
			}
		}
		//fmt.Printf("cleared\n")
		
		themap.move_count[cellwith_a[i].y][cellwith_a[i].x] = 0
		//fancy_print_count()
		for i:=0;i<1000;i++{
			explore_marked_n_cells(i)
			if themap.move_count[endy][endx] > 0 {
				break
			}
		}

		//fmt.Printf("--%v\n", themap.move_count[endy][endx])

		if themap.move_count[endy][endx] == -1 {
			//unreachable from this 
		} else if thestepsneeded > themap.move_count[endy][endx] {
			thestepsneeded = themap.move_count[endy][endx]
			best_posy = cellwith_a[i].y
			best_posx = cellwith_a[i].x
		}

		//fancy_print_count()
	}

	fmt.Printf("v %v\n", thestepsneeded)
	fmt.Printf("y %v\n", best_posy)
	fmt.Printf("x %v\n", best_posx)
}
