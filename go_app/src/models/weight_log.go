package models

import (
	time "main/src/time"
)

type (
	// WeightLog 定义原始的数据库字段
	WeightLog struct {
		ID        int64          `json:"id,omitempty" db:"id" valid:"-"`
		Weight    float64        `json:"weight,omitempty" db:"weight" valid:"-"`
		UserID    int64          `json:"user_id,omitempty" db:"user_id" valid:"-"`
		CreatedAt time.TimeStamp `json:"created_at,omitempty" db:"created_at" valid:"-"`
		UpdatedAt time.TimeStamp `json:"updated_at,omitempty" db:"updated_at" valid:"-"`
	}
)
