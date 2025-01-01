package dto


type Token struct{
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refrsh_token"`
}