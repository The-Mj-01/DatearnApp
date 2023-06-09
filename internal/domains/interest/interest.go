package interest

import "context"

type InterestRepositoryInterface interface {
	GetAllInterest(id *uint, name *string, limit *int, offset int) *[]Interest
	CreateInterest(name string) (*Interest, error)
	UpdateInterest(oldInterest, newInterest *Interest) (*Interest, error)
	DeleteInterest(interest *Interest) (*Interest, error)
}

type InterestServiceInterface interface {
	GetAllInterest(id *uint, name *string, limit *int, offset int) (*[]Interest, error)
	CreateInterest(name string) (*Interest, error)
	UpdateInterest(id *uint, name string) (*Interest, error)
	DeleteInterest(interestId *uint) (*Interest, error)
}

type InterestUseCaseInterface interface {
	GetAllInterest(ctx context.Context, token string, request *InterestGetRequest) (*[]Interest, error)
	CreateInterest(ctx context.Context, token string, request *InterestCreateRequest) (*Interest, error)
	UpdateInterest(ctx context.Context, token string, request *InterestUpdateRequest) (*Interest, error)
	DeleteInterest(ctx context.Context, token string, request *InterestDeleteRequest) (*Interest, error)
}
