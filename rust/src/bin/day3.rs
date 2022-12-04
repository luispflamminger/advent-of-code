// UNFINISHED

fn main() {
    let input = include_str!("../../../input/day3.txt");

    let res: Vec<char> = input
        .lines()
        .map(|line| {
            let (comp1, comp2) = line.split_at(line.len() / 2);
            return comp1.chars().filter(|&char| comp2.contains(char));
        })
        .collect();

    println!("First Part: {:?}", res)
}
