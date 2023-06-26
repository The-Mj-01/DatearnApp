package interest

type InterestService struct {
	repo InterestRepositoryInterface
}

func NewInterestService(repo InterestRepositoryInterface) InterestServiceInterface {
	return &InterestService{
		repo: repo,
	}
}

func (s *InterestService) GetAllInterest(id *uint, name *string, limit *int, offset int) (*[]Interest, error) {
	interest := s.repo.GetAllInterest(id, name, limit, offset)
	if len(*interest) == 0 {
		return nil, InterestNotFound
	}

	return interest, nil
}

func (i InterestService) CreateInterest(name string) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}

func (i InterestService) UpdateInterest(id *uint, name string) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}

func (i InterestService) DeleteInterest(interestId *uint) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}
