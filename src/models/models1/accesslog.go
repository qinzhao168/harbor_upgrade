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

package models1

import (
	"time"
	"harbor_upgrade/src/db"
)

//http://gorm.io/docs/create.html

//	ApiProtocol       string    `gorm:"column:api_protocol"`
// AccessLog holds information about logs which are used to record the actions that user take to the resourses.
type AccessLog struct {
	LogID     int       `gorm:"primary_key" json:"log_id"`
	UserID    int       `gorm:"column:user_id"  json:"user_id"`
	ProjectID int64     `gorm:"column:project_id"  json:"project_id"`
	RepoName  string    `gorm:"column:repo_name" json:"repo_name"`
	RepoTag   string    `gorm:"column:repo_tag" json:"repo_tag"`
	GUID      string    `gorm:"column:GUID"  json:"guid"`
	Operation string    `gorm:"column:operation" json:"operation"`
	OpTime    time.Time `gorm:"column:op_time" json:"op_time"`
}

func (c *AccessLog) TableName() string {
	return "access_log"
}

func (c *AccessLog) GetAccessLogs() ([]AccessLog, error) {

	var accesslogs []AccessLog

	err := db.GetDB1().Table(c.TableName()).Scan(&accesslogs).Error
	return accesslogs, err
}
