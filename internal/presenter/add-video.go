package presenter

type AddVideo struct {
	Name    string `json:"name" validate:"required"`
	Url     string `json:"url" validate:"required"`
	Size    int    `json:"size" validate:"required"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Content string `json:"content"`
}
