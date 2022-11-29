package posts

import (
	"regexp"
	"sort"
	"strings"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	"github.com/jinzhu/copier"
	"golang.org/x/exp/slices"
)

type usecase struct {
	postsRepo         domain.PostsRepository
	userRepo          domain.UsersRepository
	subscriptionsRepo domain.SubscriptionsRepository
	imgUseCase        domain.ImageUseCase
}

func New(p domain.PostsRepository, u domain.UsersRepository, i domain.ImageUseCase, s domain.SubscriptionsRepository) domain.PostsUseCase {
	return &usecase{
		postsRepo:         p,
		userRepo:          u,
		subscriptionsRepo: s,
		imgUseCase:        i,
	}
}

func (u usecase) RenderHTML(content []byte, blur bool) (string, error) {
	r := regexp.MustCompile(`\[img\|.{121}.?]`)
	bytesImages := r.FindAll(content, -1)

	if len(bytesImages) == 0 {
		return string(content), nil
	}

	m := regexp.MustCompile(`wsrv.nl`)
	if blur {
		img := bytesImages[0]
		if !m.Match(img) {
			// make special error for this
			return "", domain.ErrBadRequest
		}
		img = img[5 : len(img)-1]
		if img[len(img)-1] == '/' {
			img = img[:len(img)-1]
		}

		idx := strings.LastIndex(string(img[:len(img)-1]), "/")
		if idx == -1 {
			return "", domain.ErrBadRequest
		}

		tmp := append([]byte("blur_"), img[idx+1:]...)
		img = append(img[:idx+1], tmp...)
		img = append(append([]byte(`<img src="`), img...), []byte(`" class="post-content__image">`)...)

		return string(img), nil
	}

	for i, img := range bytesImages {
		if !m.Match(img) {
			// make special error for this
			return "", domain.ErrBadRequest
		}
		img = img[5 : len(img)-1]
		if img[len(img)-1] == '/' {
			img = img[:len(img)-1]
		}
		bytesImages[i] = append(append([]byte(`<img src="`), img...), []byte(`" class="post-content__image">`)...)

		content = r.ReplaceAll(content, bytesImages[i])
	}

	return string(content), nil
}

func (u usecase) GetPostsByFilter(userID, authorID uint64) ([]models.Post, error) {
	r := make([]models.Post, 0)
	var err error
	validate := true

	switch {
	case authorID == 0:
		if r, err = u.postsRepo.GetPostsBySubscriptions(userID); err != nil {
			return nil, err
		}
		validate = false
	case authorID > 0:
		if r, err = u.postsRepo.GetAllByUserID(authorID); err != nil {
			return nil, err
		}
		if authorID == userID {
			validate = false
		}
	default:
		return nil, domain.ErrInternal
	}

	for i, post := range r {
		if validate {
			as, errSubscriptions := u.subscriptionsRepo.GetSubscriptionByUserAndAuthorID(userID, authorID)
			if errSubscriptions != nil {
				return nil, errSubscriptions
			}

			if as.Tier >= post.Tier {
				r[i].IsAllowed = true
			}
		} else {
			r[i].IsAllowed = true
		}

		if r[i].Content, err = u.RenderHTML([]byte(r[i].ContentTemplate), !r[i].IsAllowed); err != nil {
			return nil, err
		}

		author, errGetAuthor := u.userRepo.GetByID(post.UserID)
		if errGetAuthor != nil {
			return nil, errGetAuthor
		}

		tags, getTagsErr := u.GetTagsByPostID(r[i].ID)
		if getTagsErr != nil {
			return nil, getTagsErr
		}
		tagsStr := u.ConvertTagsToStrSlice(tags)

		r[i].Tags = tagsStr
		r[i].Author.UserID = author.ID
		r[i].Author.Username = author.Username
		if r[i].Author.ImgPath, err = u.imgUseCase.GetImage(author.Avatar); err != nil {
			return nil, err
		}

		if r[i].LikesNum, err = u.GetLikesNum(post.ID); err != nil {
			return nil, domain.ErrInternal
		}
		r[i].IsLiked = u.IsPostLiked(userID, post.ID)
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i].ID > r[j].ID
	})

	return r, nil
}

