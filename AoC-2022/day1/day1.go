package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func my_atoi(t string) int{
	tmp:=0
	for i:=0; i<len(t) && int('0')<= int(t[i]) && int(t[i]) <=int('9'); i++{
		tmp *= 10;
		tmp += int(t[i]) - int('0')
		//fmt.Printf("\t%v %v %c\n", i, tmp, t[i])
	}
	return  tmp
}


func find_most1(){
	fmt.Printf("This is advent of code day1 puzzle\n")
	f, err := os.Open("/home/garid/Documents/advent/day1/input.txt")
    check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.ReadString('\n')
	check(err)
	fmt.Printf("5 bytes %s\n", string(b4))

	one_elf_sum :=0
	max_elf_sum :=0

	for i:=0;;i++{
		b4, err = r4.ReadString('\n')
		if err == io.EOF{
			fmt.Printf("End of file")
			break
		}
		check(err)
		if b4[0] != '\n'{
			x:= my_atoi(b4)
			fmt.Printf("%d %d %v %s", i,  len(b4), x, string(b4))
			one_elf_sum += x
		} else {
			fmt.Printf("%d %d ------------- %s", i,  len(b4), string(b4))
			if one_elf_sum > max_elf_sum {
				max_elf_sum = one_elf_sum
			}
			one_elf_sum = 0
		}

	}
	fmt.Printf("max one elf %v\n", max_elf_sum)

}



func find_most3(){
	fmt.Printf("This is advent of code day1 puzzle\n")
	f, err := os.Open("/home/garid/Documents/advent/day1/input.txt")
    check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.ReadString('\n')
	check(err)
	fmt.Printf("5 bytes %s\n", string(b4))

	one_elf_sum :=0
	var most3 [3] int

	for i:=0;;i++{
		b4, err = r4.ReadString('\n')
		if err == io.EOF{
			fmt.Printf("End of file")
			break
		}
		check(err)
		if b4[0] != '\n'{
			x:= my_atoi(b4)
			fmt.Printf("%d %d %v %s", i,  len(b4), x, string(b4))
			one_elf_sum += x
		} else {
			fmt.Printf("%d %d ------------- %s", i,  len(b4), string(b4))
			if one_elf_sum > most3[0] {
				//this means this elf has at least more than the 3rd elf
				if one_elf_sum > most3[1]{
					//this means this elf has at least more than 2nd elf
					if one_elf_sum > most3[2]{

						most3[0] = most3[1]
						most3[1] = most3[2]
						most3[2] = one_elf_sum
					} else{
						most3[0] = most3[1]
						most3[1] = one_elf_sum
					}
				} else {
					most3[0] = one_elf_sum
					// this means this elf doesn't have more than 2nd elf
				}
			}
			one_elf_sum = 0
		}

	}
	fmt.Printf("sum of max 3 elf %v\n",
		most3[0] + most3[1] + most3[2])

}


func main(){
	find_most1() //for part 1
	find_most3() //for part 2
}
