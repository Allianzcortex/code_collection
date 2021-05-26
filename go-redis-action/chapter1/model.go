package chapter1

import "github.com/go-redis/redis/v8"

type ArticleRepo struct {
	Conn *redis.Client
}

// The same idea :
// 1. use ArticleRepo struct as a wrapper
// 2. use New*() as a method to initialize
func NewArticleRepo(conn *redis.Client) *ArticleRepo {
	return &ArticleRepo{conn}
}


