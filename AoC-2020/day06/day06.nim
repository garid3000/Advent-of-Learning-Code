import std/rdstdin

var input_line: string
var tmp_group_unique_chars: seq[char] = @[]
var a_single_group_answers: seq[string] = @[]

var sum_of_unique_chars_in_group = 0
var sum_of_chars_that_in_all_answers_in_1_group = 0

proc count_chars_that_appears_in_all_elements(a_group_answers: seq[string]): seq[char] = 
  var tmp: seq[char] = @[]
  for each_char in a_group_answers[0]:
    tmp.add(each_char)

  for a_person_answer in a_group_answers:
    var tmptmp: seq[char] = @[]
    for each_char in tmp:
      if each_char in a_person_answer:
        tmptmp.add(each_char)
    tmp = tmptmp

  return tmp


while true:
  let retisok = readLineFromStdin("", input_line)
  if not retisok:
    break
  if input_line == "":
    echo "---------------------", tmp_group_unique_chars,  "\t", tmp_group_unique_chars.len(), "\t", count_chars_that_appears_in_all_elements(a_single_group_answers)
    sum_of_unique_chars_in_group += tmp_group_unique_chars.len()
    sum_of_chars_that_in_all_answers_in_1_group += count_chars_that_appears_in_all_elements(a_single_group_answers).len()
    tmp_group_unique_chars = @[]
    a_single_group_answers = @[]
  else:
    echo input_line
    a_single_group_answers.add(input_line)
    for each_char_in_line in input_line:
      if not (each_char_in_line in tmp_group_unique_chars):
        tmp_group_unique_chars.add(each_char_in_line)

echo "final (part-1): ", sum_of_unique_chars_in_group
echo "final (part-2): ", sum_of_chars_that_in_all_answers_in_1_group
