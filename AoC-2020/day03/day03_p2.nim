import os
let debug = false #true

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
var map_y_len = the_map.len()
var map_x_width = the_map[0].len()

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


proc try_different_slopes(dx: int, dy: int): int = 
  var coord = Coord(x:1, y:1)
  var count_os = 0
  var count_xs = 0

  draw_the_map()
  while coord.y <= map_y_len:
    var tmp_line = the_map[coord.y - 1]
    if tmp_line[coord.x - 1] == '.':
      change_cur_pos_symbol(coord.x, coord.y, 'O')
      count_os += 1
    else:
      change_cur_pos_symbol(coord.x, coord.y, 'X')
      count_xs += 1

    coord.y = coord.y + dy
    coord.x = (coord.x - 1 + dx) mod map_x_width + 1
    sleep(10)

    change_cur_pos(0, 20)
    echo ' '

  result = count_xs

proc main = 
  change_cur_pos(0, 30)
  var s1 =  try_different_slopes(1, 1) 
  var s2 =  try_different_slopes(3, 1) 
  var s3 =  try_different_slopes(5, 1) 
  var s4 =  try_different_slopes(7, 1) 
  var s5 =  try_different_slopes(1, 2) 
  stdout.write(s1, " * ", s2, " * ",s3, " * ",s4, " * ", s5, " = ", s1*s2*s3*s4*s5)

main()
