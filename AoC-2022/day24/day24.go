package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "time"
)


var (
	round = 1
	newgrid = [7][30][125] byte {};
	finaly, finalx = 0, 0
	inity, initx = 0, 0
	horsize, versize = 0, 0
	stage = 1  //part 1 or part 2
)


// new_parsing ...
func new_parsing(ithrow int, line string)  {
	for j,L:=0,len(line); j<L; j++ {
		switch line[j] {
		case '^':
			newgrid[0][ithrow][j] = '^'
		case '>':
			newgrid[1][ithrow][j] = '>'
		case 'v':
			newgrid[2][ithrow][j] = 'v'
		case '<':
			newgrid[3][ithrow][j] = '<'
		default:
			;
		}

		if ithrow == 0 && '.' == line[j] {
			inity, initx = ithrow, j
		} else if '.' == line[j]  && line[2] == '#' && line[3] == '#'{ // this should be the last line
			finaly, finalx = ithrow, j
			horsize, versize = L-1, ithrow

			// making frames
			for kk:=0; kk<7; kk++ {
				for i:=0;i<versize;i++{
					newgrid[kk][i][0] = '#'
					newgrid[kk][i][horsize] = '#'
				}
				
				for j:=0;j<horsize;j++{
					newgrid[kk][0][j] = '#'
					newgrid[kk][versize][j] = '#'
				}
			}
			// making initial position of
			newgrid[5][inity][initx] = 'o'
			newgrid[6][inity][initx] = ' '
			newgrid[4][inity][initx] = ' '
			newgrid[6][finaly][finalx] = ' '
			newgrid[5][finaly][finalx] = ' '
			newgrid[6][finaly][finalx] = ' '
			newgrid[4][finaly][finalx] = ' '
		}
	}
}

// printing_newgrid ...
func printing_newgrid()  {
	// fmt.Printf("\033[2J");

	fmt.Printf("\033[%d;%dH Round:\t%d", 0, 0, round)
	for kk:=0; kk<6; kk+=2 {
		for ii:=0; ii<=versize; ii++ {
			for jj:=0; jj<=horsize; jj++ {
				// fmt.Printf("%c", newgrid[kk][ii][jj])
				fmt.Printf("\033[%d;%dH%c", 3 + kk/2*(versize+2) + ii, 3 + jj,     newgrid[kk][ii][jj])
			}
			for jj:=0; jj<=horsize; jj++ {
				// fmt.Printf("%c", newgrid[kk+1][ii][jj])
				fmt.Printf("\033[%d;%dH%c", 3+ kk/2*(versize+2) + ii, 3+ jj+horsize + 5, newgrid[kk+1][ii][jj])
			}
		}
	}


	// for ii:=0; ii<=versize; ii++ {
	// 	for jj:=0; jj<=horsize; jj++ {
	// 		fmt.Printf("\033[%d;%dH%c", 3+ 2*(versize+2) + ii, 3+ jj+horsize*2 + 10, newgrid[6][ii][jj])
	// 	}
	// }
}



func moveblizzard_up()  { //^
	tmp := make([]byte, horsize, horsize)
	for j:=0;j<horsize; j++{tmp[j] = newgrid[0][1][j]}

	for ithrow:=1; ithrow<versize-1; ithrow++{
		for j:=0;j<horsize; j++{
			newgrid[0][ithrow][j] = newgrid[0][ithrow+1][j]
		}
	}
	for j:=0;j<horsize; j++{newgrid[0][versize-1][j] = tmp[j]}
}



func moveblizzard_down()  { // V
	tmp := make([]byte, horsize, horsize)
	for j:=0;j<horsize; j++{tmp[j] = newgrid[2][versize-1][j]}

	for ithrow:=versize-1; ithrow>1; ithrow--{
		for j:=0;j<horsize; j++{
			newgrid[2][ithrow][j] = newgrid[2][ithrow-1][j]
		}
	}
	for j:=0;j<horsize; j++{newgrid[2][1][j] = tmp[j]}
}

func moveblizzard_left()  { //<
	tmp := make([]byte, versize, versize)
	for i:=0;i<versize; i++{
		tmp[i] = newgrid[3][i][1] // remembers left most column
	}
	for jthcol:=1; jthcol<horsize-1; jthcol++{ // set jth column with jthcol+1
		for i:=0; i<versize; i++{
			newgrid[3][i][jthcol] = newgrid[3][i][jthcol+1]  
		}
	}
	for i:=0; i<versize; i++{newgrid[3][i][horsize-1] = tmp[i]}
}



func moveblizzard_right()  { //>
	tmp := make([]byte, versize, versize)
	for i:=0;i<versize; i++{
		tmp[i] = newgrid[1][i][horsize-1] // remembers rightmost most column
	}
	for jthcol:=horsize-1; jthcol>1; jthcol--{ // set jth column with jthcol+1
		for i:=0; i<versize; i++{
			newgrid[1][i][jthcol] = newgrid[1][i][jthcol-1]  
		}
	}
	for i:=0; i<versize; i++{newgrid[1][i][1] = tmp[i]}
}

