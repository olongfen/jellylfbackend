package services

import (
	"context"
	"jellylfbackend/global"
	"jellylfbackend/model"
)

func GetWorksWithPersonID(ctx context.Context, personID uint) ([]*model.Work, error) {
	var (
		works []*model.Work
	)
	db := global.DB.WithContext(ctx).Model(&model.Work{}).Where("person_id = ?", personID).Order("created_at desc").Find(&works)
	return works, db.Error
}
