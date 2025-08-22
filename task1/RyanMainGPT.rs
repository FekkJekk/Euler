fn main() {
    let mut z = 0;

    for x in 1..1000 {
        if x % 3 == 0 || x % 5 == 0 {
            z += x;
        }
    }

    println!("{}", z);
}
