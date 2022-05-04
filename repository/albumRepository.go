package repository

import "example/web-service-gin/domain"

type AlbumRepository struct {
	albums []domain.Album
}

func NewAlbumRepository() *AlbumRepository {
	var albums = []domain.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	repository := new(AlbumRepository)
	repository.albums = albums

	return repository
}

func (r *AlbumRepository) List() []domain.Album {
	return r.albums
}

func (r *AlbumRepository) Create(album domain.Album) interface{} {
	r.albums = append(r.albums, album)
	return album
}

func (r *AlbumRepository) Get(id string) *domain.Album {
	for _, a := range r.albums {
		if a.ID == id {
			return &a
		}
	}

	return nil
}
