#[warn(unused_imports)]
extern crate party_worker_service;
// use dotenv::dotenv;
use std::{sync::Arc, thread};
use waitgroup::WaitGroup;

use mongodb::sync::Collection;
use party_worker_service::{
    datasource::{connect_meilisearch, connect_mongo_db, connect_nats},
    model::{Party},
    workers::{party_create, party_delete, party_join, party_unjoin, party_update},
};

#[tokio::main]
async fn main() -> std::io::Result<()> {
    // dotenv().ok().expect("env  is not initialized");
    let mongo_client = connect_mongo_db();
    let nc = connect_nats();
    let mc = Arc::new(connect_meilisearch());

    mc.health().await.unwrap();

    let party_coll: Collection<Party> = mongo_client.database("party").collection("party");

    let wg = WaitGroup::new();

    let cnc = nc.clone();
    let w = wg.worker();
    let c_party_coll = party_coll.clone();

    let nmc = mc.clone();

    thread::spawn(move || {
        party_create(cnc, c_party_coll, nmc);
        drop(w);
    });

    let cnc = nc.clone();
    let w = wg.worker();
    let c_party_coll = party_coll.clone();

    let nmc = mc.clone();

    thread::spawn(move || {
        party_update(cnc, c_party_coll, nmc);
        drop(w);
    });

    let w = wg.worker();
    let c_party_coll = party_coll.clone();
    let cnc = nc.clone();
    let nmc = mc.clone();

    thread::spawn(move || {
        party_delete(cnc, c_party_coll, nmc);
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

    println!("all workers started");

    wg.wait().await;

    Ok(())
}