func calc_empty_grid(){
	for ii:=1; ii<versize; ii++{
		for jj:=1; jj<horsize; jj++ {
			newgrid[4][ii][jj] = ' '
			for kk:=0; kk <4; kk++{
				if newgrid[kk][ii][jj] != ' ' {
					newgrid[4][ii][jj] = '#'
				}
			}
		}
	}
}

// abs ...
func abs(v int) int {
	if v > 0 {return v;} else {return -v;}
}

func explore(){
	// this should always executed after the calculating empty  (i.e. movealbe) cells

	//step 1 remove positions now has blizards
	for ii:=1; ii<versize-1; ii++{
		for jj:=1; jj<horsize-1; jj++ {
			if newgrid[4][ii][jj] == '#' && newgrid[5][ii][jj] == 'o' {
				newgrid[5][ii][jj] = '.'
			}
		}
	}

	// printing_newgrid()
	// time.Sleep(time.Millisecond * 10)

	// populate future grid (i.e. calc which cell elves can move into)
	for ii:=0; ii<versize+1; ii++{ 
		for jj:=1; jj<horsize; jj++ {
			if newgrid[5][ii][jj] == 'o' || newgrid[5][ii][jj] == '.'  { // this round
				if ii > 0{
				if newgrid[4][ii-1][jj] == ' ' {newgrid[6][ii-1][jj] = 'o'}
				}
				if newgrid[4][ii+1][jj] == ' ' {newgrid[6][ii+1][jj] = 'o'}
				if newgrid[4][ii][jj-1] == ' ' {newgrid[6][ii][jj-1] = 'o'}
				if newgrid[4][ii][jj+1] == ' ' {newgrid[6][ii][jj+1] = 'o'}
				if newgrid[4][ii][jj]   == ' ' {newgrid[6][ii][jj]   = 'o'}
			}
		}
	}

	// printing_newgrid()
	// time.Sleep(time.Millisecond * 10)

	// move into those cells calculated last step
	//////////////////////////////////////////////////////////
	// for ii:=1; ii<versize+1;ii++{  // here			    //
	// 	for jj:=1; jj<horsize; jj++ {					    //
	// 		newgrid[5][ii][jj] = newgrid[6][ii][jj]		    //
	// 		if ii != versize{							    //
	// 			newgrid[6][ii][jj] = ' '				    //
	// 		}else if (ii != finaly && jj != finalx) {	    //
	// 			newgrid[6][ii][jj] = '#'				    //
	// 		}											    //
	// 	}												    //
	// }												    //
	//////////////////////////////////////////////////////////


	for ii:=1; ii<versize; ii++ {
		for jj:=1; jj<horsize; jj++{
			newgrid[5][ii][jj] = newgrid[6][ii][jj]
			newgrid[6][ii][jj] = ' '
		}
		newgrid[5][inity][initx] = newgrid[6][inity][initx]
		newgrid[5][finaly][finalx] = newgrid[6][finaly][finalx]
	}
}


func check_stage1_or_2(){
	if newgrid[5][finaly][finalx] == 'o' && stage == 1{
		stage++;

		for ii:=1; ii<versize;ii++{  
			for jj:=1; jj<horsize; jj++ {
				newgrid[5][ii][jj] = ' '
			}
		}
		newgrid[5][inity][initx] = ' '
		newgrid[6][inity][initx] = ' '
	}

	if newgrid[5][inity][initx] == 'o' && stage == 2{
		// panic(2)
		stage++;

		for ii:=1; ii<versize;ii++{  
			for jj:=1; jj<horsize; jj++ {
				newgrid[5][ii][jj] = ' '
			}
		}
		newgrid[5][finaly][finalx] = ' '
		newgrid[6][finaly][finalx] = ' '
	}

	if newgrid[5][finaly][finalx] == 'o' && stage == 3{
		panic(3)
	}
}


func oneround(){
	moveblizzard_up()
	moveblizzard_down()
	moveblizzard_left()
	moveblizzard_right()
	calc_empty_grid()
	explore()

	// printing_newgrid()
	check_stage1_or_2()

	printing_newgrid()
	// time.Sleep(time.Millisecond * 10)
	round++
}

func main() {
	for kk:=0; kk<7; kk++ {
		for ii:=0; ii<30; ii++{
			for jj:=0; jj<125; jj++{
				newgrid[kk][ii][jj] = ' '
			}
		}
	}
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
		new_parsing(i, line)
	}
	fmt.Printf("%v %v\n", horsize, versize)

	fmt.Printf("\033[2J");


	printing_newgrid()
	// time.Sleep(time.Millisecond * 1000)
	for ;;{
		oneround()
		// time.Sleep(time.Millisecond * 100)
	}
}
