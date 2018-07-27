// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models5

import (
	"time"
)

// WatchItem ...
type WatchItem struct {
	ID           int64     `gorm:"primary_key" json:"id"`
	PolicyID     int64     `gorm:"column:policy_id" json:"policy_id"`
	Namespace    string    `gorm:"column:namespace" json:"namespace"`
	OnDeletion   bool      `gorm:"column:on_deletion" json:"on_deletion"`
	OnPush       bool      `gorm:"column:on_push" json:"on_push"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

//TableName ...
func (w *WatchItem) TableName() string {
	return "replication_immediate_trigger"
}
