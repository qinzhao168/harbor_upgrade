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
	"harbor_upgrade/src/db"
)

// Member holds the details of a member.
type Member struct {
	ProjectID    int64  `gorm:"column:project_id" json:"project_id"`
	EntityID     int    `gorm:"column:entity_id" json:"entity_id"`
	EntityType   string `gorm:"column:entity_type" json:"entity_type"`
	Role         int    `gorm:"column:role" json:"role"`
	CreationTime time.Time    `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time     `gorm:"column:update_time" json:"update_time"`
}

func (c *Member) TableName() string {
	return "project_member"
}

func (c *Member) GetMembers() ([]Member, error) {

	var members []Member

	err := db.GetDB5().Table(c.TableName()).Scan(&members).Error
	return members, err
}

func (c *Member) InsertOne() error {

	err := db.GetDB5().Create(&c).Error

	return err
}
