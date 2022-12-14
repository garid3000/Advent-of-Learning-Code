package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strings"

	//"strings"
	"strconv"
	//"log"
)

type node struct {
	childs []node;
	size int;   //at lowest level it can be used
	val int;   //at lowest level it can be used
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func get_n_tabs(n int) string{
	tmp := ""
	for i:=0; i<n; i++{
		tmp += "  "
	}
	return tmp
}


// func compare_2trees(a node, b node, depth int) bool {
// 	aisnum := (len(a.childs)==0) && (a.size==1)
// 	bisnum := (len(b.childs)==0) && (b.size==1)
// 	depth_tabs := get_n_tabs(depth)
// 	if aisnum && bisnum {
// 		ret:= a.val <= b.val
// 		fmt.Printf("%s-Compare: %d vs %d \t%v\n", depth_tabs, a.val, b.val, ret)
// 		return ret
// 	} else if  (!aisnum) && (bisnum) {
// 		fmt.Printf("%s-Compare: %d vs %d \t%v\n", depth_tabs, a.val, b.val, ret)
// 		("- Mixed types; convert left to [9] and retry comparison")

// 	}
// }




func fancy_print(a_nod node, depth int){
	depth_str := get_n_tabs(depth)
	fmt.Printf("%s{\n", depth_str)

	if a_nod.size == 1 && len(a_nod.childs) == 0{
		fmt.Printf("%s %d\n", depth_str, a_nod.val)
	} else if a_nod.size >= 1{
		for i:=0; i<a_nod.size; i++{
			fancy_print(a_nod.childs[i], depth + 1)
		}
	}
	fmt.Printf("%s}\n", depth_str)
}

func fancy_print_1(a_nod node){
	fmt.Print("[")

	if a_nod.size == 1 && len(a_nod.childs) == 0{
		fmt.Printf("%d, ", a_nod.val)
	} else if a_nod.size >= 1{
		for i:=0; i<a_nod.size; i++{
			fancy_print_1(a_nod.childs[i])
		}
	}
	fmt.Printf("]")
}


func recursive(expression string, ith int) (node, int){
	this_node := node{}
	if ith == len(expression) - 1 {
		return this_node, ith
	} 

	current_number_str := ""
	current_child_node := node{}
	last_type := 'i'
	// 'i' initial
	// 'N' last one was a node
	// 'n' last one was a number
	             
	for i:=ith; i<len(expression); i++{
		if expression[i] == '[' {
			var endindex int
			current_child_node, endindex= recursive(expression, i + 1)
			i = endindex
			last_type = 'N'
			if (i == len(expression)-1) {
				this_node.childs = append(this_node.childs, current_child_node)
				this_node.size++
				//fmt.Printf("APpedn: %v\n", this_node)
				return this_node, len(expression)-1
			}

			//this_node.childs = append(this_node.childs, childnode)
			//this_node.size++
		} else if expression[i] == ']' {
			////fmt.Printf("%d,%c\n", expression[i], last_type)
			if last_type == 'N'{
				this_node.childs = append(this_node.childs, current_child_node)
				this_node.size++
				//fmt.Printf("APpedn: %v\n", this_node)
			}else if last_type == 'n'{
				v, e := strconv.Atoi(current_number_str)
				//fmt.Printf("\nconverted end %d\n", v)
				//fmt.Printf("%v\n", this_node)
				current_number_str = ""
				if e != nil{panic(123)}
				tmp_child_nod := node{}
				tmp_child_nod.size = 1
				tmp_child_nod.val = v
				this_node.childs = append(this_node.childs, tmp_child_nod)
				////fmt.Printf("%v\n", tmp_child_nod)
				////fmt.Printf("%v\n", this_node)
				
				this_node.size++
				//fmt.Printf("%v\n", this_node)
			} else if last_type == 'i' {
				this_node.childs = append(this_node.childs, node{})
				this_node.size++
			}

			//fmt.Printf("return %v\n", this_node)
			return this_node, i
		} else if (expression[i] >= '0') && (expression[i] <= '9') {
			current_number_str += string(expression[i])
			last_type = 'n'
		} else if expression[i] == ',' {
			////fmt.Printf("%d,%c\n", expression[i], last_type)
			if last_type == 'N'{
				this_node.childs = append(this_node.childs, current_child_node)
				this_node.size++
			}else if last_type == 'n'{
				v, e := strconv.Atoi(current_number_str)
				//fmt.Printf("\nconverted mid %d\n", v)
				if e != nil{panic(123)}
				tmp_child_nod := node{}
				tmp_child_nod.size = 1
				tmp_child_nod.val = v
				this_node.childs = append(this_node.childs, tmp_child_nod)
				////fmt.Printf("%v\n", tmp_child_nod)
				////fmt.Printf("%v\n", this_node)
				this_node.size++
				
			} else if last_type == 'i' {
				this_node.childs = append(this_node.childs, node{})
				this_node.size++
			}

			current_number_str = ""
			////fmt.Printf("%v %v\n", i, this_node)

		} else {panic(404);}

		//fmt.Printf("%d-%c-%v\n", i, expression[i], this_node)
	}
	panic(123123)
}




//func string2node(expresssion string) node {//	//remove the parenthesis
//	
//	v, err := strconv.Atoi(expresssion)
//	if err == nil{ // probably expresssion is integer in string
//		tmp := node{len:1, val:v}
//		return tmp
//	} // else expresssion is probably another list of expressions
//
//	// need to split expression with ,
//
//
//	return node{}
//}

func string_layer(expression string) (string, string) {						  
	layer := ""																	  
	commas := ""																  
	tmp := 0																	  
	for _, char := range expression{											  
		if char == '[' {														  
			tmp++																  
		} //else																  
																				  
		//if char == ','{commas += ","} else {commas += " "}					  
																				  
		layer += string(byte(tmp) + 'a')										  
		if char == ',' && tmp == 1 {commas += ","} else {commas += " "}			  
																				  
		if char == ']'{															  
			tmp--																  
		}																		  
	}																			  
	return layer, commas														  
}																			  


func split_list(alist string) ([]string, int){
	var x []string;
	_, commas := string_layer(alist)
	an_element := ""

	if !(alist[0]=='[' && alist[len(alist)-1]==']') {
		fmt.Printf("Not a list: %v", alist)
		panic(99)
	} else if (alist == "[]"){
		return []string{}, 0
	}
	for i,n:= 1, len(alist)-1; i<n; i++ {
		if commas[i] == ','{
			x = append(x, an_element)
			an_element = ""
		} else {
			an_element += string(alist[i])
		}
	}
	x = append(x, an_element)

	return x, len(x)
}

func compare2str(a, b string, tab_depth int) int {//a,b are value {num, or list}
	// i need give -1 not correct, +1 correct 0 for equal
	//fmt.Printf("%v- Compare %v vs %v\n", get_n_tabs(tab_depth), a, b)
	var is_a_num, is_b_num bool
	va, ea := strconv.Atoi(a)
	vb, eb := strconv.Atoi(b)
	if (ea == nil) {is_a_num = true}
	if (eb == nil) {is_b_num = true}

	if is_a_num && is_b_num { // both are numbers //case1 
		////fmt.Printf("\t\t%v<=%v %v\n", va, vb, va <= vb)
		if va == vb {
			return 0
		}else if va < vb {
			//fmt.Printf("%v- Left side is smaller, so inputs are in the right order\n", get_n_tabs(tab_depth+1))
			return 1
		}else{
			//fmt.Printf("%v- Right side is smaller, so inputs are not in the right order\n", get_n_tabs(tab_depth+1))
			return -1
		}
	} else if (!is_a_num) && (!is_b_num) { //both are list
		a_elements, a_element_size := split_list(a)
		b_elements, b_element_size := split_list(b)

		for i:=0; i<min(a_element_size, b_element_size); i++{
			tmpret:= compare2str(a_elements[i], b_elements[i], tab_depth + 1)
			if tmpret != 0 {
				// //fmt.Printf("returning %v\n", tmpret)
				return tmpret
			} // remaining lists are irrelavant apperantly
		}

		if (a_element_size > b_element_size) { // right must have longer values
			//fmt.Printf("%v- Right side ran out of items, so inputs are not in the right order\n",get_n_tabs(tab_depth+1))
			return -1 //
		} else if (a_element_size < b_element_size) {
			//fmt.Printf("%v- Left side ran out of items, so inputs are in the right order\n",	get_n_tabs(tab_depth+1))
			return 1
		} 
		return 0
	} else if (is_a_num) && (!is_b_num) { // a is num, b is list
		//make a: num -> list conversion
		newa := "[" + a + "]"
		//fmt.Printf("%s- Mixed types; convert left to %s and retry comparison\n", get_n_tabs(tab_depth+1), newa)
		tmpret := compare2str(newa, b, tab_depth + 1)
		if tmpret != 0 {return tmpret} // remaining lists are irrelavant apperantly
	} else if (!is_a_num) && (is_b_num) { // a is list, b is num
		//make b: num -> list conversion
		newb := "[" + b + "]"
		//fmt.Printf("%s- Mixed types; convert right to %s and retry comparison\n", get_n_tabs(tab_depth+1), newb)
		tmpret := compare2str(a, newb, tab_depth + 1)
		if tmpret != 0 {return tmpret} // remaining lists are irrelavant apperantly
	}
	return 0
}

var packets []string;

func print_all_packets () {
	for i:=0; i<len(packets); i++ {
		fmt.Printf("%d\t%s\n", i, packets[i])
	}
}

func get_index_in_packets(search string) int {
	for i:=0; i<len(packets); i++ {
		if packets[i] == search {
			return i
		}
	}
	panic(999)
}



func buble_sort(){
	// sorts packets variable
	size_of_packets := len(packets)
	for i:=0; i<size_of_packets-1; i++ {
		for j:=0; j<size_of_packets - i - 1; j++{
			//fmt.Printf("%d %d %d\n", size_of_packets, i, j)
			if compare2str(packets[j], packets[j+1], 0) < 0{
				packets[j], packets[j+1] = packets[j+1], packets[j]
			}
		}
	}
	
}
																				  

func main() {
	fmt.Printf("AoC2022 day 13\n")

	//fname := "/home/garid/Documents/advent/AoC-2022/day13/test"
	fname := "/home/garid/Documents/advent/AoC-2022/day13/input.txt"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	reader := bufio.NewReader(file)
	sum := 0 
	rightorderindex := []int{}
	for i := 0; ; i++ {
		line1, ret1 := reader.ReadString('\n')
		line2, ret2 := reader.ReadString('\n')
		line3, ret3 := reader.ReadString('\n')

		line1 = line1[:len(line1)-1]
		line2 = line2[:len(line2)-1]
		packets = append(packets, line1)
		packets = append(packets, line2)

		
		fmt.Printf("== Pair %v ==\n", i + 1)
		//fmt.Println(line1)
		//fmt.Println(line2)
		comparision := compare2str(line1, line2, 0)
		if (comparision >= 0) {
			sum += (i + 1)
			rightorderindex = append(rightorderindex, i+1)
		}
		fmt.Printf("comp: %v\n", comparision)


		if ret3 == io.EOF || ret2 == io.EOF || ret1 == io.EOF{
			fmt.Printf("\nFile has ended. Total %d lines.\n", i)
			fmt.Printf("\nFile has ended. Total %d lines.\n %v", i, line3)
			break
		}
	}


	fmt.Printf("sum: %v\n", sum) // this is part1


	//part2 additional2 paackets
	packets = append(packets, "[[2]]")
	packets = append(packets, "[[6]]")
	//fmt.Printf("packets %v %v\n", len(packets), packets )
	//print_all_packets()
	//fmt.Printf("right: %v\n", rightorderindex) // this is part1
	buble_sort()

	//print_all_packets()
	//fmt.Printf("packets %v\n", packets)

    indexof_2 := get_index_in_packets("[[2]]") + 1
    indexof_6 := get_index_in_packets("[[6]]") + 1

	fmt.Printf("%d\n", indexof_2)	
	fmt.Printf("%d\n", indexof_6)	
	fmt.Printf("product: %d\n", indexof_6 * indexof_2)	
}
