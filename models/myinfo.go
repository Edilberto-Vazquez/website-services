package models

type Profile struct {
	Image        string   `json:"image" bson:"image"`
	Name         string   `json:"name" bson:"name"`
	Description  string   `json:"description" bson:"description"`
	Email        string   `json:"email" bson:"email"`
	Technologies []string `json:"technologies" bson:"technologies"`
	Resume       string   `json:"resume" bson:"resume"`
}

type Project struct {
	Name         string   `json:"name" bson:"name"`
	Description  string   `json:"description" bson:"description"`
	Repository   string   `json:"repository" bson:"repository"`
	Technologies []string `json:"technologies" bson:"technologies"`
	Website      string   `json:"website" bson:"website"`
}

type Projects struct {
	List []Project `json:"projects" bson:"projects"`
}

type Dates struct {
	From string `json:"from" bson:"from"`
	To   string `json:"to" bson:"to"`
}

type Job struct {
	Position    string `json:"position" bson:"position"`
	Company     string `json:"company" bson:"company"`
	Location    string `json:"location" bson:"location"`
	Description string `json:"description" bson:"description"`
	Dates       Dates  `json:"dates" bson:"dates"`
}

type Jobs struct {
	List []Job `json:"jobs" bson:"jobs"`
}
