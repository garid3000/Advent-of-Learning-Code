import std/rdstdin
import std/strutils

proc sorting(list: seq[int]): seq[int] = 
    result.add(list[0])
    for e_init_list in list[1..^1]:
        var e_init_list_insert = false
        for i, e_sorted_int in result:
            if e_init_list < e_sorted_int:
                result.insert(e_init_list, i)
                e_init_list_insert = true
                break
        if not e_init_list_insert:
            result.add(e_init_list)
        echo result
    return

proc jolt_difference_count(sorted_list: seq[int], diff: int): int=
    var val = 0
    for adapJolt in sorted_list:
        if (adapJolt - val) == diff:
            result += 1
        val = adapJolt


proc main = 
    var adapters: seq[int]
    var linestr: string

    while true:
        var ret = readLineFromStdin("", linestr)
        if not ret:
            break
        adapters.add(parseInt(linestr))
        stdout.write(linestr, "\n")
    echo adapters, "\t", adapters.len()
    var sorted_adapters = sorting(adapters)
    sorted_adapters.add(sorted_adapters[^1]+3)
    echo sorted_adapters, "\t", sorted_adapters.len()

    var counted_1jolt_diff = jolt_difference_count(sorted_adapters, 1) 
    var counted_3jolt_diff = jolt_difference_count(sorted_adapters, 3)

    echo "jolt_difference_count(sorted_adapters, 1)", counted_1jolt_diff
    echo "jolt_difference_count(sorted_adapters, 3)", counted_3jolt_diff
    echo "multiplied = ", counted_1jolt_diff * counted_3jolt_diff

    sorted_adapters.insert(0, 0)
    echo "final ", sorted_adapters
    var diffstring: string
    for i in countup(1, sorted_adapters.len()-1):
        diffstring.add($(sorted_adapters[i] - sorted_adapters[i-1]))
    echo diffstring

    # some math trick
    diffstring = diffstring.replace("3", " ")
    echo diffstring

    diffstring = diffstring.replace("1111", "7   ")
    echo diffstring

    diffstring = diffstring.replace("111", "4  ")
    echo diffstring

    diffstring = diffstring.replace("11", "2 ")
    echo diffstring

    diffstring = diffstring.replace("1", " ")
    echo diffstring

    var product = 1
    for each_pool in diffstring.split(" "):
        if each_pool != "":
            product *= parseInt(each_pool)
    echo "final product for part 2: = ", product

main()


#  3 11 3 => 11 => 2 possibilties
#  3 111 3 => 111 => 4 possibilties
#  3 1111 3 => 111 => 7 possibilties
