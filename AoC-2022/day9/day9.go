package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)



type Coord struct {
	x, y int
}

var (
	head_cur_coor = Coord{200,160}
	tail_cur_coor = Coord{200,160}

	maxheadx = 200 // i know this is stupid  // i should have extract this value dynamically
	minheadx = 200
	maxheady = 160
	minheady = 160

	coordinates_that_tail_visited [300][300] bool
)

func move_head(dir byte){
	switch(dir){
		case 'U':
			head_cur_coor.y += 1 
		case 'L':
			head_cur_coor.x -= 1 
		case 'R':
			head_cur_coor.x += 1 
		case 'D':
			head_cur_coor.y -= 1 
		default:
			fmt.Printf("Bad dir:=%c", dir)
			panic(1)
	}
}

func abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}

func is_tail_touched_head(tcoor, hcoor Coord) bool{
	if abs(tcoor.x - hcoor.x) > 1{
		return false
	}
	if abs(tcoor.y - hcoor.y) > 1{
		return false
	}
	return true
}



func follow_tail(){
	//if is_tail_touched_head() {
	//	return //wihtout moving
	//}
	ver_dis := head_cur_coor.y - tail_cur_coor.y
	hor_dis := head_cur_coor.x - tail_cur_coor.x

	if (abs(ver_dis) < 2) && (abs(hor_dis) < 2) {
		return
		//no movement required
	} else if (abs(ver_dis) == 2) && (hor_dis == 0) {
		//only move vertically
		if ver_dis > 0 {tail_cur_coor.y += 1}
		if ver_dis < 0 {tail_cur_coor.y -= 1}
	} else if (abs(hor_dis) == 2) && (ver_dis == 0) {
		//only move horizontally
		if hor_dis > 0 {tail_cur_coor.x += 1}
		if hor_dis < 0 {tail_cur_coor.x -= 1}
	} else {
		//now the diagonale
		tmp_coor0 := Coord{tail_cur_coor.x-1, tail_cur_coor.y-1}
		tmp_coor1 := Coord{tail_cur_coor.x-1, tail_cur_coor.y+1}
		tmp_coor2 := Coord{tail_cur_coor.x+1, tail_cur_coor.y-1}
		tmp_coor3 := Coord{tail_cur_coor.x+1, tail_cur_coor.y+1}

		if is_tail_touched_head(head_cur_coor, tmp_coor0) {tail_cur_coor = tmp_coor0; return}
		if is_tail_touched_head(head_cur_coor, tmp_coor1) {tail_cur_coor = tmp_coor1; return}
		if is_tail_touched_head(head_cur_coor, tmp_coor2) {tail_cur_coor = tmp_coor2; return}
		if is_tail_touched_head(head_cur_coor, tmp_coor3) {tail_cur_coor = tmp_coor3; return}

		// if left here panic
		fmt.Printf("Couldn't find better diagonale move for tail\n")
		fmt.Printf("Positions: head %v\ttail %v",
			head_cur_coor, tail_cur_coor,
		)
		panic(1)
	} 
	
}

func log_maxmin_of_head(){
	if head_cur_coor.x > maxheadx {maxheadx = head_cur_coor.x}
	if head_cur_coor.x < minheadx {minheadx = head_cur_coor.x}
	if head_cur_coor.y > maxheady {maxheady = head_cur_coor.y}
	if head_cur_coor.y < minheady {minheady = head_cur_coor.y}
}


func move(dir byte, steps int){
	for i:=0; i<steps; i++{
		move_head(dir) // move the head
		follow_tail()
		coordinates_that_tail_visited[tail_cur_coor.y][tail_cur_coor.x] = true
		log_maxmin_of_head()
	}
}


func main(){
	fmt.Printf("AoC2022\n");
	fname := "/home/garid/Documents/advent/AoC-2022/day9/input.txt"
	file, err:= os.Open(fname)
	if err != nil{
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	var dir byte
	var steps int

	reader := bufio.NewReader(file)
	for i:=0;;i++{
		line, ret := reader.ReadString('\n')
		if ret == io.EOF{
			fmt.Printf("File has ended\n")
			break
		}
		fmt.Sscanf(line, "%c %d\n", &dir, &steps)
		fmt.Printf("%d\t%c-%d| %v %v |->", i, dir, steps, head_cur_coor, tail_cur_coor)
		move(dir, steps)
		fmt.Printf("\t|%v %v\n",  head_cur_coor, tail_cur_coor)
	}
	//finished instructions

	fmt.Printf("x:%d to %d\n", minheadx, maxheadx)
	fmt.Printf("y:%d to %d\n", minheady, maxheady)

	fmt.Printf("current head %v\n", head_cur_coor)
	fmt.Printf("current tail %v\n", tail_cur_coor)
	//count tail visited
	tailvisit:=0
	for i:=0; i<300; i++{
		for j:=0; j<300; j++{
			if coordinates_that_tail_visited[i][j] {
				tailvisit++;
			}
		}
	}

	fmt.Printf("tail has visited %v positions\n", tailvisit) //this is part1

}
