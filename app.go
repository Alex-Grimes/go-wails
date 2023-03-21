package main

import (
	"context"
	"encoding/json"
	"fmt"
)

const BaseUrl = "https://api.github.com"

var githubResponse APIResponse

type APIResponse []interface{}
type Gist struct {
	Description string      `json:"description"`
	Public      bool        `json:"public"`
	Files       interface{} `json:"files"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) GetPublicRepositories() (APIResponse, error) {
	url := fmt.Sprintf("%s/repositories", BaseUrl)
	response, err := MakeGetRequest(url, "")

	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &githubResponse)
	return githubResponse, nil
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
