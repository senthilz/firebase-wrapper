package firebase

import (
	"context"
	"errors"

	firebasev4 "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// App is the entry point for Firebase services
type App struct {
	auth *auth.Client
}

type Config struct {
	CredentialsFile string
	ProjectID       string
}

// New initializes the Firebase app with services
func New(ctx context.Context, cfg Config) (*App, error) {
	if cfg.CredentialsFile == "" {
		return nil, errors.New("credentials file path is required")
	}

	fbApp, err := firebasev4.NewApp(ctx, &firebasev4.Config{
		ProjectID: cfg.ProjectID,
	}, option.WithCredentialsFile(cfg.CredentialsFile))
	if err != nil {
		return nil, err
	}

	authClient, err := fbApp.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &App{
		auth: authClient,
	}, nil
}

// Auth returns the authentication service
func (a *App) Auth() *Auth {
	return &Auth{client: a.auth}
}

// Auth wraps Firebase Authentication
type Auth struct {
	client *auth.Client
}

func (a *Auth) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	if uid == "" {
		return nil, errors.New("uid cannot be empty")
	}
	return a.client.GetUser(ctx, uid)
}
