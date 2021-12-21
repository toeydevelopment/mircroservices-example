extern crate party_worker_service;
use std::time::Instant;
use std::{str::FromStr, thread};
use waitgroup::WaitGroup;

use dotenv::dotenv;
use mongodb::bson::doc;
use mongodb::sync::Collection;
use party_worker_service::{
    datasource::{connect_mongo_db, connect_nats},
    model::Party,
    workers::{party_create, party_delete, party_join, party_unjoin, party_update},
};

#[tokio::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok().expect("env  is not initialized");
    let mongo_client = connect_mongo_db();
    let nc = connect_nats();

    let _ = mongo_client;

    println!("{}", nc.client_ip().unwrap());

    let party = Party::empty();

    let party_coll: Collection<Party> = mongo_client.database("party").collection("party");

    let wg = WaitGroup::new();

    let cnc = nc.clone();
    let w = wg.worker();
    
    let c_party_coll = party_coll.clone();
    thread::spawn(move || {
        party_create(cnc, c_party_coll);
        drop(w);
    });

    let cnc = nc.clone();
    let w = wg.worker();
    let c_party_coll = party_coll.clone();

    thread::spawn(move || {
        party_update(cnc, c_party_coll);
        drop(w);
    });

    let w = wg.worker();
    let c_party_coll = party_coll.clone();
    let cnc = nc.clone();

    thread::spawn(move || {
        party_delete(cnc, c_party_coll);
        drop(w);
    });

    let w = wg.worker();
    let c_party_coll = party_coll.clone();
    let cnc = nc.clone();

    thread::spawn(move || {
        party_join(cnc, c_party_coll);
        drop(w);
    });

    let w = wg.worker();
    let c_party_coll = party_coll.clone();
    let cnc = nc.clone();

    thread::spawn(move || {
        party_unjoin(cnc, c_party_coll);
        drop(w);
    });

    wg.wait().await;

    party_coll.insert_one(party, None).unwrap();

    Ok(())
}
