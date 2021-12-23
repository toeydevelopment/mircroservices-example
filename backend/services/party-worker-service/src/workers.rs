use futures::executor::block_on;
#[warn(unused_variables, unused_must_use)]
use meilisearch_sdk::client::Client;
use std::str::FromStr;
use std::sync::Arc;
use std::{io::ErrorKind, time::Duration};

use crate::model::{JoinOrUnJoinRequest, Party, PartyRequest, PartySearcher};
use bson::doc;
use mongodb::sync::Collection;

const SERVICE_NAME: &str = "PARTY_WORKER_SERVICE";

pub fn party_create(nc: nats::Connection, _coll: Collection<Party>, _mc: Arc<Client>) {
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
                            .map_err(|err| println!("insert data failed: {:?}", err))
                            .ok();

                        let idx = PartySearcher::from_party(party.clone());

                        block_on(async {
                            _mc.index("party")
                                .add_or_update(&[idx], Some("id"))
                                .await
                                .unwrap();
                        });

                        party.id.to_string()
                    })
                    .map_err(|err| {
                        println!("something went wrong decode message failed {:?}", err);
                        ""
                    })
                    .unwrap();

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

pub fn party_update(nc: nats::Connection, _coll: Collection<Party>, _mc: Arc<Client>) {
    let _update_sub = nc.queue_subscribe("party.update", SERVICE_NAME).unwrap();
    loop {
        match _update_sub.next_timeout(Duration::from_millis(500)) {
            Ok(msg) => {
                let data = msg.data.as_slice();

                let id = serde_json::from_slice(data)
                    .map(|r: PartyRequest| {
                        let id = &r.clone().id.unwrap().to_owned()[..];
                        let party = r.to_party();

                        let v = bson::to_document(&party).unwrap();

                        _coll
                            .update_one(
                                doc! {
                                    "_id": bson::oid::ObjectId::from_str(id).unwrap(),
                                },
                                v,
                                None,
                            )
                            .map_err(|err| println!("update data failed: {:?}", err))
                            .ok();

                        let idx = PartySearcher::from_party(party.clone());

                        block_on(async {
                            _mc.index("party")
                                .add_or_update(&[idx], Some("id"))
                                .await
                                .unwrap();
                        });

                        party.id.to_string()
                    })
                    .map_err(|err| {
                        println!("something went wrong decode message failed {:?}", err);
                        ""
                    })
                    .unwrap();

                msg.respond(id).unwrap();
            }

            Err(err) => {
                if err.kind() != ErrorKind::TimedOut {
                    break;
                }
            }
        };
    }

    _update_sub.close().unwrap();
}

pub fn party_delete(nc: nats::Connection, _coll: Collection<Party>, _mc: Arc<Client>) {
    let _delete_sub = nc.queue_subscribe("party.delete", SERVICE_NAME).unwrap();

    loop {
        match _delete_sub.next_timeout(Duration::from_millis(500)) {
            Ok(msg) => {
                let data = msg.data.as_slice();

                serde_json::from_slice(data)
                    .map(|r: String| {
                        // block_on(async {
                        //     _mc.index("party").delete_document(r.clone()).await.unwrap();
                        // });

                        _coll
                            .update_one(
                                doc! {
                                    "_id": bson::oid::ObjectId::from_str(&r[..]).unwrap(),
                                },
                                doc! {
                                    "created_at": bson::DateTime::now(),
                                },
                                None,
                            )
                            .map_err(|err| println!("update data failed: {:?}", err))
                            .ok();
                    })
                    .map_err(|err| {
                        println!("something went wrong decode message failed {:?}", err);
                        ""
                    })
                    .unwrap();

                msg.ack().ok();
            }

            Err(err) => {
                if err.kind() != ErrorKind::TimedOut {
                    break;
                }
            }
        };
    }

    _delete_sub.close().unwrap();
}

pub fn party_join(nc: nats::Connection, _coll: Collection<Party>) {
    let _join_sub = nc.queue_subscribe("party.join", SERVICE_NAME).unwrap();
    loop {
        match _join_sub.next_timeout(Duration::from_millis(500)) {
            Ok(msg) => {
                let data = msg.data.as_slice();

                serde_json::from_slice(data)
                    .map(|r: JoinOrUnJoinRequest| {
                        let id = &r.id[..];

                        _coll
                            .update_one(
                                doc! {
                                    "_id": bson::oid::ObjectId::from_str(id).unwrap(),
                                },
                                doc! {
                                    "$addToSet": {
                                        "joined": r.user_email
                                    }
                                },
                                None,
                            )
                            .map_err(|err| println!("join failed: {:?}", err))
                            .ok();
                    })
                    .map_err(|err| {
                        println!("something went wrong decode message failed {:?}", err);
                        ""
                    })
                    .unwrap();

                msg.ack().ok();
            }

            Err(err) => {
                if err.kind() != ErrorKind::TimedOut {
                    break;
                }
            }
        };
    }

    _join_sub.close().unwrap();
}

pub fn party_unjoin(nc: nats::Connection, _coll: Collection<Party>) {
    let _unjoin_sub = nc.queue_subscribe("party.unjoin", SERVICE_NAME).unwrap();
    loop {
        match _unjoin_sub.next_timeout(Duration::from_millis(500)) {
            Ok(msg) => {
                let data = msg.data.as_slice();

                serde_json::from_slice(data)
                    .map(|r: JoinOrUnJoinRequest| {
                        let id = &r.id[..];

                        _coll
                            .update_one(
                                doc! {
                                    "_id": bson::oid::ObjectId::from_str(id).unwrap(),

                                },
                                doc! {
                                    "$pull": {
                                        "joined": r.user_email
                                    }
                                },
                                None,
                            )
                            .map_err(|err| println!("join failed: {:?}", err))
                            .ok();
                    })
                    .map_err(|err| {
                        println!("something went wrong decode message failed {:?}", err);
                        ""
                    })
                    .unwrap();

                msg.ack().ok();
            }

            Err(err) => {
                if err.kind() != ErrorKind::TimedOut {
                    break;
                }
            }
        };
    }
    _unjoin_sub.close().unwrap();
}
