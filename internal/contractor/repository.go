package contractor

import "context"

type repository interface {
	Save(ctx context.Context, contractor Contractor) (Contractor, error)
	FindById(ctx context.Context, id int64) (Contractor, error)
	/*Update(ctx context.Context, contractor Contractor) error*/
}
