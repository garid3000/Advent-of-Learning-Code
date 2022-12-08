package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var arr[99][99] int;

//func visible_tree_1line(aline [99] int) int {
//	sum := 0
//	current_max := -1
//	for i:=0; i<99; i++{
//		if aline[i] > current_max{
//			current_max = aline[i]
//			sum++;
//		}
//	}
//
//	for i:=98; i>-1; i++{
//		if aline[i] > current_max{
//			current_max = aline[i]
//			sum++;
//		}
//	}
//
//	return sum
//}

func scene_score(i int, j int) int{
	c_up := 0
	//to up
	for ii:=i-1; ii>=0; ii--{
		if arr[ii][j] >= arr[i][j]{
			c_up++
			//fmt.Printf("to up not arr[%d,%d]=%d\n", ii, j, arr[ii][j])
			break
		} else {
			c_up++
		}
	}

	//to down
	c_down :=0
	for ii:=i+1; ii<99; ii++{
		if arr[ii][j] >= arr[i][j]{
			c_down++
			//fmt.Printf("to down not arr[%d,%d]=%d\n", ii, j, arr[ii][j])
			break
		} else {c_down++}
	}

	//to left
	c_left :=0
	for jj:=j-1; jj>=0; jj--{
		if arr[i][jj] >= arr[i][j]{
			c_left++
			//fmt.Printf("to left not arr[%d,%d]=%d\n", i, jj, arr[i][jj])
			break
		} else {c_left++}
	} 


	//to right
	c_right :=0
	for jj:=j+1; jj<99; jj++{
		if arr[i][jj] >= arr[i][j]{
			c_right++
			//fmt.Printf("to right not arr[%d,%d]=%d\n", i, jj, arr[i][jj])
			break
		} else {c_right++}
	}

	return c_up * c_down * c_right * c_left
}










func isit_visible(i int, j int) int{
	state := 1
	//to up
	for ii:=i-1; ii>=0; ii--{
		if arr[ii][j] >= arr[i][j]{
			state = 0
			//fmt.Printf("to up not arr[%d,%d]=%d\n", ii, j, arr[ii][j])
			break
		}
	}
	if state >0 {return state}

	//to down
	state =1
	for ii:=i+1; ii<99; ii++{
		if arr[ii][j] >= arr[i][j]{
			state = 0
			//fmt.Printf("to down not arr[%d,%d]=%d\n", ii, j, arr[ii][j])
			break
		}
	}
	if state >0 {return state}

	//to left
	state =1
	for jj:=j-1; jj>=0; jj--{
		if arr[i][jj] >= arr[i][j]{
			state = 0
			//fmt.Printf("to left not arr[%d,%d]=%d\n", i, jj, arr[i][jj])
			break
		}
	}
	if state >0 {return state}


	//to right
	state =1
	for jj:=j+1; jj<99; jj++{
		if arr[i][jj] >= arr[i][j]{
			state = 0
			//fmt.Printf("to right not arr[%d,%d]=%d\n", i, jj, arr[i][jj])
			break
		}
	}
	if state >0 {return state}

	return state
}

func count_visible_trees() int {
	sum:=0
	for i:=1; i<98; i++{
		for j:=1; j<98; j++{
			sum += isit_visible(i,j)
		}
	}
	return sum
}

func highest_scene_score() int {
	highest:=0
	for i:=1; i<98; i++{
		for j:=1; j<98; j++{
			tmp:= scene_score(i, j)
			if tmp > highest{
				highest = tmp
			}
		}
	}
	return highest
}



func main(){
	fmt.Printf("AoC2022\n");
	fname := "/home/garid/Documents/advent/AoC-2022/day8/input.txt"
	file, err:= os.Open(fname)
	if err != nil{
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	reader := bufio.NewReader(file)
	for i:=0;;i++{
		line, ret := reader.ReadString('\n')
		if ret == io.EOF{
			fmt.Printf("File has ended\n")
			break
		}
		fmt.Printf("%d\t%s\t%v\n", i, line[:len(line)-1], len(line)-1)

		for j,L:=0,len(line)-1;j<L; j++{
			arr[i][j] = int(line[j] - '0') 
		}
	}
	//fmt.Println(
	//	count_visible_trees())
    x := 1
	y := 2
	fmt.Printf("arr[%d,%d]=%d\n", x,y, arr[x][y])
	fmt.Println(isit_visible(1,2))


	fmt.Println("visible tree:",
		count_visible_trees() + 99 * 4 - 4)


	fmt.Println("highest tree: ",
		highest_scene_score())

}
