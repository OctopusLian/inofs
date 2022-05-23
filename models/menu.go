package models

type Tree struct {
	Id       int
	AuthName string
	UrlFor   string
	Weight   int
	Children []*Tree
}
