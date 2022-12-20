package main

import (
"bufio"
"fmt"
"io"
//"log"
"os"
"strings"
)

type Valve struct {
name string;
rate int;
tunnells []string;
//state bool;
}


var (
valves = make([]Valve, 0, 100);
current_room = "AA"
max_released = 0
max_released_steps = ""
)

func room_label2index(label string) int {
	for i,ith_room := range valves {
		if ith_room.name == label {
			return i
		}
	}
	fmt.Printf("Couldn't find room labeled <%s> panicing\n", label)
	panic(3)
}


func line_parser(line string){
	// example "Valve GL has flow rate=0; tunnels lead to valves AF, CQ"
	lineparts := strings.Split(line, ";")
	if len(lineparts) != 2 {panic(1)}
	//first part
	var valve_id string
	var valve_rate int;
	fmt.Sscanf(lineparts[0],
		"Valve %s has flow rate=%d",
		&valve_id, &valve_rate);

	//second parts
	this_valve_leads_to := strings.Split(lineparts[1], ", ")
	this_valve_leads_to[0] = strings.Split(this_valve_leads_to[0], " ")[5]
	//if len(lineparts) != 2 {panic(2)}

	//leadsto := strings.Split(lineparts[1], ", ")

	//fmt.Printf("%v\t%v\t%v\n", valve_id, valve_rate, lineparts)
	valves = append(valves,
		Valve{
			name:valve_id,
			rate:valve_rate,
			tunnells:this_valve_leads_to,
			//state:false,
		},
	)
}

func sum_flow_rate(paralel_universe_valve_states []bool) int {
	sum := 0
	for i,ith_valve := range valves {
		if paralel_universe_valve_states[i] {
			sum += ith_valve.rate
		}
	}
	return sum
}

func note_max_release_pressure(pu_released_press int, pu_steps string){
	if pu_released_press > max_released {
		max_released = pu_released_press
		max_released_steps = pu_steps
	}
}

func is_all_visited(pu_valve_states []bool) bool {
	tmp := true
	for _, ith_state := range pu_valve_states {
		tmp = ith_state  && tmp
	}
	return tmp
}


func explore(i_thmin int, pu_valve_states []bool, pu_cur_room string, pu_cur_released_press int, pu_steps_leads_here string)  {
	//calc flow rate from previous state
	pu_cur_released_press += sum_flow_rate(pu_valve_states)
	if i_thmin == 30 {
		note_max_release_pressure(pu_cur_released_press, pu_steps_leads_here)
		//fmt.Printf("%v\n\n", pu_steps_leads_here)
		return 
	}
	//test
	if is_all_visited(pu_valve_states) {
		return 
	}

	
	//not yet: -> go to the other rooms OR open this valve(if not openned yet)
	curRoomIndex := room_label2index(pu_cur_room)
	if pu_valve_states[curRoomIndex] == false { // not opened yet
		if valves[curRoomIndex].rate > 0 { // I don't bother to open valves if its rate is 0
			//clone this parellel universe
			//new_pu_valve_states := pu_valve_states //other variables are same, excepts the time
			new_pu_valve_states := make([]bool, len(valves)) //other variables are same, excepts the time
			for ith,v := range pu_valve_states {
				new_pu_valve_states[ith] = v
			}
			new_pu_valve_states[curRoomIndex] = true
			new_pu_steps_leads_here := fmt.Sprintf("%s%d open %s released %d\n",
				pu_steps_leads_here, i_thmin, pu_cur_room, pu_cur_released_press)
			
			//explore the new parallel universe
			explore(i_thmin+1, new_pu_valve_states, pu_cur_room, pu_cur_released_press, new_pu_steps_leads_here)
			//println("opening", pu_cur_room, new_pu_steps_leads_here)
		}
	}

	for _,ith_room_that_leadsfromhere := range valves[curRoomIndex].tunnells {
		//clone this universe
		//new_pu_valve_states := pu_valve_states //other variables are same, excepts the time
		new_pu_steps_leads_here := fmt.Sprintf("%s%d move %s->%s released %d\t%v\n",
			pu_steps_leads_here, i_thmin, pu_cur_room, ith_room_that_leadsfromhere, pu_cur_released_press, pu_valve_states)

		//explore the new parallel universe
		explore(i_thmin+1, pu_valve_states, ith_room_that_leadsfromhere, pu_cur_released_press, new_pu_steps_leads_here)
	}
}

func explore_nostep_log(i_thmin int, pu_valve_states []bool, pu_cur_room string, pu_cur_released_press int)  {
	//calc flow rate from previous state
	pu_cur_released_press += sum_flow_rate(pu_valve_states)
	if i_thmin == 30 {
		note_max_release_pressure(pu_cur_released_press, "none")
		//fmt.Printf("%v\n\n", pu_steps_leads_here)
		return 
	}
	//test
	if is_all_visited(pu_valve_states) {
		note_max_release_pressure(pu_cur_released_press, "none")
		return 
	}

	
	//not yet: -> go to the other rooms OR open this valve(if not openned yet)
	curRoomIndex := room_label2index(pu_cur_room)
	if pu_valve_states[curRoomIndex] == false { // not opened yet
		if valves[curRoomIndex].rate > 0 { // I don't bother to open valves if its rate is 0
			//clone this parellel universe
			//new_pu_valve_states := pu_valve_states //other variables are same, excepts the time
			new_pu_valve_states := make([]bool, len(valves)) //other variables are same, excepts the time
			for ith,v := range pu_valve_states {
				new_pu_valve_states[ith] = v
			}
			new_pu_valve_states[curRoomIndex] = true
			//new_pu_steps_leads_here := fmt.Sprintf("%s%d open %s released %d\n",
			//	pu_steps_leads_here, i_thmin, pu_cur_room, pu_cur_released_press)
			
			//explore the new parallel universe
			explore_nostep_log(i_thmin+1, new_pu_valve_states, pu_cur_room, pu_cur_released_press) //, new_pu_steps_leads_here)
			//println("opening", pu_cur_room, new_pu_steps_leads_here)
		}
	}

	for _,ith_room_that_leadsfromhere := range valves[curRoomIndex].tunnells {
		//clone this universe
		//new_pu_valve_states := pu_valve_states //other variables are same, excepts the time
		//new_pu_steps_leads_here := fmt.Sprintf("%s%d move %s->%s released %d\t%v\n",
		//	pu_steps_leads_here, i_thmin, pu_cur_room, ith_room_that_leadsfromhere, pu_cur_released_press, pu_valve_states)

		//explore the new parallel universe
		explore_nostep_log(i_thmin+1, pu_valve_states, ith_room_that_leadsfromhere, pu_cur_released_press) //, new_pu_steps_leads_here)
	}
}



func main() {
	fmt.Printf("AoC2022 day 16\n")

	//fname := "/home/garid/Documents/advent/AoC-2022/day16/input.txt" ;
	fname := "/home/garid/Documents/advent/AoC-2022/day16/test"    

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
		fmt.Print(line, "\n")
		line_parser(line)
	}
	fmt.Printf("%v\n", valves)
	fmt.Printf("%v\n", room_label2index("AA"))

	initial_valve_state := make([]bool, len(valves))
	// explore(
	// 	1,
	// 	initial_valve_state,
	// 	"AA",
	// 	0,
	// 	"beginning\n",
	// )


	explore_nostep_log(
		1,
		initial_valve_state,
		"AA",
		0,
	)

	fmt.Printf("Finished following are the steps\n")
	fmt.Printf("%v\n%v\n", max_released_steps, max_released)
}
