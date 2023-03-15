// Provides functionality for CRUD operations

package model

import model "pdfPro/model/entity"

func CreateUserAccount(user *model.UserAccount) (ok bool, err error) {
	return createUserAccountImpl(user)
}

func createUserAccountImpl(user *model.UserAccount) (ok bool, err error) {
	result := dbGorm.Create(user)
	err = result.Error
	if err != nil {
		return false, err
	}
	return true, nil
}
