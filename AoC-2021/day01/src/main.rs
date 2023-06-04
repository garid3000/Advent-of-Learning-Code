// use std::env;
use std::fs;

fn main() {
    let mut guess = String::new();
    println!("Press enter");
    std::io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    let filepath = "./src/input.txt";
    //let filepath = "./src/test.txt";

    let mut oldervalue: i32 = -999;
    let mut counter: u32 = 0;

    for line in fs::read_to_string(filepath).unwrap().lines() {
        let value: i32 = line.trim().parse().expect("Failed while converting to int");

        if oldervalue == -999 {
            println!("{value} (N/A - no previous measurmerement)");
        } else if oldervalue < value {
            println!("{value} (increased)");
            counter = counter + 1;
        } else if oldervalue > value {
            println!("{value} (decreased)");
        } else {
            println!("{value} (no change)");
        }

        oldervalue = value;
    }

    println!("Number of incremented measurements {counter}");

    // part 2: moving average with with 3 entry window
    let mut cur_moving_sum: i32;
    let mut old_moving_sum: i32 = -999;

    let file_string = fs::read_to_string(filepath).unwrap();
    //let lines = filestr.lines();
    
    let mut window = [0, 0, 0];
    counter = 0;

    for (i, line) in file_string.lines().enumerate() {
        let value: i32 = line.trim().parse().expect("Failed while converting to int");
        //println!("{i}, {line}");
        
        window = shift(window, value);
        if i == 0 || i == 1 {
            continue;
        }
        
        cur_moving_sum = sum3arr(window);

        if old_moving_sum == -999 {
            println!("{cur_moving_sum} (N/A - no previous measurmerement)");
        } else if old_moving_sum < cur_moving_sum {
            println!("{cur_moving_sum} (increased)");
            counter = counter + 1;
        } else if old_moving_sum > cur_moving_sum {
            println!("{cur_moving_sum} (decreased)");
        } else {
            println!("{cur_moving_sum} (no change)");
        }

        old_moving_sum = cur_moving_sum;
    }

    println!("Number of incremented measurements {counter}");
}


fn shift(array: [i32; 3], newmember: i32) -> [i32; 3] {
    let mut newarr: [i32; 3] = [0,0,0];
    newarr[0] = array[1];
    newarr[1] = array[2];
    newarr[2] = newmember;
    newarr
}

fn sum3arr(array: [i32; 3]) -> i32{
    array[0] + array[1] + array[2]
}
