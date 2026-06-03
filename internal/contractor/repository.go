package contractor

import "context"

type repository interface {
	Save(ctx context.Context, contractor Contractor) (Contractor, error)
}
