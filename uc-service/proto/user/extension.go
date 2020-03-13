package go_micro_srv_uc

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (uc *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
