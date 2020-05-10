package model

type UserConnection struct {
	TotalCount int     `json:"totalCount"`
	Nodes      []*User `json:"nodes"`
}
