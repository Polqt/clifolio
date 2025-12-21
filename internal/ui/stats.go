package ui

import (
	"clifolio/internal/services"
	"clifolio/internal/ui/components"

)

type statsModel struct {
	stats 		*services.GitHubStats
	loading 	bool
	err			error
	spin 		components.SpinnerComponent
	width		int
	height		int
	username	string
}

type statsLoadedMsg struct {
	stats *services.GitHubStats
}

type statsErrorMsg struct {
	err error
}

type statsTickMsg struct{}

func StatsModel(username string) *statsModel {
	return &statsModel{
		
	}
}