package pkg

import (
	"fmt"
	//"sync"
	//"time"
	//"harbor_upgrade/src/config"
	"harbor_upgrade/src/models/models1"
	"harbor_upgrade/src/models/models5"
)

//用户表没有改变
func SyncUser() {
	fmt.Printf(" Starting SyncUser=============>>")
	m1_users := new(models1.User)

	users, err := m1_users.GetUsers()
	if err != nil {

		fmt.Errorf("m1_users.GetUser failed err:%v\n", err)

		return

	}

	m5_user := new(models5.User)
	for _, user := range users {
		m5_user.Username = user.Username
		m5_user.Password = user.Password
		m5_user.UserID = user.UserID
		m5_user.Deleted = user.Deleted
		m5_user.Email = user.Email
		m5_user.UpdateTime = user.UpdateTime
		m5_user.Salt = user.Salt
		m5_user.CreationTime = user.CreationTime
		m5_user.Comment = user.Comment
		m5_user.ResetUUID = user.ResetUUID
		m5_user.Realname = user.Realname
		m5_user.HasAdminRole = user.HasAdminRole

		err = m5_user.InsertOne(m5_user)
		if err != nil {
			fmt.Printf("m5_user.InsertOne failed err:%v\n", err)
			return
		}

	}

	fmt.Printf(" end Success SyncUser=============>>")

}

func SyncAccessLog() {
	fmt.Printf(" Starting SyncAccessLog=============>>")
	m1_accesslog := new(models1.AccessLog)

	accessLogs, err := m1_accesslog.GetAccessLogs()
	if err != nil {
		fmt.Errorf("m1_accesslog.GetAccessLogs failed err:%v\n", err)
		return

	}

	m5_accesslog := new(models5.AccessLog)
	for _, accesslog := range accessLogs {
		m5_accesslog.LogID = accesslog.LogID
		u, err := new(models1.User).GetUser(accesslog.UserID)
		if err != nil {
			fmt.Errorf("new(models1.User).GetUser failed err:%v\n", err)
			return
		}

		m5_accesslog.Username = u.Username
		m5_accesslog.ProjectID = accesslog.ProjectID
		m5_accesslog.RepoName = accesslog.RepoName
		m5_accesslog.RepoTag = accesslog.RepoTag
		m5_accesslog.GUID = accesslog.GUID
		m5_accesslog.Operation = accesslog.Operation
		m5_accesslog.OpTime = accesslog.OpTime

		err = m5_accesslog.InsertOne()
		if err != nil {
			fmt.Printf("m5_accesslog.InsertOne failed err:%v\n", err)
			return
		}

	}

	fmt.Printf(" end success SyncAccessLog=============>>")
}

//同步镜像
//public 1
func SyncProjectAndMetadata() {
	fmt.Printf(" Starting SyncProjectAndMetadata=============>>")
	m1_project := new(models1.Project)

	projects, err := m1_project.GetProject()
	if err != nil {
		fmt.Errorf("m1_project.GetProject failed err:%v\n", err)
		return
	}

	m5_project := new(models5.Project)
	for _, project := range projects {

		m5_project.Deleted = project.Deleted
		m5_project.Name = project.Name
		m5_project.ProjectID = project.ProjectID
		m5_project.OwnerID = project.OwnerID
		m5_project.CreationTime = project.UpdateTime
		m5_project.UpdateTime = project.UpdateTime
		err = m5_project.InsertOne()
		if err != nil {
			fmt.Printf("m5_project.InsertOne failed err:%v\n", err)
			return
		}

		m5_metadata := new(models5.ProjectMetadata)
		m5_metadata.ProjectID = project.ProjectID
		m5_metadata.Name = "public"
		m5_metadata.UpdateTime = project.UpdateTime
		m5_metadata.CreationTime = project.UpdateTime
		m5_metadata.Deleted = project.Deleted
		if project.Public == 1 {
			m5_metadata.Value = "true"
		} else {
			m5_metadata.Value = "false"
		}

		err = m5_metadata.InsertOne()
		if err != nil {
			fmt.Errorf("m5_metadata.InsertOne failed err:%v\n", err)
			return
		}

	}

	fmt.Printf(" end success SyncProjectAndMetadata=============>>")
}

func Syncproject_member() {
	fmt.Printf(" Starting Syncproject_member=============>>")
	m1_projectmember := new(models1.ProjectMember)

	proMembers, err := m1_projectmember.GetProjectMember()
	if err != nil {
		fmt.Errorf("m1_projectmember.GetProjectMember failed err:%v\n", err)
		return
	}

	m5_project := new(models5.Member)
	for _, project := range proMembers {

		m5_project.ProjectID = project.ProjectID

		m5_project.UpdateTime = project.CreationTime
		m5_project.CreationTime = project.CreationTime
		m5_project.EntityID = project.UserId
		m5_project.EntityType = "u"
		m5_project.Role = project.Role

		err = m5_project.InsertOne()
		if err != nil {
			fmt.Printf(" member m5_project.InsertOne failed err:%v\n", err)
			return
		}

	}

	fmt.Printf(" end success Syncproject_member=============>>")
}

func SyncRepo() {
	fmt.Printf(" Starting SyncRepo=============>>")
	m1_projectmember := new(models1.RepoRecord)

	proMembers, err := m1_projectmember.GetRepoRecord()
	if err != nil {
		fmt.Errorf("m1_projectmember.GetProjectMember failed err:%v\n", err)
		return
	}

	m5_project := new(models5.RepoRecord)
	for _, project := range proMembers {

		m5_project.ProjectID = project.ProjectID
		m5_project.RepositoryID = project.RepositoryID
		m5_project.Name = project.Name
		m5_project.Description = project.Description
		m5_project.PullCount = project.PullCount
		m5_project.StarCount = project.StarCount
		m5_project.CreationTime = project.CreationTime
		m5_project.UpdateTime = project.UpdateTime
		err = m5_project.InsertOne()
		if err != nil {
			fmt.Printf(" member m5_project.InsertOne failed err:%v\n", err)
			return
		}

	}
	fmt.Printf(" end success SyncRepo=============>>")
}
