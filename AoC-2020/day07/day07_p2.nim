import std/rdstdin
import std/strutils

type
    Bag = object
        name: string
        subbags: seq[tuple[sname:string, scout:int]]
        inner_bags_count: int

var input_line: string
var all_bags: seq[Bag]

proc parse_each_sub_sentence(sub_sentence: string): tuple[sname:string, scout:int]=
    if sub_sentence == "no other bags.":
        return
    var num = parseInt(sub_sentence.split(' ')[0])
    var color = sub_sentence.split(' ')[1] & " " & sub_sentence.split(' ')[2]
    # stdout.write sub_sentence, '\t'
    result = (color, num)

proc parse_raw_sub_bags_seq(sub_bags_raw_str: string): seq[tuple[sname:string, scout:int]] =
    var sub_sentences = sub_bags_raw_str.split(", ")
    for each_sentence in sub_sentences:
        var tmp = parse_each_sub_sentence(each_sentence)
        if tmp.sname != "": # to reduce the empty bags
            result.add(tmp)
    return


proc read_and_populate_bags(debugging: bool): seq[Bag] = 
    while true:
        let ret = readLineFromStdin("", input_line)
        if not ret:
            break
        var bags_and_subbags = input_line.split(" bags contain ")
        assert bags_and_subbags.len() == 2
        var tmp_bag_name = bags_and_subbags[0]
        var tmp_subbags_str_raw = bags_and_subbags[1]

        if debugging:
            echo  tmp_bag_name, "\t", parse_raw_sub_bags_seq(tmp_subbags_str_raw)
        result.add(
            Bag(
                name: tmp_bag_name, 
                subbags: parse_raw_sub_bags_seq(tmp_subbags_str_raw), 
                inner_bags_count: -1
            )
        )
    return

proc print_all_bags(list_of_bags: seq[Bag]) =
    for eachbag in list_of_bags:
        echo eachbag

proc get_sub_bags_from_A_bag(Astr: string): seq[tuple[sname:string, scout:int]] = 
    for each_bag in all_bags:
        if Astr == each_bag.name:
            return  each_bag.subbags

proc get_sub_bags_count_inside_Astr(Astr: string): int = 
    for each_bag in all_bags:
        if Astr == each_bag.name:
            return  each_bag.inner_bags_count

proc set_sub_bags_count_inside_Astr(Astr: string, counted_bag: int) = 
    for i, each_bag in all_bags:
        if Astr == each_bag.name:
            all_bags[i].inner_bags_count = counted_bag

proc string_times_int(astr: string, times: int): string = 
    for i in countup(0, times):
        result = result & astr

proc is_A_bag_contain_eventually_B_bag(Astr: string, Bstr: string, depth: int): bool =
    var sub_bags = get_sub_bags_from_A_bag(Astr)
    echo string_times_int("===", depth), " startingi to check ", Astr, " have ", Bstr, " inside =================>>>", sub_bags, " <<<<<<<<<<<<<<"

    for each_subbag_name_and_count  in sub_bags:
        echo string_times_int("===", depth), " ", each_subbag_name_and_count.sname 
        if each_subbag_name_and_count.sname == Bstr:
            echo string_times_int("===", depth), " ", Astr, " contains ", Bstr
            return true

    # walk thru each sub bags
    for each_subbag_name_and_count  in sub_bags:
        echo string_times_int("+++", depth), " ", each_subbag_name_and_count
        if is_A_bag_contain_eventually_B_bag(each_subbag_name_and_count.sname, Bstr, depth + 1):
            echo string_times_int("---", depth), " ", each_subbag_name_and_count.sname, " contains ", Bstr
            return true
        else:
            echo string_times_int("---", depth), " ", each_subbag_name_and_count.sname, " never contains ", Bstr
    
    return false

proc count_inner_bags_inside_A_bag(Astr: string, depth:int): int= 
    result = get_sub_bags_count_inside_Astr(Astr)
    if result  != -1: #meeans already counted it
        return

    # here mean not counted yet
    result = 0 

    var sub_bags = get_sub_bags_from_A_bag(Astr)
    if sub_bags.len() == 0:
        set_sub_bags_count_inside_Astr(Astr, 0)
        return 0 # means emtpy bag

    #result += sub_bags.len()

    echo string_times_int("===", depth), " startingi to check ", Astr, " have inside =======>>>", sub_bags, " <<<<<<<<<<<<<<"
    for each_subbag_name_and_count in sub_bags:
        echo string_times_int("---", depth), get_sub_bags_count_inside_Astr(each_subbag_name_and_count.sname)
        result += count_inner_bags_inside_A_bag(each_subbag_name_and_count.sname, 2) * each_subbag_name_and_count.scout + each_subbag_name_and_count.scout
        #         -----------------------------inside of subbag x  count of subbag------------------------------------  + subbag itself x count of that subbag=



all_bags = read_and_populate_bags(false)
print_all_bags(all_bags)

## var counting_countainment_of_shiny_gold = 0
#for eachbag in all_bags:
#    #counting_countainment_of_shiny_gold += int(is_A_bag_contain_eventually_B_bag(eachbag.name, "shiny gold", 0))
#    var tmp = count_inner_bags_inside_A_bag(eachbag.name, 0)
#    echo eachbag.inner_bags_count, tmp
#
echo count_inner_bags_inside_A_bag("shiny gold", 0)
#print_all_bags(all_bags)
