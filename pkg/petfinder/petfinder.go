package petfinder

type Config struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
}

type Client struct {
	baseUrl      string
	clientId     string
	clientSecret string
	authToken    string
}

type Pager struct {
	Limit int
	Page  int
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

func NewClient(baseUrl string, clientId string, clientSecret string) *Client {
	return &Client{
		baseUrl:      baseUrl,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}
