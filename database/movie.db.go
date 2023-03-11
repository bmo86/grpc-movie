package database

import (
	"grpc/models"
)

func (i *Instance_conn) CreteMovie(movie *models.Movie) (uint, error) {
	err := i.db.Create(&movie)
	return movie.ID, err.Error
}

func (i *Instance_conn) Getmovie(id int) (*models.Movie, error) {

	var movie models.Movie

	err := i.db.First(&movie, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}
