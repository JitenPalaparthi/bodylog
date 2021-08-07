package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BodyLog struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty" mapstructure:"_id"`
	UserId       primitive.ObjectID `json:"userId" bson:"_userId" mapstructure:"_userId"`
	LogDateTime  string             `json:"logDateTime" bson:"logDateTime" mapstructure:"logDateTime"`
	Image        string             `json:"image" bson:"image" mapstructure:"image"`
	ImageX       string             `json:"imageX" bson:"imageX" mapstructure:"imageX"`
	ImageY       string             `json:"imageY" bson:"imageY" mapstructure:"imageY"`
	BodyPoints   []BodyPoint        `json:"bodyPoints" bson:"bodyPoints" mapstructure:"bodyPoints"`
	EnteredBy    string             `json:"enteredBy" bson:"enteredBy" mapstructure:"enteredBy"`
	Status       string             `json:"status" bson:"status" mapstructure:"status"`
	LastModified string             `json:"lastModified" bson:"lastModified" mapstructure:"lastModified"`
}

type BodyPoint struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty" mapstructure:"_id"`
	OffSetX        string             `json:"offSetX" bson:"offSetX" mapstructure:"offSetX"`
	OffSetY        string             `json:"offSetY" bson:"offSetY" mapstructure:"offSetY"`
	ProblemType    string             `json:"problemType" bson:"problemType" mapstructure:"problemType"`
	Description    string             `json:"description" bson:"description" mapstructure:"description"`
	LogOn          string             `json:"logOn" bson:"logOn" mapstructure:"logOn"`
	ProblemStartOn string             `json:"problemStartOn" bson:"problemStartOn" mapstructure:"problemStartOn"`
	ProblemEndOn   string             `json:"problemEndOn" bson:"problemEndOn" mapstructure:"problemEndOn"`
	Status         string             `json:"status" bson:"status" mapstructure:"status"`
	LastModified   string             `json:"lastModified" bson:"lastModified" mapstructure:"lastModified"`
}
