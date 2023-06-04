fn main() {
    let mut input = String::new();
    println!("Enter: 0 for test, 1 for data");
    std::io::stdin()
        .read_line(&mut input)
        .expect("Shouldn't fail at stdin");
    let input: i32 = input.trim().parse().expect("not int");

    let mut inputpath = String::new();
    if input == 0 {
        inputpath.push_str("./src/test.txt")
    } else if input == 1 {
        inputpath.push_str("./src/test.txt")
    } else {
        print!("Bad int, enter 1 or 0");
        return;
    }
}
