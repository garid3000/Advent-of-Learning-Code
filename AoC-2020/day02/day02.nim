import std/strutils

let filepath = "input.txt"
# let filepath = "test.txt"

proc day02_part01() = 
  var ok_count = 0

  for line in lines filepath:
    let str_parts = split(line, ' ')
    let numrange = split(str_parts[0], '-')
    let min_range = parseInt(numrange[0])
    let max_range = parseInt(numrange[1])
    let eachchar = str_parts[1][0]
    let password = str_parts[2]
    if min_range <= count(password, eachchar):
      if count(password, eachchar) <= max_range:
        echo ok_count, " ", line, " ok"
        ok_count+=1
    
  echo ok_count

proc day02_part02() = 
  var ok_count = 0

  for line in lines filepath:
    let str_parts = split(line, ' ')
    let numvals = split(str_parts[0], '-')
    let index0 = parseInt(numvals[0]) - 1
    let index1 = parseInt(numvals[1]) - 1
    let eachchar = str_parts[1][0]
    let password = str_parts[2]
    var tmp_char_occur_count = 0
    if password[index0] != eachchar:
      tmp_char_occur_count += 1
    if password[index1] != eachchar:
      tmp_char_occur_count += 1

    if tmp_char_occur_count == 1:
      echo ok_count, " ", line, " ok"
      ok_count+=1
    
  echo ok_count

day02_part01()
day02_part02()
