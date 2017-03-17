package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ProductName struct {
	ID   *string `bson:"id" json:"id"`
	Name *string `bson:"name" json:"name"`
}

type ProductHeader struct {
	ID                     *string        `bson:"id"`
	Name                   *string        `bson:"name"`
	PlanningReleaseDate    *time.Time     `bson:"planningReleaseDate" `
	RealReleaseDate        *time.Time     `bson:"realReleaseDate" `
	Status                 *string        `bson:"status"`
	CreatorID              *string        `bson:"creatorId"`
	ProductManagerObjectID *bson.ObjectId `bson:"productManagerObjectId"`
	ProductManagerID       *string        `bson:"productManagerId"`
}

type ProductHeader_Get struct {
	ID                  *string    `json:"id"`
	Name                *string    `json:"name"`
	PlanningReleaseDate *time.Time `json:"planningReleaseDate" `
	RealReleaseDate     *time.Time `json:"realReleaseDate" `
	Status              *string    `json:"status"`
	CreatorID           *string    `json:"creatorId"`
	ProductManagerID    *string    `json:"productManagerId"`
	ProductManager      *string    `json:"productManager"`
}

type Product struct {
	OID                        bson.ObjectId   `bson:"_id"`
	ID                         *string         `bson:"id"`
	Name                       *string         `bson:"name" `
	Description                *string         `bson:"description" `
	CreatedTime                *time.Time      `bson:"createdTime"`
	CreatorObjectID            *bson.ObjectId  `bson:"creatorObjectId"`
	CreatorID                  *string         `bson:"creatorId"`
	ProductManagerObjectID     *bson.ObjectId  `bson:"productManagerObjectId"`
	ProductManagerID           *string         `bson:"productManagerId"`
	MarketingManagerObjectID   *bson.ObjectId  `bson:"marketingManagerObjectId"`
	MarketingManagerID         *string         `bson:"marketingManagerId"`
	DevelopmentManagerObjectID *bson.ObjectId  `bson:"developmentManagerObjectId"`
	DevelopmentManagerID       *string         `bson:"developmentManagerId"`
	OtherExecutorObjectIDs     []bson.ObjectId `bson:"otherExecutorObjectIds"`
	OtherExecutorIDs           []string        `bson:"otherExecutorIds"`
	PlanningReleaseDate        *time.Time      `bson:"planningReleaseDate" `
	RealReleaseDate            *time.Time      `bson:"realReleaseDate" `
	Percent                    *int            `bson:"percent"`
	Status                     *string         `bson:"status"`
}

type Product_Get struct {
	OID                  bson.ObjectId `json:"_id"`
	ID                   *string       `json:"id"`
	Name                 *string       `json:"name" `
	Description          *string       `json:"description" `
	CreatedTime          *time.Time    `json:"createdTime"`
	CreatorID            *string       `json:"creatorId"`
	Creator              *string       `json:"creator"`
	ProductManagerID     *string       `json:"productManagerId"`
	ProductManager       *string       `json:"productManager"`
	MarketingManagerID   *string       `json:"marketingManagerId"`
	MarketingManager     *string       `json:"marketingManager"`
	DevelopmentManagerID *string       `json:"developmentManagerId"`
	DevelopmentManager   *string       `json:"developmentManager"`
	OtherExecutorIDs     []string      `json:"otherExecutorIds"`
	OtherExecutors       *string       `json:"otherExecutors"`
	PlanningReleaseDate  *time.Time    `json:"planningReleaseDate" `
	RealReleaseDate      *time.Time    `json:"realReleaseDate" `
	Percent              *int          `json:"percent"`
	Status               *string       `json:"status"`
}

type Product_Post struct {
	ID                         *string        `json:"id"`
	Name                       *string        `json:"name" `
	Description                *string        `json:"description" `
	CreatedTime                *time.Time     `json:"createdTime"`
	CreatorObjectID            *bson.ObjectId `json:"creatorObjectId"`
	CreatorID                  *string        `json:"creatorId"`
	ProductManagerObjectID     *bson.ObjectId `json:"productManagerObjectId"`
	ProductManagerID           *string        `json:"productManagerId"`
	MarketingManagerObjectID   *bson.ObjectId `json:"marketingManagerObjectId"`
	MarketingManagerID         *string        `json:"marketingManagerId"`
	DevelopmentManagerObjectID *bson.ObjectId `json:"developmentManagerObjectId"`
	DevelopmentManagerID       *string        `json:"developmentManagerId"`
	OtherExecutorIDs           []string       `json:"otherExecutorIds"`
	PlanningReleaseDate        *time.Time     `json:"planningReleaseDate" `
	RealReleaseDate            *time.Time     `json:"realReleaseDate" `
	Percent                    *int           `json:"percent"`
	Status                     *string        `json:"status"`
}
