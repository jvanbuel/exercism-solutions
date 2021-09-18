use chrono::Duration;

use chrono::{DateTime, Utc};

// Returns a Utc DateTime one billion seconds after start.
pub fn after(start: DateTime<Utc>) -> DateTime<Utc> {
    return start
        .checked_add_signed(Duration::seconds(1000000000))
        .unwrap();
}
