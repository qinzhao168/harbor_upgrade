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
)

// RepPolicy is the model for a replication policy, which associate to a project and a target (destination)
type RepPolicy struct {
	ID           int64  `gorm:"column:id" json:"id"`
	ProjectID    int64  `gorm:"column:project_id" json:"project_id"`
	TargetID     int64  `gorm:"column:target_id" json:"target_id"`
	Name         string `gorm:"column:name" json:"name"`
	Enabled      int       `gorm:"column:enabled" json:"enabled"`
	Description  string    `gorm:"column:description" json:"description"`
	CronStr      string    `gorm:"column:cron_str" json:"cron_str"`
	StartTime    time.Time `gorm:"column:start_time" json:"start_time"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	Deleted      int       `gorm:"column:deleted" json:"deleted"`
}

// RepJob is the model for a replication job, which is the execution unit on job service, currently it is used to transfer/remove
// a repository to/from a remote registry instance.
type RepJob struct {
	ID           int64    `gorm:"column:id" json:"id"`
	Status       string   `gorm:"column:status" json:"status"`
	Repository   string   `gorm:"column:repository" json:"repository"`
	SrcRepo      string   `gorm:"column:src_repo" json:"src_repo"`
	PolicyID     int64    `gorm:"column:policy_id" json:"policy_id"`
	Operation    string   `gorm:"column:operation" json:"operation"`
	Tags         string   `gorm:"column:tags" json:"-"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

// RepTarget is the model for a replication targe, i.e. destination, which wraps the endpoint URL and username/password of a remote registry.
type RepTarget struct {
	ID           int64     `gorm:"column:id" json:"id"`
	URL          string    `gorm:"column:url" json:"endpoint"`
	Name         string    `gorm:"column:name" json:"name"`
	Username     string    `gorm:"column:username" json:"username"`
	Password     string    `gorm:"column:password" json:"password"`
	Type         int       `gorm:"column:target_type" json:"type"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

//TableName is required by by beego orm to map RepTarget to table replication_target
func (r *RepTarget) TableName() string {
	return "replication_target"
}

//TableName is required by by beego orm to map RepJob to table replication_job
func (r *RepJob) TableName() string {
	return "replication_job"
}

//TableName is required by by beego orm to map RepPolicy to table replication_policy
func (r *RepPolicy) TableName() string {
	return "replication_policy"
}
