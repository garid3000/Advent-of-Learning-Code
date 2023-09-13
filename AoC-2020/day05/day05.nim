import std/rdstdin

var line:string
var row_str: string
var col_str: string
var remember_highest_seatid: int = 0
var remember_lowest_seatid: int = 999999
var seats: seq[int]

proc binstr_2_int(a_string: string, one_char: char): int = 
  var val = 0
  var tmp_string = a_string
  while tmp_string.len() > 0:
    val *= 2
    if tmp_string[0] == one_char:
      val += 1
    tmp_string = tmp_string[1..^1]
  
  return val

while true:
  let ok = readLineFromStdin("How are you?", line)
  if not ok: break
  assert line.len == 10
  row_str = line[0..6]
  col_str = line[7..9]

  var val_row = binstr_2_int(row_str, 'B')
  var val_col = binstr_2_int(col_str, 'R')
  var seatID = val_row * 8 + val_col
  if remember_highest_seatid < seatID:
    remember_highest_seatid = seatID
  if remember_lowest_seatid > seatID:
    remember_lowest_seatid = seatID
  seats.add(seatID)
  echo row_str, "+", col_str, "\trow:", val_row, "\tcol:", val_col, '\t', seatID, '\t', remember_highest_seatid 

for i in countup(remember_lowest_seatid, remember_highest_seatid):
  if not(i in seats):
    echo "this is probably your seat = ", i , "this didn't show up in the seats"
