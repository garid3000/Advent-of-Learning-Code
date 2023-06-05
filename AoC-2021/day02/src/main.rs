use std::fs;

fn test_or_real() -> String {
    let mut input = String::new();
    println!("Enter: 0 for test, 1 for data");
    std::io::stdin()
        .read_line(&mut input)
        .expect("Shouldn't fail at stdin");
    let input: i32 = input.trim().parse().expect("not int");

    let mut inputpath = String::new();
    if input == 0 {
        inputpath.push_str("./data/test.txt");
    } else if input == 1 {
        inputpath.push_str("./data/input.txt");
    } else {
        print!("Bad int, enter 1 or 0");
    }

    inputpath
}

#[derive(Debug)]
struct Coord {
    x: i32,
    y: i32,
}



fn main() {
    let datafile = test_or_real();

    let mut pos = Coord { x: 0, y: 0 };

    for line in fs::read_to_string(datafile).unwrap().lines() {
        //let value: i32 = line.trim().parse().expect("Failed while converting to int");
        let parts = line.split(' ');

        let mut dir: i8 = -1;
        for (i, s) in parts.enumerate() {
            print!("{i} => {s} \t| ");
            if i == 0 {
                if s == "forward" {
                    dir = 0;
                } else if s == "down" {
                    dir = 1; 
                } else if s == "up" {
                    dir = 2;
                } else {
                    dir = -1;
                }
            } else if i == 1 {
                let value: i32 = s.trim().parse().expect("Failed while converting to int");
                if dir == 0 {
                    pos.x = pos.x + value;
                } else if dir == 1 {
                    pos.y = pos.y + value;
                } else if dir == 2 {
                    pos.y = pos.y - value;
                } else {
                    println!("Dir shouldn't get == {dir}");
                }
                print!("pos {:?}", pos);
            } else {
                println!("SHouldn't be here");
                return;
            }
        }
        println!("");
    }
    println!(
        "final {:?}, mult = {}", 
        pos, pos.x * pos.y
        );
}
