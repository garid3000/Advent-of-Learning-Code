package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	//"strconv"

	"strconv"
	//"log"
	"strings"

	//"gopkg.in/ini.v1"
)


type Coord struct {
	y, x int;
}

var (
	x_max = -9999
    x_min =  9999
	y_max = -9999
	y_min =  9999
	grid = [200][1000]int{}
	x_offset = 0
	y_offset = 0
)

func write2file(fname string) {
    f, err := os.Create(fname)
	if err != nil { panic(123) }
	tmprow := [501]byte{}
	tmprow[500] = '\n'
	
	for yy:=0; yy<175; yy++{
		for xx:=250; xx<750; xx++{
			switch (grid[yy][xx]) {
			case 1:
				tmprow[xx-250] = '#'
			case 0:
				tmprow[xx-250] = '.'
			case 2:
				tmprow[xx-250] = 'o'
			default:
				panic(123)
			}
		}
		_, err := f.WriteString(string(tmprow[:]))
		if err != nil { panic(1234)}
	}
}

func print_grid(fname string){
	tmpstr := ""
	for yy:=0; yy<200; yy++{
		for xx:=0; xx<1000; xx++{
			if grid[yy][xx] == 1{
				tmpstr += "#"
				//fmt.Printf("#")
			} else if yy==0 && xx==100 { 
				tmpstr += "+"
				// fmt.Printf("+")
			} else  if grid[yy][xx] == 2{
				tmpstr += "o"
				// fmt.Printf("o")
			} else {
				tmpstr += "."
				// fmt.Printf(".")
			} 
		}
		tmpstr += "\n"
		// fmt.Printf("\n")
	}

	if len(fname) > 0 {
		err := os.WriteFile(fname, []byte(tmpstr), 0644)
		if err != nil {panic(err)}
	}

}


func coor_parser(coor_str string) Coord{
	var yy, xx int;
	fmt.Sscanf(coor_str, "%d,%d", &xx, &yy)
	xx -= x_offset
	yy -= y_offset

	if xx > x_max {x_max=xx}
	if xx < x_min {x_min=xx}
	if yy > y_max {y_max=yy}
	if yy < y_min {y_min=yy} //loggin max min x y

	return Coord{x:xx, y:yy}
}

func inc_val(initial, end int) int{
	if initial > end {return -1} else {return 1}
}

func draw_line(p1, p2 Coord){
	if p1.y == p2.y {         //vertical line
		inc := inc_val(p1.x, p2.x)
		for xx:=p1.x; xx!=p2.x; xx += inc {
			grid[p1.y][xx] = 1
		}
		grid[p2.y][p2.x] = 1
	} else if p1.x == p2.x { //horizontal line
		inc := inc_val(p1.y, p2.y)
		for yy:=p1.y; yy!=p2.y; yy += inc {
			grid[yy][p1.x] = 1
		}
		grid[p2.y][p2.x] = 1
	} else {
		fmt.Printf("wtf: %v-%v\n", p1, p2)
		panic(1)
	}
}

func path_parser(pathstr string) {
	points_str := strings.Split(pathstr, " -> ")
	fmt.Println(points_str)
	for i,size:=0,len(points_str); i<size-1; i++ {
		a_coord := coor_parser(points_str[i])
		b_coord := coor_parser(points_str[i+1])
		draw_line(a_coord, b_coord)

		fmt.Printf("\t%v\t%v\n", a_coord, b_coord)
	}
}

func movesand(sand_pos Coord) (Coord, bool){
	if sand_pos.y >= 199 {return sand_pos, true;} // panic(420)} // sand went to abyss


	//move the sand
	newPos := Coord{} //new coord

	//check below coordinate
	newPos.x = sand_pos.x
	newPos.y = sand_pos.y + 1
	if grid[newPos.y][newPos.x] == 0 { //below is empty
		return movesand(newPos)
	}
	

	//check diagonale left
	newPos.x = sand_pos.x - 1
	newPos.y = sand_pos.y + 1
	if grid[newPos.y][newPos.x] == 0 { //below is empty
		return movesand(newPos)
	}


	//check diagonale left
	newPos.x = sand_pos.x + 1
	newPos.y = sand_pos.y + 1
	if grid[newPos.y][newPos.x] == 0 { //below is empty
		return movesand(newPos)
	}

	return sand_pos, false
}

func put_sands(){
	initialsandpos := Coord{y:0, x:500} //500,0 , xoffset -100, so x=100
	for i:=1;;i++ {
		end_sand_pos, ret := movesand(initialsandpos)
		if ret == false {
			grid[end_sand_pos.y][end_sand_pos.x] = 2
			//print_grid("grid/"+strconv.Itoa(i))
			write2file("grid/"+strconv.Itoa(i))
		} else {
			fmt.Printf("abyss %d\n", i-1)
			break
		}

		if end_sand_pos.x == initialsandpos.x &&  end_sand_pos.y == initialsandpos.y {
			fmt.Printf("Full %d\n", i)
			break
		}

	}

}

func main() {
	fmt.Printf("AoC2022 day 14\n")

	//fname := "/home/garid/Documents/advent/AoC-2022/day14/test"
	fname := "/home/garid/Documents/advent/AoC-2022/day14/input.txt"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	reader := bufio.NewReader(file)
	count := 0 
	for i := 0; ; i++ {
		line, ret1 := reader.ReadString('\n')
		if ret1 == io.EOF{
			fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			break
		}

		line = line[:len(line)-1]
		fmt.Println(line)
		path_parser(line)
	}
	print_grid("grid/0")
	//put_sands()
	
	fmt.Printf("count: %v\n", count) // this is part1
	fmt.Printf("y_max %v\n", y_max) // this is part1
	fmt.Printf("y_min %v\n", y_min) // this is part1
	fmt.Printf("x_max %v\n", x_max) // this is part1
	fmt.Printf("x_min %v\n", x_min) // this is part1


	//putting the horizontal

	for xx:=0; xx < 1000; xx++ {
		grid[y_max + 2][xx] = 1
	}
	put_sands()

}
