import std/rdstdin
import std/strutils
import std/sequtils

#proc calc_simple_expr(simple_exp: string): int=
#    # calculates expression that has no parenthesis
#    var splitted_exp = simple_exp.split(" ")
#    assert splitted_exp.len() mod 2 == 1
#    result = parseInt(splitted_exp[0])
#
#    for i in countup(1, splitted_exp.len()-1, 2):
#        # echo splitted_exp[i], splitted_exp[i+1]
#        assert splitted_exp[i] in @["+", "*"]
#        if splitted_exp[i] == "+":
#            result += parseInt(splitted_exp[i+1])
#        elif splitted_exp[i] == "*":
#            result *= parseInt(splitted_exp[i+1])
#    return

proc calc_simple_expr(simple_exp: string): int=
    # calculates expression that has no parenthesis
    var splitted_exp = simple_exp.split(" ")

    # 1st the + operations
    while "+" in splitted_exp:
        for i,each_element in splitted_exp:
            if each_element == "+":
                var tmp = parseInt(splitted_exp[i-1]) +  parseInt(splitted_exp[i+1])
                splitted_exp[i-1] = $tmp
                splitted_exp.delete(i..i+1)
                break

    while "*" in splitted_exp:
        for i,each_element in splitted_exp:
            if each_element == "*":
                var tmp = parseInt(splitted_exp[i-1]) *  parseInt(splitted_exp[i+1])
                splitted_exp[i-1] = $tmp
                splitted_exp.delete(i..i+1)
                break
    assert splitted_exp.len() == 1
    result = parseInt(splitted_exp[0])



proc find_1st_range_of_max_depth(depmask:seq[int], maxdepth:int): tuple[i0:int, i1:int]=
    var is_cur_max_depth = false
    for i, each_depth in depmask:
        if (is_cur_max_depth == false) and (each_depth == maxdepth):
            result.i0 = i
            is_cur_max_depth = true
        if (is_cur_max_depth == true) and (each_depth != maxdepth):
            result.i1 = i - 1
            return
        if i == depmask.len() - 1:
            result.i1 = i
            return

proc calculate_depth_mask(input_expr: string): tuple[depmask:seq[int], maxdepth: int] = 
    # reead input_expr and outputs maximum depth
    result.depmask = newSeq[0](input_expr.len())

    var cur_depth: int = 0
    #var max_depth: int = 0

    # calculate depth mask
    for i, each_char in input_expr:
        if each_char == '(':
            cur_depth += 1
        result.depmask[i] = cur_depth
        if each_char == ')':
            cur_depth -= 1

        # calculating max depth
        if result.maxdepth < cur_depth:
            result.maxdepth = cur_depth

proc simplify_expression(input_expr: string): string=
    #echo "=> ", input_expr, "|\t" #, (range_i0, range_i1), "|\t", input_expr[(range_i0+1)..(range_i1-1)]
    var (exp_depthmask, exp_maxdepth) = calculate_depth_mask(input_expr)
    #var (start_index_maxdepht, end_index_maxdepht) = find_1st_range_of_max_depth(exp_maxdepth, exp_depthmask)
    if exp_maxdepth == 0:
        result = $(calc_simple_expr(input_expr))
        return


    var (range_i0, range_i1) = find_1st_range_of_max_depth(exp_depthmask, exp_maxdepth)

    # echo (range_i0, range_i1)
    var simplified_string_max_depth = simplify_expression(input_expr[(range_i0+1)..(range_i1-1)])
    result = input_expr[0..(range_i0-1)] & simplified_string_max_depth & input_expr[(range_i1+1)..^1]
    return

proc finish_simplyfying(input_expr:string): string = 
    result = input_expr
    while true:
        var (_, exp_maxdepth) = calculate_depth_mask(result)
        if exp_maxdepth == 0:
            result = $calc_simple_expr(result)
            break
        result = simplify_expression(result)
        # echo result


proc main=
    var input_line: string
    var sum_after_eval = 0
    while true:
        var ret = readLineFromStdin("", input_line)
        if not ret:
            break
        
        # var tmp = parse_line(input_line.replace(" ", ""))
        # var tmp = parse_line(input_line)
        #var tmp = calculate_depth_mask(input_line)
        #echo  input_line, "\t", tmp
        var eval = parseInt(finish_simplyfying(input_line))
        echo eval
        sum_after_eval += eval
    
    echo "sum = ", sum_after_eval

main()
