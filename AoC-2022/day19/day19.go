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

type blueprint struct {
	id int
	ore__bot_ore_cost int
	clay_bot_ore_cost int
	obs_bot_ore__cost int
	obs_bot_clay_cost int
	geod_bot_ore_cost int
	geod_bot_obs_cost int
}

type mystate struct{
	num_ore__bot int
	num_clay_bot int
	num_obsi_bot int
	num_geod_bot int

	produced_ore int
	produced_clay int
	produced_obsi int
	produced_goed int

	steps [24]int //0means nothing, 1means orebot, 2 means clay, 3 means obs, 4 means geo
}

func production(mycurrentstate mystate) mystate{
	tmp := mycurrentstate
	tmp.produced_ore += tmp.num_ore__bot
	tmp.produced_clay += tmp.num_clay_bot
	tmp.produced_obsi += tmp.num_obsi_bot
	tmp.produced_goed += tmp.num_geod_bot
	return tmp
}

func available_branches(mycurrentstate mystate, mycurrentbp blueprint) [5]int {
	//0means nothing, 1means orebot, 2 means clay, 3 means obs, 4 means geode
	tmp:= [5]int{1, 0, 0, 0, 0}
	if  mycurrentstate.produced_ore  >= mycurrentbp.ore__bot_ore_cost {tmp[1]=1}
	if  mycurrentstate.produced_ore  >= mycurrentbp.clay_bot_ore_cost {tmp[2]=1}
	if (mycurrentstate.produced_ore  >= mycurrentbp.obs_bot_ore__cost) &&
	   (mycurrentstate.produced_clay >= mycurrentbp.obs_bot_clay_cost) {tmp[3]=1}
	if (mycurrentstate.produced_ore  >= mycurrentbp.geod_bot_ore_cost) &&
	   (mycurrentstate.produced_obsi >= mycurrentbp.geod_bot_obs_cost) {tmp[4]=1}
	
	return tmp
}

var (
	blueprints = make([]blueprint, 0, 40)
	max_produced_geod = 0
	steps_that_produced_max_geod = [24]int{}
)

func parse_out_blueprint(line string) {
	var id_, cost_ore_bot, cost_clay_bot, cost_obsidian_bot_ore, cost_obsidian_bot_clay, cost_geode_bot_ore, cost_geode_bot_obsidian int;
	fmt.Sscanf(
		line,
		"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		 &id_, &cost_ore_bot, &cost_clay_bot, &cost_obsidian_bot_ore, &cost_obsidian_bot_clay, &cost_geode_bot_ore, &cost_geode_bot_obsidian,
	)
	blueprints = append(
		blueprints,
		blueprint{
			id:id_,
			ore__bot_ore_cost: cost_ore_bot, 
			clay_bot_ore_cost: cost_clay_bot,
			obs_bot_ore__cost: cost_obsidian_bot_ore,
			obs_bot_clay_cost: cost_obsidian_bot_clay,
			geod_bot_ore_cost: cost_geode_bot_ore,
			geod_bot_obs_cost: cost_geode_bot_obsidian,
		},
	)
}

func build_bots(botnum int, curstate mystate, bp blueprint) mystate {
	clonestate := curstate
	switch(botnum){
	case 0:
	case 1:
		clonestate.num_ore__bot++
		clonestate.produced_ore -= bp.ore__bot_ore_cost
	case 2:
		clonestate.num_clay_bot++
		clonestate.produced_ore -= bp.clay_bot_ore_cost
	case 3:
		clonestate.num_obsi_bot++
		clonestate.produced_ore  -= bp.obs_bot_ore__cost
		clonestate.produced_clay -= bp.obs_bot_clay_cost
	case 4:
		clonestate.num_geod_bot++
		clonestate.produced_ore  -= bp.geod_bot_ore_cost
		clonestate.produced_obsi -= bp.geod_bot_obs_cost
	default:
		fmt.Printf("Buildots fail %v", botnum)
		panic(1)
	}
	return clonestate
}

func explore(bp blueprint, curstate mystate, cur_time, max_time int) {
	// if curstate.produced_goed > 0{
	// 	fmt.Printf("%v %v\n", cur_time, curstate)
	// }

	if cur_time == max_time {
		//check the if this time it produce more
		if max_produced_geod < curstate.produced_goed {
			max_produced_geod = curstate.produced_goed
			steps_that_produced_max_geod = curstate.steps
		}
		return //curstate
	}
	//else check the all possible variables
	mycurent_available_branches := available_branches(curstate, bp)
	// for i, val := range mycurent_available_branches {
	// 	if val == 1 {
	// 		pu_state_after_build_bots := build_bots(i, curstate, bp)
	// 		pu_state_after_produciton := production(pu_state_after_build_bots)
	// 		explore(bp, pu_state_after_produciton, cur_time+1, max_time)
	// 	}
	// }

	for i:=4; i>=0; i-- {
		if mycurent_available_branches[i] == 1 {
			state_produced := production(curstate)
			pu_state_after_build_bots := build_bots(i, state_produced, bp)
			explore(bp, pu_state_after_build_bots, cur_time+1, max_time)
		}
		
	}
	return 
}

func max_production(bp blueprint, tmax int) int {
	// here we're going to calc bp
	var produced_geode int
	mycurstate := mystate{num_ore__bot: 1}
	explore(bp, mycurstate, 1, tmax)
	produced_geode = max_produced_geod
	
	// fmt.Printf("%v  %v\n", mycurstate, mycurent_available_branches)
	return produced_geode
}


func main() {
	fmt.Printf("AoC2022 day 19\n")

	//fname := "/home/garid/Documents/advent/AoC-2022/day19/input.txt" ;
	fname := "/home/garid/Documents/advent/AoC-2022/day19/test"

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
		fmt.Printf("%d\t%v\n", i, line)
		parse_out_blueprint(line)
	}


	fmt.Printf("%v\n", blueprints)
	fmt.Println("Now for the finding out each max amount obs production")
	maxes := make([]int, len(blueprints), len(blueprints))
	for i:=0; i<len(blueprints); i++ {
		maxes[i] = max_production(blueprints[i], 24)
		fmt.Printf("Max Production: %v\n", maxes[i])
		break
	}
	
}
