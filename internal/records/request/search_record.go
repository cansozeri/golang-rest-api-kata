package request

type SearchRecord struct {
	StartDate string `json:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate   string `json:"endDate" validate:"required,datetime=2006-01-02"`
	MinCount  int    `json:"minCount" validate:"required"`
	MaxCount  int    `json:"maxCount" validate:"required,gtecsfield=MinCount"`
}
