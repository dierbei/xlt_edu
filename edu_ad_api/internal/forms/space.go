package forms

type SpaceForm struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	SpaceKey string `json:"space_key"`
}

type AdForm struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	SpaceID     int32  `json:"spaceId"`
	Keyword     string `json:"keyword"`
	HTMLContent string `json:"htmlContent"`
	Text        string `json:"text"`
	Img         string `json:"img"`
	Link        string `json:"link"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Status      int32  `json:"status"`
	Priority    int32  `json:"priority"`
}
