// internal/api/handlers/handlers.go
package handlers

import (
	"encoding/json"
	"nba_api/internal/api/dto"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type PlayerHandler struct {
}

func NewHandler() *PlayerHandler {
	return &PlayerHandler{}
}

func (h *PlayerHandler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) error {
	playersData := []map[string]string{
		{
			"id":      "0a322e71-6696-43e6-88ae-114b1dc7da01",
			"name":    "Lebron James",
			"number":  "6",
			"height":  "2.06",
			"team_id": "aa5e5a81-fb2a-4004-915d-315781e523a6",
		},
		{
			"id":      "4f2ee54b-ef18-4a69-a71c-df309fc1ffeb",
			"name":    "Stephen Curry",
			"number":  "30",
			"height":  "1.91",
			"team_id": "a19cfc09-8e7d-4daa-8ce3-7d098df88f19",
		},
	}

	var players []dto.PlayerDTO
	for _, playerData := range playersData {
		id, err := uuid.Parse(playerData["id"])
		if err != nil {
			http.Error(w, "Erro na conversão dos dados", http.StatusInternalServerError)
			return err
		}

		name := playerData["name"]

		teamID, err := uuid.Parse(playerData["team_id"])
		if err != nil {
			http.Error(w, "Erro na conversão dos dados", http.StatusInternalServerError)
			return err
		}

		number, err := strconv.Atoi(playerData["number"])
		if err != nil {
			http.Error(w, "Erro na conversão dos dados", http.StatusInternalServerError)
			return err
		}

		height, err := strconv.ParseFloat(playerData["height"], 32)
		if err != nil {
			http.Error(w, "Erro na conversão dos dados", http.StatusInternalServerError)
			return err
		}

		player := dto.NewPlayerDTO(id, teamID, name, number, height)
		players = append(players, *player)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(players); err != nil {
		http.Error(w, "Erro na codificação JSON", http.StatusInternalServerError)
		return err
	}

	return nil
}

