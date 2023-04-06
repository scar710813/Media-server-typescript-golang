package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dan6erbond/jolt-server/graph/generated"
	"github.com/dan6erbond/jolt-server/pkg/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// RateTv is the resolver for the rateTv field.
func (r *mutationResolver) RateTv(ctx context.Context, tmdbID string, rating float64) (*models.Review, error) {
	if rating > MaxReview {
		return nil, fmt.Errorf("rating cannot be higher than 5 stars")
	} else if rating < MinReview {
		return nil, fmt.Errorf("rating must be at least 1 star")
	}

	//nolint:revive
	tmdbId, err := strconv.ParseInt(tmdbID, 10, 64)
	if err != nil {
		return nil, err
	}

	tv, err := r.tvService.GetOrCreateTvByTmdbID(int(tmdbId))

	if err != nil {
		return nil, err
	}

	user, err := r.authService.GetUser(ctx)

	if err != nil {
		return nil, err
	}

	var ratings []models.Review

	err = r.db.Model(&tv).Where("created_by_id = ?", user.ID).Association("Reviews").Find(&ratings)

	if err != nil {
		return nil, err
	}

	var tvRating models.Review
	if len(ratings) > 0 {
		tvRating = ratings[0]

		tvRating.Rating = rating

		err = r.db.Save(&tvRating).Error

		if err != nil {
			return nil, err
		}
	} else {
		tvRating = models.Review{
			Rating:    rating,
			CreatedBy: *user,
		}

		err = r.db.Model(&tv).Association("Reviews").Append(&tvRating)

		if err != nil {
			return nil, err
		}
	}

	return &tvRating, nil
}

// ReviewTv is the resolver for the reviewTv field.
func (r *mutationResolver) ReviewTv(ctx context.Context, tmdbID string, review string) (*models.Review, error) {
	user, err := r.authService.GetUser(ctx)

	if err != nil {
		return nil, err
	}

	//nolint:revive
	tmdbId, err := strconv.ParseInt(tmdbID, 10, 64)
	if err != nil {
		return nil, err
	}

	tv, err := r.tvService.GetOrCreateTvByTmdbID(int(tmdbId))

	if err != nil {
		return nil, err
	}

	var reviews []models.Review

	err = r.db.Model(tv).Where("created_by_id = ?", user.ID).Association("Reviews").Find(&reviews)

	if err != nil {
		return nil, err
	}

	var tvReview models.Review
	if len(reviews) > 0 {
		tvReview = reviews[0]

		tvReview.Review = review

		err = r.db.Save(&tvReview).Error

		if err != nil {
			return nil, err
		}
	} else {
		tvReview.CreatedBy = *user
		tvReview.Review = review

		err = r.db.Model(&tv).Association("Reviews").Append(&tvReview)

		if err != nil {
			return nil, err
		}
	}

	return &tvReview, nil
}

// Tv is the resolver for the tv field.
func (r *queryResolver) Tv(ctx context.Context, id *string, tmdbID *string) (*models.Tv, error) {
	if id != nil {
		var tv models.Tv

		//nolint:revive
		dbID, err := strconv.ParseInt(*tmdbID, 10, 64)
		if err != nil {
			return nil, err
		}

		err = r.db.First(&tv, dbID).Error
		if err != nil {
			return nil, err
		}

		return &tv, nil
	}

	if tmdbID == nil {
		return nil, gqlerror.Errorf("one of id and tmdbId must be given")
	}

	//nolint:revive
	tmdbId, err := strconv.ParseInt(*tmdbID, 10, 64)
	if err != nil {
		return nil, err
	}

	//TODO: adapt for Tv
	_, err = r.movieService.GetOrCreateMovieByTmdbID(int(tmdbId))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Rating is the resolver for the rating field.
func (r *tvResolver) Rating(ctx context.Context, obj *models.Tv) (float64, error) {
	var ratings []models.Review

	err := r.db.Model(&obj).Association("Reviews").Find(&ratings)

	if err != nil {
		return 0, err
	}

	var (
		sum        float64
		numRatings int64
	)
	//nolint:wsl
	for _, val := range ratings {
		if val.Rating == 0 {
			continue
		}
		sum += val.Rating
		numRatings++
	}

	var rating float64
	if sum > 0 {
		rating = sum / float64(numRatings)
	} else {
		rating = 0
	}
	return rating, nil
}

// Reviews is the resolver for the reviews field.
func (r *tvResolver) Reviews(ctx context.Context, obj *models.Tv) ([]*models.Review, error) {
	var reviews []*models.Review

	err := r.db.Model(&obj).Association("Reviews").Find(&reviews)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

// UserReview is the resolver for the userReview field.
func (r *tvResolver) UserReview(ctx context.Context, obj *models.Tv) (*models.Review, error) {
	user, err := r.authService.GetUser(ctx)

	if err != nil {
		return nil, err
	}

	var reviews []models.Review

	err = r.db.Model(&obj).Where("created_by_id = ?", user.ID).Association("Reviews").Find(&reviews)

	if err != nil {
		return nil, err
	}

	if len(reviews) == 0 {
		//nolint:nilnil // nil represents null in GraphQL
		return nil, nil
	}

	return &reviews[0], nil
}

// AvailableOnJellyfin is the resolver for the availableOnJellyfin field.
func (r *tvResolver) AvailableOnJellyfin(ctx context.Context, obj *models.Tv) (bool, error) {
	panic(fmt.Errorf("not implemented: AvailableOnJellyfin - availableOnJellyfin"))
}

// Genres is the resolver for the genres field.
func (r *tvResolver) Genres(ctx context.Context, obj *models.Tv) ([]string, error) {
	return []string(obj.Genres), nil
}

// Watched is the resolver for the watched field.
func (r *tvResolver) Watched(ctx context.Context, obj *models.Tv) (bool, error) {
	panic(fmt.Errorf("not implemented: Watched - watched"))
}

// AddedToWatchlist is the resolver for the addedToWatchlist field.
func (r *tvResolver) AddedToWatchlist(ctx context.Context, obj *models.Tv) (bool, error) {
	panic(fmt.Errorf("not implemented: AddedToWatchlist - addedToWatchlist"))
}

// Tv returns generated.TvResolver implementation.
func (r *Resolver) Tv() generated.TvResolver { return &tvResolver{r} }

type tvResolver struct{ *Resolver }
