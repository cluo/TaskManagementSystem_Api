package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type TaskHeader struct {
	ID                      *string        `bson:"id" json:"id"`
	Name                    *string        `bson:"name" json:"name"`
	Resume                  *string        `bson:"resume" json:"resume"`
	PlanningBeginDate       *time.Time     `bson:"planningBeginDate" json:"planningBeginDate"`
	PlanningEndDate         *time.Time     `bson:"planningEndDate" json:"planningEndDate"`
	RealBeginDate           *time.Time     `bson:"realBeginDate" json:"realBeginDate"`
	RealEndDate             *time.Time     `bson:"realEndDate" json:"realEndDate"`
	Status                  *string        `bson:"status" json:"status"`
	PrimaryExecutorObjectID *bson.ObjectId `bson:"primaryExecutorObjectId" json:"primaryExecutorObjectId"`
	PrimaryExecutor         *string        `bson:"primaryExecutor" json:"primaryExecutor"`
}

type Task struct {
	ID     *string `bson:"id" json:"id"`
	Name   *string `bson:"name" json:"name"`
	Resume *string `bson:"resume" json:"resume"`

	Description     *string        `bson:"description" json:"description"`
	CustomerContact *string        `bson:"customerContact" json:"customerContact"`
	CreatedTime     *time.Time     `bson:"createdTime" json:"createdTime"`
	CreatorObjectID *bson.ObjectId `bson:"creatorObjectId" json:"creatorObjectId"`
	Creator         *string        `bson:"creator" json:"creator"`

	PrimarySellerObjectID *bson.ObjectId `bson:"primarySellerObjectId" json:"primarySellerObjectId"`
	PrimarySeller         *string        `bson:"primarySeller" json:"primarySeller"`

	PrimaryOCObjectID       *bson.ObjectId  `bson:"primaryOCObjectId" json:"primaryOCObjectId"`
	PrimaryOC               *string         `bson:"primaryOC" json:"primaryOC"`
	PrimaryExecutorObjectID *bson.ObjectId  `bson:"primaryExecutorObjectId" json:"primaryExecutorObjectId"`
	PrimaryExecutor         *string         `bson:"primaryExecutor" json:"primaryExecutor"`
	OtherExecutorObjectIds  []bson.ObjectId `bson:"otherExecutorObjectIds" json:"otherExecutorObjectIds"`
	OtherExecutors          []string        `bson:"otherExecutors" json:"otherExecutors"`

	RequiringBeginDate    *time.Time     `bson:"requiringBeginDate" json:"requiringBeginDate"`
	RequiringEndDate      *time.Time     `bson:"requiringEndDate" json:"requiringEndDate"`
	PlanningBeginDate     *time.Time     `bson:"planningBeginDate" json:"planningBeginDate"`
	PlanningEndDate       *time.Time     `bson:"planningEndDate" json:"planningEndDate"`
	RealBeginDate         *time.Time     `bson:"realBeginDate" json:"realBeginDate"`
	RealEndDate           *time.Time     `bson:"realEndDate" json:"realEndDate"`
	Percent               *int           `bson:"percent" json:"percent"`
	Status                *string        `bson:"status" json:"status"`
	ParentProductObjectID *bson.ObjectId `bson:"parentProductObjectId" json:"parentProductObjectId"`
	ParentProduct         *string        `bson:"parentProduct" json:"parentProduct"`
	ParentProjectObjectID *bson.ObjectId `bson:"parentProjectObjectId" json:"parentProjectObjectId"`
	ParentProject         *string        `bson:"parentProject" json:"parentProject"`
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
