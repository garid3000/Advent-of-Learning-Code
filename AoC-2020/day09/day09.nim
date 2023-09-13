import os
import std/rdstdin
import std/strutils
import std/strformat

type
    Combo_index = object
        combo_found: bool
        ith: int
        jth: int

var input_line: string
if paramCount() != 1:
    echo "bad input"

var num_preamble: int = parseInt(paramStr(1))
echo "====", num_preamble, "=========", paramCount()

var lastest_preamble_list: seq[int] = newSeq[int](num_preamble)
var all_numbers_list: seq[int]
var the_outlier_num: int


for i in countup(0, num_preamble-1):
    var ret = readLineFromStdin("", input_line)
    if not ret:
        break
    lastest_preamble_list[i] = parseInt(input_line)
    all_numbers_list.add(parseInt(input_line))

proc term_clean =
    stdout.write("\x1bc")

proc term_goto(x: int, y: int) =
    # because ansi-term (starts from (1,1))
    var new_x = x + 1
    var new_y = y + 1
    stdout.write(fmt("\x1b[{new_y};{new_x}H"))

term_clean()
while true:
    var ret = readLineFromStdin("", input_line)
    if not ret:
        break
    #echo input_line
    var new_value = parseInt(input_line)
    var combo_i_j = Combo_index(combo_found:false)
    # check whether this is sum of 2 numbers
    for i in countup(0, num_preamble-1):
        for j in countup(0, num_preamble-1):
            if i == j:
                break
            if new_value == lastest_preamble_list[i] + lastest_preamble_list[j]:
                combo_i_j.combo_found = true
                combo_i_j.ith = i
                combo_i_j.jth = j
                break

    if not combo_i_j.combo_found:
        term_goto(0, num_preamble + 2)
        stdout.write(&"couldn't find this {new_value}    \n")
        the_outlier_num = new_value
        # break

    for i in countup(0, num_preamble-1):
        term_goto(0, i)
        stdout.write(&"{i:02d}\t{lastest_preamble_list[i]}\t")
        if (combo_i_j.ith == i) or (combo_i_j.jth == i):
            stdout.write(&" <------- ")
        else:
            stdout.write("           ")

    stdout.write(&"{new_value}    \n")
    # sleep(10)


    
    #for i in countup(0, num_preamble-2):
    #    lastest_preamble_list[i] = lastest_preamble_list[i+1]
    #    #parseInt(input_line)

    lastest_preamble_list[0..^2] = lastest_preamble_list[1..^1]
    lastest_preamble_list[^1] = parseInt(input_line)
    all_numbers_list.add(new_value)


#echo(all_numbers_list, " ", all_numbers_list.len())
echo(all_numbers_list.len())


proc slice_sum_min_max(ith: int, jth: int): tuple[sum:int, min:int, max:int] =
    result.min = 999999999999999999
    for i in countup(ith, jth):
        result.sum += all_numbers_list[i]
        if result.max <  all_numbers_list[i]:
            result.max = all_numbers_list[i]
        if result.min >  all_numbers_list[i]:
            result.min = all_numbers_list[i]
    return

for i in countup(0, all_numbers_list.len()-1):
    for j in countup(i + 1, all_numbers_list.len()-1):
        var tmp_sum_min_max = slice_sum_min_max(i, j)
        if tmp_sum_min_max.sum == the_outlier_num:
            echo i, "->",j, "\t", the_outlier_num, "=outlier\t" ,tmp_sum_min_max, "\t the sum of max min = ", tmp_sum_min_max.max + tmp_sum_min_max.min
