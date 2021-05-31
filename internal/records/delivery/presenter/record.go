package presenter

import "time"

type RecordPresenter struct {
	Code    int       `json:"code"`
	Msg     string    `json:"msg"`
	Records []*Record `json:"records"`
}

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}
