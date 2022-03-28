package domain

type Notification struct {
	Id         int32      `json:"id"`
	Reason     string     `json:"reason"`
	Repository Repository `json:"repository"`
	Url        string     `json:"html_url"`
}

type Repository struct {
	ID       int32  `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}
