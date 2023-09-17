import std/rdstdin
import std/strutils
import std/math

type
    Memory = object
        registers: seq[string]
        vals: seq[int]
var memory: Memory

proc int_to_str36bit(a:int):string = 
    result = ""
    var tmp = a
    for i in countup(0, 35):
        result = $(tmp mod 2) & result
        tmp = tmp div 2
    return

# proc str36bit_to_int(s:string):int = 
#     for i in countup(1, 36):
#         result += int(s[^i] == '1') * (2 ^ (i - 1))
#     return


proc parse_assign_cmd(inStr: string): tuple[mem:int, val:int] = 
    result.val = parseInt(inStr.split(" = ")[^1])
    result.mem = parseInt(inStr.split(" = ")[0].split('[')[1].split(']')[0])

proc peculiar_bit_masking_memory(mask:string, binval:string): string =
    result = binval
    for i in countup(0, 35):
        if mask[i] == '1':
            result[i] = '1'
        elif mask[i] == 'X':
            result[i] = 'X'

proc indice_of_floating_bits(glob: string): seq[int] =
    for i, eachChar in glob:
        if eachChar == 'X':
            result.add(i)
    return

proc assign_single_register_val(single_reg: string, val:int) = 
    assert not ('X' in single_reg)

    var ith = find(memory.registers, single_reg)
    if ith == -1:
        memory.registers.add(single_reg)
        memory.vals.add(val)
    else:
        memory.vals[ith] = val


proc assign_int_val_to_all_reg_in_glob(glob: string, val: int) = 
    var floaters = indice_of_floating_bits(glob)
    var lcl_glob = glob
    if floaters.len() == 1:
        lcl_glob[floaters[0]] = '0'
        assign_single_register_val(lcl_glob, val)
        lcl_glob[floaters[0]] = '1'
        assign_single_register_val(lcl_glob, val)
    else: #  2+  floaters
        lcl_glob[floaters[0]] = '0'
        assign_int_val_to_all_reg_in_glob(lcl_glob, val)
        lcl_glob[floaters[0]] = '1'
        assign_int_val_to_all_reg_in_glob(lcl_glob, val)

    return

proc main =
    var inStr:  string
    var mask:   string

    while true:
        var ret = readLineFromStdin("", inStr)
        if not ret:
            break

        stdout.write("\n", inStr, "\t")
        if "mask" in inStr:
            mask = inStr.split(" = ")[^1]
            stdout.write(mask)
        elif "mem" in inStr:
            var parsed_cmd = parse_assign_cmd(inStr)
            var bin_mem = int_to_str36bit(parsed_cmd.mem)
            var masked_applied_mem_glob = peculiar_bit_masking_memory(mask, bin_mem)
            #var dec_val_after_mask = str36bit_to_int(masked_applied_tmp)
            stdout.write(parsed_cmd)
            stdout.write("\t\t", bin_mem, "\t", masked_applied_mem_glob)
            assign_int_val_to_all_reg_in_glob(masked_applied_mem_glob, parsed_cmd.val)

        assert memory.vals.len() == memory.registers.len()

    for i, eachmem in memory.registers:
        echo i, "\t", eachmem, "\t <== ", memory.vals[i]

    echo "\n\nsum of vals in reg =", memory.vals.sum()

main()
