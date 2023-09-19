import std/rdstdin
import std/strutils
import std/sequtils
import std/math
import sugar

type
  OkRange = object
    minVal: int
    maxVal: int

  SingleField = object
    name: string
    okRanges: seq[OkRange]

  Ticket = object
    vals: seq[int]

var
  allFields     :seq[SingleField]
  myTicket      :Ticket 
  nearbyTickets :seq[Ticket]

proc isThisInRange(tmpRange: OkRange, val: int): bool = 
  return (tmpRange.minVal <= val) and (val <= tmpRange.maxVal)

proc isThisInRange(tmpField: SingleField, val: int): bool = 
  for each_sep_range in tmpField.okRanges:
    if each_sep_range.isThisInRange(val):
      return true
  return false

proc parse_single_line_field(a: string): SingleField = 
  result.name = a.split(": ")[0]
  #result.okRanges =
  #stdout.write(result.name, "\t")
  result.okRanges = a.split(": ")[1].split(" or ").map(
    #(x:string)->seq[int]=>(x.split("-").map(parseInt))
    (x:string)->OkRange=>(OkRange(minVal:parseInt(x.split("-")[0]), maxVal:parseInt(x.split("-")[1])))
    )
  echo result

proc populate_all_fields(): seq[SingleField] = 
  var inStr: string
  while true:
    var ret = readLineFromStdin("", inStr)
    if not ret:
      quit("SHouldn't be EOF in here")
    if inStr == "":
      return
    result.add(parse_single_line_field(inStr))

proc populate_my_ticket(): Ticket = 
  var inStr: string
  var flagged = false
  while true:
    var ret = readLineFromStdin("", inStr)
    if not ret:
      quit("SHouldn't be EOF in here")
    if flagged:
      result.vals = inStr.split(",").map(parseInt)
      return
    if inStr == "":
      if flagged == false:
        quit("didn't got my ticket")
      return
    if inStr == "your ticket:":
      flagged = true
      continue

proc populate_nearby_tickets(): seq[Ticket] = 
  var inStr: string
  var flagged = false
  while true:
    var ret = readLineFromStdin("", inStr)
    if not ret:
      return #quit("SHouldn't be EOF in here")
    if flagged:
      result.add(Ticket(vals: inStr.split(",").map(parseInt)))
    if not flagged:
      if inStr == "":
        continue
        #quit("didn't got others ticket")
      if inStr == "nearby tickets:":
        flagged = true
        continue
   
proc isTicketsValsOK(tmpTicket: Ticket): bool = 
  for each_val_in_tmp_ticket in tmpTicket.vals:
    var val_appeared_flag = false
    for each_field in allFields:
      if each_field.isThisInRange(each_val_in_tmp_ticket):
        val_appeared_flag = true
        break
    if not val_appeared_flag:
      return false
  return true

proc mockData(tmp_seq_tickets: seq[Ticket]) = 
  echo "nearby tickets:"
  for eachticket in tmp_seq_tickets:
    for eacval in eachticket.vals:
      stdout.write(eacval, ",")
    stdout.write "\n"



proc main = 
  allFields     = populate_all_fields()
  myTicket      = populate_my_ticket()
  nearbyTickets = populate_nearby_tickets()
  echo allFields, "\n", myTicket, "\n", nearbyTickets, "=====", nearbyTickets.len()

  nearbyTickets = filter(nearbyTickets, isTicketsValsOK)
  echo nearbyTickets.len()

  echo "-----------------------------------\n\n\n\n"
  var tickets_ith_col_and_valid_fields_numbers: seq[seq[int]]
  for tickets_ith_col in 0..<myTicket.vals.len():
    var tmp_all_val_in_ith_col = nearbyTickets.map(
      (x: Ticket) -> int => x.vals[tickets_ith_col]
    )
    #echo tickets_ith_col, "--", tmp_all_val_in_ith_col
    var tmp_valid_field_id = allFields.filter(
      (x: SingleField) -> bool => (
        tmp_all_val_in_ith_col.map(
          (xx:int) -> int => int(x.isThisInRange(xx))
        ).sum() == tmp_all_val_in_ith_col.len()) # tmp_all_val_in_ith_col.map
    ).map(
      (x: SingleField) -> (int) => find(allFields, x)
    )
    echo tickets_ith_col, tmp_valid_field_id
    tickets_ith_col_and_valid_fields_numbers.add(tmp_valid_field_id)

  var asdf: seq[int] = newSeq[int](allFields.len())
  for i in 0..<allFields.len():
    asdf[i] = -1
  echo asdf

  while true:
    if asdf.map(
      (x: int) -> int => int(x == -1)
    ).sum() == 0:
      break
    for i, availfields in tickets_ith_col_and_valid_fields_numbers:
      if availfields.len() == 1:
        asdf[i] = availfields[0]
        echo asdf
        for j in 0..<tickets_ith_col_and_valid_fields_numbers.len():
          var tmp_i = tickets_ith_col_and_valid_fields_numbers[j].find(availfields[0])
          if tmp_i != -1:
            tickets_ith_col_and_valid_fields_numbers[j].delete(tmp_i)


  var productval = 1
  for i, each_vals in myTicket.vals:
    echo i, "\t", each_vals, "\t", allFields[asdf[i]]
    if "departure" in allFields[asdf[i]].name :
      productval *= each_vals

  echo "final prod val: ", productval

main()
