package domain

type TweetsRepo interface {
	AddTweet(Tweet) (*Tweet, error)
	GetAllTweets(int) ([]Tweet, error)
}

type postgresTweetsRepo struct {
}

func NewTweetsRepo() TweetsRepo {
	return &postgresTweetsRepo{}
}

func (r *postgresTweetsRepo) AddTweet(tweet Tweet) (*Tweet, error) {
	return &tweet, nil
}

func (r *postgresTweetsRepo) GetAllTweets(userId int) ([]Tweet, error) {
	return []Tweet{
		{Id: 1, UserId: 2, Text: "hello tweet"},
	}, nil
}
