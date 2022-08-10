// Create package handlers here ...
package handlers

import (
	profiledto "dumbmerch/dto/profile"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Declare handlerProfile struct here ...
type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

// HandlerProfile function here ...
func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

// Create GetProfile method here ...
func (h *handlerProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var profile models.Profile
	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProfile(profile)}
	json.NewEncoder(w).Encode(response)
}

// Create convertResponseProfile function here ...
func convertResponseProfile(u models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{
		ID:      u.ID,
		Phone:   u.Phone,
		Gender:  u.Gender,
		Address: u.Address,
		UserID:  u.UserID,
		User:    u.User,
	}
}
