import os
import std/rdstdin
import std/strformat
import std/strutils

var input_line: string
var the_original_global_coded_instructions: seq[string]
var global_coded_instructions: seq[string]
var global_coded_instructions_visit_counter: seq[int]
var global_accumulator: int
var global_instruction_pos: int
# var count_exec_instruction = 0

type
    Execution_end_type = enum
        exec_looped, exec_not_finished, exec_finished, exec_cur_out_of_range


proc term_clean =
    stdout.write("\x1bc")

proc term_goto(x: int, y: int) =
    # because ansi-term (starts from (1,1))
    var new_x = x + 1
    var new_y = y + 1
    stdout.write(fmt("\x1b[{new_y};{new_x}H"))

proc parse_each_line_instruction: Execution_end_type =
    var ins_str = global_coded_instructions[global_instruction_pos]
    var (str_ope, int_val) = (ins_str.split(' ')[0], parseInt(ins_str.split(' ')[1]))

    global_coded_instructions_visit_counter[global_instruction_pos] += 1
    if global_coded_instructions_visit_counter[global_instruction_pos] > 1:
        return exec_looped

    if   str_ope == "nop": # nothing happens
        global_instruction_pos += 1
    elif str_ope == "acc":
        global_accumulator += int_val
        global_instruction_pos += 1
    elif str_ope == "jmp":
        global_instruction_pos += int_val
        if (global_instruction_pos < 0) or (global_instruction_pos > global_coded_instructions.len()):
            return exec_cur_out_of_range
        
    # count_exec_instruction+=1
    if global_instruction_pos == global_coded_instructions.len():
        return exec_finished
    else:
        return exec_not_finished

proc populate_global_code_instructions: seq[string] = 
    while true:
        var ret = readLineFromStdin("", input_line)
        if not ret:
            break
        result.add(input_line)

    #global_coded_instructions_visit_counter = newseq[0](result.len())

proc part1(debug: bool): tuple[isCorrect: bool, accVal: int] = 
    while true:
        var tmp_position_before_exec = global_instruction_pos
        var ret = parse_each_line_instruction()

        if debug: 
            term_clean()
            term_goto(0, tmp_position_before_exec)
            echo tmp_position_before_exec, "=>  ", global_coded_instructions[tmp_position_before_exec], "\t accVal=", global_accumulator
            
            term_goto(0, 20);
            echo global_coded_instructions_visit_counter
            sleep 1000

        if ret == exec_looped:
            #term_clean()
            return (false, global_accumulator)
        elif ret == exec_finished:
            #term_clean()
            return (true, global_accumulator)

    # return (false, global_accumulator)
    # echo "final acc val= ", global_accumulator


proc duplicate_global_coded_instructions_from_orig = 
    global_coded_instructions = newSeq[""](the_original_global_coded_instructions.len())
    for i in countup(0, the_original_global_coded_instructions.len()-1):
        global_coded_instructions[i] = the_original_global_coded_instructions[i]
    global_coded_instructions_visit_counter = newSeq[0](the_original_global_coded_instructions.len())
    global_accumulator = 0
    global_instruction_pos = 0

proc part2(debug: bool) = 
    for i in countup(0, the_original_global_coded_instructions.len()-1):
        duplicate_global_coded_instructions_from_orig()
        if not("acc" in global_coded_instructions[i]):

            if "jmp" in global_coded_instructions[i]:
                global_coded_instructions[i] = global_coded_instructions[i].replace("jmp", "nop")
            elif "nop" in global_coded_instructions[i]:
                global_coded_instructions[i] = global_coded_instructions[i].replace("nop", "jmp")

            #stdout.write(i, " ", $global_coded_instructions, "|\t")
            var tmp = part1(false)
            echo i,  "\t-->", tmp
            if tmp.isCorrect:
                echo "i'm breaking from searching"
                break
        

        #global_coded_instructions[i] = the_original_global_coded_instructions[i]



the_original_global_coded_instructions = populate_global_code_instructions()
duplicate_global_coded_instructions_from_orig()
echo part1(false)
part2(false)
