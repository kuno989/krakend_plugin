package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Key struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId    string `json:"-"`
	CreatedAt string `bson:"created_at" json:"created_at"`
	ExpiresAt string `bson:"expires_at" json:"expires_at"`
	ApiKey    string `bson:"api_key" json:"api_key"`
}

type ResponseMessage struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
