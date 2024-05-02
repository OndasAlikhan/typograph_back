package service

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
	"typograph_back/src/repository/repository_interface"
	"typograph_back/src/service/service_interface"
)

type UserService struct {
	repository  repository_interface.UserRepositoryInterface
	roleService service_interface.RoleServiceInterface
	jwtService  service_interface.JWTServiceInterface
}

func NewUserService(repo repository_interface.UserRepositoryInterface) *UserService {
	return &UserService{
		repository:  repo,
		roleService: NewRoleService(),
		jwtService:  NewJWTService(),
	}
}

func (us *UserService) GetAll() ([]*model.User, error) {
	return us.repository.GetAll()
}

func (us *UserService) GetById(id uint) (*model.User, error) {
	return us.repository.GetById(id)
}

func (us *UserService) GetByIds(ids []uint) ([]*model.User, error) {
	return us.repository.GetByIds(ids)
}

func (us *UserService) GetByEmail(email string) (*model.User, error) {
	return us.repository.GetByEmail(email)
}

func (us *UserService) Create(request *dto.UserStoreRequest) (*model.User, error) {
	passwordHash, err := us.jwtService.GenerateHash(request.Password)
	if err != nil {
		return nil, err
	}

	role, err := us.roleService.GetBySlug("user")
	if err != nil {
		return nil, err
	}

	user := model.User{Email: request.Email, Name: request.Name, Password: passwordHash, RoleID: role.ID}

	return us.repository.Save(user)
}

func (us *UserService) UpdatePassword(id uint, request *dto.UserUpdatePasswordRequest) (*model.User, error) {
	user, err := us.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	user.Password, err = us.jwtService.GenerateHash(request.Password)
	if err != nil {
		return nil, err
	}

	return us.repository.Save(*user)
}

func (us *UserService) Delete(id uint) error {
	return us.repository.Delete(id)
}

func (us *UserService) HasPermission(id uint, permissionSlug string) bool {
	return us.repository.HasPermission(id, permissionSlug)
}
