package service

import (
	"context"
	"log"

	"forum/internal/entity"
	"forum/internal/repository"
)

type postService struct {
	postRepo repository.PostRepo
	catRepo  repository.CategoryRepo
}

func NewPostService(pRepo repository.PostRepo, catRepo repository.CategoryRepo) PostService {
	log.Println("| | post service is done!")
	return &postService{
		postRepo: pRepo,
		catRepo:  catRepo,
	}
}

func (p *postService) CreatePost(ctx context.Context, post entity.Post) (int64, error) {
	createdPostID, err := p.postRepo.CreatePost(ctx, post)
	if err != nil {
		return -1, err
	}

	if err := p.catRepo.CreateCategory(ctx, uint64(createdPostID), post.Categories); err != nil {
		// p.postRepo.DeletePost(ctx, createdPostID)
		return -1, err
	}

	return createdPostID, nil
}

func (p *postService) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	posts, err := p.postRepo.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(posts); i++ {
		posts[i].Categories, err = p.catRepo.GetCategoriesByPostID(ctx, posts[i].ID)
		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}

func (p *postService) GetPostByID(ctx context.Context, postID uint64) (entity.Post, error) {
	post, err := p.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return entity.Post{}, err
	}

	post.Categories, err = p.catRepo.GetCategoriesByPostID(ctx, postID)
	if err != nil {
		return entity.Post{}, err
	}

	return post, nil
}
