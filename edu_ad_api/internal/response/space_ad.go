package response

type SpaceAdDTO struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	SpaceKey   string  `json:"spaceKey"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
	IsDel      int     `json:"isDel"`
	AdDTOList  []AdDTO `json:"adDTOList"`
}

type SpaceDTO struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	SpaceKey   string `json:"spaceKey"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	IsDel      int32  `json:"isDel"`
}

type AdDTO struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	SpaceID               int    `json:"spaceId"`
	Keyword               string `json:"keyword"`
	HTMLContent           string `json:"htmlContent"`
	Text                  string `json:"text"`
	Img                   string `json:"img"`
	Link                  string `json:"link"`
	StartTime             string `json:"startTime"`
	EndTime               string `json:"endTime"`
	CreateTime            string `json:"createTime"`
	UpdateTime            string `json:"updateTime"`
	Status                int    `json:"status"`
	Priority              int    `json:"priority"`
	StartTimeFormatString string `json:"startTimeFormatString"`
	EndTimeFormatString   string `json:"endTimeFormatString"`
}
