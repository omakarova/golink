package mymodels

type NewUser struct {
	Username          string `json:"username"`
	Password       string `json:"password"`
}

type UserInfo struct {
	Username        string `json:"username"`
	LinksCount       int `json:"linkscount"`
}