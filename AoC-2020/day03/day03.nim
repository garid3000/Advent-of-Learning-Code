import os
let debug = false #$true

type
  Coord = object
    x: int
    y: int

var fname =  "test.txt"
if not debug:
  fname = "input.txt"

var the_map: seq[string]

for line in lines fname:
  echo line, " " ,len(line)
  the_map.add(line)

echo "ready?"
###############################################################################
var coord = Coord(x:1, y:1)
var map_y_len = the_map.len()
var map_x_width = the_map[0].len()
var count_os = 0
var count_xs = 0

proc change_cur_pos(posx: int, posy: int) = 
  stdout.write("\x1b[", posy, ';', posx, 'H')

proc change_cur_pos_symbol(posx: int, posy: int, shape_char: char) = 
  change_cur_pos(posx, posy)
  stdout.write(shape_char)

proc draw_the_map = 
  stdout.write("\x1bc")
  for line in lines fname:
    echo line, " " ,len(line)
    the_map.add(line)

draw_the_map()

while coord.y <= map_y_len:
  var tmp_line = the_map[coord.y - 1]
  if tmp_line[coord.x - 1] == '.':
    change_cur_pos_symbol(coord.x, coord.y, 'O')
    count_os += 1
  else:
    change_cur_pos_symbol(coord.x, coord.y, 'X')
    count_xs += 1

  coord.y = coord.y + 1
  coord.x = (coord.x - 1 + 3) mod map_x_width + 1
  sleep(10)

  change_cur_pos(0, 20)
  echo ' '


change_cur_pos(0, 20)
echo "count_os:\t", count_os
echo "count_xs:\t", count_xs
