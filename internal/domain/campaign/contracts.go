package campaign

type NewCampaignDTO struct {
	Name      string   `json:"name"`
	Content   string   `json:"content"`
	Emails    []string `json:"emails"`
	CreatedBy string   `json:"created_by"`
}

type GetCampaignDTO struct {
	ID      string    `json:"id"`
	Send    []Contact `json:"send"`
	Content string    `json:"content"`
	Status  string    `json:"status"`
	Amount  int       `json:"amount"`
}
