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

// UserGroupTable is the name of table in DB that holds the user object
const UserGroupTable = "user_group"

// UserGroup ...
type UserGroup struct {
	ID          int    `gorm:"primary_key" json:"id,omitempty"`
	GroupName   string `gorm:"column:group_name" json:"group_name,omitempty"`
	GroupType   int    `gorm:"column:group_type" json:"group_type,omitempty"`
	LdapGroupDN string `gorm:"column:ldap_group_dn" json:"ldap_group_dn,omitempty"`
}

// TableName ...
func (u *UserGroup) TableName() string {
	return UserGroupTable
}
