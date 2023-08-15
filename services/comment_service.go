package services

import (
	"errors"
	"strconv"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type CommentService struct {
}

// 获取所有未删除的评论
func (s *CommentService) GetVideoComment(video_id string) ([]models.Comment, error) {

	id, err := strconv.ParseInt(video_id, 10, 64)
	var Comments *[]models.Comment
	if err != nil {
		return *Comments, errors.New("视频不存在")
	}
	config.DB.Where("video_id = ? AND deleted_at is NULL", uint(id)).Find(&Comments)
	return *Comments, nil
}
