# import std/strutils
import std/sequtils
import sugar

type
  Cycle = ref object
    elements: seq[int]
    cur: int
    desVal: int
    des: int
    picked3: seq[int]
    maxCupVal: int

proc findDes(cyc: Cycle, posVal: int): int = 
  var possibleVal = posVal
  if possibleVal == 0:
    possibleVal = cyc.maxCupVal

  if possibleVal in cyc.picked3:
    return findDes(cyc, possibleVal - 1)
  else: 
    return possibleVal


proc next_element_on_this_val(cyc: Cycle, curVal: int): int = 
  var tmpcur = cyc.elements.find(curVal)
  return (tmpcur + 1) mod cyc.elements.len()

proc pick3cups(cyc: Cycle): seq[int] = 
  var curVal = cyc.elements[cyc.cur]
  for i in 1..3:
    var next_val_index = cyc.next_element_on_this_val(curVal)
    result.add( cyc.elements[next_val_index] )
    cyc.elements.delete(next_val_index)

proc rotate_until_this_val_at_this_index(cyc: Cycle, thisVal, thisIndex: int) =
  while true:
    if cyc.elements[thisIndex] == thisVal:
      break
    var firstVal = cyc.elements[0]
    cyc.elements.delete(0)
    cyc.elements.add(firstVal)

proc put3_after_des(cyc: Cycle, older_val_curcup, older_cur_index: int)= 
  cyc.des = cyc.elements.find(cyc.desVal)

  cyc.elements.insert( cyc.picked3[0], cyc.des + 1 )
  cyc.elements.insert( cyc.picked3[1], cyc.des + 2 )
  cyc.elements.insert( cyc.picked3[2], cyc.des + 3 )

  cyc.rotate_until_this_val_at_this_index(older_val_curcup, older_cur_index)
  return

proc print_cups_with_selectedness(cyc: Cycle)=
  stdout.write("cups: ")
  for i in 0..<cyc.elements.len():
    if i == cyc.cur:
      stdout.write("(", cyc.elements[i], ")")
    else:
      stdout.write(" ", cyc.elements[i], " ")
  stdout.write("\n")

proc print_final_lables_except_1(cyc: Cycle)=
  var index_with_lbl_1 = cyc.elements.find(1)
  for i in 1..<cyc.elements.len():
    var ii = (index_with_lbl_1 + i) mod cyc.elements.len()
    stdout.write(cyc.elements[ii])
  stdout.write("\n")

proc move(cyc: Cycle)= 
  var val_cur_cup = cyc.elements[cyc.cur] 
  var cyc_cur = cyc.cur
  cyc.print_cups_with_selectedness()

  cyc.picked3 = cyc.pick3cups()
  echo "pick up: ", cyc.picked3, " ", cyc.elements
  
  cyc.desVal  = cyc.findDes(val_cur_cup - 1)
  echo "destinationVal: ", cyc.desVal

  cyc.put3_after_des(val_cur_cup, cyc_cur)
  echo "after putting destinatino"
  cyc.cur = (cyc.cur + 1) mod cyc.elements.len()

proc main() = 
  var cyc = Cycle()
  var inStr = stdin.readLine()
  cyc.elements = inStr.map( (x: char) -> int => int(x) - int('0') )
  cyc.maxCupVal = cyc.elements.max()

  for i in 1..100:
    echo "\n-- move ", i, " --"
    move(cyc)

  echo inStr, " => " , cyc.elements
  cyc.print_final_lables_except_1()


main()
