package models

import "time"

type AuthModel struct {
	ID        string `bson:"id" json:"id"`
	Name      string `bson:"name" json:"name"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"-"`
	ProjectID string `bson:"project_id" json:"project_id"`
	Status    string `bson:"status" json:"status"`

	// NOTE: internal usage
	TokenID   string    `bson:"token_id" json:"-"`   // used for auth token
	Catalog   string    `bson:"catalog" json:"-"`    // used for auth token
	ExpiredAt time.Time `bson:"expired_at" json:"-"` // used for auth token
}

func (auth *AuthModel) IsValidToken() bool {
	return auth.TokenID != "" && auth.ExpiredAt.After(time.Now())
}

func (auth *AuthModel) WithToken(token string, expiredAt time.Time) *AuthModel {
	auth.TokenID = token
	auth.ExpiredAt = expiredAt

	return auth
}
