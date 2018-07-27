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

//RepoTable is the table name for repository
const RepoTable = "repository"

// RepoRecord holds the record of an repository in DB, all the infors are from the registry notification event.
type RepoRecord struct {
	RepositoryID int64     `gorm:"primary_key" json:"repository_id"`
	Name         string    `gorm:"column:name" json:"name"`
	ProjectID    int64     `gorm:"column:project_id"  json:"project_id"`
	Description  string    `gorm:"column:description" json:"description"`
	PullCount    int64     `gorm:"column:pull_count" json:"pull_count"`
	StarCount    int64     `gorm:"column:star_count" json:"star_count"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

//TableName is required by by beego orm to map RepoRecord to table repository
func (p *RepoRecord) TableName() string {
	return RepoTable
}

// RepositoryQuery : query parameters for repository
type RepositoryQuery struct {
	Name        string
	ProjectIDs  []int64
	ProjectName string
	LabelID     int64
	Pagination
}

func (p *RepoRecord) InsertOne() (error) {

	err := db.GetDB5().Create(p).Error

	return err
}
