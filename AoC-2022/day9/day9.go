package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Coord struct {
	x, y int
}

var (
	//chain[0] = Coord{200,160}
	//tail_cur_coor = Coord{200,160}

	chain = [10]Coord{}

	maxheadx = 200 // i know this is stupid  // i should have extract this value dynamically
	minheadx = 200
	maxheady = 160
	minheady = 160

	coordinates_that_tail_visited [300][300][10]bool
)

func move_head(dir byte) {
	switch dir {
	case 'U':
		chain[0].y += 1
	case 'L':
		chain[0].x -= 1
	case 'R':
		chain[0].x += 1
	case 'D':
		chain[0].y -= 1
	default:
		fmt.Printf("Bad dir:=%c", dir)
		panic(1)
	}
}

//adf
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// test documentation
func is_tail_touched_head(tcoor, hcoor Coord) bool {
	if abs(tcoor.x-hcoor.x) > 1 {
		return false
	}
	if abs(tcoor.y-hcoor.y) > 1 {
		return false
	}
	return true
}

func follow_the_lead(lead_coor, follower_coor Coord) Coord {
	//if is_tail_touched_head() {
	//	return //wihtout moving
	//}
	ver_dis := lead_coor.y - follower_coor.y
	hor_dis := lead_coor.x - follower_coor.x
	newFollower_coor := Coord{follower_coor.x, follower_coor.y}

	if (abs(ver_dis) < 2) && (abs(hor_dis) < 2) {
		return newFollower_coor
		//no movement required
	} else if (abs(ver_dis) == 2) && (hor_dis == 0) {
		//only move vertically
		if ver_dis > 0 {
			newFollower_coor.y += 1
		}
		if ver_dis < 0 {
			newFollower_coor.y -= 1
		}
		return newFollower_coor
	} else if (abs(hor_dis) == 2) && (ver_dis == 0) {
		//only move horizontally
		if hor_dis > 0 {
			newFollower_coor.x += 1
		}
		if hor_dis < 0 {
			newFollower_coor.x -= 1
		}
		return newFollower_coor
	} else {
		//now the diagonale
		tmp_coor0 := Coord{follower_coor.x - 1, follower_coor.y - 1}
		tmp_coor1 := Coord{follower_coor.x - 1, follower_coor.y + 1}
		tmp_coor2 := Coord{follower_coor.x + 1, follower_coor.y - 1}
		tmp_coor3 := Coord{follower_coor.x + 1, follower_coor.y + 1}

		if is_tail_touched_head(lead_coor, tmp_coor0) {
			return tmp_coor0
		}
		if is_tail_touched_head(lead_coor, tmp_coor1) {
			return tmp_coor1
		}
		if is_tail_touched_head(lead_coor, tmp_coor2) {
			return tmp_coor2
		}
		if is_tail_touched_head(lead_coor, tmp_coor3) {
			return tmp_coor3
		}

		// if left here panic
		fmt.Printf("Couldn't find better diagonale move for tail\n")
		fmt.Printf("Positions: head %v\ttail %v",
			lead_coor, follower_coor,
		)
		panic(1)
	}

}

func log_maxmin_of_head() {
	if chain[0].x > maxheadx {
		maxheadx = chain[0].x
	}
	if chain[0].x < minheadx {
		minheadx = chain[0].x
	}
	if chain[0].y > maxheady {
		maxheady = chain[0].y
	}
	if chain[0].y < minheady {
		minheady = chain[0].y
	}
}

func move(dir byte, steps int) {
	for i := 0; i < steps; i++ {
		move_head(dir)            // move the head
		for j := 1; j < 10; j++ { //for the chain
			chain[j] = follow_the_lead(chain[j-1], chain[j])
			//tail_cur_coor = follow_the_lead(chain[0], tail_cur_coor)
			coordinates_that_tail_visited[chain[j].y][chain[j].x][j] = true
		}
		log_maxmin_of_head()
	}
}

func main() {
	//initial coordinates of chain
	for i := 0; i < 10; i++ {
		chain[i].x = 200
		chain[i].y = 160
	}

	fmt.Printf("AoC2022\n")
	fname := "/home/garid/Documents/advent/AoC-2022/day9/input.txt"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}

	var dir byte
	var steps int

	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		line, ret := reader.ReadString('\n')
		if ret == io.EOF {
			fmt.Printf("File has ended\n")
			break
		}
		fmt.Sscanf(line, "%c %d\n", &dir, &steps)
		fmt.Printf("%d\t%c-%d| %v %v |->", i, dir, steps, chain[0], chain[9])
		move(dir, steps)
		fmt.Printf("\t|%v %v\n", chain[0], chain[9])
	}
	//finished instructions

	fmt.Printf("x:%d to %d\n", minheadx, maxheadx)
	fmt.Printf("y:%d to %d\n", minheady, maxheady)

	fmt.Printf("current head %v\n", chain[0])
	fmt.Printf("current tail %v\n", chain[9])
	//count tail visited
	visit_of_1th := 0
	visit_of_9th := 0
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			if coordinates_that_tail_visited[i][j][1] {
				visit_of_1th++
			}
			if coordinates_that_tail_visited[i][j][9] {
				visit_of_9th++
			}
		}
	}

	fmt.Printf("1-th has visited %v positions\n", visit_of_1th) //this is part1
	fmt.Printf("9-th has visited %v positions\n", visit_of_9th) //this is part2

}
