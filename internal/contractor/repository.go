package contractor

import "context"

type saverRepository interface {
	Save(ctx context.Context, contractor Contractor) (Contractor, error)
}

type getterRepository interface {
	FindByID(ctx context.Context, id int64) (Contractor, error)
}

type updaterRepository interface {
	Update(ctx context.Context, contractor Contractor) (Contractor, error)
}

type renamerRepository interface {
	getterRepository
	updaterRepository
}
