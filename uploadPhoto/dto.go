package uploadphoto

type GetListPhoto struct {
	Limit    int    `json:"limit"`
	SearchBy string `json:"search_by"`
	Keyword  string `json:"keyword"`
}

type SavePhoto struct {
	Title   string `json:"title"`
	Caption string `json:"caption"`
	URL     string `json:"photo_url"`
}

type EditPhoto struct {
	PhotoID    int64  `json:"photo_id"`
	NewTitle   string `json:"new_title"`
	NewCaption string `json:"new_caption"`
}
