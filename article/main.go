package article

import "realworld-backend/shared"

type Article struct {
	Id shared.ID `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Body string `json:"body"`
	Slug string `json:"slug"`
}

type ArticleInput struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Body string `json:"body"`
}

// type UpdateArticleInput struct {
// 	Id string `json:"id"`
// 	Title string `json:"title"`
// 	Description string `json:"description"`
// 	Body string `json:"body"`
// }



type ArticleRepository interface {
	Create(*Article) error
	Update(*Article) error
	FindBySlug(string) *Article
	FindAll() []Article
}