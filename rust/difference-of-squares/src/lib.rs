pub fn square_of_sum(n: u32) -> u32 {
    return { 1..=n }
        .reduce(|x, y| x + y)
        .expect("Something went wrong!")
        .pow(2);
}

pub fn sum_of_squares(n: u32) -> u32 {
    return { 1..=n }
        .reduce(|x, y| x + y.pow(2))
        .expect("Something went wrong!");
}

pub fn difference(n: u32) -> u32 {
    return square_of_sum(n) - sum_of_squares(n);
}
