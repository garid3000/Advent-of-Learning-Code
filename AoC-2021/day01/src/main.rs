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
        let value: i32 = line.trim()
            .parse()
            .expect("Failed while converting to int");

        if oldervalue == -999 {
            println!("{value} (N/A - no previous measurmerement)");
        } else if oldervalue < value{
            println!("{value} (increased)");
            counter = counter + 1;
        } else if oldervalue > value{
            println!("{value} (decreased)");
        } else {
            println!("Shouldn't be here");
        }

        oldervalue = value;

    }

    println!("Number of incremented measurements {counter}");
}
