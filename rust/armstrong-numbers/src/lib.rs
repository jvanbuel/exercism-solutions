use std::convert::TryInto;

pub fn is_armstrong_number(num: u32) -> bool {
    let digits: Vec<u32> = num
        .to_string()
        .chars()
        .map(|d| d.to_digit(10).unwrap())
        .collect();

    let armstrong_check = digits
        .iter()
        .map(|d| d.pow(digits.len().try_into().unwrap()))
        .reduce(|x, y| x + y)
        .expect("Something went wrong while calculating the armstrong check!");
    return num == armstrong_check;
}
