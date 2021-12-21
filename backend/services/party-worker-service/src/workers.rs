#[warn(unused_variables, unused_must_use)]
use std::{io::ErrorKind, time::Duration};

use crate::model::{Party, PartyRequest};
use mongodb::sync::Collection;

const SERVICE_NAME: &str = "PARTY_WORKER_SERVICE";

pub fn party_create(nc: nats::Connection, _coll: Collection<Party>) {
    let create_sub = nc.queue_subscribe("party.create", SERVICE_NAME).unwrap();

    loop {
        match create_sub.next_timeout(Duration::from_millis(500)) {
            Ok(msg) => {
                let data = msg.data.as_slice();

                let id = serde_json::from_slice(data)
                    .map(|r: PartyRequest| {
                        let party = r.to_party();

                        _coll
                            .insert_one(&party, None)
                            .map_err(|err| println!("insert data failed: {:?}", err)).ok();

                        party.id.to_string()
                    })
                    .map_err(|err| {
                        println!("something went wrong decode message failed {:?}", err);
                        ""
                    }).unwrap();

                msg.respond(id).unwrap();
            }

            Err(err) => {
                if err.kind() != ErrorKind::TimedOut {
                    break;
                }
            }
        };
    }

    create_sub.close().unwrap();
}

pub fn party_update(nc: nats::Connection, _coll: Collection<Party>) {
    let _update_sub = nc.queue_subscribe("party.update", SERVICE_NAME).unwrap();
}

pub fn party_delete(nc: nats::Connection, _coll: Collection<Party>) {
    let _delete_sub = nc.queue_subscribe("party.delete", SERVICE_NAME).unwrap();
}

pub fn party_join(nc: nats::Connection, _coll: Collection<Party>) {
    let _join_sub = nc.queue_subscribe("party.join", SERVICE_NAME).unwrap();
}

pub fn party_unjoin(nc: nats::Connection, _coll: Collection<Party>) {
    let _unjoin_sub = nc.queue_subscribe("party.unjoin", SERVICE_NAME).unwrap();
}
