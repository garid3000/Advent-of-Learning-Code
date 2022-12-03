package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}


func winornot(oponentsign byte, mysign byte) int{
	o:= int(oponentsign) - int('A')
	m:= int(mysign)      - int('X') + 3
	//oponent
	// 0 rock
	// 1 paper
	// 2 scissor

	// me
	// 3 rock
	// 4 paper
	// 5 scissor


	switch (m-o) %3 {
	case 0:
		return 3 //draw
	case 1:
		return 6 //win
	case 2:
		return 0 //lose
	default:
		panic(1)
	}
	//3 -> 0 mean equal
	//4,1 -> 1 mean I won
	//5,2 -> 2 I lost
}

func calc_score(oponentsign byte, mysign byte) int{
	return (winornot(oponentsign, mysign) + int(mysign) - int('X') + 1)
}



func part1(){
	f, err := os.Open("/home/garid/Documents/advent/day2/input.txt")
    check(err)

	sumscore:=0
	r4 := bufio.NewReader(f)
	for {
		b4, ret := r4.ReadString('\n')
		if ret == io.EOF{
			fmt.Print("file has ended.")
			break
		}
		tmp:= calc_score(b4[0], b4[2])
		fmt.Printf("%c, %c \t = %v\n", b4[0], b4[2], tmp)
		sumscore +=tmp
	}
	fmt.Printf("sum %v\n", sumscore)
}


func p2_get_mysign(oponentsign byte, me byte) byte{
	o:= int(oponentsign) - int('A')
	//oponent
	// 0 rock
	// 1 paper
	// 2 scissor

	// me
	// x = lose
	// y = draw
	// z = win


	switch me {
	case 'X':  //lose
		return byte((o+2)%3 + int('X'))
	case 'Y':  //draw
		return  byte((o)%3 + int('X'))
	case 'Z':
		return byte((o+1)%3 + int('X'))
	default:
		panic(1)
	}


}

func part2(){
	f, err := os.Open("/home/garid/Documents/advent/day2/input.txt")
    check(err)

	sumscore:=0
	r4 := bufio.NewReader(f)
	for {
		b4, ret := r4.ReadString('\n')
		if ret == io.EOF{
			fmt.Print("file has ended.")
			break
		}

		fmt.Printf("%c, %c \t => \n", b4[0], b4[2])
		my_sing := p2_get_mysign(b4[0], b4[2])
		fmt.Printf("my supposed, %c\n", my_sing)
		tmp:= calc_score(b4[0], my_sing)
		fmt.Printf("%c, %c \t = %v\n", b4[0], my_sing, tmp)
		sumscore +=tmp
	}
	fmt.Printf("sum %v\n", sumscore)
}

func main(){
	part2()
}
