package usecases

var GoogleEntry GoogleEntryInterface

type GoogleEntryInterface interface {
}

type googleEntry struct {
}

func NewGoogleEntry() GoogleEntryInterface {
	GoogleEntry = &googleEntry{}
	return GoogleEntry
}

func (f *googleEntry) Authentication() error {
	return nil
}

func (f *googleEntry) CreateCampaign() error {
	return nil
}

func (f *googleEntry) UploadCustomAudience() error {
	return nil
}

func (f *googleEntry) MapAudienceCampaign() error {
	return nil
}
