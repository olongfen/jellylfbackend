package services

import (
	"context"

	"jellylfbackend/global"
	"jellylfbackend/model"
)

func GetPersonOneWithEmail(ctx context.Context, email string) (*model.Person, error) {
	var (
		person *model.Person
	)
	db := global.DB.WithContext(ctx).Model(&model.Person{}).Where("email = ?", email).First(&person)
	return person, db.Error
}
