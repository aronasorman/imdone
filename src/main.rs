use envconfig::Envconfig;

#[derive(Envconfig)]
struct Config {
    #[envconfig(from = "PUSHOVER_API_TOKEN")]
    pub pushover_api_token: String,
    #[envconfig(from = "PUSHOVER_USER_KEY")]
    pub pushover_user_key: String,
}

fn main() {
    let config = Config::init_from_env().unwrap();
    let client = reqwest::blocking::Client::new();
    let resp = client
        .post("https://api.pushover.net/1/messages.json")
        .query(&[
            ("token", config.pushover_api_token),
            ("user", config.pushover_user_key),
            ("priority", "1".to_string()),
            ("message", "Your task is done!".to_string()),
        ])
        .send()
        .unwrap();

    println!("Message sent. Response: {:?}", resp.text());
}
