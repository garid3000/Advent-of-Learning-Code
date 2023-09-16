import std/rdstdin
import std/strutils

type
    Direciton = enum
        north, west, east, south
        # north = y+, south = y-
        # east = x+,  west = x-
        #

    Ship = object
        x: int
        y: int
        dir: Direciton

proc rotate_ship(ship: Ship, rot_dir_char: char, rot_degree: int): Ship =
    result = ship
    assert rot_dir_char in @['R', 'L']
    assert rot_degree in @[0, 90, 180, 270]
    
    var rot_int = (rot_degree div 90)
    #echo "rotatioon  ", rot_dir_char, "\t", rot_int
    for i in countup(0, rot_int-1):
        if rot_dir_char == 'R':
            if result.dir == south: result.dir = west
            elif result.dir == west:  result.dir = north
            elif result.dir == north: result.dir = east
            elif result.dir == east:  result.dir = south
        if rot_dir_char == 'L':
            if result.dir == south: result.dir = east
            elif result.dir == east:  result.dir = north
            elif result.dir == north: result.dir = west
            elif result.dir == west:  result.dir = south

proc forward_ship(ship: Ship, forward_val: int): Ship =
    result = ship
    var
        dx: int
        dy: int
    if result.dir == south: (dx, dy)  = (0, -1)
    if result.dir == west:  (dx, dy)  = (-1, 0)
    if result.dir == north: (dx, dy)  = (0,  1)
    if result.dir == east:  (dx, dy)  = (1,  0)

    result.x += dx * forward_val
    result.y += dy * forward_val

proc nsew(ship: Ship, dir_char: char, movement_val: int):Ship = 
    result = ship
    assert dir_char in @['N', 'S', 'E', 'W']
    if dir_char == 'N':
        result.y += movement_val
    if dir_char == 'S':
        result.y -= movement_val
    if dir_char == 'E':
        result.x += movement_val
    if dir_char == 'W':
        result.x -= movement_val

proc parse_each_cmd(curship: Ship, cmd: string): Ship =
    #result = curship
    if cmd[0] in @['N', 'S', 'E', 'W']:
        result = nsew(curship, cmd[0], parseInt(cmd[1..^1]) )
    elif cmd[0] in @['R', 'L']:
        result = rotate_ship(curship, cmd[0], parseInt(cmd[1..^1]) )
    elif cmd[0] == 'F':
        result = forward_ship(curship, parseInt(cmd[1..^1]))

proc main = 
    var str_line: string
    var ship = Ship(x:0, y:0, dir:east)
    echo ship
    while true:
        var ret = readLineFromStdin("", str_line)
        if not ret:
            break
        ship = parse_each_cmd(ship, str_line)
        echo str_line, "\t", ship

    echo "Manhattan distance: ",  abs(ship.x) + abs(ship.y)

main()
