package posts

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"
	"github.com/microcosm-cc/bluemonday"

	"github.com/jinzhu/copier"
	"golang.org/x/exp/slices"
)

type usecase struct {
	postsMicroservice         domain.PostsMicroservice
	userMicroservice          domain.UsersMicroservice
	subscriptionsMicroservice domain.SubscriptionMicroservice
	imgUseCase                domain.ImageUseCase
}

func New(p domain.PostsMicroservice, u domain.UsersMicroservice, i domain.ImageUseCase, s domain.SubscriptionMicroservice) domain.PostsUseCase {
	return &usecase{
		postsMicroservice:         p,
		userMicroservice:          u,
		subscriptionsMicroservice: s,
		imgUseCase:                i,
	}
}

func blurContent(content string) string {
	r := regexp.MustCompile(`"https://wsrv\.nl/\?url=.{100}.?"`)

	img := r.Find([]byte(content))

	if len(img) == 0 {
		return ""
	}

	idx := bytes.LastIndex(img, []byte(`/`))
	toReplace := string(img[:idx+1]) + "blur_" + string(img[idx+1:])

	return fmt.Sprintf(`<img src=%s class="post-content__image">`, toReplace)
}

func SanitizeContent(content string, blur bool) string {
	p := bluemonday.UGCPolicy()

	p.AllowAttrs("class").OnElements("img")

	if blur {
		content = blurContent(content)
	}

	return p.Sanitize(content)
}

func (u usecase) GetPostsByFilter(userID, authorID uint64) ([]models.Post, error) {
	r := make([]models.Post, 0)
	var err error
	validate := true

	switch {
	case authorID == 0:
		if r, err = u.postsMicroservice.GetPostsBySubscriptions(userID); err != nil {
			return nil, err
		}
		validate = false
	case authorID > 0:
		if r, err = u.postsMicroservice.GetAllByUserID(authorID); err != nil {
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
			as, errSubscriptions := u.subscriptionsMicroservice.GetSubscriptionByUserAndAuthorID(userID, authorID)
			if errSubscriptions != nil {
				return nil, errSubscriptions
			}

			if as.Tier >= post.Tier {
				r[i].IsAllowed = true
			}
		} else {
			r[i].IsAllowed = true
		}

		r[i].Content = SanitizeContent(post.Content, !r[i].IsAllowed)

		author, errGetAuthor := u.userMicroservice.GetByID(post.UserID)
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

	for index, post := range r {
		comments, err := u.GetCommentsByPostID(post.ID)
		if err != nil {
			return nil, err
		}
		r[index].CommentsNum = uint64(len(comments))
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i].ID < r[j].ID
	})

	return r, nil
}

func (u usecase) GetPostByID(postID, userID uint64) (models.Post, error) {
	r, err := u.postsMicroservice.GetPostByID(postID)
	if err != nil {
		return models.Post{}, err
	}

	author, errGetAuthor := u.userMicroservice.GetByID(r.UserID)
	if errGetAuthor != nil {
		return models.Post{}, err
	}

	as, errSubscriptions := u.subscriptionsMicroservice.GetSubscriptionByUserAndAuthorID(userID, author.ID)
	if errSubscriptions != nil {
		return models.Post{}, errSubscriptions
	}

	if as.Tier >= r.Tier {
		r.IsAllowed = true
	}

	r.Content = SanitizeContent(r.Content, !r.IsAllowed)

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

	comments, err := u.GetCommentsByPostID(postID)
	if err != nil {
		return models.Post{}, err
	}
	r.CommentsNum = uint64(len(comments))

	return r, nil
}

func (u usecase) Create(post models.Post, userID uint64) (models.Post, error) {
	post.UserID = userID
	var err error
	post, err = u.postsMicroservice.Create(post)
	if err != nil {
		return models.Post{}, err
	}

	if err = u.CreateTags(post.Tags, post.ID); err != nil {
		return models.Post{}, err
	}

	post.Content = SanitizeContent(post.Content, false)

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

	updatePost.Content = SanitizeContent(updatePost.Content, false)

	return updatePost, u.postsMicroservice.Update(updatePost)
}

func (u usecase) DeleteByID(postID uint64) error {
	err := u.DeleteTagDeps(postID)
	if err != nil {
		return err
	}
	return u.postsMicroservice.DeleteByID(postID)
}

func (u usecase) GetLikesByPostID(postID uint64) ([]models.Like, error) {
	return u.postsMicroservice.GetAllLikesByPostID(postID)
}

func (u usecase) GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error) {
	return u.postsMicroservice.GetLikeByUserAndPostID(userID, postID)
}

func (u usecase) LikePost(userID, postID uint64) error {
	return u.postsMicroservice.CreateLike(userID, postID)
}

func (u usecase) UnlikePost(userID, postID uint64) error {
	return u.postsMicroservice.DeleteLikeByID(userID, postID)
}

func (u usecase) GetLikesNum(postID uint64) (uint64, error) {
	likes, err := u.GetLikesByPostID(postID)
	if err != nil {
		return 0, err
	}
	return uint64(len(likes)), nil
}

func (u usecase) IsPostLiked(userID, postID uint64) bool {
	if l, err := u.GetLikeByUserAndPostID(userID, postID); err != nil || l == (models.Like{}) {
		return false
	}
	return true
}

func (u usecase) CreateTags(tagNames []string, postID uint64) error {
	for _, tagName := range tagNames {
		tag, err := u.postsMicroservice.GetTagByName(tagName)
		tagID := tag.ID
		if err != nil {
			tagID, err = u.postsMicroservice.CreateTag(tagName)
			if err != nil {
				return err
			}
		}
		if err = u.postsMicroservice.CreateDepTag(postID, tagID); err != nil {
			return err
		}
	}
	return nil
}

func (u usecase) GetTagsByPostID(postID uint64) ([]models.Tag, error) {
	tagDeps, err := u.postsMicroservice.GetTagDepsByPostId(postID)
	if err != nil {
		return nil, err
	}

	var tags []models.Tag
	for _, dep := range tagDeps {
		tag, tagErr := u.postsMicroservice.GetTagById(dep.TagID)
		if tagErr != nil {
			return tags, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (u usecase) DeleteTagDeps(postID uint64) error {
	tagDeps, err := u.postsMicroservice.GetTagDepsByPostId(postID)
	if err != nil {
		return err
	}
	for _, tagDep := range tagDeps {
		if err = u.postsMicroservice.DeleteDepTag(tagDep); err != nil {
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
		err = u.postsMicroservice.DeleteDepTag(tagDep)
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

func (u usecase) CreateComment(comment models.Comment) (models.Comment, error) {
	id, date, err := u.postsMicroservice.CreateComment(comment)
	if err != nil {
		return models.Comment{}, err
	}
	comment.ID = id
	comment.DateCreated = date
	return comment, nil
}

func (u usecase) GetCommentsByPostID(postID uint64) ([]models.Comment, error) {
	return u.postsMicroservice.GetCommentsByPostID(postID)
}

func (u usecase) GetCommentByID(commentID uint64) (models.Comment, error) {
	return u.postsMicroservice.GetCommentByID(commentID)
}

func (u usecase) UpdateComment(commentID uint64, commentMsg string) (models.Comment, error) {
	comment, err := u.GetCommentByID(commentID)
	if err != nil {
		return models.Comment{}, err
	}
	comment.Content = commentMsg
	err = u.postsMicroservice.UpdateComment(comment)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (u usecase) DeleteComment(commentID uint64) error {
	return u.postsMicroservice.DeleteCommentByID(commentID)
}
