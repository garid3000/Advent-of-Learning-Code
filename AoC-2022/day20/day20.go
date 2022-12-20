package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	init_array = make([]int, 0, 5000)
	index_arr = make([]int, 0, 5000)
)

func whereIs(val int) (index int) {
	for i, arrVal := range init_array{
		if val == arrVal {
			index = i
			return
		}
	}
	index = -1
	return 
}

func abs(x int) int{
	if x < 0 {return -x}
	return x
}

func sign(x int) int{
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func findIndex(index int) (indexofindex int){
	// fmt.Printf("\ttrying to find %d from %v\n", index, index_arr)
	for i,indexval := range index_arr{
		if indexval == index  {
			indexofindex = i
			return 
		}
	}
	fmt.Printf("PANICING           \t\t\t%v\n", index_arr)
	panic(1)
}

func fancy_printing(){
	fmt.Printf("This Is fancy: ")
	for i, L:=0, len(init_array); i<L; i++{
		fmt.Printf("%d, ", init_array[findIndex(i)])
	}
	fmt.Printf("\t\t\t%v\n", index_arr)

}

func shuffle(){
	for i,L:=0,len(index_arr); i<L; i++ {
		//shuffle each one by one
		value := init_array[i] //we need to shuffle this ammount
		signvalue := sign(value)
		value = abs(value)
		value = value % len(index_arr)
		// value = signvalue * value // just trimming the over lapper
		// current_pos := index_arr[i]
		// new_pos:= current_pos + value // new position
		fmt.Printf("shuffling %d %v num=%v value=%d sing=%v\n", i, index_arr[i], init_array[i], value, signvalue)
		// fancy_printing()

		if signvalue == 1 {
			for ;value > 0; { //while value isn't 0
				// fmt.Printf("> %d %d %v\n", i, value,  index_arr)
				new_pos := index_arr[i] + 1
				if new_pos >= L-1 {
					// new_pos = L-1 //remove 1 from all index
					for ii:=0; ii<L; ii++{
						if index_arr[ii] < index_arr[i] {
							index_arr[ii]++;
						}
					}
					index_arr[i] = 0;
				} else {
					// fmt.Printf("\tnew_pos: %d\n", new_pos)
					// this is new position of init_array[i]
					// the previoues index_arr[i] + 1 shoulbe go down by one
					index_new_pos := findIndex(new_pos)
					index_arr[index_new_pos] -=1;
					index_arr[index_new_pos] +=L;
					index_arr[index_new_pos] %= L;
					index_arr[i] = new_pos;
				}
				

				value--;
				// fancy_printing();
			}
			// fmt.Printf("-=-%v   %v\n", value,  index_arr)

		} else if signvalue == -1 {
			for ;value > 0; { //while value isn't 0
				new_pos := index_arr[i] - 1
				if new_pos <= 0 {
					// new_pos = L //remove 1 from all index
					for ii:=0; ii<L; ii++{
						if index_arr[ii] > index_arr[i] {
							index_arr[ii]--;
						}
					}
					index_arr[i] = L-1;
				} else {
					// this is new position of init_array[i]
					// the previoues index_arr[i] + 1 shoulbe go down by one
					index_new_pos := findIndex(new_pos)
					index_arr[index_new_pos] +=1;
					index_arr[index_new_pos] %= L;

					index_arr[i] = new_pos;
				}
				value--;
				// fancy_printing(); 
			}
			// fmt.Printf("-=-%v   %v\n", value,  index_arr)

		}
		fmt.Println()
		//fmt.Printf("F");fancy_printing()
	}
}

func main() {
	fmt.Printf("AoC2022 day 20\n")

	fname := "/home/garid/Documents/advent/AoC-2022/day20/input.txt" ;
	//fname := "/home/garid/Documents/advent/AoC-2022/day20/test"

	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		var val int;
		line, ret1 := reader.ReadString('\n')
		if ret1 == io.EOF {
			fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			break
		}
		fmt.Sscanf(line, "%d\n", &val)
		init_array = append(init_array, val)
		index_arr = append(index_arr, i)
		fmt.Printf("%d\t%d\t%d\n", i, whereIs(val), val,)
		// line = line[:len(line)-1]

		// fmt.Printf("%d\t%v\n", i, line)
		
		// parse_out_blueprint(line)
	}


	fmt.Printf("%v %v %v\n", init_array, len(init_array), cap(init_array))
	fmt.Printf("%v %v %v\n", index_arr, len(index_arr), cap(index_arr))
	fancy_printing()
	shuffle()

	fmt.Printf("%v %v %v\n", init_array, len(init_array), cap(init_array))
	fmt.Printf("")

	fancy_printing()

	index_arr_index_of_0 := whereIs(0)
	fmt.Printf("0 is at %v\n", index_arr_index_of_0 )
	fmt.Printf("+1000th val is  %v\n", init_array[index_arr_index_of_0 + 1000])
	fmt.Printf("+2000th val is  %v\n", init_array[index_arr_index_of_0 + 2000])
	fmt.Printf("+3000th val is  %v\n", init_array[index_arr_index_of_0 + 3000])

}
