package main

type Record struct {
	Browsers []string `json:"browsers,omitempty"`
	Company  string   `json:"company,omitempty"`
	Country  string   `json:"country,omitempty"`
	Email    string   `json:"email,omitempty"`
	Job      string   `json:"job,omitempty"`
	Name     string   `json:"name,omitempty"`
	Phone    string   `json:"phone,omitempty"`
}
