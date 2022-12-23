package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Coord struct {
	y, x int
}

type Elf struct {
	pos Coord;
	new Coord;
}

var (
	elves = make([]Elf, 0, 3000)
	num_elves = 0
	round = 0
	sequence = [4]byte{'n', 's', 'w', 'e'}
	dy_seq = [4]int{-1, 1,  0, 0}
	dx_seq = [4]int{0,  0, -1, 1}
)

func abs(x int) int {
	if x > 0 {
		return x 
	}
	return -x
}

func empty_spaces() int{
	M_x, M_y := -9999, -9999 // refers to max x max y
	m_x, m_y :=  9999,  9999 // refers to min x min y

	for i:=0; i<num_elves; i++ {
		if elves[i].pos.x > M_x {M_x = elves[i].pos.x}
		if elves[i].pos.x < m_x {m_x = elves[i].pos.x}
		if elves[i].pos.y > M_y {M_y = elves[i].pos.y}
		if elves[i].pos.y < m_y {m_y = elves[i].pos.y}
	}
	return (M_x - m_x + 1) * (M_y - m_y +1) - num_elves
}

//func is_ith_jth_elves_neighbors(ith, jth int) bool {
//	if abs(elves[ith].pos.x-elves[jth].pos.x)+
//		abs(elves[ith].pos.y-elves[jth].pos.y) == 1 {
//		return true
//	}
//	return false
//}

func is_an_elf_on(y, x int) bool {
	for i:=0; i<num_elves; i++ {
		if (elves[i].pos.y==y) && (elves[i].pos.x==x){
			return true
		}
	}
	return false
}


func can_go(dir byte, _nw, _nn, _ne, _sw, _sn, _se, _ee, _ww bool) bool {
	if !(_nw || _nn || _ne) && dir == 'n' {return true}
	if !(_sw || _sn || _se) && dir == 's' {return true}
	if !(_sw || _nw || _ww) && dir == 'w' {return true}
	if !(_se || _ne || _ee) && dir == 'e' {return true}
	return false
}

func consider_new_pos(ith int) {	
	_nw := is_an_elf_on(elves[ith].pos.y-1, elves[ith].pos.x-1)
	_nn := is_an_elf_on(elves[ith].pos.y-1, elves[ith].pos.x  )
	_ne := is_an_elf_on(elves[ith].pos.y-1, elves[ith].pos.x+1)

	_sw := is_an_elf_on(elves[ith].pos.y+1, elves[ith].pos.x-1)
	_sn := is_an_elf_on(elves[ith].pos.y+1, elves[ith].pos.x  )
	_se := is_an_elf_on(elves[ith].pos.y+1, elves[ith].pos.x+1)

	_ee := is_an_elf_on(elves[ith].pos.y  , elves[ith].pos.x+1)
	_ww := is_an_elf_on(elves[ith].pos.y  , elves[ith].pos.x-1)

	num_neighbors := _nw || _nn || _ne || _sw || _sn || _se || _ee || _ww

	if !num_neighbors {
		elves[ith].new = elves[ith].pos
		return
	}
	
	// propose
	for ii:=0; ii<4; ii++{
		dirchar := sequence[(round + ii)%4]
		if can_go(dirchar, _nw, _nn, _ne, _sw, _sn, _se, _ee, _ww){
			elves[ith].new.x =  elves[ith].pos.x + dx_seq[(round + ii)%4]
			elves[ith].new.y =  elves[ith].pos.y + dy_seq[(round + ii)%4]
			// fmt.Printf("%d+1-th round, %d-th elf  considers %c direction: %v->%v\n",
				// round,
				// ith,
				// dirchar,
				// elves[ith].pos,
				// elves[ith].new,
			// )
			return
		}
	}
	elves[ith].new = elves[ith].pos
}

func move_2_new(ith int) int {
	count := 0
	// count the number of other elves who consider same new position with i-th elf
	for i:=0; i<num_elves; i++{
		if i!=ith{
			if elves[i].new == elves[ith].new {
				count++
			}
		}
	}

	if count > 0{
		return 1// nothing if others considered same new postion
		//1 means not moved
	}
	if elves[ith].pos == elves[ith].new {
		return 1
		//1 means not moved
	}

	elves[ith].pos = elves[ith].new //move
	return 0
	//0 means moved
}


func execute_1_round() bool {
	// 1st half of step
	for i:=0; i<num_elves; i++{
		consider_new_pos(i)
	}
	//2nd half of step
	sum := 0
	for i:=0; i<num_elves; i++{
		sum += move_2_new(i)
	}
	round++;

	return sum == num_elves
}


func printing(shouldIprint bool) {
	if shouldIprint {
		fmt.Printf("\033[2J\033[0;0H---round:%d----", round); // clear screen
		for i:=0; i<num_elves; i++{
			fmt.Printf("\033[%d;%dH#", elves[i].pos.y+24, elves[i].pos.x+40)
		}

		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	fmt.Printf("AoC2022 day23\n")

	fname := "/home/garid/Documents/advent/AoC-2022/day23/input.txt" ;
	// fname := "/home/garid/Documents/advent/AoC-2022/day23/test"
	// fname := "/home/garid/Documents/advent/AoC-2022/day23/test_small"

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
		fmt.Printf("%d\t%v\n", i, line[:len(line)-1])
		for j,L:=0,len(line)-1; j<L; j++ {
			if line[j] == '#'{
				elves = append(elves, Elf{pos:Coord{y:i,x:j}})
			}
		}

	}
	num_elves = len(elves)
	// fmt.Printf("%v %v\n", elves, len(elves))

	printing(true)
	for i:=0;; i++ {
		if execute_1_round() {
			break
		}
		printing(true)
		fmt.Printf("\n%d\n", i)
	}
	printing(true)

	fmt.Printf("%v\n", empty_spaces())
}
