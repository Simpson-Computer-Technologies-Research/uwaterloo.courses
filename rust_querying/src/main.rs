use std::{time::{SystemTime, UNIX_EPOCH}, fs};

fn main() {
	// Parse the json data
	let data: String = fs::read_to_string("../default_data.json").expect("Unable to read file");
    let json: serde_json::Value = serde_json::from_str(&data).unwrap();

	// Create a new result array
	let mut res: String = "".to_string();

	// Get the start time
	let start: u128 = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_nanos();
	if let Some(courses) = json.as_array() {
		courses.iter().for_each(|course| {
			if let Some(desc) = course["description"].as_str() {
				if desc.to_lowercase().contains("computer") {
					let course: String = serde_json::to_string(&course).unwrap();
					res.push_str(course.as_str());
				}
			}
		});
	}

	// Print the end time and result
	let res: String = res.replace("\\", "");
	let end: u128 = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_nanos();
	let end_time: u128 = end - start;
	println!("{}", end_time);
}