import os
import std/rdstdin

proc term_cur_jump(y: int, x: int) = 
    stdout.write("\x1b[", y+1, ";", x+1, "H")

proc print_map(map_str: seq[string]) = 
    for y,each_line in map_str:
        term_cur_jump(y, 0)
        stdout.write(each_line, "\n")

proc populate_map(): seq[string] = 
    var input_line: string
    while true:
        var ret = readLineFromStdin("", input_line)
        if not ret:
            break
        result.add(input_line)

proc count_active_neighbors(y:int, x:int, map: seq[string]): int = 
    for ny in countup(max(0, y-1), min(y+1, map.len()-1)):
        for nx in countup(max(0, x-1), min(x+1, map[0].len()-1)):
            result += int(map[ny][nx] == '#')
    result -= int(map[y][x] == '#')
    return


proc update_map(prevmap: seq[string]):seq[string] = 
    result = prevmap
    for y, each_line in prevmap:
        for x, each_char in each_line:
            var neightCount = count_active_neighbors(y, x, prevmap)
            if (neightCount == 0) and prevmap[y][x] == 'L':
                result[y][x] = '#'
            elif (neightCount >= 4) and prevmap[y][x] == '#':
                result[y][x] = 'L'

proc count_all_active_seat(map: seq[string]): int= 
    for y, each_line in map:
        for x, each_char in each_line:
            if map[y][x] == '#':
                result += 1

proc main =
    stdout.write("\x1bc")
    var map = populate_map()
    var new_map: seq[string]
    while true:
        print_map( map )
        sleep(100) 
        new_map = update_map(map)
        if new_map == map:
            break
        else:
            map = new_map
    echo "\n\nfinal active seat sum: ", count_all_active_seat(new_map)

main()
