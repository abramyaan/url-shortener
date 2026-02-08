package link
import "errors"
// То, что присылает клиент
type LinkCreateRequest struct {
	URL string `json:"url"`
}

func (req *LinkCreateRequest) Validate() error {
	if req.URL == "" {
		return errors.New("url is required")
	}
	
	return nil
}

// Ответ при создании ссылки
type LinkResponse struct {
	ID   uint   `json:"id"`
	URL  string `json:"url"`
	Hash string `json:"hash"`
}

// Статистика кликов 
type StatResponse struct {
	LinkID uint   `json:"link_id"`
	Hash   string `json:"hash"`
	Clicks int    `json:"clicks"`
}