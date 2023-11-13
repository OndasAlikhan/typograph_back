package database

import (
	"errors"
	"fmt"
	"typograph_back/src/application"
	"typograph_back/src/model"

	"gorm.io/gorm"
)

func RunFixtures() {
	runRoles()
	runPermissions()
	assignPermissionsToAdmin()
	runParagraphs()
}

func runParagraphs() {
	fmt.Println("Starting runParahraphs")
	var count int64
	application.GlobalDB.Model(&model.Paragraph{}).Count(&count)

	fmt.Println("count", count)
	if count > 0 {
		fmt.Println("if count > 0")
		return
	}

	application.GlobalDB.Create(&Quotes)
}

func runRoles() {
	var roles = []model.Role{
		{Name: "Admin", Slug: "admin"},
		{Name: "User", Slug: "user"},
	}

	for _, role := range roles {
		var existingRole model.Role
		if err := application.GlobalDB.Where("slug = ?", role.Slug).First(&existingRole).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				application.GlobalDB.Create(&role)
			}
		}
	}
}

func runPermissions() {
	var permissions = []model.Permission{
		{Name: "Users index", Slug: "users-index"},
		{Name: "Roles index", Slug: "roles-index"},
		{Name: "Roles store", Slug: "roles-store"},
		{Name: "Roles show", Slug: "roles-show"},
		{Name: "Roles update", Slug: "roles-update"},
		{Name: "Roles delete", Slug: "roles-delete"},
		{Name: "Roles add permissions", Slug: "roles-add-permissions"},
		{Name: "Roles remove permissions", Slug: "roles-remove-permissions"},
		{Name: "Permissions index", Slug: "permissions-index"},
		{Name: "Permissions store", Slug: "permissions-store"},
		{Name: "Permissions show", Slug: "permissions-show"},
		{Name: "Permissions update", Slug: "permissions-update"},
		{Name: "Permissions delete", Slug: "permissions-delete"},
	}

	for _, permission := range permissions {
		var existingPermission model.Permission
		if err := application.GlobalDB.Where("slug = ?", permission.Slug).First(&existingPermission).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				application.GlobalDB.Create(&permission)
			}
		}
	}
}

func assignPermissionsToAdmin() {
	var adminRole model.Role
	if err := application.GlobalDB.Where("slug = ?", "admin").First(&adminRole).Error; err != nil {
		return
	}

	var permissions []model.Permission
	if err := application.GlobalDB.Find(&permissions).Error; err != nil {
		return
	}

	if err := application.GlobalDB.Model(&adminRole).Association("Permissions").Append(&permissions); err != nil {
		return
	}
}
