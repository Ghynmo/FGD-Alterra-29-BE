package apinews

import "context"

type Domain struct {
	Articles []SubDomain
}

type SubDomain struct {
	Title   string
	NewsUrl string
	ImgUrl  string
}

type UseCase interface {
	GetAPINews(ctx context.Context, apikey string) (Domain, error)
}

type Repository interface {
	GetAPINews(ctx context.Context, apikey string) (Domain, error)
}
