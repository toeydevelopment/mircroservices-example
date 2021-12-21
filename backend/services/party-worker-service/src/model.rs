use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct PartyRequest {
    pub name: String,
    pub image_path: Option<String>,
    pub description: Option<String>,
    pub seat_limit: Option<i64>,
    pub seat: Option<i64>,
    pub created_at: i64,
    pub updated_at: Option<i64>,
    pub deleted_at: Option<i64>,
}

impl PartyRequest {
    pub fn to_party(self) -> Party {
        let mut p = Party::empty();
        p.name = self.name;
        p.description = self.description;
        p.image_path = self.image_path;
        p.seat = self.seat;
        p.seat_limit = self.seat_limit;
        p.created_at = mongodb::bson::DateTime::from_millis(self.created_at);
        p.updated_at = self.updated_at.map(|f| mongodb::bson::DateTime::from_millis(f));
        p.deleted_at = self.deleted_at.map(|f| mongodb::bson::DateTime::from_millis(f));
        p
    }
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Party {
    #[serde(rename = "_id")]
    pub id: bson::oid::ObjectId,
    pub name: String,
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
