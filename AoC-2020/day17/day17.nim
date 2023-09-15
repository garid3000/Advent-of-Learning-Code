import std/rdstdin

var input_line: string
type
    Cell = object
        x: int
        y: int
        z: int
        # a: bool 


proc read_the_cell_at_the_beginning: seq[Cell]=
    var init_y = 0
    while true:
        var ret = readLineFromStdin("", input_line)
        if not ret:
            break

        for i, each_char in input_line:
            if each_char == '#':
                result.add( Cell(x:i, y:init_y, z:0))
        init_y += 1

        echo(input_line, ret)

proc Does_cell_active(cur_cells: seq[Cell], x:int, y:int, z:int):bool =
    for each_cell in cur_cells:
        if (each_cell.x == x) and (each_cell.y == y) and (each_cell.z == z):
            return true
    return false

proc Count_neighbor_cell(cur_cells: seq[Cell], x:int, y:int, z:int):int = 
    for dx in countup(-1, 1):
        for dy in countup(-1, 1):
            for dz in countup(-1, 1):
                if not ((dx == 0) and (dy == 0) and (dz == 0)):
                    result += int(Does_cell_active(cur_cells , x + dx, y + dy, z+dz))
    return
                    
proc max_ranges(cur_cells: seq[Cell]): tuple[x0:int, x1:int, y0:int, y1:int, z0:int, z1:int] = 
    result.x0 = cur_cells[0].x
    result.x1 = cur_cells[0].x

    result.y0 = cur_cells[0].y
    result.y1 = cur_cells[0].y

    result.z0 = cur_cells[0].z
    result.z1 = cur_cells[0].z

    for each_cell in cur_cells:
        if  result.x0 > each_cell.x:
            result.x0 = each_cell.x
        if  result.x1 < each_cell.x:
            result.x1 = each_cell.x

        if  result.y0 > each_cell.y:
            result.y0 = each_cell.y 
        if  result.y1 < each_cell.y:
            result.y1 = each_cell.y

        if  result.z0 > each_cell.z:
            result.z0 = each_cell.z 
        if  result.z1 < each_cell.z:
            result.z1 = each_cell.z

    result.x0 -= 1 
    result.x1 += 1
                
    result.y0 -= 1
    result.y1 += 1
                
    result.z0 -= 1
    result.z1 += 1

proc calculate_next_state(prev_state: seq[Cell]): seq[Cell]=
    var next_possible_range = max_ranges(prev_state)

    for new_x in countup(next_possible_range.x0, next_possible_range.x1):
        for new_y in countup(next_possible_range.y0, next_possible_range.y1):
            for new_z in countup(next_possible_range.z0, next_possible_range.z1):
                var counted_neightbors = Count_neighbor_cell(prev_state, new_x, new_y, new_z)

                if Does_cell_active(prev_state, new_x, new_y, new_z):
                    if (counted_neightbors==2) or (counted_neightbors==3):
                        result.add(Cell(x:new_x, y:new_y, z:new_z))
                else:
                    if (counted_neightbors == 3):
                        result.add(Cell(x:new_x, y:new_y, z:new_z))

proc print_cells(cur_state: seq[Cell])=
    var next_possible_range = max_ranges(cur_state)

    for new_z in countup(next_possible_range.z0, next_possible_range.z1):
        stdout.write("\nz=", new_z, "\n")
        for new_y in countup(next_possible_range.y0, next_possible_range.y1):
            for new_x in countup(next_possible_range.x0, next_possible_range.x1):
                if Does_cell_active(cur_state, new_x, new_y, new_z):
                    stdout.write('#')
                else:
                    stdout.write('.')
            stdout.write('\n')

proc count_active_cells(prev_state: seq[Cell]): int=
    var next_possible_range = max_ranges(prev_state)

    for new_x in countup(next_possible_range.x0, next_possible_range.x1):
        for new_y in countup(next_possible_range.y0, next_possible_range.y1):
            for new_z in countup(next_possible_range.z0, next_possible_range.z1):
                result += int(Does_cell_active(prev_state, new_x, new_y, new_z))

var curCells = read_the_cell_at_the_beginning()


for i in countup(1, 6):
    curCells = calculate_next_state(curCells)
    stdout.write("After ", i,  " cycle\t", count_active_cells(curCells), "\n")

