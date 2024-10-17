package usecase

type (
	ExampleUseCase interface {
		GoroutineTest() (any, error)
	}

	exampleUseCase struct {
	}
)

func NewExampleUseCase() ExampleUseCase {
	return &exampleUseCase{}
}

func (u *exampleUseCase) GoroutineTest() (any, error) {
	return "Ok", nil
}
