package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ProjectHeader struct {
	ID                     *string        `bson:"id"`
	Name                   *string        `bson:"name" `
	CreatorID              *string        `bson:"creatorId"`
	Status                 *string        `bson:"status"`
	PrimarySellerID        *string        `bson:"primarySellerId"`
	ProjectManagerObjectID *bson.ObjectId `bson:"projectManagerObjectId"`
	ProjectManagerID       *string        `bson:"projectManagerId"`
	ProductManagerID       *string        `bson:"productManagerId"`
	DevelopmentManagerID   *string        `bson:"developmentManagerId"`
	PlanningReleaseDate    *time.Time     `bson:"planningReleaseDate" `
	RealReleaseDate        *time.Time     `bson:"realReleaseDate" `
	RealAcceptanceDate     *time.Time     `bson:"realAcceptanceDate" `
}
type Project struct {
	OID                        bson.ObjectId  `bson:"_id"`
	ID                         *string        `bson:"id"`
	Name                       *string        `bson:"name" `
	Description                *string        `bson:"description" `
	CustomerContact            *string        `bson:"customerContact" `
	CreatedTime                *time.Time     `bson:"createdTime"`
	CreatorObjectID            *bson.ObjectId `bson:"creatorObjectId"`
	CreatorID                  *string        `bson:"creatorId"`
	PrimarySellerObjectID      *bson.ObjectId `bson:"primarySellerObjectId"`
	PrimarySellerID            *string        `bson:"primarySellerId"`
	RequiringAcceptanceDate    *time.Time     `bson:"requiringAcceptanceDate"`
	ProjectManagerObjectID     *bson.ObjectId `bson:"projectManagerObjectId"`
	ProjectManagerID           *string        `bson:"projectManagerId"`
	ProductManagerObjectID     *bson.ObjectId `bson:"productManagerObjectId"`
	ProductManagerID           *string        `bson:"productManagerId"`
	DevelopmentManagerObjectID *bson.ObjectId `bson:"developmentManagerObjectId"`
	DevelopmentManagerID       *string        `bson:"developmentManagerId"`
	// OtherExecutorObjectIds  []bson.ObjectId `bson:"otherExecutorObjectIds"`
	// OtherExecutorIDs        []string        `bson:"otherExecutorIds"`
	OtherExecutors          *string        `bson:"otherExecutors"`
	PlanningReleaseDate     *time.Time     `bson:"planningReleaseDate" `
	RealReleaseDate         *time.Time     `bson:"realReleaseDate" `
	RealAcceptanceDate      *time.Time     `bson:"realAcceptanceDate" `
	Status                  *string        `bson:"status"`
	RelevantProductObjectID *bson.ObjectId `bson:"relevantProductObjectId" `
	RelevantProductID       *string        `bson:"relevantProductId"`
}

type ProjectHeader_Get struct {
	ID                   *string    `json:"id"`
	Name                 *string    `json:"name" `
	CreatorID            *string    `json:"creatorId"`
	Status               *string    `json:"status"`
	PrimarySellerID      *string    `json:"primarySellerId"`
	ProjectManagerID     *string    `json:"projectManagerId"`
	ProjectManager       *string    `json:"projectManager"`
	ProductManagerID     *string    `json:"productManagerId"`
	DevelopmentManagerID *string    `json:"developmentManagerId"`
	PlanningReleaseDate  *time.Time `json:"planningReleaseDate" `
	RealReleaseDate      *time.Time `json:"realReleaseDate" `
	RealAcceptanceDate   *time.Time `json:"realAcceptanceDate" `
}

type Project_Get struct {
	ID                      *string    `json:"id"`
	Name                    *string    `json:"name" `
	Description             *string    `json:"description" `
	CustomerContact         *string    `json:"customerContact" `
	CreatedTime             *time.Time `json:"createdTime"`
	CreatorID               *string    `json:"creatorId"`
	Creator                 *string    `json:"creator"`
	PrimarySellerID         *string    `json:"primarySellerId"`
	PrimarySeller           *string    `json:"primarySeller"`
	RequiringAcceptanceDate *time.Time `json:"requiringAcceptanceDate"`
	ProjectManagerID        *string    `json:"projectManagerId"`
	ProjectManager          *string    `json:"projectManager"`
	ProductManagerID        *string    `json:"productManagerId"`
	ProductManager          *string    `json:"productManager"`
	DevelopmentManagerID    *string    `json:"developmentManagerId"`
	DevelopmentManager      *string    `json:"developmentManager"`
	// OtherExecutorIDs        []string        `json:"otherExecutorIds"`
	// OtherExecutors          []string        `json:"otherExecutors"`
	OtherExecutors      *string    `json:"otherExecutors"`
	PlanningReleaseDate *time.Time `json:"planningReleaseDate" `
	RealReleaseDate     *time.Time `json:"realReleaseDate" `
	RealAcceptanceDate  *time.Time `json:"realAcceptanceDate" `
	Status              *string    `json:"status"`
	RelevantProductID   *string    `json:"relevantProductId"`
	RelevantProduct     *string    `json:"relevantProduct"`
}

type Project_Post struct {
	ID                         *string        `json:"id"`
	Name                       *string        `json:"name" `
	Description                *string        `json:"description" `
	CustomerContact            *string        `json:"customerContact" `
	CreatedTime                *time.Time     `json:"createdTime"`
	CreatorObjectID            *bson.ObjectId `json:"creatorObjectId"`
	CreatorID                  *string        `json:"creatorId"`
	PrimarySellerObjectID      *bson.ObjectId `json:"primarySellerObjectId"`
	PrimarySellerID            *string        `json:"primarySellerId"`
	RequiringAcceptanceDate    *time.Time     `json:"requiringAcceptanceDate"`
	ProjectManagerObjectID     *bson.ObjectId `json:"projectManagerObjectId"`
	ProjectManagerID           *string        `json:"projectManagerId"`
	ProductManagerObjectID     *bson.ObjectId `json:"productManagerObjectId"`
	ProductManagerID           *string        `json:"productManagerId"`
	DevelopmentManagerObjectID *bson.ObjectId `json:"developmentManagerObjectId"`
	DevelopmentManagerID       *string        `json:"developmentManagerId"`
	// OtherExecutorObjectIds  []bson.ObjectId `json:"otherExecutorObjectIds"`
	// OtherExecutorIDs        []string        `json:"otherExecutorIds"`
	OtherExecutors          *string        `json:"otherExecutors"`
	PlanningReleaseDate     *time.Time     `json:"planningReleaseDate" `
	RealReleaseDate         *time.Time     `json:"realReleaseDate" `
	RealAcceptanceDate      *time.Time     `json:"realAcceptanceDate" `
	Status                  *string        `json:"status"`
	RelevantProductObjectID *bson.ObjectId `json:"relevantProductObjectId" `
	RelevantProductID       *string        `json:"relevantProductId"`
}
