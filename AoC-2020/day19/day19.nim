import std/strutils
import std/sequtils
import sugar

type
  RuleType = enum
    simple_char, single_series, double_series

  OneRule = object
    id: int
    ruleType: RuleType

    single_series_id: seq[int]
    double_series_ids: (seq[int], seq[int])
    comparing_char: char

var
  rulebook: seq[OneRule]
  zerothruleid: int

proc spaceSep2int(a:string): seq[int]= 
  result =  a.split(" ").map(
    (x: string)->int => x.parseInt()
  )
  

proc populate_rulebook = 
  while true:
    var tmp = stdin.readLine()
    if tmp == "":
      return
    elif '"' in tmp:
      echo tmp, "\t simple char rule" 
      rulebook.add(
        OneRule(
          id: tmp.split(":")[0].parseInt(),
          ruleType: simple_char,
          comparing_char: tmp.split('"')[1][0],
        )
      )
    elif '|' in tmp:
      echo tmp, "\t double series rule" 
      rulebook.add(
        OneRule(
          id: tmp.split(": ")[0].parseInt(),
          ruleType: double_series,
          double_series_ids: (
            tmp.split(": ")[1].split(" | ")[0].spaceSep2int(),
            tmp.split(": ")[1].split(" | ")[1].spaceSep2int()
            ),
        )
      )
    else:
      echo tmp, "\t single series rule"
      rulebook.add(
        OneRule(
          id: tmp.split(": ")[0].parseInt(),
          ruleType: single_series,
          single_series_id: tmp.split(": ")[1].spaceSep2int()
        )
      )

populate_rulebook()
echo "===========\n\n"
for eachRule in rulebook:
  if eachRule.ruleType == simple_char:
    echo eachRule.id, "\t", eachRule.comparing_char
  elif eachRule.ruleType == single_series:
    echo eachRule.id, "\t", eachRule.single_series_id
  elif eachRule.ruleType == double_series:
    echo eachRule.id, "\t", eachRule.double_series_ids

  if eachRule.id == 0:
    echo eachRule.id, "==============================================================="
    zerothruleid = eachRule.id


# proc checkAgainstRulebook(s: string, jthChar: int, ithRule: int): bool = 
#   if rulebook[ithRule].ruleType == simple_char:
#     if (s[jthChar] == rulebook[ithRule].comparing_char) and (jthChar == s.len()-1):
#       return true
#     else:
#       return false
#   if rulebook[ithRule].ruleType == single_series:
#     for ij, eachSubRule in rulebook[ithRule].single_series_id:
#       if not checkAgainstRulebook(s, jthChar + ij, eachSubRule):
#         return false
#     return true
# 
#   return
#
#

proc checkAgainstRulebook(s: string, ithRule: int): (bool, string)

proc checkSeriesOfRules(s: string, ruleIds: seq[int]): (bool, string) =
    var tmpString = s
    var ret: bool
    for i_eachSubRule in ruleIds:
      (ret, tmpString) = checkAgainstRulebook(tmpString, i_eachSubRule)
      if not ret:
        return (false, tmpString)
    return (true, tmpString)


proc checkAgainstRulebook(s: string, ithRule: int): (bool, string) = 
  if rulebook[ithRule].ruleType == simple_char:
    if (s[0] == rulebook[ithRule].comparing_char):
      return (true, s[1..^1])
    else:
      return (false, "")

  elif rulebook[ithRule].ruleType == single_series:
    return checkSeriesOfRules(s, rulebook[ithRule].single_series_id)

  elif rulebook[ithRule].ruleType == double_series:
    var tmp1 = checkSeriesOfRules(s, rulebook[ithRule].double_series_ids[0])
    if tmp1[0]:
      return tmp1
    var tmp2 = checkSeriesOfRules(s, rulebook[ithRule].double_series_ids[1])
    if tmp2[0]:
      return tmp2
    return (false, "")
  return

proc main = 
  var instr:string
  while true:
    try:
      instr = stdin.readLine()
    except:
      break
    echo instr, "\t" , instr.checkAgainstRulebook(zerothruleid)
  return

main()
