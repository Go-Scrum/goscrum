package db

type Slack struct {
	Model
	Organization   Organization
	OrganizationID string

	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	BotUserID   string `json:"bot_user_id"`
	AppID       string `json:"app_id"`
	TeamID      string `json:"team_id"`
	TeamName    string `json:"team_name"`
}
