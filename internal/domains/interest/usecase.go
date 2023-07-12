package interest

import (
	"Datearn/pkg/advancedError"
	"Datearn/pkg/userHandler"
	"context"
)

type InterestUseCase struct {
	sv        InterestServiceInterface
	decoderFn func(ctx context.Context, token string) (uint, error)
}

func NewInterestUseCase(sv InterestServiceInterface, decoderFn func(ctx context.Context, token string) (uint, error)) InterestUseCaseInterface {
	if decoderFn == nil {
		decoderFn = userHandler.ExtractUserIdFromToken
	}

	return &InterestUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

func (s *InterestUseCase) GetAllInterest(ctx context.Context, token string, request *InterestGetRequest) (*[]Interest, error) {
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return s.sv.GetAllInterest(request.Id, request.Name, request.Limit, request.Offset)
}

func (s *InterestUseCase) CreateInterest(ctx context.Context, token string, request *InterestCreateRequest) (*Interest, error) {
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return s.sv.CreateInterest(request.Name)
}

func (s *InterestUseCase) UpdateInterest(ctx context.Context, token string, request *InterestUpdateRequest) (*Interest, error) {
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}

	return s.sv.UpdateInterest(request.Id, *request.Name)

}

func (s *InterestUseCase) DeleteInterest(ctx context.Context, token string, request *InterestDeleteRequest) (*Interest, error) {
	_, err := s.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}

	return s.sv.DeleteInterest(request.Id)
}
