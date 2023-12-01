fn main() {
    let input: &str = include_str!("../../../input/day1.txt");

    let lines = input.split("\n\n");

    let mut parsed: Vec<u32> = lines
        .map(|line: &str| {
            line.split("\n")
                .flat_map(|num: &str| num.parse::<u32>())
                .sum()
        })
        .collect();

    parsed.sort_by(|a, b| b.cmp(a));

    println!("Part 1: {}", parsed[0]);
    println!("Part 2: {}", parsed[..3].iter().sum::<u32>());
}
