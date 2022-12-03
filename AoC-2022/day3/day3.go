// Advent of Code com
// link:https://adventofcode.com/2022/day/3
package main


import (
	"bufio"
	"fmt" // for printing
	"io"
	"os" // for opening file
)

func char2num(c byte) int {
	if (c >= 'a') && (c <='z'){
		return int(c - 'a' + 1)
	} else if (c >= 'A') && (c <='Z'){
		return int(c - 'A' + 27)
	} else{
		fmt.Printf("Unexpected char: %c\n", c)
		panic(int(c))
	}
}


func num2char(n int) byte {
	if (n >= 1) && (n <=26) {
		return 'a' + byte(n - 1)
	} else if (n >= 27) && (n <= 52) {
		return 'A' + byte(n - 27)
	} else {
		fmt.Printf("Unexpected int: %d\n", n)
		panic(n)
	}
}



func argmin(arr [53] int) int{
	for i,L:=0, len(arr); i<L; i++ {
		if arr[i] == 1{
			fmt.Printf("%d %c\n", i, arr[i])
			return i
		}
	}

	fmt.Printf("Array index overflow, %v  I have no idea", arr)
	panic(1)
}


func charFinder(line string) int {
	Len := len(line) - 1
    HalfLen := Len / 2  //shoul be int right?, since Len is int
	var arr1 [53] int
	var arr2 [53] int
	var arr3 [53] int

	//for i:=0, j:=Len-1; i<Halfgen; i++, j-- { apperently this won't work
	// read here https://stackoverflow.com/a/38081920/14696853
    for i, j := 0, Len-1   ; i < HalfLen;    i, j = i+1, j-1 {
		fmt.Printf("%d\t%d\t%c %c\n", i,j, line[i], line[j])
		arr1[char2num(line[i])]=1;
		arr2[char2num(line[j])]=1;
	}

	//for i:=0; i<53; i++{
	//	arr3[i] = arr1[i] * arr2[i]
	//}
	arr3 = binaryAndBetween2array(arr1, arr2)



	fmt.Printf("  ")
	for i:=byte(0); i<byte(26); i++{
		fmt.Printf(" %c", i + 'a')
	}
	for i:=byte(0); i<byte(26); i++{
		fmt.Printf(" %c", i + 'A')
	}
	fmt.Printf("\n%v\n", arr1)
	fmt.Printf("%v\n", arr2)
	fmt.Printf("%v\n", arr3)
	fmt.Printf("%c\n", num2char(argmin(arr3)))

	//return num2char( argmin(arr3))
	return argmin(arr3)
}

func line2arr(line string) [53]int{
	var arr [53] int;

	for i, L:= 0, len(line)-1; i<L; i++{
		arr[char2num(line[i])] = 1
	}

	return arr;
}

func binaryAndBetween2array(arr1 [53]int, arr2 [53]int) [53]int{
	var arr3 [53]int;
	for i:=0; i<53; i++{
		arr3[i] = arr1[i] * arr2[i]
	}
	return arr3
}


func binaryAndBetween3array(arr1 [53]int, arr2 [53]int, arr3 [53] int) [53]int{
	var arr [53]int;
	for i:=0; i<53; i++{
		arr[i] = arr1[i] * arr2[i] * arr3[i]
	}
	return arr
}


func part1(){
	fmt.Printf("Hi, This is day 3\n")
	fname := "/home/garid/Documents/advent/day3/input.txt"
	file, err := os.Open(fname)
	if err!= nil{
		fmt.Printf("file: %s not found\n", fname)
		panic(err)
	}

	reader := bufio.NewReader(file) // prepare line by line reading
	sum := 0
	for {
		line, ret := reader.ReadString('\n') // this reads incl. \n char
		if ret == io.EOF{
			fmt.Printf("File has finished")
			fmt.Printf("Sum is %v\n", sum)

			return
		}
		//the_char :=

		fmt.Printf("%s------------\n", line)
		sum+=charFinder(line)
	}

}



func part2(){
	fmt.Printf("Hi, This is day 3 part 2\n\n")
	fname := "/home/garid/Documents/advent/day3/input.txt"
	file, err := os.Open(fname)
	if err!= nil{
		fmt.Printf("file: %s not found\n", fname)
		panic(err)
	}

	reader := bufio.NewReader(file) // prepare line by line reading
	sum := 0
	var line0 string;
	var line1 string;
	var line2 string;
	for i:=0;;i++{
		line, ret := reader.ReadString('\n') // this reads incl. \n char

		if ret == io.EOF{
			fmt.Printf("File has finished")
			fmt.Printf("Sum is %v\n", sum)
			return
		}

		fmt.Printf("%s", line)
		if i%3 == 0{
			line0 = line
		} else if i%3 == 1 {
			line1 = line

		} else {
			line2 = line
			//will work every third line
			num := argmin(
				binaryAndBetween3array(
					line2arr(line0),
					line2arr(line1),
					line2arr(line2)))
			fmt.Printf("--------%c\t%d\n-----------\n", byte(num), num)
			sum += num
		}
		

		//the_char :=

		//sum+=charFinder(line)
	}

}



func main(){
	//part1()
	part2()
}
