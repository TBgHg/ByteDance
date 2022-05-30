package service

import (
	"ByteDance/cmd/favorite/repository"
	"ByteDance/utils"
)

type FavoriteRequest struct {
	UserId     int64
	Token      string
	VideoId    int64
	ActionType int32
}

func RelationAction(userId int32, videoId int32, actionType int32) (err error) {
	//更新 如果数据库没有该数据则返回IsExist = 0
	IsExist := repository.FavoriteDao.RelationUpdate(userId, videoId, actionType)

	if IsExist == 0 {
		//添加该数据
		err = repository.FavoriteDao.RelationCreate(userId, videoId)
		utils.CatchErr("添加失败", err)
	}

	return err
}
