package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type TaskHeader struct {
	ID                      *string        `bson:"id"`
	Name                    *string        `bson:"name"`
	Resume                  *string        `bson:"resume"`
	PlanningBeginDate       *time.Time     `bson:"planningBeginDate"`
	PlanningEndDate         *time.Time     `bson:"planningEndDate"`
	RealBeginDate           *time.Time     `bson:"realBeginDate"`
	RealEndDate             *time.Time     `bson:"realEndDate"`
	Status                  *string        `bson:"status"`
	PrimaryExecutorID       *string        `bson:"primaryExecutorId"`
	PrimaryExecutorObjectID *bson.ObjectId `bson:"primaryExecutorObjectId"`
	PrimaryExecutor         *string        `bson:"primaryExecutor"`
}

type Task struct {
	OID                     bson.ObjectId   `bson:"_id"`
	ID                      *string         `bson:"id"`
	Name                    *string         `bson:"name" `
	Resume                  *string         `bson:"resume" `
	Description             *string         `bson:"description" `
	CustomerContact         *string         `bson:"customerContact" `
	CreatedTime             *time.Time      `bson:"createdTime"`
	CreatorObjectID         *bson.ObjectId  `bson:"creatorObjectId"`
	CreatorID               *string         `bson:"creatorId"`
	PrimarySellerObjectID   *bson.ObjectId  `bson:"primarySellerObjectId"`
	PrimarySellerID         *string         `bson:"primarySellerId"`
	PrimaryOCObjectID       *bson.ObjectId  `bson:"primaryOCObjectId"`
	PrimaryOCID             *string         `bson:"primaryOCId"`
	PrimaryExecutorObjectID *bson.ObjectId  `bson:"primaryExecutorObjectId"`
	PrimaryExecutorID       *string         `bson:"primaryExecutorId"`
	OtherExecutorObjectIds  []bson.ObjectId `bson:"otherExecutorObjectIds"`
	OtherExecutorIDs        []string        `bson:"otherExecutorIds"`
	RequiringBeginDate      *time.Time      `bson:"requiringBeginDate"`
	RequiringEndDate        *time.Time      `bson:"requiringEndDate"`
	PlanningBeginDate       *time.Time      `bson:"planningBeginDate" `
	PlanningEndDate         *time.Time      `bson:"planningEndDate" `
	RealBeginDate           *time.Time      `bson:"realBeginDate" `
	RealEndDate             *time.Time      `bson:"realEndDate" `
	Percent                 *int            `bson:"percent"`
	Status                  *string         `bson:"status"`
	ParentProductObjectID   *bson.ObjectId  `bson:"parentProductObjectId" `
	ParentProductID         *string         `bson:"parentProductId"`
	ParentProjectObjectID   *bson.ObjectId  `bson:"parentProjectObjectId" `
	ParentProjectID         *string         `bson:"parentProjectId"`
}
type EmployeeName struct {
	Name *string `bson:"name"`
}

type ProductName struct {
	Name *string `bson:"name"`
}
type ProjectName struct {
	Name *string `bson:"name"`
}
type MaxID struct {
	ID *string `bson:"id"`
}

type TaskHeader_Get struct {
	ID                *string    `json:"id"`
	Name              *string    `json:"name"`
	Resume            *string    `json:"resume"`
	PlanningBeginDate *time.Time `json:"planningBeginDate"`
	PlanningEndDate   *time.Time `json:"planningEndDate"`
	RealBeginDate     *time.Time `json:"realBeginDate"`
	RealEndDate       *time.Time `json:"realEndDate"`
	Status            *string    `json:"status"`
	PrimaryExecutorID *string    `json:"primaryExecutorId"`
	PrimaryExecutor   *string    `json:"primaryExecutor"`
}

type Task_Get struct {
	ID                 *string    `json:"id"`
	Name               *string    `json:"name"`
	Resume             *string    `json:"resume"`
	Description        *string    `json:"description"`
	CustomerContact    *string    `json:"customerContact"`
	CreatedTime        *time.Time `json:"createdTime"`
	CreatorID          *string    `json:"creatorId"`
	Creator            *string    `json:"creator"`
	PrimarySellerID    *string    `json:"primarySellerId"`
	PrimarySeller      *string    `json:"primarySeller"`
	PrimaryOCID        *string    `json:"primaryOCId"`
	PrimaryOC          *string    `json:"primaryOC"`
	PrimaryExecutorID  *string    `json:"primaryExecutorId"`
	PrimaryExecutor    *string    `json:"primaryExecutor"`
	OtherExecutorIds   []string   `json:"otherExecutorIds"`
	OtherExecutors     []string   `json:"otherExecutors"`
	RequiringBeginDate *time.Time `json:"requiringBeginDate"`
	RequiringEndDate   *time.Time `json:"requiringEndDate"`
	PlanningBeginDate  *time.Time `json:"planningBeginDate"`
	PlanningEndDate    *time.Time `json:"planningEndDate"`
	RealBeginDate      *time.Time `json:"realBeginDate"`
	RealEndDate        *time.Time `json:"realEndDate"`
	Percent            *int       `json:"percent"`
	Status             *string    `json:"status"`
	ParentProductID    *string    `json:"parentProductId"`
	ParentProduct      *string    `json:"parentProduct"`
	ParentProjectID    *string    `json:"parentProjectId"`
	ParentProject      *string    `json:"parentProject"`
}

type Task_Post struct {
	Name                    *string         `json:"name" `
	Resume                  *string         `json:"resume" `
	Description             *string         `json:"description" `
	CustomerContact         *string         `json:"customerContact" `
	CreatedTime             *time.Time      `json:"createdTime"`
	CreatorObjectID         *bson.ObjectId  `json:"creatorObjectId"`
	CreatorID               *string         `json:"creatorId"`
	PrimarySellerObjectID   *bson.ObjectId  `json:"primarySellerObjectId"`
	PrimarySellerID         *string         `json:"primarySellerId"`
	PrimaryOCObjectID       *bson.ObjectId  `json:"primaryOCObjectId"`
	PrimaryOCID             *string         `json:"primaryOCId"`
	PrimaryExecutorObjectID *bson.ObjectId  `json:"primaryExecutorObjectId"`
	PrimaryExecutorID       *string         `json:"primaryExecutorId"`
	OtherExecutorObjectIds  []bson.ObjectId `json:"otherExecutorObjectIds"`
	OtherExecutorIDs        []string        `json:"otherExecutorIds"`
	RequiringBeginDate      *time.Time      `json:"requiringBeginDate"`
	RequiringEndDate        *time.Time      `json:"requiringEndDate"`
	PlanningBeginDate       *time.Time      `json:"planningBeginDate" `
	PlanningEndDate         *time.Time      `json:"planningEndDate" `
	RealBeginDate           *time.Time      `json:"realBeginDate" `
	RealEndDate             *time.Time      `json:"realEndDate" `
	Percent                 *int            `json:"percent"`
	Status                  *string         `json:"status"`
	ParentProductObjectID   *bson.ObjectId  `json:"parentProductObjectId" `
	ParentProductID         *string         `json:"parentProductId"`
	ParentProjectObjectID   *bson.ObjectId  `json:"parentProjectObjectId" `
	ParentProjectID         *string         `json:"parentProjectId"`
}
