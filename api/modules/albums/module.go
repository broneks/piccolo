package albums

import (
	"piccolo/api/repo"
	"piccolo/api/shared"
)

type AlbumsModule struct {
	server    *shared.Server
	albumRepo *repo.AlbumRepo
}

func New(server *shared.Server, albumRepo *repo.AlbumRepo) *AlbumsModule {
	return &AlbumsModule{
		server:    server,
		albumRepo: albumRepo,
	}
}
