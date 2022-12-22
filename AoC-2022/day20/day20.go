package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	init_array = make([]int, 0, 5000)
	index_arr  = make([]int, 0, 5000)
	index_arrf = make([]float32, 0, 5000)
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
		return 1
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
	// fmt.Printf("PANICING           \t\t\t%v\n", index_arr)
	fmt.Printf("Couldn't find      \t\t\t%v\n", index)
	fmt.Printf("PANICING           \t\t\t%v\n", index_arr)
	panic(2)
}

func fancy_printing(){
	fmt.Printf("fancy: ")
	for i, L:=0, len(init_array); i<L; i++{
		fmt.Printf("%d,\t", init_array[findIndex(i)])
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
				// if new_pos >= L-1 {
				// 	// new_pos = L-1 //remove 1 from all index
				// 	for ii:=0; ii<L; ii++{
				// 		if index_arr[ii] < index_arr[i] {
				// 			index_arr[ii]++;
				// 		}
				// 	}
				// 	index_arr[i] = 0;
				// } else {
				// 	// fmt.Printf("\tnew_pos: %d\n", new_pos)
					// this is new position of init_array[i]
					// the previoues index_arr[i] + 1 shoulbe go down by one
					index_new_pos := findIndex(new_pos)
					index_arr[index_new_pos] -=1;
					index_arr[index_new_pos] +=L;
					index_arr[index_new_pos] %= L;
					index_arr[i] = new_pos;
				// }
				

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

func copy_array_to_float() {
	for i,L:=0,len(init_array); i<L; i++{
		index_arrf[i] = float32(index_arr[i])
	}
}

func new_shuffle(isprinting bool)  {
	for i,L:=0,len(init_array); i<L; i++{
		copy_array_to_float()
		value     := init_array[i] //we need to shuffle this ammount
		signvalue := sign(value)
		value      = abs(value)
		//value = value % len(index_arr)
		newvalue := float32(index_arr[i]) +
			float32(signvalue) * (0.5) +
			float32(signvalue) * float32(value)
		//make sure new value in range of 0 -> L 

		// println(newvalue)
		for ;newvalue > float32(L); {
			newvalue -= float32(L);
		}
		for ;newvalue < 0; {
			newvalue += float32(L);
		}

		println(init_array[i], newvalue)
		// new value in in 0->L 

		fmt.Printf("\n%v %v->%v\n", init_array[i], index_arr[i], newvalue)
		if float32(index_arr[i]) < newvalue{
			for ii:=0; ii<L; ii++ {
				if index_arr[i] < index_arr[ii] && float32(index_arr[ii]) < newvalue {
					index_arr[ii]--;
				}
			}
			index_arr[i] = int(newvalue)
		} else {
			for ii:=0; ii<L; ii++ {
				if index_arr[i] > index_arr[ii] && float32(index_arr[ii]) > newvalue {
					index_arr[ii]++;
				}
			}
			index_arr[i] = int(newvalue) + 1
		}
		if isprinting{
			fmt.Printf("Move %d\t to %d-th |  ", init_array[i], index_arr[i])
			fancy_printing()
		}
	}

}



func main() {
	fmt.Printf("AoC2022 day 20\n")

	fname := "/home/garid/Documents/advent/AoC-2022/day20/input.txt" ;
	// fname := "/home/garid/Documents/advent/AoC-2022/day20/test"

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
		index_arr  = append(index_arr,    i)
		index_arrf = append(index_arrf,  float32(i))
		fmt.Printf("%d\t%d\t%d\n", i, whereIs(val), val,)
	}
	fmt.Printf("initial:|"); fancy_printing()
	new_shuffle(false)
	fmt.Printf("Last   :|"); fancy_printing()
	// shuffle()
	// fmt.Printf("%v %v %v\n", init_array, len(init_array), cap(init_array))
	// fmt.Printf("")
	// fancy_printing()
	index_arr_index_of_0 := index_arr[whereIs(0)]
	fmt.Printf("0 is at %v\n", index_arr_index_of_0 )
	index1000 :=  (index_arr_index_of_0 + 1000) % len(init_array)
	index2000 :=  (index_arr_index_of_0 + 2000) % len(init_array)
	index3000 :=  (index_arr_index_of_0 + 3000) % len(init_array)
	fmt.Printf("+1000th val is %v-> %v\n", index1000, init_array[findIndex(index1000)])
	fmt.Printf("+2000th val is %v-> %v\n", index2000, init_array[findIndex(index2000)])
	fmt.Printf("+3000th val is %v-> %v\n", index3000, init_array[findIndex(index3000)])
	fmt.Printf("len()=%d\n", len(init_array))

	fmt.Printf("value %v\n",
	 init_array[findIndex(index1000)] +
	 init_array[findIndex(index2000)] +
	 init_array[findIndex(index3000)],
	)

}
