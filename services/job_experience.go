package services

import (
	"context"

	"jellylfbackend/global"
	"jellylfbackend/model"
)

func GetJobExperienceWithPersonID(ctx context.Context, personID uint) ([]*model.JobExperience, error) {
	var (
		jobExperiences []*model.JobExperience
	)
	db := global.DB.WithContext(ctx).Model(&model.JobExperience{}).Where("person_id = ?", personID).Order("start_time desc").Find(&jobExperiences)
	return jobExperiences, db.Error

}
