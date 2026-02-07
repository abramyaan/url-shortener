package link

type LinkCreateRequest struct {
	URL string `json:"url" validate:"required,url"`

}

type LinkResponse struct {
	ID   uint   `json:"id"`
	URL  string `json:"url"`
	Hash string `json:"hash"`
}