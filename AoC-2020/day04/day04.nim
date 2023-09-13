import std/strutils

# let fpath = "test.txt"
let fpath = "input.txt"

let ok_char_for_hex_chars = @['a', 'b', 'c', 'd', 'e', 'f', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
let ok_stings_for_eye_color = @[ "amb", "blu", "brn", "gry", "grn", "hzl", "oth" ]

type
  Passport = object
    byr: string
    iyr: string
    eyr: string
    hgt: string
    hcl: string
    ecl: string
    pid: string
    cid: string

proc is_valid_passport_part_1(pport: Passport): bool = 
  if (pport.byr == "") or 
     (pport.iyr == "") or
     (pport.eyr == "") or
     (pport.hgt == "") or
     (pport.hcl == "") or
     (pport.ecl == "") or
     (pport.pid == ""): # (pport.cid != ""):
    result = false
    return
  result = true
  return

proc try_str_2_int(shouldbenumericstr: string): (bool, int) =
  if shouldbenumericstr == "":
    return (false, 0) 
  
  var int_val = 0
  for each_char in shouldbenumericstr:
    if int(each_char) >= int('0') and int(each_char) <= int('9'):
      int_val = int_val * 10 + (int(each_char) - int('0'))
    else:
      return (false, 0)
      
  return (true, int_val)


proc is_valid_passport_part_2(pport: Passport): bool = 
  # check byr
  var (ret, val) = try_str_2_int(pport.byr)
  if (not ret) or (val < 1920) or (val > 2002):
    result = false
    echo "==> bad birth"
    return

  # check iyr
  (ret, val) = try_str_2_int(pport.iyr)
  if (not ret) or (val < 2010) or (val > 2020):
    result = false # stdout.write(" ====", val, "~", pport.iyr, "==!!== ")
    echo "==> bad iyr"
    return

  # check eyr
  (ret, val) = try_str_2_int(pport.eyr)
  if (not ret) or (val < 2020) or (val > 2030):
    echo "==> bad eyr"
    result = false
    return

  if pport.hgt.len() < 2:
    result = false
    echo "==> bad hgt"
    return
  var unit_cm_or_in = pport.hgt[^2 .. ^1]
  (ret, val) = try_str_2_int(pport.hgt[0 .. ^3])

  if not ret:
    echo "==> bad hgt (bad int)"
    return false

  if unit_cm_or_in == "cm":
    if (val < 150) or (val > 193):
      echo "==> bad hgt (bad cm int)"
      return false
  elif unit_cm_or_in == "in":
    if (val < 59) or (val > 76):
      echo "==> bad hgt (bad in int)"
      return false
  else:
    echo "==> bad hgt"
  
  if pport.hcl.len() != 7:
    echo "bad hcl len"
    return false
  if pport.hcl[0] != '#':
    echo "bad hcl not # at begnin"
    return false
  for i in countup(1, 6):
    var each_char = pport.hcl[i]
    if not(each_char in ok_char_for_hex_chars):
      echo "bad hcl at", i, "th char = ", pport.hcl[i]
      return false
    #if not (each_char in ):
  
  if not (pport.ecl in ok_stings_for_eye_color):
    return false

  if pport.pid.len() != 9:
    return false

  (ret,val) = try_str_2_int(pport.pid)
  if not ret:
    return false

  return true


var passport_list: seq[Passport]
var line_splitted: seq[string]
var a_entry: seq[string]

# reading the files, and populating passport list
var tmp_passport = Passport()
for line in lines fpath:
  if line == "":
    echo "=============== "
    passport_list.add(tmp_passport)
    tmp_passport = Passport()
  else:
    line_splitted = line.split(' ')
    for each_split in line_splitted:
      a_entry = each_split.split(':')
      if   a_entry[0] == "byr":
        tmp_passport.byr = a_entry[1]
      elif a_entry[0] == "iyr":
        tmp_passport.iyr = a_entry[1]
      elif a_entry[0] == "eyr":
        tmp_passport.eyr = a_entry[1]
      elif a_entry[0] == "hgt":
        tmp_passport.hgt = a_entry[1]
      elif a_entry[0] == "hcl":
        tmp_passport.hcl = a_entry[1]
      elif a_entry[0] == "ecl":
        tmp_passport.ecl = a_entry[1]
      elif a_entry[0] == "pid":
        tmp_passport.pid = a_entry[1]
      elif a_entry[0] == "cid":
        tmp_passport.cid = a_entry[1]
      echo each_split


var ok_pass_counter_1 = 0
var ok_pass_counter_2 = 0

for iterating_passport in passport_list:
  var is_it_ok_p1 = is_valid_passport_part_1(iterating_passport)
  var is_it_ok_p2 = is_valid_passport_part_2(iterating_passport)
  ok_pass_counter_1 += int(is_it_ok_p1)
  ok_pass_counter_2 += int(is_it_ok_p2)
  echo(ok_pass_counter_1, "\t", ok_pass_counter_2, "\t", iterating_passport)
  
