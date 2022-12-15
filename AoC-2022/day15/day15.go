package main

import (
	"bufio"
	"fmt"
	"io"
	//"log"
	"os"
)

type Coord struct {
	y, x int
}

type Crop struct {
	start, end int
	center, dis int
}

type Sensor struct {
	coord Coord;
	nearest_beacon Coord;
	distance2beacon int;
}

var (
	x_max = -9999999
    x_min =  9999999
	y_max = -9999999
	y_min =  9999999
	sensors = make([]Sensor, 0, 100)
	beacons_coords = make([]Coord, 0, 100)
)


func abs(a int) int{
	if a < 0 {
		return -a;		
	}
	return a;
}

func manhattan_distance(a, b Coord) int {
	return abs(a.x - b.x) + abs(a.y - b.y)
}

func line_parser(line string) {
	sensor_coord := Coord{}
	closest_beacon_coord := Coord{}

	_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
		&sensor_coord.x,
		&sensor_coord.y,
		&closest_beacon_coord.x,
		&closest_beacon_coord.y,
	)
	if err != nil {
		panic(1)
	}
	sensors = append(sensors, Sensor{
		coord:sensor_coord,
		nearest_beacon:closest_beacon_coord,
		distance2beacon: manhattan_distance(sensor_coord, closest_beacon_coord),
	})
	beacons_coords = append(beacons_coords, closest_beacon_coord)

	//fmt.Printf("%v\t%v\tdistance:%d\n",
	//	sensor_coord,
	//	closest_beacon_coord,
	//	manhattan_distance(sensor_coord, closest_beacon_coord))
}

func get_max_min_position(){
	for _, ithsensor := range sensors {
		if x_max < ithsensor.coord.x { x_max = ithsensor.coord.x}
		// if x_max < ithsensor.nearest_beacon.x { x_max = ithsensor.nearest_beacon.x}
		if x_min > ithsensor.coord.x { x_min = ithsensor.coord.x}
		// if x_min > ithsensor.nearest_beacon.x { x_min = ithsensor.nearest_beacon.x}

		if y_max < ithsensor.coord.y { y_max = ithsensor.coord.y}
		// if y_max < ithsensor.nearest_beacon.y { y_max = ithsensor.nearest_beacon.y}
		if y_min > ithsensor.coord.y { y_min = ithsensor.coord.y}
		//if y_min > ithsensor.nearest_beacon.y { y_min = ithsensor.nearest_beacon.y}
	}
}




func count_in_a_row(yhat int) (int, bool) {
	crops := make([]Crop, 0, len(sensors))
	lowest_pos_x  :=  9999999
	highest_pos_x := -9999999

	//size_of_row := x_max-x_min +1
	//row_posibilty := make([]int, size_of_row)
	//fmt.Printf("%v %v\n", row_posibilty, len(row_posibilty))

	//calculations
	for _, ithsensor := range sensors{
		xcenter := ithsensor.coord.x// - x_min                       // we need an offset
		left_over_distance := ithsensor.distance2beacon - abs(ithsensor.coord.y - yhat)
		//fmt.Println(xcenter, left_over_distance, ithsensor )

		if left_over_distance < 0 {
			;// nothing to do, i.e ith-sensor is too far far y-hat such that 
		} else {
			starting_x := xcenter - left_over_distance
			ending_x := xcenter + left_over_distance 
			crops = append(crops,
				Crop{start:starting_x, end:ending_x, center:xcenter, dis:left_over_distance})

			if starting_x < lowest_pos_x {
				lowest_pos_x = starting_x
			}
			if ending_x > highest_pos_x {
				highest_pos_x = ending_x
			}
		}
	}



	count := 0
	for xx:=lowest_pos_x; xx<=highest_pos_x; xx++ {
		tmp := false
		//var sensor_id_char byte
		// check the sensors
		var isitoverlap bool
		//for ii, ithsensor := range sensors{
		for _, ithsensor := range sensors{
			if ithsensor.coord.x == xx  && ithsensor.coord.y == yhat{
				isitoverlap = true
				//sensor_id_char = 'a' + byte(ii)
				break
			} else if ithsensor.nearest_beacon.x == xx  && ithsensor.nearest_beacon.y == yhat{
				isitoverlap = true
				//sensor_id_char = 'a' + byte(ii)
				break
			}
			// fmt.Printf("%T %d", ii, ii)
		}

		if !isitoverlap { // not overlapped
			//for ii, ithcrop := range crops {
			for _, ithcrop := range crops {
				if ithcrop.start <= xx && xx <= ithcrop.end {
					count++
					tmp = true
					//sensor_id_char = '0' + byte(ii)
					break
				}
			}
		}
		if !isitoverlap {
			if !tmp{
				if xx >=0 && xx <= x_max{
				//if xx >=0 && xx <= 20{
					//if xx >=0 && xx <= highest_pos_x{
					fmt.Println(yhat, xx, "<---------------")
					return xx, true
				}
			}
		}
		// if !tmp{
		// 	fmt.Printf("%#")
		// } else {
		// 	fmt.Printf(".")
		// }
		
	}

	//fmt.Println()
	return count, false
}

func could_this_have_beacon(checking Coord, _print bool) bool { // false means can't put
	if checking.x < 0 || checking.y <0 {
		return false
	}
	if checking.x > x_max || checking.y > y_max {
		return false
	}
	//var tmp true;
	for i, ithsenser := range sensors{
		if _print{
			fmt.Printf("%d\t%v\t%v\t%d\t%d\n",
				i,
				checking,
				ithsenser,
				manhattan_distance(checking, ithsenser.coord),
				ithsenser.distance2beacon,
			)
		} 
		if manhattan_distance(checking, ithsenser.coord) <= ithsenser.distance2beacon  {
			return false
		}
	}
	fmt.Printf("Beacon could be here: %v ==> %v\n",
		checking,
		checking.x * 4000000 + checking.y)
	return true
}

func go_perimeter(sen_coord Coord, d int) {
	L := Coord{y:sen_coord.y,     x:sen_coord.x-d-1}
	R := Coord{y:sen_coord.y,     x:sen_coord.x+d+1}
	U := Coord{y:sen_coord.y+d+1, x:sen_coord.x}
	D := Coord{y:sen_coord.y-d-1, x:sen_coord.x}

	for i:=0; i<d+2; i++ {
		could_this_have_beacon(L, false)  
		could_this_have_beacon(R, false) 
		could_this_have_beacon(U, false) 
		could_this_have_beacon(D, false) 

		L.x++
		L.y--

		R.x--
		R.y++

		U.x++
		U.y++

		D.x--
		D.y--
	} 
}

func main() {
	fmt.Printf("AoC2022 day 15\n")

	fname := "/home/garid/Documents/advent/AoC-2022/day15/input.txt" ; yhat := 2000000
	//fname := "/home/garid/Documents/advent/AoC-2022/day15/test"      ; yhat := 10

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
		fmt.Println(line)
		line_parser(line)
	}

	fmt.Printf("%v\n", sensors)
	fmt.Printf("%v\n", beacons_coords)
	get_max_min_position()
	fmt.Printf("max y %d\n", y_max)
	fmt.Printf("min y %d\n", y_min)
	fmt.Printf("max x %d\n", x_max)
	fmt.Printf("min x %d\n", x_min)

	count, _ := count_in_a_row(yhat)
	fmt.Println(count)


	for _,ith_sensor := range sensors{
		go_perimeter(ith_sensor.coord, ith_sensor.distance2beacon)
	}

	//could_this_have_beacon(Coord{x:14, y:11}, true)


}
