use meilisearch_sdk::client::Client as MC;
use mongodb::{bson::doc, sync::Client};
pub fn connect_mongo_db() -> Client {
    let mongo_host = dotenv::var("MONGO_HOST").expect("MONGO_HOST is not initialized");
    let mongo_username = dotenv::var("MONGO_USERNAME").expect("MONGO_USERNAME is not initialized");
    let mongo_password = dotenv::var("MONGO_PASSWORD").expect("MONGO_PASSWORD is not initialized");

    let client_opt = format!(
        "mongodb://{}:{}@{}",
        mongo_username, mongo_password, mongo_host
    );

    let client = Client::with_uri_str(client_opt).expect("uri str incorrect");

    client
        .database("admin")
        .run_command(doc! {"ping": 1}, None)
        .expect("ping command failed");

    println!("MongoDV Connected successfully.");

    client
}

pub fn connect_nats() -> nats::Connection {
    let nats_host = dotenv::var("NATS_HOST")
        .expect("NATS_HOST is not initialized")
        .to_owned();

    nats::connect(&nats_host[..]).expect("nats connection failed")
}

pub fn connect_meilisearch() -> MC {
    let mei_host = dotenv::var("MEILISEARCH_HOST")
        .expect("MEILISEARCH_HOST is not initialized")
        .to_owned();

    let mei_key = dotenv::var("MEILISEARCH_KEY")
        .expect("MEILISEARCH_KEY is not initialized")
        .to_owned();

    MC::new(mei_host, mei_key)
}
