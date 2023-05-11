package datamodels

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Category struct {
    ID                  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    CategoryName        string             `json:"categoryName,omitempty" bson:"categoryName,omitempty"`

}

type Comment struct {
    ID          primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
    PostId      primitive.ObjectID  `json:"postId,omitempty" bson:"postId,omitempty"`
    AuthorId    primitive.ObjectID  `json:"authorId,omitempty" bson:"authorId,omitempty"`
    Title       string             `json:"title,omitempty" bson:"title,omitempty"`
    Content     string              `json:"content,omitempty" bson:"content,omitempty"`
    Rate        string              `json:"rate,omitempty" bson:"rate,omitempty"`
    CreatedAt   time.Time           `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type Content struct {
    ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    AuthorID    primitive.ObjectID `json:"authorId,omitempty" bson:"authorId,omitempty"`
    Title       string             `json:"title,omitempty" bson:"title,omitempty"`
    Content     string             `json:"content,omitempty" bson:"content,omitempty"`
    CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type Post struct {
    ID          primitive.ObjectID      `json:"_id,omitempty" bson:"_id,omitempty"`
    AuthorID    primitive.ObjectID      `json:"authorId,omitempty" bson:"authorId,omitempty"`
    IsPublished bool                    `json:"isPublished,omitempty" bson:"isPublished,omitempty"`
    Title       string                  `json:"title,omitempty" bson:"title,omitempty"`
    Summary     string                  `json:"summary,omitempty" bson:"summary,omitempty"`
    CreatedAt   time.Time               `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
    Contents    []primitive.ObjectID    `json:"contents,omitempty" bson:"contents,omitempty"`
    Categories  []primitive.ObjectID    `json:"categories,omitempty" bson:"categories,omitempty"`
}

type User struct {
    ID          primitive.ObjectID      `json:"_id,omitempty" bson:"_id,omitempty"`
    FirstName   string                  `json:"firstName,omitempty" bson:"firstName,omitempty"`
    LastName    string                  `json:"lastName,omitempty" bson:"lastName,omitempty"`
    Username    string                  `json:"username,omitempty" bson:"username,omitempty"`
    Password    string                  `json:"password,omitempty" bson:"password,omitempty"`
    BirthDate   time.Time               `json:"birthDate,omitempty" bson:"birthDate,omitempty"`
    Email       string                  `json:"email,omitempty" bson:"email,omitempty"`
    CreatedAt   time.Time               `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

