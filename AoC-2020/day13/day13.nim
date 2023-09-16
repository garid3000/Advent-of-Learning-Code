import std/strutils
import os
# var earliest_time: int

type
    Func_ax_b = object
        a: int
        b: int 

proc calc_func_ax_b(funcParam: Func_ax_b, x: int): int= 
    result = funcParam.a * x + funcParam.b

proc calc_series_of_funcs(funcs: seq[Func_ax_b], x: int): int =
    if funcs.len() == 1:
        return calc_func_ax_b(funcs[0], x)
    return calc_series_of_funcs(funcs[0..^2], calc_func_ax_b(funcs[^1], x))


proc find_best_bus(my_min: int, bus_schedule: seq[int]): tuple[busId:int, time:int] =
    for time in countup(my_min, my_min + bus_schedule[0]):
        stdout.write("\n", time)
        for each_bus in bus_schedule:
            if time mod each_bus == 0:
                stdout.write("\t", "D")
                result.time = time
                result.busId = each_bus
                return
            else:
                stdout.write("\t", ".")

#  proc find_part_2(bus_schedule: seq[int], bus_remainders: seq[int]): int =
#      var foundit: bool
#      var k0 = 0
#  
#      while true:
#          var time = bus_schedule[0] * k0
#          foundit = true
#          for i, bus_p in bus_schedule:
#              if ((time + bus_remainders[i]) mod bus_p) != 0:
#                  foundit = false
#                  break
#          echo time
#          if foundit: 
#              echo time
#              break
#          k0 += 1
#      return
#  
#  
#  proc find_part_2_max(bus_schedule: seq[int], bus_remainders: seq[int]): int =
#      # searches using maximum depature interval bus
#      var foundit: bool
#      var kmax = 0
#  
#      var maxbus = bus_schedule.max()
#      var rem_maxbus = bus_remainders[find(bus_schedule, maxbus)]
#      echo (maxbus, rem_maxbus) 
#      #return
#  
#      while true:
#          var time = maxbus * kmax
#          foundit = true
#          for i, bus_p in bus_schedule:
#              if ((time - rem_maxbus + bus_remainders[i]) mod bus_p) != 0:
#                  foundit = false
#                  break
#          echo time
#          if foundit: 
#              echo time
#              break
#          kmax += 1# 19
#      return
#  
#  proc test_2(busIds: seq[int], busRems: seq[int])=
#      for i, busid in busIds:
#          echo busid, "\t", busRems[i], "\t|", busid mod busIds[0], "\t", busRems[i] mod busIds[0]
#  
#      echo "\n\n"
#      for i, busid in busIds:
#          for k0 in countup(1, busIds[0]):
#              if (busid * k0 + busRems[i]) mod busIds[0] == 0:
#                  echo busid, "\t", busRems[i], "\t| k", i," = ",  k0

proc test_asfasdf(busIds: seq[int], busREms: seq[int])=
    var seq_funcs: seq[Func_ax_b]
    seq_funcs.add(Func_ax_b(a:busIds[^1], b: -busREms[^1]))

    for i in countup(0, busIds.len()-2):
        #echo "bus: ", busIds[i] , "\t", busREms[i]
        for j in countup(0, busIds[i]-1):
            var tmp = calc_series_of_funcs(seq_funcs, j)
            if (tmp mod busIds[i]) == ((999 * busIds[i] - busREms[i]) mod busIds[i]): 
                # 999 added because when busIds[i] > busREms[i], it produced negative value 
                # which resulted no function added at this revolution
                # so basically assumption is num busRem is 999 times higher than busID
                seq_funcs.add(Func_ax_b(a:busIds[i], b:j))
                break
        echo seq_funcs

    echo seq_funcs, "0->",calc_series_of_funcs(seq_funcs, 0), "this is the p2 answer <==="
    echo seq_funcs, "1->",calc_series_of_funcs(seq_funcs, 1), "this is the next answer"
    #echo seq_funcs, "2->",calc_series_of_funcs(seq_funcs, 2)
    return




proc main =
    var fpath = paramStr(1)
    var count = 0
    var earliest_min: int
    var bus_schedule_str: seq[string]
    var bus_schedule: seq[int]
    var bus_remainders: seq[int]

    for line in lines fpath:
        echo line
        if count == 0:
            earliest_min = parseInt(line)
            count += 1
        elif count == 1:
            bus_schedule_str = line.split(',')
            for i,each_element in bus_schedule_str:
                if each_element != "x":
                    bus_schedule.add(parseInt(each_element))
                    bus_remainders.add(i)
            break
    echo earliest_min, "\t", bus_schedule_str, "\t", bus_schedule, "\t", bus_remainders
    var bestBus = find_best_bus(earliest_min, bus_schedule)
    echo "\n", bestBus, "\t value =", (bestBus.time - earliest_min) * bestBus.busId, "\n\n\n"

    #var bestTime = find_part_2_max(bus_schedule, bus_remainders)
    #echo bestTime

    # test_2(bus_schedule, bus_remainders)
    test_asfasdf(bus_schedule, bus_remainders)

main()
