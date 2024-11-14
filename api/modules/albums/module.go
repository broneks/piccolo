package albums

import (
	"piccolo/api/repo"
	"piccolo/api/types"
)

type AlbumsModule struct {
	server    *types.Server
	albumRepo *repo.AlbumRepo
}

func New(server *types.Server, albumRepo *repo.AlbumRepo) *AlbumsModule {
	return &AlbumsModule{
		server:    server,
		albumRepo: albumRepo,
	}
}
