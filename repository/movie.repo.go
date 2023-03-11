package repository

import "grpc/models"

type Repo interface {
	CreteMovie(movie *models.Movie) (uint, error)
	Getmovie(id int) (*models.Movie, error)
}

var u Repo

func SetRepo(r *Repo) {
	u = *r
}

func CreteMovie(movie *models.Movie) (uint, error) {
	return u.CreteMovie(movie)
}

func Getmovie(id int) (*models.Movie, error) {
	return u.Getmovie(id)
}
