package db

import (
	"sync"

	"pgjbz.dev/restfull-api/models"
)

type database struct {
	mu sync.Mutex
	albums []models.Album
	currentId uint64
}

var db database

func GetAlbums() []models.Album {
	return db.albums
}

func AddAlbum(album *models.Album) {
	db.mu.Lock()
	defer db.mu.Unlock()
	album.ID = db.currentId
	db.currentId++
	db.albums = append(db.albums, *album)
}

func init() {
	db = database{albums: []models.Album{}, currentId: 1}
	db.albums = append(db.albums, models.Album{ID: db.currentId, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})
	db.currentId++
	db.albums = append(db.albums, models.Album{ID: db.currentId, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99})
	db.currentId++
	db.albums = append(db.albums, models.Album{ID: db.currentId, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price:39.99})
	db.currentId++
}
