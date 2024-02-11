package repository

type SampleRepository struct{}

type ISampleRepository interface {
	CreateSample() error
}

func NewSampleRepository() ISampleRepository {
	return &SampleRepository{}
}

func (r *SampleRepository) CreateSample() error {
	return nil
}
