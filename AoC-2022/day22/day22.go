package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Position struct {
	y, x int           
	direction int       // 0 up 1 right 2 down 3 left            R = ++, L --
	//apperently I didn't read the question fully there
	// there is actual numbering system, but I will use this as it is
}


var (
	// lines = make([]string, 0, 300)
	grid [300][180] byte
	cmds = make([]string, 0, 5000)
	pos  = Position{}
)

func part2_advance_1(){
	thisFace := face_identifier()
	dir := pos.direction % 4
	newPos := Position{}
	switch (dir) {
	case 0:
		if (grid[pos.y-1][pos.x] == '.') {
			pos.y = pos.y - 1
		} else if  grid[pos.y-1][pos.x] == ',' {
			switch (thisFace) {
			case 1:
				newPos = Position{y: pos.x+100, x: 1,         direction: 1}
			case 2:
				newPos = Position{y: 200,       x: pos.x-100, direction: 0}
			case 5:
				newPos = Position{y: pos.x+50,  x:51,         direction: 1}
			default:
				fmt.Printf("\npanic dir:^ , face:%d\n", thisFace)
				panic(13)
			}
			if grid[newPos.y][newPos.x] == '.' {
				pos = newPos
			}
		}
	case 1, -3:
		if (grid[pos.y][pos.x+1] == '.') {
			pos.x = pos.x + 1
		} else if  grid[pos.y][pos.x+1] == ',' {
			switch (thisFace) {
			case 2:
				newPos = Position{y:151-pos.y,    x:pos.x-50,    direction: 3}
			case 3:
				newPos = Position{y:50,           x:pos.y+50,    direction: 0}
			case 4:
				newPos = Position{y:151-pos.y,    x:pos.x+50,    direction: 3}
			case 6:
				newPos = Position{y:150,          x:pos.y-100,   direction: 0}
			default:
				fmt.Printf("\npanic dir:> , face:%d\n", thisFace)
				panic(13)
			}
			if grid[newPos.y][newPos.x] == '.' {
				pos = newPos
			}
		}
	case 2, -2:
		if (grid[pos.y+1][pos.x] == '.') {
			pos.y = pos.y + 1
		} else if grid[pos.y+1][pos.x] == ',' {
			switch (thisFace) {
			case 2:
				newPos = Position{y:pos.x-50,     x:100,         direction: 3}
			case 4:
				newPos = Position{y:pos.x+100,    x:50,          direction: 3}
			case 6:
				newPos = Position{y:1,            x:pos.x+100,   direction: 2}
			default:
				fmt.Printf("\npanic dir:v , face:%d\n", thisFace)
				panic(13)
			}

			if grid[newPos.y][newPos.x] == '.' {
				pos = newPos
			}
		}
	case 3, -1:
		if (grid[pos.y][pos.x-1] == '.') {
			pos.x = pos.x - 1
		} else if  grid[pos.y][pos.x-1] == ',' {
			switch (thisFace) {
			case 1:
				newPos = Position{y:151-pos.y,     x:1,         direction: 1}
			case 3:
				newPos = Position{y:101,           x:pos.y-50,  direction: 2}
			case 5:
				newPos = Position{y:151-pos.y,     x:51,        direction: 1}
			case 6:
				newPos = Position{y:1,            x:pos.y-100,  direction: 2}
			default:
				fmt.Printf("\npanic dir:< , face:%d\n", thisFace)
				panic(13)
			}

			if grid[newPos.y][newPos.x] == '.' {
				pos = newPos
			}
		}
	default:
		panic(13)
	}

}

func advance_1(){
	dir := pos.direction % 4
	if dir == 0 {                              //up
		if (grid[pos.y-1][pos.x] == '.') {
			pos.y = pos.y - 1
		} else if  grid[pos.y-1][pos.x] == ',' {
			for i:=299; i>=0; i--{
				if grid[i][pos.x] == '.'{
					pos.y = i
					break
				} else if grid[i][pos.x] == '#'{
					break
				}
			}
		}
	} else if dir == 1 || dir == -3 {          //righ
		if (grid[pos.y][pos.x+1] == '.') {
			pos.x = pos.x + 1
		} else if  grid[pos.y][pos.x+1] == ',' {
			for j:=0; j<180; j++{
				if grid[pos.y][j] == '.'{
					pos.x = j
					break
				} else if grid[pos.y][j] == '#'{
					break
				}
			}
		}
	} else if dir == 2 || dir == -2 {          // down
		if (grid[pos.y+1][pos.x] == '.') {
			pos.y = pos.y + 1
		} else if grid[pos.y+1][pos.x] == ',' {
			for i:=0; i<300; i++{
				if grid[i][pos.x] == '.'{
					pos.y = i
					break
				} else if grid[i][pos.x] == '#'{
					break
				}
			}
		}
	} else if dir == 3 || dir == -1 {          //left
		if (grid[pos.y][pos.x-1] == '.') {
			pos.x = pos.x - 1
		} else if  grid[pos.y][pos.x-1] == ',' {
			for j:=179; j>=0; j--{
				if grid[pos.y][j] == '.'{
					pos.x = j
					break
				} else if grid[pos.y][j] == '#'{
					break
				}
			}
		}
	}
}

