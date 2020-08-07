package main

type Data struct {
	IP       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"connect"`
	UserList []string `json:"user_list"`
}
