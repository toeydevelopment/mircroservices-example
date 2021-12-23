use meilisearch_sdk::document::Document;
use serde::{Deserialize, Serialize};
use std::str::FromStr;

#[derive(Serialize, Deserialize, Debug)]
pub struct PartySearcher {
    pub id: String,
    pub name: String,
    pub description: Option<String>,
}

impl Document for PartySearcher {
    type UIDType = String;
    fn get_uid(&self) -> &Self::UIDType {
        &self.id
    }
}

impl PartySearcher {
    pub fn from_party(p: Party) -> PartySearcher {
        PartySearcher {
            id: p.id.to_hex(),
            name: p.name,
            description: p.description,
        }
    }
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct JoinOrUnJoinRequest {
    pub id: String,
    pub user_email: String,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct PartyRequest {
    pub id: Option<String>,
    pub name: Option<String>,
    pub user_email: Option<String>,
    pub image_path: Option<String>,
    pub description: Option<String>,
    pub seat_limit: Option<i64>,
    pub seat: Option<i64>,
    pub created_at: Option<i64>,
    pub updated_at: Option<i64>,
    pub deleted_at: Option<i64>,
}

impl PartyRequest {
    pub fn to_party(self) -> Party {
        let mut p = Party::empty();
        if self.id.is_some() {
            let id = &self.id.unwrap()[..];
            p.id = bson::oid::ObjectId::from_str(id).unwrap();
        }
        if self.user_email.is_some() {
            p.owner = self.user_email.unwrap();
        }

        if self.name.is_some() {
            p.name = self.name.unwrap();
        }
        
        p.description = self.description;
        p.image_path = self.image_path;
        p.seat = self.seat;
        p.seat_limit = self.seat_limit;
        if self.created_at.is_some() {
            p.created_at = mongodb::bson::DateTime::from_millis(self.created_at.unwrap());
        }
        p.updated_at = self
            .updated_at
            .map(|f| mongodb::bson::DateTime::from_millis(f));
        p.deleted_at = self
            .deleted_at
            .map(|f| mongodb::bson::DateTime::from_millis(f));
        p
    }
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Party {
    #[serde(rename = "_id")]
    pub id: bson::oid::ObjectId,
    pub name: String,
    pub owner: String,
    pub joined: Vec<String>,
    pub image_path: Option<String>,
    pub description: Option<String>,
    pub seat_limit: Option<i64>,
    pub seat: Option<i64>,
    pub created_at: mongodb::bson::DateTime,
    pub updated_at: Option<mongodb::bson::DateTime>,
    pub deleted_at: Option<mongodb::bson::DateTime>,
}

impl Party {
    pub fn empty() -> Party {
        Party {
            id: bson::oid::ObjectId::new(),
            name: "".to_string(),
            owner: "".to_string(),
            joined: Vec::new(),
            seat: None,
            seat_limit: None,
            description: None,
            image_path: None,
            deleted_at: None,
            updated_at: None,
            created_at: bson::DateTime::now(),
        }
    }
}
