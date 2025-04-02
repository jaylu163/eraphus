package manager

import (
	"context"
	"github.com/jaylu163/eraphus/internal/hades/logging"
)

func GetHotList(ctx context.Context, cid string) {

	logging.WithFor(ctx).Infof("GetHotList: cid %v cid value:%v", "cid", cid)
}
