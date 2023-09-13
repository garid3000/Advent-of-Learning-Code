import std/strutils
let filepath = "input.txt"

var numbers: seq[int]

proc read_the_file() =
  let f = open(filepath)
  defer: f.close()

  for line in lines filepath: # kinda similar to my style of writing with python lines = file.readlines()
    echo line, "\tasfd\t"
    numbers.add(parseInt(line))

proc checkallcombos2() = 
  for i in 0..numbers.len-1:
    for j in 0..numbers.len-1:
      if numbers[i] + numbers[j] == 2020:
        echo i, ',',  j, ',', numbers[i], ',', numbers[j], ',', numbers[i]*numbers[j]
        return

proc checkallcombos3() = 
  for i in 0..numbers.len-1:
    for j in 0..numbers.len-1:
      for k in 0..numbers.len-1:
        if numbers[i] + numbers[j] + numbers[k] == 2020:
          echo i, ',',  j, ',', k, ',', numbers[i], ',', numbers[j], ',', numbers[k], ',',numbers[i]*numbers[j]*numbers[k]
          return
    

read_the_file()
echo numbers
checkallcombos2()
checkallcombos3()
