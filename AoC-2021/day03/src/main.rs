use std::fs;

fn test_or_real() -> String {
    let mut input = String::new();
    println!("Enter: 0 for test, 1 for data");
    std::io::stdin()
        .read_line(&mut input)
        .expect("failed while getting stdin");

    let input: i32 = input.trim().parse().expect("can't convert to int");

    let mut inputpath = String::new();

    if input == 0 {
        inputpath.push_str("./data/test.txt");
    } else if input == 1 {
        inputpath.push_str("./data/data.txt");
    } else {
        println!("Bad int, enter 1 or 0");
    }

    inputpath
}

#[derive(Debug)]
struct Digit {
    // for single digit
    c0: usize, // counted 0s
    c1: usize, // counted 1s
}


fn part1() {
    let fname = test_or_real();
    let mut line_width: usize = 0;
    for (i, line) in fs::read_to_string(&fname).unwrap().lines().enumerate() {
        line_width = line.trim().len();
        println!("{i:04}, {line}, {line_width}");
        break;
    }

    // let C = [ Digit{count0:0, count1:0}; line_width];
    let mut vec_digit_count: Vec<Digit> = Vec::new();
    for _ in 0..line_width {
        vec_digit_count.push(Digit { c0: 0, c1: 0 })
    }

    for (i, line) in fs::read_to_string(&fname).unwrap().lines().enumerate() {
        print!("{i:04}, {line} === ");
        for (ith_ch, ch) in line.chars().enumerate() {
            print!("{}-", ch);
            if ch == '0' {
                vec_digit_count[ith_ch].c0 += 1;
            } else if ch == '1' {
                vec_digit_count[ith_ch].c1 += 1;
            } else {
                print!("shouldn't be different than 0 and 1");
                return;
            }
        }

        println!("{:?}", vec_digit_count);
    }

    //gamma_rate value (name is from the problem text)
    //epsilon_rate value (name is from the problem text)
    let mut gamma_rate: u32 = 0;
    let mut epsilon_rate: u32 = 0;

    for i in 0..line_width {
        if vec_digit_count[i].c1 > vec_digit_count[i].c0 {
            gamma_rate = gamma_rate | (1 << (line_width-i-1));
        } else if vec_digit_count[i].c1 < vec_digit_count[i].c0 {
            epsilon_rate = epsilon_rate | (1 << (line_width-i-1));
        } else {
            print!("number of 1s and 0s are equal???");
            return;
        }
    }
    println!("gamma  = {gamma_rate:15b} = {gamma_rate}");
    println!("epsilon= {epsilon_rate:15b} = {epsilon_rate}");
    let power_consumption = gamma_rate * epsilon_rate;
    println!("{power_consumption}");
}

fn txt_in_vec_out() -> Vec<String> {
    let fname = test_or_real();
    let mut line_width: usize;
    let mut vec_data: Vec<String> = Vec::new();

    for (i, line) in fs::read_to_string(&fname).unwrap().lines().enumerate() {
        line_width = line.trim().len();
        println!("{i:04}, {line}, {line_width}");
        vec_data.push(String::from(line.trim()));
    }

    vec_data
}

fn find_number_of_occurences(cur_list: Vec<String>, ith: usize) -> Digit {
    let mut ith_digit = Digit{ c0:0, c1:0 };
    for each_element in &cur_list {
        if each_element.as_bytes()[ith] == b'1' {
            ith_digit.c1 = ith_digit.c1 + 1
        } else if each_element.as_bytes()[ith] == b'0' {
            ith_digit.c0 = ith_digit.c0 + 1
        }
    } 
    ith_digit
}

fn remove(cur_list: Vec<String>, ith_char: usize, take_bin: u8) -> Vec<String> {
    let mut new_list: Vec<String> = Vec::new();

    for v in cur_list{
        if v.as_bytes()[ith_char] == take_bin {
            new_list.push(v);
        }
    }

    new_list
}


fn common_or_1(cur_list: Vec<String>, ith: usize) -> Vec<String> {
    let ith_digit: Digit = find_number_of_occurences(cur_list, ith);
    
    let new_list = Vec::new();
    if ith_digit.c0 > ith_digit.c1 { 
    } 
    else { }

    new_list
}


fn part2(){
    let vec_data = txt_in_vec_out();

}
















fn main() {
    part1();
    part2();
}

