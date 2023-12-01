# import std/strutils
import std/sequtils
import sugar

const cycSize = 1_000_000
const iterNum = 10_000_000

type
  Cycle = ref object
    #elements: seq[int]
    elements: array[cycSize, int] # 1..1_000_000
    curVal: int
    curIndex: int
    desVal: int
    desIndex: int

    picked3_buffer: array[3, int]
    #maxCupVal: int

proc findDesVal(cyc: Cycle, pos_des_val: int): int {.inline.} = 
  result = pos_des_val

  while true:
    if result in cyc.picked3_buffer:
      result -= 1
      if result == 0:
        result = cycSize # cyc.maxCupVal
    else: 
      return

proc initialize_elements(cyc: Cycle, initial_seq: seq[int]) = 
  for i in 0..<cycSize:
    cyc.elements[i] = i+1

  for i, eachVal in initial_seq:
    cyc.elements[i] = eachVal


proc pick3cups(cyc: Cycle) {.inline.}= 
  ## puts curIndex+1, curIndex+2, curIndex+3, val into buffer
  cyc.picked3_buffer[0..<3] = cyc.elements[cyc.curIndex+1..cyc.curIndex+3]

proc move(cyc: Cycle) {.inline.}= 
  #var round_cur_val = cyc.elements[cyc.curIndex] 
  #var round_cur_ind = cyc.curIndex
  cyc.curVal = cyc.elements[cyc.curIndex]                               # O(1)
  #cyc.print_cups_with_selectedness()

  cyc.pick3cups()                                                       # O(1)

  #echo "pick up: ", cyc.picked3, " ", cyc.elements
  # 
  cyc.desVal = cyc.findDesVal(cyc.curVal - 1)                           # O(1) worst 3
  cyc.desIndex = cyc.elements.find(cyc.desVal)                          # O(N) worst n


  if cyc.desIndex > cyc.curIndex + 3:
    #shift counter-clockwise
    #cyc.elements[cyc.curIndex+1..cyc.desIndex-3] = cyc.elements[cyc.curIndex+4..cyc.desIndex]
    for i in cyc.curIndex+1..cyc.desIndex-3:
      cyc.elements[i] = cyc.elements[i+3]
    
    cyc.elements[cyc.desIndex-2..cyc.desIndex] = cyc.picked3_buffer[0..2]
  elif cyc.desIndex < cyc.curIndex:
    #shift clockwise
    #cyc.elements[cyc.desIndex+4..cyc.curIndex+3] = cyc.elements[cyc.desIndex+1..cyc.curIndex] 
    #this need god damn sh*t
    #for i in cyc.curIndex+3..cyc.desIndex+4: # reverse aka countdown
    #  cyc.elements[i] = cyc.elements[i-3]
    for i in cyc.curIndex+1..cycSize-4:
      cyc.elements[i] = cyc.elements[i+3]

    cyc.elements[cyc.desIndex+1..cyc.desIndex+3] = cyc.picked3_buffer[0..2]
  else:
    echo "cyc.desIndex=", cyc.desIndex 
    echo "cyc.curIndex=", cyc.curIndex 
    quit("something wrong")
   
  cyc.curIndex =  (cyc.curIndex + 1) mod cycSize
  #echo "destinationVal: ", cyc.desVal

  #cyc.put3_after_des(val_cur_cup, cyc_cur)
  #echo "after putting destinatino"
  #cyc.cur = (cyc.cur + 1) mod cyc.elements.len()

proc main() = 
  var cyc = Cycle()
  var inStr = stdin.readLine()
  var initSeq = inStr.map( (x: char) -> int => int(x) - int('0') )
  cyc.initialize_elements(initSeq)
  #cyc.maxCupVal = 1_000_000 # cyc.maxCupVal = cyc.elements.max()

  for i in 1..iterNum:
    #echo "\n-- move ", i, " --"
    if i mod 1_000 == 0:
      echo "i=", i, "\t", i * 100 / iterNum  , "%"
    cyc.move()

  echo inStr, " => " , cyc.elements[0..20], " max: "

  var next1index =  (cyc.curIndex + 1) mod cycSize
  var next2index =  (cyc.curIndex + 2) mod cycSize
  echo "next1val", cyc.elements[next1index]
  echo "next2val", cyc.elements[next2index]
  echo "multiplying",  cyc.elements[next1index] * cyc.elements[next2index]
  #cyc.print_final_lables_except_1()

main()
