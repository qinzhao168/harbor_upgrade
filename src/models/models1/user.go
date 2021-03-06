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

// User holds the details of a user.
type User struct {
	UserID       int    `gorm:"primary_key" json:"user_id"`
	Username     string `gorm:"column:username" json:"username"`
	Email        string `gorm:"column:email" json:"email"`
	Password     string `gorm:"column:password" json:"password"`
	Realname     string `gorm:"column:realname" json:"realname"`
	Comment      string `gorm:"column:comment" json:"comment"`
	Deleted      int    `gorm:"column:deleted" json:"deleted"`
	HasAdminRole int       `gorm:"column:sysadmin_flag" json:"has_admin_role"`
	ResetUUID    string    `gorm:"column:reset_uuid" json:"reset_uuid"`
	Salt         string    `gorm:"column:salt:" json:"-"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

func (c *User) TableName() string {
	return "user"
}

func (c *User) GetUsers() ([]User, error) {

	var users []User

	err := db.GetDB1().Table(c.TableName()).Find(&users, " user_id not in(1,2)").Error
	
	return users, err
}

func (c *User) GetUser(userID int) (User, error) {
	var user User
	err := db.GetDB1().Where("user_id = ?", userID).First(&user).Error
	return user, err
}
