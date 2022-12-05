// Advent of Code com
// link:https://adventofcode.com/2022/day/4

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func isoverlap(elf1start int, elf1end int, elf2start int, elf2end int) bool {
	if       elf2start >= elf1start && elf2start <= elf1end {
		return true
	} else if   elf2end >= elf1start && elf2end <= elf1end {
		return true
	} else if elf1start >= elf2start && elf1start <= elf2end {
		return true
	} else if   elf1end >= elf2start && elf1end <= elf2end {
		return true
	} else{
		return false
	}

}

func is1contain2(elf1start int, elf1end int, elf2start int, elf2end int) bool {
	return (elf1start <= elf2start && elf1end >= elf2end)
}

func isfulloverlap(elf1start int, elf1end int, elf2start int, elf2end int) bool {
	return is1contain2(elf1start, elf1end, elf2start, elf2end) || is1contain2(elf2start, elf2end, elf1start, elf1end)
}

func parser(line string) [4]int{
	//var elf1_s, elf1_e, elf2_s, elf2_e int;
	var arr [4] int;
	stage := 0
	for i,L := 0, len(line)-1; i<L; i++ {
		if line[i] == '-' || line[i] == ',' {
			stage++;
		} else if line[i] >= '0'  && line[i] <='9' {
			arr[stage] = arr[stage] * 10;
			arr[stage] = arr[stage] + int(line[i] - '0');
		} else {
			fmt.Printf("Bad char %c %d\n", line[i], int(line[i])) 
			panic(3)
		}

	}

	fmt.Printf("%v\t%s", arr, line)
	return arr

	//		??_n, ret := fmt.Fscanln(reader, "%d-%d,%d-%d", &elf1_s, &elf1_e, &elf2_s, &elf2_e)

	//fmt.Printf("%d %d\t %d %d %d %d %d\n",
	//	i, _n, elf1_s, elf1_e, elf2_s, elf2_e, sum)
}


func part1_2(){
	fmt.Printf("This is day4\n");
	fname := "/home/garid/Documents/advent/AoC-2022/day4/input.txt"
	file, err := os.Open(fname)
	if err != nil{
		fmt.Printf("file %s is not found\n", fname);
		panic(2)
	}

	reader := bufio.NewReader(file)
	//read by each linea
	for i, fullcontain,overlap:=0,0,0 ;;i++{
		line, ret := reader.ReadString('\n')
		//fmt.Printf("%d %s", i, line, )
		if ret == io.EOF{
			fmt.Printf("File has ended  \n")
			fmt.Printf("Full contain: %d \n", fullcontain)
			fmt.Printf("Overlap %d \n", overlap)
			return 
		}
		elf1se2se := parser(line);

		if isfulloverlap(
			elf1se2se[0],
			elf1se2se[1],
			elf1se2se[2],
			elf1se2se[3]){
			fullcontain++;
		}
		if isoverlap(
			elf1se2se[0],
			elf1se2se[1],
			elf1se2se[2],
			elf1se2se[3]){
			overlap++;
		}
	}
}





func part1_2_Sscanf(){
	fmt.Printf("This is day4\n");
	fname := "/home/garid/Documents/advent/AoC-2022/day4/input.txt"
	file, err := os.Open(fname)
	if err != nil{
		fmt.Printf("file %s is not found\n", fname);
		panic(2)
	}

	reader := bufio.NewReader(file)
	//read by each linea

	var elf1_s, elf1_e, elf2_s, elf2_e int;
	for i, fullcontain,overlap:=0,0,0 ;;i++{
		line, ret := reader.ReadString('\n')
		//fmt.Printf("%d %s", i, line, )
		//n, ret := fmt.Fscanln(reader, "%d-%d,%d-%d\n",
		//	&elf1_s, &elf1_e, &elf2_s, &elf2_e)
		n, ret := fmt.Sscanf(line, "%d-%d,%d-%d\n",
			&elf1_s, &elf1_e, &elf2_s, &elf2_e)
		fmt.Printf("%d, %d-%d,%d-%d, %d\t\n", i,
			elf1_s, elf1_e, elf2_s, elf2_e, n)
		
		if ret == io.EOF{
			fmt.Printf("File has ended  \n")
			fmt.Printf("Full contain: %d \n", fullcontain)
			fmt.Printf("Overlap %d \n", overlap)
			return 
		}
		//elf1se2se := parser(line);

		if isfulloverlap(
			elf1_s,
			elf1_e,
			elf2_s,
			elf2_e){
			fullcontain++;
		}
		if isoverlap(
			elf1_s,
			elf1_e,
			elf2_s,
			elf2_e){
			overlap++;
		}
	}
}





func main(){
	//part1_2();  My own implementation of extracting values from string
	part1_2_Sscanf();
}

