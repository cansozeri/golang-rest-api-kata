package request

type SearchRecordRequest struct {
	StartDate string `json:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate   string `json:"endDate" validate:"required,datetime=2006-01-02,gtecsfield=StartDate"`
	MinCount  int    `json:"minCount" validate:"required"`
	MaxCount  int    `json:"maxCount" validate:"required,gtecsfield=MinCount"`
}
