package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	// "time"
	//"strings"
)

var (
	grid          = [2022*4][7]byte{} // 
	max_columns   = [7]int{}
)


func min_max() (int,int) {
	min := max_columns[0]
	max := max_columns[0]
	for i:=1; i<7; i++ {
		if max_columns[i] > max {max = max_columns[i]}
		if max_columns[i] < min {min = max_columns[i]}
	}
	return min, max
}


func set_grid(){
	for i:=0; i<8088; i++ {
		for j:=0; j<7; j++ {
			grid[i][j] = '.'
		}
	}
}


//func set_grid_expected(){
//	for i:=0; i<8088; i++ {
//		for j:=0; j<7; j++ {
//			grid_expected[i][j] = '.'
//		}
//	}
//}

//change here
//func get_highest_peak() int{
//	var sum int
//	for i:=0; i<2022; i++{
//		sum = 0
//		for j:=0; j<7; j++ {
//			sum += int(grid[i][j]) - int('.')
//		}
//		if sum == 0 {
//			return i
//		}
//	}
//	fmt.Printf("grid overflow?")
//	panic(1)
//}

func draw_grid(from_higher, to_lower int, label string) {
	fmt.Print("\033[2J") //Clear screen
	fmt.Printf("\033[%d;%dH", 0, 0) // Set cursor position


	for h:=from_higher; h>=to_lower && h>=0; h-- {
		fmt.Printf(" %06d |%s|\n", h, grid[h])
	}
	fmt.Printf(" %06d +-------+\n%s\n%v\n", -1, label, max_columns)
    //time.Sleep(time.Millisecond)

}


func insert_shape_1() int{
	//highest_peak := get_highest_peak()
	_, highest_peak := min_max()
	if highest_peak != 0 {
		highest_peak += 1
	}
	grid[highest_peak + 3][0] = '.'
	grid[highest_peak + 3][1] = '.'
	grid[highest_peak + 3][2] = '@'
	grid[highest_peak + 3][3] = '@'
	grid[highest_peak + 3][4] = '@'
	grid[highest_peak + 3][5] = '@'
	grid[highest_peak + 3][6] = '.'
	return highest_peak
}

func insert_shape_2() int{
	//highest_peak := get_highest_peak()
	_, highest_peak := min_max()
	highest_peak += 1
	grid[highest_peak + 3][3] = '@'
	grid[highest_peak + 4][2] = '@'
	grid[highest_peak + 4][3] = '@'
	grid[highest_peak + 4][4] = '@'
	grid[highest_peak + 5][3] = '@'
	return highest_peak
}



func insert_shape_3() int{
	//highest_peak := get_highest_peak()
	_, highest_peak := min_max()
	highest_peak += 1
	grid[highest_peak + 3][2] = '@'
	grid[highest_peak + 3][3] = '@'
	grid[highest_peak + 3][4] = '@'
	grid[highest_peak + 4][4] = '@'
	grid[highest_peak + 5][4] = '@'
	return highest_peak
}


func insert_shape_4() int{
	//highest_peak := get_highest_peak()
	_, highest_peak := min_max()
	highest_peak += 1
	grid[highest_peak + 3][2] = '@'
	grid[highest_peak + 4][2] = '@'
	grid[highest_peak + 5][2] = '@'
	grid[highest_peak + 6][2] = '@'
	return highest_peak
}


func insert_shape_5() int{
	//highest_peak := get_highest_peak()
	_, highest_peak := min_max()
	highest_peak += 1
	grid[highest_peak + 3][2] = '@'
	grid[highest_peak + 4][2] = '@'
	grid[highest_peak + 3][3] = '@'
	grid[highest_peak + 4][3] = '@'
	return highest_peak
}



func is_dropping_rock_on_wall(sideofwall int) bool {
	//incase of the height limits
	end_lower_height, start_higher_height := min_max()
	end_lower_height = 0
	start_higher_height += 10
	if start_higher_height > 8088 {start_higher_height = 8088}
	if end_lower_height    < 0 {end_lower_height = 0}

	state := false
	if sideofwall == 1 {
		for h:=start_higher_height; h >= end_lower_height; h-- {
			if grid[h][6] == '@' {
				state = true
			}
		}
		return state
	} else if sideofwall == -1 {
		for h:=start_higher_height; h >= end_lower_height; h-- {
			if grid[h][0] == '@' {
				state = true
			}
		}
		return state
	} else  {
		fmt.Printf("Wrong sideofwall val: %d", sideofwall)
		panic(2)
	}
}



