package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "strconv"
	// "time"
	//"strings"
)

type droplet struct{
	x,y,z int
}

var(
	alldroplets = make([]droplet, 0, 3000)
	surfaces = 0
	correct_surfaces = 0

	max_x=0
	min_x=0
	max_y=0
	min_y=0
	max_z=0
	min_z=0

	grid      = [40][40][40]int{}
	explored = [40][40][40]int{}
	count_explored = 0
)

func abs(x int)int{
	if x < 0 {return -x}
	return x
}

func appendDroplet(newDroplet droplet){
	tmp := 6
	for _, droplet := range alldroplets{
		if abs(droplet.x - newDroplet.x) == 1 && droplet.y == newDroplet.y && droplet.z == newDroplet.z {
			tmp-=2;
		} else if abs(droplet.y - newDroplet.y) == 1 && droplet.x == newDroplet.x && droplet.z == newDroplet.z {
			tmp-=2;
		} else if abs(droplet.z - newDroplet.z) == 1 && droplet.x == newDroplet.x && droplet.y == newDroplet.y {
			tmp-=2;
		}
	}
	surfaces += tmp
	alldroplets = append(alldroplets, newDroplet)

	if newDroplet.x > max_x {max_x = newDroplet.x}
	if newDroplet.x < min_x {min_x = newDroplet.x}
	if newDroplet.y > max_y {max_y = newDroplet.y}
	if newDroplet.y < min_y {min_y = newDroplet.y}
	if newDroplet.z > max_z {max_z = newDroplet.z}
	if newDroplet.z < min_z {min_z = newDroplet.z}

	grid[newDroplet.x+10][newDroplet.y+10][newDroplet.z+10] = 1
}

func explore(x,y,z int){
	if grid[x][y][z] == 1{
		return 
	} else if explored[x][y][z] == 1{
		return // already explored
	} else {
		explored[x][y][z] = 1
		count_explored++;
		// set current cursor as explored and explore neighbors
		//explore 8 neighbors
		if x > 0  {explore(x-1, y, z)}
		if x < 39 {explore(x+1, y, z)}

		if y > 0  {explore(x, y-1, z)}
		if y < 39 {explore(x, y+1, z)}
		
		if z > 0  {explore(x, y, z-1)}
		if z < 39 {explore(x, y, z+1)}
	}

}

func isThis_xyz_adroplet(x, y, z int) bool {
	for _, droplet := range alldroplets{
		if droplet.x == x && droplet.y ==y && droplet.z == z{
			return true
		}
	}
	return false
}

func count_outer_surface(){
	var gx, gy, gz int
	for _, droplet := range alldroplets{
		gx, gy, gz =  droplet.x+10, droplet.y+10, droplet.z+10
		if grid[gx+1][gy][gz]==0 && explored[gx+1][gy][gz]==1 {correct_surfaces++}
		if grid[gx-1][gy][gz]==0 && explored[gx-1][gy][gz]==1 {correct_surfaces++}

		if grid[gx][gy+1][gz]==0 && explored[gx][gy+1][gz]==1 {correct_surfaces++}
		if grid[gx][gy-1][gz]==0 && explored[gx][gy-1][gz]==1 {correct_surfaces++}

		if grid[gx][gy][gz+1]==0 && explored[gx][gy][gz+1]==1 {correct_surfaces++}
		if grid[gx][gy][gz-1]==0 && explored[gx][gy][gz-1]==1 {correct_surfaces++}
	}

	// for x:=0; x<40; x++ {
	// 	for y:=0; y<40; y++ {
	// 		for z:=0; z<40; z++ {
	// 		}
	// 	}
	// }
}


func main() {
	fmt.Printf("AoC2022 day 18\n")

	fname := "/home/garid/Documents/advent/AoC-2022/day18/input.txt" ;
	// fname := "/home/garid/Documents/advent/AoC-2022/day18/test"
	// fname := "/home/garid/Documents/advent/AoC-2022/day18/test1"

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

		//line = line[:len(line)-1]
		var xx,yy,zz int
		fmt.Sscanf(line, "%d,%d,%d\n", &xx, &yy, &zz)
		newDroplet := droplet{x:xx, y:yy, z:zz}
		appendDroplet(newDroplet)
	}
	fmt.Printf("%v\n", alldroplets)
	fmt.Printf("Surface:%v\n", surfaces)
	fmt.Println("min, max:")
	fmt.Println(min_x,max_x)
	fmt.Println(min_y,max_y)
	fmt.Println(min_z,max_z)

	explore(0, 0, 0)
	
	fmt.Printf("explored: %v of %v\n", count_explored, 40*40*40)
	count_outer_surface()
	fmt.Printf("outer surface %v\n", correct_surfaces) 
}
