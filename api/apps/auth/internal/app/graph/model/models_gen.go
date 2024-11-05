// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/kloudlite/api/pkg/repos"
)

type InviteCode struct {
	ID         repos.ID `json:"id"`
	Name       string   `json:"name"`
	InviteCode string   `json:"inviteCode"`
}

type Mutation struct {
}

type OAuthProviderStatus struct {
	Provider string `json:"provider"`
	Enabled  bool   `json:"enabled"`
}

type Query struct {
}

type RemoteLogin struct {
	Status     string  `json:"status"`
	AuthHeader *string `json:"authHeader,omitempty"`
}

type Session struct {
	ID           repos.ID `json:"id"`
	UserID       repos.ID `json:"userId"`
	UserEmail    string   `json:"userEmail"`
	LoginMethod  string   `json:"loginMethod"`
	UserVerified bool     `json:"userVerified"`
}

type User struct {
	ID             repos.ID               `json:"id"`
	Name           string                 `json:"name"`
	Email          string                 `json:"email"`
	Avatar         *string                `json:"avatar,omitempty"`
	Invite         string                 `json:"invite"`
	Verified       bool                   `json:"verified"`
	Approved       bool                   `json:"approved"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	Joined         string                 `json:"joined"`
	ProviderGitlab map[string]interface{} `json:"providerGitlab,omitempty"`
	ProviderGithub map[string]interface{} `json:"providerGithub,omitempty"`
	ProviderGoogle map[string]interface{} `json:"providerGoogle,omitempty"`
}

func (User) IsEntity() {}