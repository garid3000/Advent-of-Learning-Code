import std/rdstdin
import std/strutils
import std/sequtils
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
   
iterator all_invalid_values(validFields: seq[SingleField], checkingTickets: seq[Ticket]):int=
  for eachTicket in checkingTickets:
    for i,eachVal in eachTicket.vals:
      var val_appeared_at_least_one_of_field = false
      for each_field in validFields:
        if each_field.isThisInRange(eachVal):
          val_appeared_at_least_one_of_field = true
          break
      if not val_appeared_at_least_one_of_field:
        echo eachVal
        yield eachVal

proc main = 
  var allFields = populate_all_fields()
  var myTicket = populate_my_ticket()
  var nearbyTickets = populate_nearby_tickets()
  echo allFields, "\n", myTicket, "\n", nearbyTickets

  var ticket_scan_err_rate = 0
  for each_invalid_val in all_invalid_values(allFields, nearbyTickets):
    ticket_scan_err_rate += each_invalid_val
  echo "ticket_scan_err_rate", ticket_scan_err_rate

main()
