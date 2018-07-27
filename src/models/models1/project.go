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

//	ApiProtocol       string    `gorm:"column:api_protocol"`
// Project holds the details of a project.
type Project struct {
	ProjectID    int64     `gorm:"primary_key" json:"project_id"`
	OwnerID      int       `gorm:"column:owner_id" json:"owner_id"`
	Name         string    `gorm:"column:name" json:"name"`
	Deleted      int       `gorm:"column:deleted" json:"deleted"`
	Public       int    `gorm:"column:public" json:"public"` //1公有的 0私有的
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

func (rp *Project) TableName() string {
	return "project"
}

func (c Project) GetProject() ([]Project, error) {

	var projects []Project
	err := db.GetDB1().Table(c.TableName()).Find(&projects, " project_id not in(1)").Error

	//err := db.GetDB1().Table(c.TableName()).Scan(&projects).Error
	return projects, err
}

// ProjectSorter holds an array of projects
type ProjectSorter struct {
	Projects []Project
}

// Len returns the length of array in ProjectSorter
func (ps *ProjectSorter) Len() int {
	return len(ps.Projects)
}

// Less defines the comparison rules of project
func (ps *ProjectSorter) Less(i, j int) bool {
	return ps.Projects[i].Name < ps.Projects[j].Name
}

// Swap swaps the position of i and j
func (ps *ProjectSorter) Swap(i, j int) {
	ps.Projects[i], ps.Projects[j] = ps.Projects[j], ps.Projects[i]
}

type ProjectMember struct {
	ProjectID    int64     `gorm:"column:project_id" json:"project_id"`
	UserId       int       `gorm:"column:user_id" json:"user_id"`
	Role         int       `gorm:"column:role" json:"role"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

func (rp *ProjectMember) TableName() string {
	return "project_member"
}

func (rp *ProjectMember) GetProjectMember() ([]ProjectMember, error) {

	var projectMembers []ProjectMember

	err := db.GetDB1().Table(rp.TableName()).Find(&projectMembers, " project_id not in(1)").Error
	return projectMembers, err
}
