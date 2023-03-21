package shared

type Pagination struct {
	Count             float64 `json:"count"`
	NextPage          string  `json:"next_page"`
	NextPageToken     string  `json:"next_page_token"`
	PreviousPage      string  `json:"previous_page"`
	PreviousPageToken string  `json:"previous_page_token"`
}
