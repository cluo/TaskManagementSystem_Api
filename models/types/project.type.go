package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ProjectHeader struct {
	ID                      *string        `bson:"id"`
	Name                    *string        `bson:"name"`
	PlanningBeginDate       *time.Time     `bson:"planningBeginDate"`
	PlanningEndDate         *time.Time     `bson:"planningEndDate"`
	RealBeginDate           *time.Time     `bson:"realBeginDate"`
	RealEndDate             *time.Time     `bson:"realEndDate"`
	Status                  *string        `bson:"status"`
	RefuseStatus            *string        `bson:"refuseStatus"`
	CreatorID               *string        `bson:"creatorId"`
	PrimarySellerID         *string        `bson:"primarySellerId"`
	PrimaryOCID             *string        `bson:"primaryOCId"`
	PrimaryExecutorID       *string        `bson:"primaryExecutorId"`
	PrimaryExecutorObjectID *bson.ObjectId `bson:"primaryExecutorObjectId"`
	PrimaryExecutor         *string        `bson:"primaryExecutor"`
}

type Project struct {
	OID                     bson.ObjectId  `bson:"_id"`
	ID                      *string        `bson:"id"`
	Name                    *string        `bson:"name" `
	Description             *string        `bson:"description" `
	CustomerContact         *string        `bson:"customerContact" `
	CreatedTime             *time.Time     `bson:"createdTime"`
	CreatorObjectID         *bson.ObjectId `bson:"creatorObjectId"`
	CreatorID               *string        `bson:"creatorId"`
	PrimarySellerObjectID   *bson.ObjectId `bson:"primarySellerObjectId"`
	PrimarySellerID         *string        `bson:"primarySellerId"`
	PrimaryOCObjectID       *bson.ObjectId `bson:"primaryOCObjectId"`
	PrimaryOCID             *string        `bson:"primaryOCId"`
	PrimaryExecutorObjectID *bson.ObjectId `bson:"primaryExecutorObjectId"`
	PrimaryExecutorID       *string        `bson:"primaryExecutorId"`
	// OtherExecutorObjectIds  []bson.ObjectId `bson:"otherExecutorObjectIds"`
	// OtherExecutorIDs        []string        `bson:"otherExecutorIds"`
	OtherExecutors        *string        `bson:"otherExecutors"`
	RequiringEndDate      *time.Time     `bson:"requiringEndDate"`
	PlanningBeginDate     *time.Time     `bson:"planningBeginDate" `
	PlanningEndDate       *time.Time     `bson:"planningEndDate" `
	RealBeginDate         *time.Time     `bson:"realBeginDate" `
	RealEndDate           *time.Time     `bson:"realEndDate" `
	Percent               *int           `bson:"percent"`
	Status                *string        `bson:"status"`
	RefuseStatus          *string        `bson:"refuseStatus"`
	ParentProductObjectID *bson.ObjectId `bson:"parentProductObjectId" `
	ParentProductID       *string        `bson:"parentProductId"`
}
type ProjectHeader_Get struct {
	ID                *string    `json:"id"`
	Name              *string    `json:"name"`
	PlanningBeginDate *time.Time `json:"planningBeginDate"`
	PlanningEndDate   *time.Time `json:"planningEndDate"`
	RealBeginDate     *time.Time `json:"realBeginDate"`
	RealEndDate       *time.Time `json:"realEndDate"`
	Status            *string    `json:"status"`
	RefuseStatus      *string    `json:"refuseStatus"`
	CreatorID         *string    `json:"creatorId"`
	PrimarySellerID   *string    `json:"primarySellerId"`
	PrimaryOCID       *string    `json:"primaryOCId"`
	PrimaryExecutorID *string    `json:"primaryExecutorId"`
	PrimaryExecutor   *string    `json:"primaryExecutor"`
}

type Project_Get struct {
	ID                *string    `json:"id"`
	Name              *string    `json:"name"`
	Description       *string    `json:"description"`
	CustomerContact   *string    `json:"customerContact"`
	CreatedTime       *time.Time `json:"createdTime"`
	CreatorID         *string    `json:"creatorId"`
	Creator           *string    `json:"creator"`
	PrimarySellerID   *string    `json:"primarySellerId"`
	PrimarySeller     *string    `json:"primarySeller"`
	PrimaryOCID       *string    `json:"primaryOCId"`
	PrimaryOC         *string    `json:"primaryOC"`
	PrimaryExecutorID *string    `json:"primaryExecutorId"`
	PrimaryExecutor   *string    `json:"primaryExecutor"`
	// OtherExecutorIds  []string   `json:"otherExecutorIds"`
	OtherExecutors    *string    `json:"otherExecutors"`
	RequiringEndDate  *time.Time `json:"requiringEndDate"`
	PlanningBeginDate *time.Time `json:"planningBeginDate"`
	PlanningEndDate   *time.Time `json:"planningEndDate"`
	RealBeginDate     *time.Time `json:"realBeginDate"`
	RealEndDate       *time.Time `json:"realEndDate"`
	Percent           *int       `json:"percent"`
	Status            *string    `json:"status"`
	RefuseStatus      *string    `json:"refuseStatus"`
	ParentProductID   *string    `json:"parentProductId"`
	ParentProduct     *string    `json:"parentProduct"`
}

type Project_Post struct {
	ID                      *string        `json:"id"`
	Name                    *string        `json:"name" `
	Description             *string        `json:"description" `
	CustomerContact         *string        `json:"customerContact" `
	CreatedTime             *time.Time     `json:"createdTime"`
	CreatorObjectID         *bson.ObjectId `json:"creatorObjectId"`
	CreatorID               *string        `json:"creatorId"`
	PrimarySellerObjectID   *bson.ObjectId `json:"primarySellerObjectId"`
	PrimarySellerID         *string        `json:"primarySellerId"`
	PrimaryOCObjectID       *bson.ObjectId `json:"primaryOCObjectId"`
	PrimaryOCID             *string        `json:"primaryOCId"`
	PrimaryExecutorObjectID *bson.ObjectId `json:"primaryExecutorObjectId"`
	PrimaryExecutorID       *string        `json:"primaryExecutorId"`
	// OtherExecutorObjectIds  []bson.ObjectId `json:"otherExecutorObjectIds"`
	// OtherExecutorIDs        []string        `json:"otherExecutorIds"`
	OtherExecutors        *string        `json:"otherExecutors"`
	RequiringEndDate      *time.Time     `json:"requiringEndDate"`
	PlanningBeginDate     *time.Time     `json:"planningBeginDate" `
	PlanningEndDate       *time.Time     `json:"planningEndDate" `
	RealBeginDate         *time.Time     `json:"realBeginDate" `
	RealEndDate           *time.Time     `json:"realEndDate" `
	Percent               *int           `json:"percent"`
	Status                *string        `json:"status"`
	RefuseStatus          *string        `json:"refuseStatus"`
	RefuseReason          *string        `json:"refuseReason"`
	ParentProductObjectID *bson.ObjectId `json:"parentProductObjectId" `
	ParentProductID       *string        `json:"parentProductId"`
}
