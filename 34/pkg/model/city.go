package model

type City struct {
	Id         uint64 `csv:"id" json:"-"`
	Name       string `csv:"name" json:"name"`
	Region     string `csv:"region" json:"region"`
	District   string `csv:"district" json:"district"`
	Population uint   `csv:"population" json:"population"`
	Foundation uint   `csv:"foundation" json:"foundation"`
}
