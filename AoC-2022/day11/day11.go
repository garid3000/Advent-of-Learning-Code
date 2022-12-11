package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type Monkey struct {
	ith         int
	items       []int64
	division    int
	if_true_to  int
	if_false_to int
	count       int64
	op_str      string
	op_int      int
}

var monkeys []Monkey

var modula int64 = 1

func input_ext_0_monkey_number(line string) int {
	var tmp int
	fmt.Sscanf(line, "Monkey %d:\n", &tmp)
	return tmp
}

func find_char(line string, ch byte) int {
	for i, L := 0, len(line); i < L; i++ {
		if line[i] == ch {
			return i
		}
	}
	fmt.Printf("could find : in %s", line)
	panic(1)
}

func input_ext_1_items(line string) []int64 {
	var vals []int64
	var tmp int

	for i, L := find_char(line, ':')+1, len(line); i < L; i++ {
		if line[i] == ' ' {
		} else if line[i] == ',' || line[i] == '\n'{
			vals = append(vals, int64(tmp))
			tmp = 0
		} else if (line[i] >= '0') && (line[i] <= '9') {
			tmp *= 10
			tmp += int(line[i] - '0')
		} else {
			fmt.Printf("Bad input in |%s| item-%d=%c\n", line, i, line[i])
		}
	}

	//vals = append(vals, tmp)
	return vals
}

func input_ext_2_op_str(line string) (string, int){
	var tmp string
	state_int_exists := false
	tmp_int := 0 
	for i, L := find_char(line, '=')+1, len(line); i < L; i++ {
		if line[i] == ' '  || line[i] == '\n' || line[i] == 'l' || line[i] == 'd'{
		} else if (line[i] >= '0') && (line[i] <= '9') {
			if state_int_exists == false {
				tmp += "i" 
			}
			state_int_exists = true
			tmp_int *= 10
			tmp_int += int(line[i] - '0')
		} else if (line[i] == 'o' || line[i] == '+' || line[i] == '-' || line[i] == '*' || line[i] == '/') {
			tmp += string(line[i])
		} 
	}
	return tmp, tmp_int
}

func input_ext_3_get_to(line string) int {
	var tmp int

	for i, L := find_char(line, 'y')+1, len(line); i < L; i++ {
		if line[i] == ' ' || line[i] == '\n' {
		} else if (line[i] >= '0') && (line[i] <= '9') {
			tmp *= 10
			tmp += int(line[i] - '0')
		} else {
			fmt.Printf("Bad input in |%s| item-%d=%c\n", line, i, line[i])
		}
	}
	return tmp
}


func ith_monkey_inspect(ith int){
	if ith >= len(monkeys){
		fmt.Printf("ith %d >= %d of monkeys", ith, len(monkeys))
		panic(1)
	}
	for i,L:= 0, len(monkeys[ith].items); i<L; i++ {
		monkeys[ith].count++
		//calculating new
		var v1, v2, new int64
		var newowner_ith int
		if monkeys[ith].op_str[0] == 'o' {
			v1 = monkeys[ith].items[i]
		} else if monkeys[ith].op_str[0] == 'i' {
			v1 = int64(monkeys[ith].op_int)
		} else {panic(3)}

		if monkeys[ith].op_str[2] == 'o' {
			v2 = monkeys[ith].items[i]
		} else if monkeys[ith].op_str[2] == 'i' {
			v2 = int64(monkeys[ith].op_int)
		} else {panic(3)}

		switch (monkeys[ith].op_str[1]) {
		case '+':
			new = v1 + v2
		case '-':
			new = v1 - 2
		case '*':
			new = v1 * v2
		case '/':
			new = v1 / v2
		default:
			panic(4)
		}
		new = new % modula

		//new = new / 3

		if new % int64(monkeys[ith].division) == 0 {
			newowner_ith = monkeys[ith].if_true_to
		} else {
			newowner_ith = monkeys[ith].if_false_to
		}
		/*
		fmt.Printf("Monkey #%d, %d->%d,remainder(%d),  sendto %d\n",
			ith,
			monkeys[ith].items[i],
			new,
			new % monkeys[ith].division,
			newowner_ith,
		)*/
		monkeys[newowner_ith].items = append(monkeys[newowner_ith].items, new)

	}
	monkeys[ith].items = [] int64{}
}

func print_monkeys(){
	//fmt.Printf("------------------------------------------------\n")
	for i,L := 0, len(monkeys); i<L;  i++{
		fmt.Printf("Monkey %d:\t%d\t%v \n", i, monkeys[i].count, monkeys[i].items)
	}
	fmt.Printf("-----------------------------------------------\n")
}

func round() {
	for i,L := 0, len(monkeys); i<L;  i++{
		ith_monkey_inspect(i)
		//fmt.Printf("%v\n", monkeys)
	}
}

func main() {
	fmt.Printf("AoC2022 day 10\n")

	//fname := "/home/garid/Documents/advent/AoC-2022/day11/test"
	fname := "/home/garid/Documents/advent/AoC-2022/day11/input.txt"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	reader := bufio.NewReader(file)
	var tmp Monkey
	for i := 0; ; i++ {
		line, ret := reader.ReadString('\n')
		if ret == io.EOF {
			fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			break
		}
		//line = line[:len(line)-1]

		fmt.Printf("%d\t%s\t\t\t", i, line)
		switch i % 7 {
		case 0:
			fmt.Printf("s")
			tmp.ith = input_ext_0_monkey_number(line)
		case 1:
			fmt.Printf("i")
			tmp.items = input_ext_1_items(line)
		case 2:
			fmt.Printf("o")
			tmp.op_str, tmp.op_int = input_ext_2_op_str(line)
		case 3:
			fmt.Printf("t")
			tmp.division = input_ext_3_get_to(line)
			modula *= int64(tmp.division)
		case 4:
			fmt.Printf("1")
			tmp.if_true_to = input_ext_3_get_to(line)
		case 5:
			fmt.Printf("0")
			tmp.if_false_to = input_ext_3_get_to(line)
			fmt.Printf("\n%v", tmp)
			monkeys = append(monkeys, tmp)
		default:

		}
	}
	//finished instructions



	fmt.Printf("%v\n", monkeys) // this is part1
	for i:=0; i<10000; i++ {
		if i%1000 == 0 {
			fmt.Printf("%d\n", i)
			print_monkeys()
		} else if i == 20 || i == 1{
			fmt.Printf("%d\n", i)
			print_monkeys()
		}

		round()
	}


	fmt.Printf("%v\n", 10000) 
	print_monkeys()
	//fmt.Printf("%v", monkeys) //this from this ans is the multiple of max 2 items
	var inspections []int
	for i,L:=0,len(monkeys); i<L; i++ {
		inspections = append(inspections, int(monkeys[i].count))
	}
	fmt.Printf("unsorted inspections: %v\n", inspections)
	sort.Ints(inspections)
	fmt.Printf("  sorted inspections: %v\n", inspections)
	fmt.Printf("monkey business %v\n", inspections[len(monkeys)-1] * inspections[len(monkeys)-2])
}
