package model

type City struct {
	Id         uint64 `csv:"id"`
	Name       string `csv:"name"`
	Region     string `csv:"region"`
	District   string `csv:"district"`
	Population uint   `csv:"population"`
	Foundation uint16 `csv:"foundation"`
}
