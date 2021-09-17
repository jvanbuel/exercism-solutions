// This stub file contains items which aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut counter = HashMap::new();
    for key in magazine {
        *counter.entry(key).or_insert(0) += 1
    }

    for key in note {
        *counter.entry(key).or_insert(0) -= 1
    }

    counter.retain(|_, v| *v < 0);

    return counter.len() == 0;
}
