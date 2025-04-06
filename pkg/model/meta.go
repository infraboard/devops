package model

import "time"

func NewMeta() *Meta {
	return &Meta{
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

type Meta struct {
	Id       string    `json:"id" bson:"_id"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	CreateBy string    `json:"create_by" bson:"create_by"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
	UpdateBy string    `json:"update_by" bson:"update_by"`
}

func (m *Meta) SetId(id string) *Meta {
	m.Id = id
	return m
}
