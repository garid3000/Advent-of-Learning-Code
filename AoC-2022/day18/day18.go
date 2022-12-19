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
	fmt.Printf("%v\n", surfaces)
}