func (u usecase) GetPostByID(postID, userID uint64) (models.Post, error) {
	r, err := u.postsRepo.GetPostByID(postID)
	if err != nil {
		return models.Post{}, err
	}

	author, errGetAuthor := u.userRepo.GetByID(r.UserID)
	if errGetAuthor != nil {
		return models.Post{}, err
	}

	as, errSubscriptions := u.subscriptionsRepo.GetSubscriptionByUserAndAuthorID(userID, author.ID)
	if errSubscriptions != nil {
		return models.Post{}, errSubscriptions
	}

	if as.Tier >= r.Tier {
		r.IsAllowed = true
	}

	if r.Content, err = u.RenderHTML([]byte(r.ContentTemplate), !r.IsAllowed); err != nil {
		return models.Post{}, err
	}

	tags, err := u.GetTagsByPostID(r.ID)
	if err != nil {
		return models.Post{}, err
	}
	tagsStr := u.ConvertTagsToStrSlice(tags)

	r.Tags = tagsStr
	r.Author.UserID = author.ID
	r.Author.Username = author.Username
	if r.Author.ImgPath, err = u.imgUseCase.GetImage(author.Avatar); err != nil {
		return models.Post{}, err
	}

	if r.LikesNum, err = u.GetLikesNum(postID); err != nil {
		return models.Post{}, errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	r.IsLiked = u.IsPostLiked(userID, postID)

	return r, nil
}

func (u usecase) Create(post models.Post, userID uint64) (models.Post, error) {
	post.UserID = userID
	var err error
	post.ID, err = u.postsRepo.Create(post)
	if err != nil {
		return models.Post{}, err
	}

	if err = u.CreateTags(post.Tags, post.ID); err != nil {
		return models.Post{}, err
	}

	if post.Content, err = u.RenderHTML([]byte(post.ContentTemplate), false); err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (u usecase) Update(post models.Post, postID uint64) (models.Post, error) {
	var err error

	updatePost, err := u.GetPostByID(postID, post.UserID)
	if err != nil {
		return models.Post{}, err
	}

	if err = copier.CopyWithOption(&updatePost, &post, copier.Option{IgnoreEmpty: true}); err != nil {
		return models.Post{}, err
	}

	if err = u.UpdateTags(post.Tags, postID); err != nil {
		return models.Post{}, err
	}

	if updatePost.ContentTemplate, err = u.RenderHTML([]byte(updatePost.ContentTemplate), false); err != nil {
		return models.Post{}, err
	}

	return updatePost, u.postsRepo.Update(updatePost)
}

func (u usecase) DeleteByID(postID uint64) error {
	err := u.DeleteTagDeps(postID)
	if err != nil {
		return err
	}
	return u.postsRepo.DeleteByID(postID)
}

func (u usecase) GetLikesByPostID(postID uint64) ([]models.Like, error) {
	return u.postsRepo.GetAllLikesByPostID(postID)
}

func (u usecase) GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error) {
	return u.postsRepo.GetLikeByUserAndPostID(userID, postID)
}

func (u usecase) LikePost(userID, postID uint64) error {
	return u.postsRepo.CreateLike(userID, postID)
}

func (u usecase) UnlikePost(userID, postID uint64) error {
	return u.postsRepo.DeleteLikeByID(userID, postID)
}

func (u usecase) GetLikesNum(postID uint64) (uint64, error) {
	likes, err := u.GetLikesByPostID(postID)
	if err != nil {
		return 0, err
	}
	return uint64(len(likes)), nil
}

func (u usecase) IsPostLiked(userID, postID uint64) bool {
	if _, err := u.GetLikeByUserAndPostID(userID, postID); err != nil {
		return false
	}
	return true
}

func (u usecase) CreateTags(tagNames []string, postID uint64) error {
	for _, tagName := range tagNames {
		tag, err := u.postsRepo.GetTagByName(tagName)
		tagID := tag.ID
		if err != nil {
			tagID, err = u.postsRepo.CreateTag(tagName)
			if err != nil {
				return err
			}
		}
		if err = u.postsRepo.CreateDepTag(postID, tagID); err != nil {
			return err
		}
	}
	return nil
}

func (u usecase) GetTagsByPostID(postID uint64) ([]models.Tag, error) {
	tagDeps, err := u.postsRepo.GetTagDepsByPostId(postID)
	if err != nil {
		return nil, err
	}

	var tags []models.Tag
	for _, dep := range tagDeps {
		tag, tagErr := u.postsRepo.GetTagById(dep.TagID)
		if tagErr != nil {
			return tags, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (u usecase) DeleteTagDeps(postID uint64) error {
	tagDeps, err := u.postsRepo.GetTagDepsByPostId(postID)
	if err != nil {
		return err
	}
	for _, tagDep := range tagDeps {
		if err = u.postsRepo.DeleteDepTag(tagDep); err != nil {
			return err
		}
	}
	return nil
}

func (u usecase) UpdateTags(tagNames []string, postID uint64) error {
	postTags, err := u.GetTagsByPostID(postID)
	if err != nil {
		return err
	}
	copyToTags := u.ConvertTagsToStrSlice(postTags)

	for _, tagName := range tagNames {
		if slices.Contains(copyToTags, tagName) {
			index := slices.Index(copyToTags, tagName)
			copyToTags = append(copyToTags[:index], copyToTags[index+1:]...)
			continue
		}
		if err = u.CreateTags(append(make([]string, 0), tagName), postID); err != nil {
			return err
		}
	}

	for _, tagName := range copyToTags {
		idx := slices.IndexFunc(postTags, func(t models.Tag) bool { return t.TagName == tagName })
		tagDep := models.TagDep{
			PostID: postID,
			TagID:  postTags[idx].ID,
		}
		err = u.postsRepo.DeleteDepTag(tagDep)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u usecase) ConvertTagsToStrSlice(tags []models.Tag) []string {
	var tagsStr []string
	for _, tag := range tags {
		tagsStr = append(tagsStr, tag.TagName)
	}
	return tagsStr
}
