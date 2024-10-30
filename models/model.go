package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  *string            `bson:"username,omitempty" json:"username,omitempty"`
	Email     *string            `bson:"email,omitempty" json:"email,omitempty"`
	Password  *string            `bson:"password,omitempty" json:"password,omitempty"`
	Token     *string            `bson:"token,omitempty" json:"token,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	UserID    string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
}

type Post struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Title       *string            `bson:"title,omitempty" json:"title,omitempty"`
	Content     *string            `bson:"content,omitempty" json:"content,omitempty"`
	Description *string            `bson:"description,omitempty" json:"description,omitempty"`
	Author      *string            `bson:"author,omitempty" json:"author,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	PostID      string             `bson:"post_id,omitempty" json:"post_id,omitempty"`
}

// future imporvements...
// type Comment struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
// 	PostID    primitive.ObjectID `bson:"post_id,omitempty" json:"post_id,omitempty"`
// 	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
// 	Content   string             `bson:"content,omitempty" json:"content,omitempty"`
// 	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
// 	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
// 	CommentID string             `bson:"comment_id,omitempty" json:"comment_id,omitempty"`
// }

// type Tag struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
// 	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
// 	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
// 	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
// 	TagID     string             `bson:"tag_id,omitempty" json:"tag_id,omitempty"`
// }

// type PostTag struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
// 	PostID    primitive.ObjectID `bson:"post_id,omitempty" json:"post_id,omitempty"`
// 	TagID     primitive.ObjectID `bson:"tag_id,omitempty" json:"tag_id,omitempty"`
// 	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
// 	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
// 	PostTagID string             `bson:"post_tag_id,omitempty" json:"post_tag_id,omitempty"`
// }
