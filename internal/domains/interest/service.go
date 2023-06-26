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

func (s *InterestService) CreateInterest(name string) (*Interest, error) {
	if name == "" {
		return nil, NameNotFound
	}

	return s.repo.CreateInterest(name)
}

func (s *InterestService) UpdateInterest(id *uint, name string) (*Interest, error) {
	if name == "" {
		return nil, NameNotFound
	}
	newInterest := &Interest{
		Name: name,
	}

	interest, err := s.GetAllInterest(id, nil, nil, 0)

	if err != nil {
		return nil, InterestNotFound
	}

	return s.repo.UpdateInterest(&(*interest)[0], newInterest)

}

func (i InterestService) DeleteInterest(interestId *uint) (*Interest, error) {
	//TODO implement me
	panic("implement me")
}
