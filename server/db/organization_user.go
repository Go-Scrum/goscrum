package db

type OrganizationUser struct {
	OrganizationID string
	Organization   Organization
	UserID         string
	User           User
	Role           string
}
