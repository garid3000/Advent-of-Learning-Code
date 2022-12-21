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
	// below part-of-struc is for part2	
	expression_based_on_humn string
	expression_based_on_humn_ready bool
}

var (
	monkeys      = make([]monkey, 0, 2100)
	orig_monkeys = make([]monkey, 0, 2100)
	global_argval1 int
	global_argval2 int
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


func get_value_of_monkey_p2(_id string) (int, bool) {
	independent_from_humn :=true
	ith := get_index_of_monkey(_id)
	if monkeys[ith].ready {
		if monkeys[ith].id != "humn" {
			return monkeys[ith].val, independent_from_humn
		} else {
			return monkeys[ith].val, false //this is the humand
		}
	} else {
		argval1, argval1_is_independent_from_humn := get_value_of_monkey_p2(monkeys[ith].arg1)
		argval2, argval2_is_independent_from_humn := get_value_of_monkey_p2(monkeys[ith].arg2)

		if (argval1_is_independent_from_humn == false) || (argval2_is_independent_from_humn ==false){
			//one of it is dependent from humn
			return 0, false
		}
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
		return thisvalue, true
	}
}


func root_check() bool {//part 2
	ith := get_index_of_monkey("root")
	global_argval1 = get_value_of_monkey(monkeys[ith].arg1)
	global_argval2 = get_value_of_monkey(monkeys[ith].arg2)
	return (global_argval1 == global_argval2)
}

/*
func testing() {
	// orig_monkeys = monkeys
	//for humn:=0;;humn++{
	for humn:=3200779769000;;humn++{

		// monkeys = orig_monkeys //clone from original

		for j, m:= range orig_monkeys{monkeys[j]=m}
		// fmt.Printf("$$%v\n", monkeys[get_index_of_monkey("humn")].val)
		monkeys[get_index_of_monkey("humn")].val = humn //set human value to
		monkeys[get_index_of_monkey("humn")].ready = true //set human value to
		// fmt.Printf("%v\n", monkeys)
		if root_check(){
			fmt.Printf("humn val %d\n", humn)
			break
		}
		if humn%1000 == 0{
			fmt.Printf("humn %d \t%d\t%d", humn, global_argval1, global_argval2)
		}

		// // monkeys = orig_monkeys //clone from original
		// for j, m:= range orig_monkeys{monkeys[j]=m}
		// // fmt.Printf("==%v\n", monkeys[get_index_of_monkey("humn")].val)
		// monkeys[get_index_of_monkey("humn")].val = -humn //set human value to
		// monkeys[get_index_of_monkey("humn")].ready = true //set human value to
		// // fmt.Printf("==%v\n", monkeys[get_index_of_monkey("humn")].val)
		// if root_check(){
		// 	fmt.Printf("humn val %d\n", -humn)
		// 	break
		// }
		if humn%1000 == 0{
			fmt.Printf("\t%d\t%d\n", global_argval1, global_argval2)
		}
	}
}*/

func get_value_or_expression_of_monkey(_id string) string {
	//this is after 1 time root
	ith := get_index_of_monkey(_id)
	if monkeys[ith].ready {
		if monkeys[ith].id != "humn" {
			return strconv.Itoa(monkeys[ith].val)
		} else {
			return "humn" //this is the humand
		}
	} else {
		if monkeys[ith].expression_based_on_humn_ready {
			return  monkeys[ith].expression_based_on_humn
		} else {
			argval1 := get_value_or_expression_of_monkey(monkeys[ith].arg1)
			argval2 := get_value_or_expression_of_monkey(monkeys[ith].arg2)

			monkeys[ith].expression_based_on_humn = fmt.Sprintf(
				"(%s %c %s)",
				argval1,
				monkeys[ith].op,
				argval2,
			)
			monkeys[ith].expression_based_on_humn_ready = true
			return monkeys[ith].expression_based_on_humn
		}
	}
	
}

func expression_creating(){
	f, _ := os.Create("out_stage1.txt")
	w := bufio.NewWriter(f)

	var line string
	for _, m := range monkeys {
		if m.ready {
			line = fmt.Sprintf("%s: %d\n", m.id, m.val)
		} else {
			line = fmt.Sprintf("%s: %s %c %s\n", m.id, m.arg1, m.op, m.arg2)
		}
	 	w.WriteString(line)
	}
	w.Flush()
}


func main() {
	// fmt.Printf("AoC2022 day21\n")

	fname := "/home/garid/Documents/advent/AoC-2022/day21/input.txt" ;
	// fname := "/home/garid/Documents/advent/AoC-2022/day21/test"

	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		line, ret1 := reader.ReadString('\n')
		if ret1 == io.EOF {
			// fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			break
		}
		line = line[:len(line)-1]
		add_monkey(line)
		// fmt.Printf("%d\t%s\n", i, line)
	}

	// println("saf")
	get_value_of_monkey_p2("root")
	// println("saf")
	orig_monkeys = make([]monkey, len(monkeys), 2100)
	for i, m:= range monkeys{orig_monkeys[i]=m}
	sum :=0
	for _,m := range monkeys {
		// fmt.Printf("%d\t%v\n",i, m)
		if m.ready {sum++}
	}
	// println(sum)
	// fmt.Printf("root: %v\n", get_value_of_monkey("root"))
	expression_creating()

	fmt.Printf(
		"%s-%s\n",
		get_value_or_expression_of_monkey(
			monkeys[get_index_of_monkey("root")].arg1,
		),
		get_value_or_expression_of_monkey(
			monkeys[get_index_of_monkey("root")].arg2,
		),
	)
}
