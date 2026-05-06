package inmates

import (
	"context"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Repository interface {
	ById(ctx context.Context, id int, userID int64) (error, context.Context, *entity.InmatesResp)
	Create(ctx context.Context, in entity.InmatesFlat) (error, context.Context, int)
	Delete(ctx context.Context, id int, userID int64) (error, context.Context)
	List(ctx context.Context, limit, offset int, userID int64) (error, context.Context, []entity.InmatesList)
	Update(ctx context.Context, in entity.InmatesFlat, userID int64) (error, context.Context, int)
}
