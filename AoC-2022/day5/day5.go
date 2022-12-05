package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
)
var str_stacks [9] string;


func line2_ninechararr(line string) [9] byte{
	var arr [9] byte;
	for i,j:=0,1;i<9;i,j=i+1,j+4{
		arr[i] = line[j]
	} 
	return arr
}

func str_inverter(str string) string{
	outstr := ""
	for i,L := 0, len(str); i<L; i++{
		outstr += string(str[L-i-1])
	}
	return outstr
}

func move1crate(src_index int, des_index int){
	//get last item of src_index=th stack
	ch := str_stacks[src_index][len(str_stacks[src_index])-1]
	// remove this item from str_stack[src_index]
	str_stacks[src_index] = str_stacks[src_index][:len(str_stacks[src_index])-1]
	//add ch to desitionation stack
	str_stacks[des_index] += string(ch)
}


func moveNcrates(src_index int, des_index int, num int){
	for i:=0; i<num;i++{
		move1crate(src_index, des_index)
	}
}


func moveNcrates_with9001(src_index int, des_index int, num int){
	tmp := str_stacks[src_index][len(str_stacks[src_index]) - num:]
	str_stacks[src_index] = str_stacks[src_index][:len(str_stacks[src_index]) - num]
	str_stacks[des_index] += tmp

}
func printLast_aka_top_chars(){
	for i:=0; i<9; i++ {
		if len(str_stacks[i]) == 0 {
			fmt.Print('_')
		}

		fmt.Printf("%c",
			str_stacks[i] [len(str_stacks[i])-1])
	}

}

func main(){
	fmt.Printf("AOC 2022, day5\n")
	fname := "/home/garid/Documents/advent/AoC-2022/day5/input.txt"
	file, err := os.Open(fname)
	if err != nil{
		fmt.Printf("file %s is not found\n", fname);
		panic(2)
	}


	reader := bufio.NewReader(file)
	for i:=0;;i++{
		line, ret := reader.ReadString('\n')
		fmt.Printf("%d\t%v\t%s", i, ret, line)
		if line == " 1   2   3   4   5   6   7   8   9 \n"{
			fmt.Printf("now operations\n")
			line, ret = reader.ReadString('\n')

			for j:=0;j<9;j++{
				str_stacks[j] += "_"
			}


			break
		} else {
			//populate

			fmt.Printf("            ")
			crates_arr := line2_ninechararr(line)
			for j:=0;j<9;j++{
				//fmt.Printf("-%c- ", crates_arr[j])
				if crates_arr[j] != ' '{
					str_stacks[j] += string(crates_arr[j])
				}
			}

		}
	}
	fmt.Println(str_stacks)
	for j:=0; j<9; j++{
		str_stacks[j] = str_inverter(str_stacks[j])
	}
	fmt.Println(str_stacks)


	fmt.Print("\n\nStarting to move:")
	var src_stack, num_crates ,des_stack int;
	for i:=0;;i++{
		line, ret := reader.ReadString('\n')

		if ret == io.EOF{
			fmt.Printf("EOF\n")
			break
		}
		n, ret1 := fmt.Sscanf(line, "move %d from %d to %d",
			&num_crates, &src_stack, &des_stack)

		if ret1 != nil{
			fmt.Printf("panic at sscanf %v %v\n", n, ret1)
			panic(ret1)
		}

		fmt.Printf("%d\t%v %v %v %v\t----%s", i, ret,
			num_crates, src_stack,  des_stack, line)

		fmt.Println(str_stacks)
		//moveNcrates(src_stack-1, des_stack-1, num_crates)          //acts as part1 
		moveNcrates_with9001(src_stack-1, des_stack-1, num_crates)   //acts as part2
		fmt.Println(str_stacks)
		fmt.Println()
	}

	printLast_aka_top_chars()
}