func initialize_grid(){
	for i:=0; i<300; i++ {
		for j:=0; j<180; j++ {
			grid[i][j] = ','
		}
	}
}

func set_row(row int, line string) {
	for j, chr := range line {
		if chr == '\n' {
			panic(1)
		}
		if line[j] == ' '{ 
			grid[row][j+1] = ','
		} else {
			grid[row][j+1] = line[j]
		}
	}
}

func print_grid(row_from, row_end int) {
	for i:=row_from; i<=row_end; i++ {
		for j:=0; j<160; j++ {
			if i == pos.y && j == pos.x {
				dir := pos.direction % 4
				switch(dir) {
				case 0:
					fmt.Printf("^")
				case 1, -3:
					fmt.Printf(">")
				case 2, -2:
					fmt.Printf("v")
				case 3, -1:
					fmt.Printf("<")
				default:
					fmt.Printf("\n%c\n", pos.direction)
					panic(123)
				}
			}else {
				fmt.Printf("%c", grid[i][j])
			}
		}
		fmt.Println()
	}
}


func print_around_pos(rowSize int){
	fmt.Printf("\033[2;0H")
	row_from := pos.y - rowSize / 2
	row_to := pos.y + rowSize / 2
	
	if row_from < 0 {row_from, row_to = 0, rowSize}
	if row_to > 299 {row_from, row_to = 299-rowSize, 299}
	print_grid(row_from, row_to)
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

func set_initial_position()(int, int) {
	for j:=0; j<180; j++ {
		if grid[1][j] == '.'  {
			pos.y = 1
			pos.x = j
			pos.direction = 1
			return 1, j
		}
	}
	panic(2)
}

func execute_a_cmd(cmd_str string, ith_cmd int, part1or2 int){
	fmt.Printf("\033[1;0H")
	fmt.Printf("executing %s\t\t%d of %d    ", cmd_str, ith_cmd, len(cmds))

	val, err := strconv.Atoi(cmd_str)
	if err == nil { //it's advance type command
		for i:=0; i<val; i++ {
			if part1or2 == 1 {
				advance_1()
			} else if part1or2 == 2{
				part2_advance_1()
			}
			// print_around_pos(60)
			// time.Sleep(time.Millisecond * 1)
		}
	} else { // it's rotate type command
		if cmd_str == "R" {
			pos.direction++
		} else if cmd_str == "L" {
			pos.direction--
		} else {
			panic(3)
		}
	}


	print_around_pos(60)
    time.Sleep(time.Millisecond)

}

func isit_within(x, m, M int) bool {
	return m <= x && x <= M
}

func face_identifier() int {
	if isit_within(pos.y,  1,   50)  && isit_within(pos.x, 51,  100)  {return 1}
	if isit_within(pos.y,  1,   50)  && isit_within(pos.x, 101, 150)  {return 2}
	if isit_within(pos.y,  51,  100) && isit_within(pos.x, 51,  100)  {return 3}
	if isit_within(pos.y,  101, 150) && isit_within(pos.x, 51,  100)  {return 4}
	if isit_within(pos.y,  101, 150) && isit_within(pos.x, 1,   50)   {return 5}
	if isit_within(pos.y,  151, 200) && isit_within(pos.x, 1,   50)   {return 6}

	panic(23)
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
	fmt.Printf("\033[2J")
	set_initial_position()
	for i, cmd := range cmds {
		execute_a_cmd(cmd, i, 2)
	}

	var facing int
	switch(pos.direction % 4) {
	case 0:
		facing = 3
	case 1, -3:
		facing = 0
	case 2, -2:
		fmt.Printf("v")
		facing = 1
	case 3, -1:
		fmt.Printf("<")
		facing = 2
	default:
		fmt.Printf("\n%c\n", pos.direction)
		panic(123)
	}
	final_pass := pos.y * 1000 + 4*pos.x +facing
		fmt.Printf("final pass:%d\n", final_pass)

}
