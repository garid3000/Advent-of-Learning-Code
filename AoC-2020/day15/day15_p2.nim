import os
import std/strutils
import std/sequtils
import sugar
# import std/sets

# must count the argv before use
if paramCount() != 3:
  quit("bad num of cmd args\n./day15 1,1,2 true 2020") 
 
var print_mode = (paramStr(2) == "true")
var numseq = paramStr(1).split(',').map( (x:string) -> int => parseInt(x) )
var iterMaxNum = parseInt(paramStr(3))

#var current_iter = numseq.len()-1
#var last_appearance:seq[int] = @[0, 1, 2]
var last_appearance:seq[int] = toSeq(0..(numseq.len()-1))

#var next_num = 0

echo numseq, " ", print_mode, " ", iterMaxNum, "|", last_appearance
echo "------------------------------------------------------------"
###############################################################################
var stackarray_last_ocur: array[30000000, int]

for i in 0..<30000000:
  stackarray_last_ocur[i] = -1


proc add_val(v: int, currentIndex:int): int = 
  result = stackarray_last_ocur[v]
  if result == -1:
    result = 0
    stackarray_last_ocur[v] = currentIndex
    return
  result = currentIndex - stackarray_last_ocur[v]
  stackarray_last_ocur[v] = currentIndex


var next_val:int

for i in 1..numseq.len():
  next_val = add_val(numseq[i-1], i)
  if print_mode:
    echo i, "\t", numseq[i-1], "\t", next_val, "\t", 
         stackarray_last_ocur[0], ",",stackarray_last_ocur[1], ",",stackarray_last_ocur[2], ",",stackarray_last_ocur[3], ",", stackarray_last_ocur[4], ",",
         stackarray_last_ocur[5], ",",stackarray_last_ocur[6], ",",stackarray_last_ocur[7], ",",stackarray_last_ocur[8], ",", stackarray_last_ocur[9], ","


for i in numseq.len()+1..iterMaxNum-1:
  if print_mode:
    stdout.write(i, "\t", next_val) 
  next_val = add_val(next_val, i)
  if print_mode:
    echo "\t", next_val, "\t",
         stackarray_last_ocur[0], ",",stackarray_last_ocur[1], ",",stackarray_last_ocur[2], ",",stackarray_last_ocur[3], ",", stackarray_last_ocur[4], ",",
         stackarray_last_ocur[5], ",",stackarray_last_ocur[6], ",",stackarray_last_ocur[7], ",",stackarray_last_ocur[8], ",", stackarray_last_ocur[9], ","

echo iterMaxNum, "===>", next_val

#proc find_last_appearance(v: int): int =
#  #for i in countdown(intlist.len()-2, 0):
#  #  if intlist[i] == v:
#  #    return intlist.len()-1-i
#  var ith = find(numseq, v)
#  if ith == -1:
#    return -1
#  return last_appearance[ith] 
#
#
#proc add_val_at_turn() =
#  if print_mode:
#    echo current_iter + 1, " ", next_num , "\t", numseq, "\t", last_appearance
#  current_iter+=1
#  var ith = find(numseq, next_num)
#  if ith == -1:
#    numseq.add(next_num)
#    last_appearance.add(current_iter)
#    next_num = 0
#  else:
#    #echo current_iter,"===",last_appearance[ith], "===", ith
#    next_num = current_iter - last_appearance[ith]
#    last_appearance[ith] = current_iter
#    #echo current_iter, "--",last_appearance[ith]
#
#
#
#
#for i in numseq.len()..iterMaxNum-2:
#  add_val_at_turn()
#  # var lastval = numseq[^1]
#  # var newval = find_last_appearance(numseq, lastval)
#  # numseq.add(newval)
#  # if print_mode:
#  #   echo numseq
#
#echo numseq[^5..^1], "last val ", numseq[^2], "\t", next_num
