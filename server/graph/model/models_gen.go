// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type FeedItem interface {
	IsFeedItem()
}

type Media interface {
	IsMedia()
}

type AddToWatchlistInput struct {
	TmdbID    string    `json:"tmdbId"`
	MediaType MediaType `json:"mediaType"`
}

type CreateRecommendationInput struct {
	TmdbID                  string    `json:"tmdbId"`
	MediaType               MediaType `json:"mediaType"`
	RecommendationForUserID string    `json:"recommendationForUserId"`
	Message                 string    `json:"message"`
}

type RefreshTokenResult struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignInResult struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignInWithJellyfinInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TMDBSearchResult struct {
	Page         int     `json:"page"`
	Results      []Media `json:"results"`
	TotalPages   int     `json:"totalPages"`
	TotalResults int     `json:"totalResults"`
}

type ToggleWatchedInput struct {
	TmdbID    string    `json:"tmdbId"`
	MediaType MediaType `json:"mediaType"`
}

type MediaType string

const (
	MediaTypeMovie MediaType = "MOVIE"
	MediaTypeTv    MediaType = "TV"
)

var AllMediaType = []MediaType{
	MediaTypeMovie,
	MediaTypeTv,
}

func (e MediaType) IsValid() bool {
	switch e {
	case MediaTypeMovie, MediaTypeTv:
		return true
	}
	return false
}

func (e MediaType) String() string {
	return string(e)
}

func (e *MediaType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MediaType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MediaType", str)
	}
	return nil
}

func (e MediaType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
