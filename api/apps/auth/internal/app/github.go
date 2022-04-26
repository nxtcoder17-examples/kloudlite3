package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	oauthGithub "golang.org/x/oauth2/github"
	"kloudlite.io/apps/auth/internal/domain"
	"kloudlite.io/pkg/errors"
	fn "kloudlite.io/pkg/functions"
)

type githubI struct {
	cfg          *oauth2.Config
	ghCli        *github.Client
	ghCliForUser func(ctx context.Context, token *oauth2.Token) *github.Client
	webhookUrl   string
}

func (gh *githubI) SearchRepos(ctx context.Context, accToken *domain.AccessToken, q string, org string, page int, size int) (*github.RepositoriesSearchResult, error) {
	rsr, _, err := gh.ghCliForUser(ctx, accToken.Token).Search.Repositories(ctx, fmt.Sprintf("%s org:%s", q, org), &github.SearchOptions{})
	if err != nil {
		return nil, errors.NewEf(err, "could not search repositories")
	}
	return rsr, nil
}

func (gh *githubI) Authorize(_ context.Context, state string) (string, error) {
	csrfToken := fn.Must(fn.CleanerNanoid(32))
	b64state, err := fn.Json.ToB64Url(map[string]string{"csrf": csrfToken, "state": state})
	if err != nil {
		return "", errors.NewEf(err, "could not JSON marshal oauth State into []byte")
	}
	return gh.cfg.AuthCodeURL(
		b64state,
		oauth2.SetAuthURLParam("allow_signup", strconv.FormatBool(true)),
	), nil
}

func (gh *githubI) ListInstallations(ctx context.Context, accToken *domain.AccessToken) ([]*github.Installation, error) {
	i, _, err := gh.ghCliForUser(ctx, accToken.Token).Apps.ListUserInstallations(ctx, &github.ListOptions{})
	if err != nil {
		return nil, errors.NewEf(err, "could not list user installations")
	}
	return i, nil
}

func (gh *githubI) ListRepos(ctx context.Context, accToken *domain.AccessToken, org string, page int, size int) ([]*github.Repository, error) {
	repos, _, err := gh.ghCliForUser(ctx, accToken.Token).Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
		Sort: "updated",
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: page,
		},
	})
	if err != nil {
		return nil, errors.NewEf(err, "could not list repos by organisations")
	}
	return repos, nil

	// repos, _, err := gh.ghCliForUser(ctx, accToken.Token).Apps.ListUserRepos(ctx, installationId, &github.ListOptions{
	// 	Page:    page,
	// 	PerPage: size,
	// })
	// if err != nil {
	// 	return nil, errors.NewEf(err, "could not list user repositories")
	// }
	// return repos, nil
}

func (gh *githubI) AddWebhook(ctx context.Context, accToken *domain.AccessToken, repoUrl string) error {
	sp := strings.Split(repoUrl, "/")
	hookUrl := fmt.Sprintf("%s?tokenId=%s", gh.webhookUrl, accToken.Id)
	hookName := "kloudlite-pipeline"

	hook, res, err := gh.ghCliForUser(ctx, accToken.Token).Repositories.CreateHook(ctx, sp[0], sp[1], &github.Hook{
		Config: map[string]interface{}{
			"url":          hookUrl,
			"content_type": "json",
		},
		Events: []string{"push"},
		Active: fn.NewBool(true),
		Name:   &hookName,
	})
	if err != nil {
		fmt.Println(res.Status, res.Body)
		// ASSERT: github returns 422 only if hook already exists on the repository
		if res.StatusCode == 422 {
			fmt.Printf("Hook: %+v\n", hook)
			return nil
		}
		return errors.NewEf(err, "could not create github webhook")
	}
	fmt.Printf("Hook: %+v\n", hook)

	return nil
}

func (gh *githubI) Callback(ctx context.Context, code, state string) (*github.User, *oauth2.Token, error) {
	token, err := gh.cfg.Exchange(ctx, code)
	if err != nil {
		return nil, nil, errors.NewEf(err, "could not exchange the token")
	}
	c := gh.cfg.Client(ctx, token)
	c2 := github.NewClient(c)
	u, _, err := c2.Users.Get(ctx, "")
	if err != nil {
		return nil, nil, errors.NewEf(err, "could nog get authenticated user from github")
	}
	return u, token, nil
}

func (gh *githubI) GetToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	return gh.cfg.TokenSource(ctx, token).Token()
}

func (gh *githubI) GetInstallationToken(ctx context.Context, accToken *domain.AccessToken, repoUrl string) (string, error) {
	gitRepo := strings.Split(repoUrl, "/")
	inst, _, err := gh.ghCli.Apps.FindRepositoryInstallation(ctx, gitRepo[0], gitRepo[1])
	if err != nil {
		return "", errors.NewEf(err, "could not fetch repository installation")
	}
	it, _, err := gh.ghCli.Apps.CreateInstallationToken(ctx, *inst.ID, &github.InstallationTokenOptions{})
	fmt.Println(it)
	if err != nil {
		return "", errors.NewEf(err, "failed to get installation token")
	}
	return it.GetToken(), err
}

func (gh *githubI) GetAppToken() {
}

func (gh *githubI) GetRepoToken() {
	panic("not implemented") // TODO: Implement
}

type GithubOAuth interface {
	GithubConfig() (clientId, clientSecret, callbackUrl, githubAppId, githubAppPKFile string)
}

func fxGithub(env *Env) domain.Github {
	clientId, clientSecret, callbackUrl, ghAppId, ghAppPKFile := env.GithubConfig()
	cfg := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     oauthGithub.Endpoint,
		RedirectURL:  callbackUrl,
		Scopes:       []string{"user:email", "admin:org"},
	}
	privatePem, err := ioutil.ReadFile(ghAppPKFile)
	if err != nil {
		panic(errors.NewEf(err, "reading github app PK file"))
	}

	appId, _ := strconv.ParseInt(ghAppId, 10, 64)
	itr, err := ghinstallation.NewAppsTransport(http.DefaultTransport, appId, privatePem)
	if err != nil {
		panic(errors.NewEf(err, "creating app transport"))
	}

	ghCliForUser := func(ctx context.Context, token *oauth2.Token) *github.Client {
		ts := oauth2.StaticTokenSource(token)
		return github.NewClient(oauth2.NewClient(ctx, ts))
	}

	ghCli := github.NewClient(&http.Client{Transport: itr, Timeout: time.Second * 30})

	return &githubI{
		cfg:          &cfg,
		ghCli:        ghCli,
		ghCliForUser: ghCliForUser,
		webhookUrl:   env.GithubWebhookUrl,
	}
}
