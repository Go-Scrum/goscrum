package model

type SlackUserConnection struct {
	TotalCount int          `json:"totalCount"`
	Nodes      []*SlackUser `json:"nodes"`
}
