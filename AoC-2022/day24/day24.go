package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type explored_cell struct {
	y, x int;
	arived_at int;
	available bool;
}

type blizzard struct {
	y, x int;
	dir byte;
	//up 0, right 1, down 2, left 3
}

var (
	round = 0 
	grid = [30][125] byte{}
	blizzards = make([]blizzard, 0, 3750)
	num_blizzard = 0
	explored_cells = make([]explored_cell, 0, 3750)
	finaly, finalx = 0, 0
	inity, initx = 0, 0
)

func print_full_grid() {
	for ii:=0; ii<30; ii++{
		for jj:=0; jj<125; jj++{
			fmt.Printf("%c", grid[ii][jj])
		}
		fmt.Println()
	}
}


func parse_line(irow int, line string){
	for j,L:=0,len(line); j<L; j++ {
		grid[irow][j] = line[j]
		if '^'  == line[j] || '>'  == line[j] || 'v'  == line[j] || '<'  == line[j] {
			blizzards = append(blizzards, blizzard{y:irow, x:j, dir:line[j]})
		} else if irow == 0 && '.' == line[j] {
			inity, initx = irow, j
		} else if '.' == line[j] {
			finaly, finalx = irow, j
		}
	}

}

func run_1_blizzard(ith int) { //run ith blizzards
	newx, newy := blizzards[ith].x,  blizzards[ith].y
	switch(blizzards[ith].dir){
	case '^':
		if grid[blizzards[ith].y-1][blizzards[ith].x] == '#' {
			for i:=29; i>0; i-- {
				if grid[i][blizzards[ith].x] != '#'{
					newy = i
					break
				}
			}
		} else {
			newy = blizzards[ith].y-1;
		}
	case '>':
		if grid[blizzards[ith].y][blizzards[ith].x+1] == '#' {
			for j:=0; j<125; j++ {
				if grid[blizzards[ith].y][j] != '#'{
					newx = j
					break
				}
			}
		} else {
			newx = blizzards[ith].x+1;
		}

		// fmt.Printf("\n %d=newx, %d=old\t\t %c", newx, blizzards[ith], grid[blizzards[ith].y][blizzards[ith].x+1]) 
	case 'v':
		if grid[blizzards[ith].y+1][blizzards[ith].x] == '#' {
			for i:=0; i<30; i++ {
				if grid[i][blizzards[ith].x] != '#'{
					newy = i
					break
				}
			}
		} else {
			newy = blizzards[ith].y+1;
		}
	case '<':
		if grid[blizzards[ith].y][blizzards[ith].x-1] == '#' {
			for j:=124; j>0; j-- {
				if grid[blizzards[ith].y][j] != '#'{
					newx = j
					break
				}
			}
		} else {
			newx = blizzards[ith].x-1;
		}
	default:
		panic(99)
	}

	// fmt.Printf("%v %v %c ---> %v %v\n", blizzards[ith].y, blizzards[ith].x, blizzards[ith].dir, newy, newx)
	blizzards[ith].x = newx;
	blizzards[ith].y = newy;
	//remove availle explored cell at newx newy
	
	for _, ecell := range explored_cells{
		if ecell.x == newx && ecell.y == newy{
			ecell.available = false
			// fmt.Printf("removing %v", ecell)
		}
	}
}

func run_1_round() {
	print_blizzards()
	// time.Sleep(time.Millisecond * 3000)
	time.Sleep(time.Millisecond)
	for i := range blizzards {
		run_1_blizzard(i)
	}

	exploring()
	round++
}

func currentavailable()int{
	sum :=0
	for _, ecell := range explored_cells{
		if ecell.available{ sum++}
	}
	return sum
}

func print_blizzards() {
	fmt.Printf("\033[2J");
	fmt.Printf("\033[%d;%dH\tround:%d, explored %d, avail %d", 0, 0, round, len(explored_cells), currentavailable())
	for _,b := range blizzards {
		fmt.Printf("\033[%d;%dH%c", b.y+1, b.x+1, b.dir)
	}

	for _, ecell := range explored_cells {
		if ecell.available {
			fmt.Printf("\033[%d;%dH%c", ecell.y+30, ecell.x+1, '@')
		} else {
			fmt.Printf("\033[%d;%dH%c", ecell.y+30, ecell.x+1, '-')
		}
	}
}




func canImove2thiscell(y, x int) bool {
	if x < 0 || y <0 {return false}
	for _,b := range blizzards {
		if b.x == x && b.y == y {
			return false
		}
	}
	if grid[y][x] == '#' {
		return false
	}
	return true
}

func visit_or_makeavail_thisCell(_y, _x int) {
	for _, thiscell := range explored_cells {
		if thiscell.y == _y && thiscell.x == _x {
			// i have visited this cell before
			thiscell.available = true
			return
		}
	}
	// i haven't visited this cell
	explored_cells = append(
		explored_cells,
		explored_cell{
			y: _y,
			x: _x,
			arived_at: round,
			available: true,
		},
	) 

	if _y == finaly && _x == finalx {
		for ;; {
			fmt.Printf("%v\n", round + 2)
		}
	}
}

func exploring(){
	for _, ecell := range explored_cells {
		if ecell.available {
			if canImove2thiscell(ecell.y-1, ecell.x  ){ visit_or_makeavail_thisCell(ecell.y-1, ecell.x  )}
			if canImove2thiscell(ecell.y+1, ecell.x  ){ visit_or_makeavail_thisCell(ecell.y+1, ecell.x  )}
			if canImove2thiscell(ecell.y  , ecell.x-1){ visit_or_makeavail_thisCell(ecell.y  , ecell.x-1)}
			if canImove2thiscell(ecell.y  , ecell.x+1){ visit_or_makeavail_thisCell(ecell.y  , ecell.x+1)}
		}
	}
}


func main() {
	for ii:=0; ii<30; ii++{for jj:=0; jj<125; jj++{grid[ii][jj] = '#'}}
	fmt.Printf("AoC2022 day24\n")
	fname := "/home/garid/Documents/advent/AoC-2022/day24/input.txt" ;
	// fname := "/home/garid/Documents/advent/AoC-2022/day24/test"
	// fname := "/home/garid/Documents/advent/AoC-2022/day24/test_small"

	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}
	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		line, ret1 := reader.ReadString('\n')
		if ret1 == io.EOF {
			fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			break
		}
		line = line[:len(line)-1]
		fmt.Printf("%d\t%v\n", i, line )
		parse_line(i, line)
	}
	num_blizzard = len(blizzards)
	print_full_grid()
	// fmt.Printf("len of blizzards:%v\n", blizzards)
	fmt.Printf("len of blizzards:%v\n", len(blizzards))
	fmt.Printf("journey starts at: %v, %v\n", inity, initx)
	fmt.Printf("journey ends   at: %v, %v\n", finaly, finalx)

	explored_cells = append(
		explored_cells,
		explored_cell{
			y:inity,
			x:initx,
			arived_at: 0,
			available: true,
		},
	)


	for i:=0; i<1000; i++ {
		run_1_round()
	}

}
