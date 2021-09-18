#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut stack = Vec::new();
    for input in inputs {
        match input {
            &CalculatorInput::Value(n) => stack.push(n),
            other => {
                let two_values = (stack.pop(), stack.pop());
                match two_values {
                    (None, _) => return None,
                    (_, None) => return None,
                    (Some(y), Some(x)) => match other {
                        &CalculatorInput::Add => stack.push(x + y),
                        &CalculatorInput::Subtract => stack.push(x - y),
                        &CalculatorInput::Multiply => stack.push(x * y),
                        &CalculatorInput::Divide => stack.push(x / y),
                        _ => return None,
                    },
                };
            }
        };
    }
    if stack.len() != 1 {
        return None;
    }
    return stack.pop();
}
