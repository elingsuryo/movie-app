package dto

type GetMovieByIDRequest struct {
	ID int64 `param:"id" validate:"required"` //json, param movie/ I=param, query ?id_pengguna=asdsasort
}

type CreateMovieRequest struct {
	Title       string `json:"title" validate:"required"` //json, param movie/ I=param, query ?id_pengguna=asdsasort
	Year        int64  `json:"year" validate:"required"`
	Director    string `json:"director" validate:"required"`
	Description string `json:"description"`
}

type UpdateMovieRequest struct {
	ID          int64  `param:"id" validate:"required"`
	Title       string `json:"title" validate:"required"` //json, param movie/ I=param, query ?id_pengguna=asdsasort
	Year        int64  `json:"year" validate:"required"`
	Director    string `json:"director" validate:"required"`
	Description string `json:"description"`
}