package contractor

import "context"

type creatorRepository interface {
	Save(ctx context.Context, contractor Contractor) (Contractor, error)
}

type renamerRepository interface {
	FindByID(ctx context.Context, id int64) (Contractor, error)
	Update(ctx context.Context, contractor Contractor) (Contractor, error)
}
