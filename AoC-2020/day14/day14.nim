import std/rdstdin
import std/strutils
import std/math

type
    Memory = object
        registers: seq[string]
        vals: seq[int]

proc int_to_str36bit(a:int):string = 
    result = ""
    var tmp = a
    for i in countup(0, 35):
        result = $(tmp mod 2) & result
        tmp = tmp div 2
    return

proc str36bit_to_int(s:string):int = 
    for i in countup(1, 36):
        result += int(s[^i] == '1') * (2 ^ (i - 1))
    return


proc parse_assign_cmd(inStr: string): tuple[mem:string, val:int] = 
    result.val = parseInt(inStr.split(" = ")[^1])
    result.mem = inStr.split(" = ")[0].split('[')[1].split(']')[0]

proc peculiar_bit_masking(mask:string, binval:string): string =
    result = ""
    for i in countup(0, 35):
        if mask[i] == 'X':
            result = result & $binval[i]
        elif mask[i] == '0':
            result = result & "0"
        elif mask[i] == '1':
            result = result & "1"

proc main =
    var inStr:  string
    var mask:   string
    var memory: Memory

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
            var binval = int_to_str36bit(parsed_cmd.val)
            var masked_applied_tmp = peculiar_bit_masking(mask, binval)
            var dec_val_after_mask = str36bit_to_int(masked_applied_tmp)
            stdout.write(parsed_cmd)
            stdout.write("\t\t", binval, "\t", masked_applied_tmp, "\t", dec_val_after_mask)

            if not (parsed_cmd.mem in memory.registers):
                memory.registers.add(parsed_cmd.mem)
                memory.vals.add(dec_val_after_mask)
            else:
                var ith = find(memory.registers, parsed_cmd.mem)
                memory.vals[ith] = dec_val_after_mask
            # stdout.write("\t|", memory)


        assert memory.vals.len() == memory.registers.len()

    echo "\nsum of vals in reg =", memory.vals.sum()

main()
