package location

type CountryService struct {
	repo CountryRepositoryInterface
}

type CityService struct {
	repo CityRepositoryInterface
}

func NewCountryService(repo CountryRepositoryInterface) CountryServiceInterface {
	return &CountryService{
		repo: repo,
	}
}

func (c CountryService) GetAllCountries(name *string, limit *int, offset int) (*[]Country, error) {
	countries := c.repo.GetAllCountries(name, limit, offset)
	if len(*countries) == 0 {
		return nil, CountryNotFound
	}

	return countries, nil
}

func NewCityService(repo CityRepositoryInterface) CityServiceInterface {
	return &CityService{
		repo: repo,
	}
}

func (c CityService) GetAllCities(name *string, limit *int, offset int) (*[]City, error) {
	cities := c.repo.GetAllCities(name, limit, offset)
	if len(*cities) == 0 {
		return nil, CityNotFound
	}

	return cities, nil
}
