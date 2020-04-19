package Models

type Member struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (b *Member) TableName() string {
	return "member"
}
