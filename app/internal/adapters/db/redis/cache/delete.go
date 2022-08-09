package cache

import (
	"context"
	"errors"
	"strconv"

	apperrors "balance-service/app/pkg/errors"
)

var ErrNoID = errors.New("no id")

func (c cacheStorage) DeleteBalance(ctx context.Context, ids ...int64) (err error) {
	client := c.client.WithContext(ctx)

	switch len(ids) {

	case 0:
		return apperrors.NewInternal(ErrNoID, "arguments")
	case 1:
		err = client.Del(
			strconv.FormatInt(ids[0], 10),
		).Err()
	case 2:
		err = client.Del(
			strconv.FormatInt(ids[0], 10),
			strconv.FormatInt(ids[1], 10),
		).Err()
	default:
		keys := make([]string, 0, len(ids))
		for _, id := range ids {
			keys = append(keys, strconv.FormatInt(id, 10))
		}
		err = client.Del(keys...).Err()
	}

	if err != nil {
		return apperrors.NewInternal(nil, "redis.Del")
	}

	return nil
}
