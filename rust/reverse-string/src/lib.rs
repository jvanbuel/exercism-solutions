pub fn reverse(input: &str) -> String {
    match input.chars().count() {
        0 => String::new(),
        1 => input.to_string(),
        _ => {
            let index = input.char_indices().nth(1).unwrap().0;
            let mut reversed = reverse(&input[index..]);
            reversed.push(input.chars().nth(0).unwrap());
            return reversed;
        }
    }
}
