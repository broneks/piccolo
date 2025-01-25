package photoservice

import (
	"context"
	"piccolo/api/consts"
	"piccolo/api/types"
)

func (svc *PhotoService) GetUserFileStorage(ctx context.Context, userId string) (*types.TotalFileStorage, error) {
	var err error

	usedMB, err := svc.photoRepo.GetUserFileStorage(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &types.TotalFileStorage{
		UsedMB:         usedMB,
		UsedPercentage: (usedMB / consts.TotalUserFileStorageLimitMB) * 100,
		TotalMB:        consts.TotalUserFileStorageLimitMB,
	}, nil
}
