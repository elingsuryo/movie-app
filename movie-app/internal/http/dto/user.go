package dto

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}

type GetAllUserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

type GetUserByIDRequest struct {
	ID int64 `param:"id" validate:"required"` //json, param movie/ I=param, query ?id_pengguna=asdsasort
}

type UserUpdateRequest struct {
	ID       int64  `param:"id" validate:"required"`       //json, param movie/ I=param, query ?id_pengguna=asdsasort
	Username string `json:"username" validate:"required"`  //json, param movie/ I=param, query ?id_pengguna=asdsasort
	FullName string `json:"full_name" validate:"required"` //json, param movie/ I=param, query ?id_pengguna=asdsasort
	Role     string `json:"role" validate:"required"`      //json, param movie/ I=param, query ?id_pengguna=asdsasort
	Password string `json:"password" validate:"required"`  //json, param movie/ I=param, query ?id_pengguna=asdsasort
}

type UserCreateRequest struct {
	Username string `json:"username" validate:"required"`  //json, param movie/ I=param, query ?id_pengguna=asdsasort
	FullName string `json:"full_name" validate:"required"` //json, param movie/ I=param, query ?id_pengguna=asdsasort
	Role     string `json:"role" validate:"required"`      //json, param movie/ I=param, query ?id_pengguna=asdsasort
	Password string `json:"password" validate:"required"`  //json, param movie/ I=param, query ?id_pengguna=asdsasortS
}

type RequestResetPassword struct {
	Username string `json:"username" validate:"required"`
}

type ResetPasswordRequest struct {
	Token    string `param:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type VerifyEmailRequest struct {
	Token string `param:"token" validate:"required"`
}