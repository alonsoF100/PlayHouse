package handlers

type UserHandlerInfo struct {
	Nickname string
	Password string
	Balance  int
	Rating   int
}

type UserResponse struct {
	Nickname string `json:"nickname"`
	Balance  int    `json:"balance"`
	Rating   int    `json:"rating"`
}

type UserHandlerMoney struct {
	Nickname string
	Money    int
}
type UserHandlerGame struct {
	Nickname string
	Color    string
	Bet      int
}
