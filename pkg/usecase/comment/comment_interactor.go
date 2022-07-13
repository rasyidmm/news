package comment

import "news/pkg/domain/repository"

type CommentInteractor struct {
	repo repository.CommentRepository
	out  CommentOutputPort
}

func NewCommentInteractor(r repository.CommentRepository, o CommentOutputPort) *CommentInteractor {
	return &CommentInteractor{
		repo: r,
		out:  o,
	}
}
