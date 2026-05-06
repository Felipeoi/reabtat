package inmates

import (
	"context"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type UseCase interface {
	ById(ctx context.Context, id int) (error, context.Context, *entity.InmatesResp)
	Create(ctx context.Context, in entity.InmatesFlat) (error, context.Context, *entity.InmatesResp)
	Delete(ctx context.Context, id int) (error, context.Context)
	List(ctx context.Context, limit, offset int) (error, context.Context, []entity.InmatesList)
	Update(ctx context.Context, in entity.InmatesFlat) (error, context.Context, *entity.InmatesResp)
}
