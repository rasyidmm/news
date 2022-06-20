package user

type UserInputPort interface {
	UserCreate(interface{}) (interface{}, error)
	UserList(interface{}) (interface{}, error)
	UserGetById(interface{}) (interface{}, error)
	UserUpdate(interface{}) (interface{}, error)
	UserDelete(interface{}) (interface{}, error)
}

type UserOutputPort interface {
	UserCreateResponse(interface{}) (interface{}, error)
	UserListResponse(interface{}) (interface{}, error)
	UserGetByIdResponse(interface{}) (interface{}, error)
	UserUpdateResponse(interface{}) (interface{}, error)
	UserDeleteResponse(interface{}) (interface{}, error)
}
