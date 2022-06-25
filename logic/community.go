package logic

import (
	"awesomeProject/dao/mysql"
	"awesomeProject/models"
)

//GetCommunityList 查找所有community
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
