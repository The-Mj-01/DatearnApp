package interest

type InterestService struct {
	repo InterestRepositoryInterface
}

func NewInterestService(repo InterestRepositoryInterface) InterestServiceInterface {
	return &InterestService{
		repo: repo,
	}
}

func (i InterestService) GetAllInterest(id *uint, name *string, limit *int, offset int) (*[]Interest, error) {
	//TODO implement me
	panic("implement me")
}

func (i InterestService) CreateInterest(name string) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}

func (i InterestService) UpdateInterest(id *uint, name string) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}

func (i InterestService) DeleteInterest(socialId *uint) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}
