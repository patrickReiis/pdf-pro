// Provides functionality for CRUD operations

package model

import (
	model "pdfPro/model/entity"
)

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

func DeleteUserAccountByEmail(email string) (ok bool, err error) {
	return deleteUserAccountByEmailImpl(email)
}

func deleteUserAccountByEmailImpl(email string) (ok bool, err error) {

	var user = &model.UserAccount{Email: email}

	result := dbGorm.Where(user).First(user)
	if err := result.Error; err != nil {
		return false, err
	}

	result = dbGorm.Unscoped().Delete(user)
	if err = result.Error; err != nil {
		return false, err
	}

	return true, nil

}
