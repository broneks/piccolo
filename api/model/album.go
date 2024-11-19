package model

import (
	"piccolo/api/helper"

	"github.com/jackc/pgx/v5/pgtype"
)

type Album struct {
	Id                 pgtype.Text        `json:"id"`
	UserId             pgtype.Text        `json:"userId"`
	Name               pgtype.Text        `json:"name"`
	Description        pgtype.Text        `json:"description"`
	CoverPhotoId       pgtype.Text        `json:"coverPhotoId"`
	IsShareLinkEnabled pgtype.Bool        `json:"isShareLinkEnabled"`
	ReadAccessHash     pgtype.Text        `json:"readAccessHash"`
	CreatedAt          pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt          pgtype.Timestamptz `json:"-"`
}

func (album *Album) SetReadAccessHash() {
	if album.IsShareLinkEnabled.Bool {
		album.ReadAccessHash = pgtype.Text{String: helper.GenerateRandomHash(), Valid: true}
	} else {
		album.ReadAccessHash = pgtype.Text{Valid: false} // nil
	}
}
