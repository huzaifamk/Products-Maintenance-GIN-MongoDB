package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Object struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Form_Fields []interface{}      `json:"form_fields" bson:"form_fields"`
}

type Form_Fields struct {
	Id            string        `json:"id,omitempty" bson:"id,omitempty"`
	Name          string        `json:"name,omitempty" bson:"name,omitempty"`
	Label         string        `json:"label,omitempty" bson:"label,omitempty"`
	Placeholder   string        `json:"placeholder,omitempty" bson:"placeholder,omitempty"`
	Step          string        `json:"step,omitempty" bson:"step,omitempty"`
	Validations   []Validations `json:"validations,omitempty" bson:"validations,omitempty"`
	Fields_Width  string        `json:"fields_width,omitempty" bson:"fields_width,omitempty"`
	Value         string        `json:"value,omitempty" bson:"value,omitempty"`
	Field_Confirm string        `json:"field_confirm,omitempty" bson:"field_confirm,omitempty"`
	Format        string        `json:"format,omitempty" bson:"format,omitempty"`
	Error_Message string        `json:"error_message,omitempty" bson:"error_message,omitempty"`
	Options       string        `json:"options,omitempty" bson:"options,omitempty"`
	Created_At    time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Updated_At    time.Time     `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Validations struct {
	Required          []Required   `json:"required,omitempty" bson:"required,omitempty"`
	Pattern           []Pattern    `json:"pattern,omitempty" bson:"pattern,omitempty"`
	Min_Length        []Min_Length `json:"min_length,omitempty" bson:"min_length,omitempty"`
	Max_Length        []Max_Length `json:"max_length,omitempty" bson:"max_length,omitempty"`
	Min_Date          string       `json:"min_date,omitempty" bson:"min_date,omitempty"`
	Max_Date          string       `json:"max_date,omitempty" bson:"max_date,omitempty"`
	Identity_Validate bool         `json:"identity_validate,omitempty" bson:"identity_validate,omitempty"`
}

type Required struct {
	Value   bool   `json:"value,omitempty" bson:"value,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Pattern struct {
	Value   string `json:"value,omitempty" bson:"value,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Min_Length struct {
	Value   int    `json:"value,omitempty" bson:"value,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Max_Length struct {
	Value   int    `json:"value,omitempty" bson:"value,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}
