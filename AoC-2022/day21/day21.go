package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

type monkey struct {
	id       string
	ready    bool
	arg1     string
	arg2     string
	op       byte
	val      int
}

var (
	monkeys = make([]monkey, 0, 2100)
)

func add_monkey(line string){
	str_splits := strings.Split(line, " ")
	str_id := str_splits[0]
	// fmt.Printf("%v\t%v\n",  len(str_splits) , str_splits)

	if len(str_splits) == 2{
		value, ret := strconv.Atoi(str_splits[1])
		if ret != nil {
			panic(1)
		}
		monkeys = append(monkeys,
			monkey{
				id:     str_id[:4],
				ready:  true,
				val:    value,
			},
		)
	} else if len(str_splits) == 4 {
		monkeys = append(monkeys,
			monkey{
				id:     str_id[:4],
				ready:  false,
				arg1:   str_splits[1],
				arg2:   str_splits[3],
				op:     str_splits[2][0],
			},
		)

	} else {
		panic(2)
	}
}

func get_index_of_monkey(_id string) int{
	for i, m := range monkeys{
		if m.id == _id {
			return i
		}
	}
	fmt.Printf("id %s didn't found\n", _id)
	panic(3)
}

func get_value_of_monkey(_id string) int {
	ith := get_index_of_monkey(_id)
	if monkeys[ith].ready {
		return monkeys[ith].val
	} else {
		argval1 := get_value_of_monkey(monkeys[ith].arg1)
		argval2 := get_value_of_monkey(monkeys[ith].arg2)
		thisvalue := 0
		switch monkeys[ith].op{
			case '+':
			thisvalue = argval1 + argval2
			case '-':
			thisvalue = argval1 - argval2
			case '*':
			thisvalue = argval1 * argval2
			case '/':
			thisvalue = argval1 / argval2
			default:
			fmt.Printf("%v %c\n", monkeys[ith], monkeys[ith].op)
			panic(4)
		}
		monkeys[ith].val = thisvalue
		monkeys[ith].ready = true
		return thisvalue
	}
}




func main() {
	fmt.Printf("AoC2022 day21\n")

	fname := "/home/garid/Documents/advent/AoC-2022/day21/input.txt" ;
	//fname := "/home/garid/Documents/advent/AoC-2022/day21/test"

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
		add_monkey(line)
		// fmt.Printf("%d\t%s\n", i, line)
	}


	// for i,m := range monkeys {
	// 	fmt.Printf("%d\t%v\n",i, m)
	// }
	fmt.Printf("root: %v\n", get_value_of_monkey("root"))
}