func moveSide(dir byte){
	end_lower_height, start_higher_height := min_max()
	start_higher_height += 10
	end_lower_height = 0
	//start_higher_height := initial_height + 10
	//end_lower_height    := initial_height - 10
	if start_higher_height > 8088 {start_higher_height = 8088}
	//if end_lower_height    < 0    {end_lower_height = 0}

	if dir == '>' {
		if is_dropping_rock_on_wall(1){
			return // can move right, cuz it's on rightmost wall
		}

		//check does it can move right (if not return from this function)
		for h:=start_higher_height; h >= end_lower_height; h-- {
			for c:=0; c<6; c++{  //0,1,2,3,4,5 and not 6
				if grid[h][c] == '@' { // if the block is the moving block
					if  grid[h][c+1] == '#' { // there's unmovable rock block on right
						return
					}
				}
			}
		}
		//comming here means moving block can move right
		for h:=start_higher_height; h >= end_lower_height; h-- {
			for c:=6; c>=1; c--{  //6, 5, ... 1 (not 0)
				if grid[h][c-1] == '@' {
					grid[h][c-1] = '.'
					grid[h][c] = '@'
				}
			}
		}
		return


	} else if dir == '<' {
		if is_dropping_rock_on_wall(-1){
			return // can move left, cuz it's on leftmost wall
		}

		//check does it can move left (if not return from this function)
		for h:=start_higher_height; h >= end_lower_height; h-- {
			for c:=1; c<7; c++{  //1,2,3,4,5,6 and not 0
				if grid[h][c] == '@' { // if the block is the moving block
					if  grid[h][c-1] == '#' { // there's unmovable rock block on left
						return
					}
				}
			}
		}
		//comming here means moving block can move right
		for h:=start_higher_height; h >= end_lower_height; h-- {
			for c:=0; c<=5; c++{  // 0,1,2,3,4,5, (not 6)
				if grid[h][c+1] == '@' {
					grid[h][c+1] = '.'
					grid[h][c] = '@'
				}
			}
		}
		return





	} else {
		fmt.Printf("unknown dir char: %c\n", dir)
		panic(1)
	}

}


func moveDown() bool{ // returns if the rock has froze
	end_lower_height, start_higher_height := min_max()
	end_lower_height = 0
	start_higher_height = start_higher_height + 10
	// end_lower_height    := initial_height - 10
	// if start_higher_height > 8088 {start_higher_height = 8088}
	// if end_lower_height    < 0    {end_lower_height = 0}

	// all items are movable down 1 
	for h:=start_higher_height; h >= end_lower_height; h-- {
		for c:=0; c<7; c++{
			if grid[h][c] == '@' {
				//fmt.Printf("hello");time.Sleep(time.Millisecond * 1000)
				if h==0{ // first layer
					//fmt.Printf("floor");time.Sleep(time.Millisecond * 20)
					freeze_movable_rock()
					return true
				} else if grid[h-1][c] == '#' {
					freeze_movable_rock()
					//fmt.Printf("rock");time.Sleep(time.Millisecond * 20)
					return true
				}
			}
		}
	}
	// it can move down
	for h:=end_lower_height; h <= start_higher_height; h++ {
		for c:=0; c<7; c++{
			if grid[h+1][c] == '@' {
				grid[h+1][c] = '.'
				grid[h][c] = '@'
			}
			//fmt.Printf("%v %v\n",  h, c);  time.Sleep(time.Millisecond * 100)
		}
	}
	//fmt.Printf("hasdofjadsklfj;adslkjf %v %v \n", end_lower_height, start_higher_height);  time.Sleep(time.Millisecond * 100)
	return false
}

func freeze_movable_rock(){
	end_lower_height, start_higher_height := min_max()
	end_lower_height = 0
	start_higher_height = start_higher_height + 10

	for h:=end_lower_height; h <= start_higher_height; h++ {
		for c:=0; c<7; c++{
			if grid[h][c]  == '@' {
				grid[h][c] = '#'
			}
		}
	}
	//update min max


	for c:=0; c<7; c++{
		for h:=start_higher_height+10; h >= end_lower_height; h-- {
			if grid[h][c] == '#'{
				max_columns[c] = h
				break
			}
		}
	}
}
var i_command_cursor =0
var commands = ""

func readCommand() byte {
	//if i_command_cursor == len(commands)  {
	//fmt.Printf("Panic")
	//panic(0)
	//}
	tmp := commands[i_command_cursor % len(commands)]
	//tmp = '<'
	i_command_cursor++ 
	return tmp
}


func insert_shape12345(j int) {
	switch j%5 {
		case 0:
			insert_shape_1()
		case 1:
			insert_shape_2()
		case 2:
			insert_shape_3()
		case 3:
			insert_shape_4()
		case 4:
			insert_shape_5()
		default:
			panic(99)
	}
}

func main() {
	fmt.Printf("AoC2022 day 17\n")
	set_grid()

	fname := "/home/garid/Documents/advent/AoC-2022/day17/input.txt" ;
	//fname := "/home/garid/Documents/advent/AoC-2022/day17/test"

	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	reader := bufio.NewReader(file)
	// var commands string
	for i := 0; ; i++ {
		line, ret1 := reader.ReadString('\n')
		if ret1 == io.EOF {
			fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			break
		}

		commands = line[:len(line)-1]
		fmt.Print(line, "\n")
		//line_parser(line)
	}
	
	//for j:=0;j<2022;j++{
	for j:=0;j<2022;j++{
		insert_shape12345(j) 
		_, M := min_max()
		draw_grid(M+10, M-30, "insert new")

		dir := readCommand()
		moveSide(dir)
		draw_grid(M+10, M-30, "moved     " + string(dir) + strconv.Itoa(i_command_cursor-1))

		for ;; {
			ret := moveDown()
			draw_grid(M+10, M-30, "down      j=" + strconv.Itoa(j))
			// if j > 1660 {
			// 	time.Sleep(time.Millisecond * 1000)
			// }

			if ret  {
				break
			}

			dir = readCommand()
			moveSide(dir)
			draw_grid(M+10, M-30, "moved     " + string(dir)+ strconv.Itoa(i_command_cursor-1) + "j=" + strconv.Itoa(j))
			//if j > 1660 {
			//	time.Sleep(time.Millisecond * 1000)
			//}
		}
		
		fmt.Printf("Placed %v\n", j)
	}
	//m,M := min_max()
	//fmt.Printf("max: %v %v\n", m, M )

}
