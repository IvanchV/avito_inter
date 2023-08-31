package handlers

import (
	"avito/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserService interface {
	ChangeUserSegment(id int, add []string, delete []string) error
	GetUserSegment(id int) ([]string, error)
}

type UserHandler struct {
	service UserService
}

func newUserHandler(router *mux.Router, service UserService) {
	u := &UserHandler{service: service}
	router.HandleFunc("/user_segment/{id}", makeHTTPHandleFunc(u.getUserSegment))
	router.HandleFunc("/change_segment/{id}", makeHTTPHandleFunc(u.changeUserSegment))
}

// @Summary ChangeUserSegment
// @Description Сhange users segments
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User_id"
// @Param arrays body models.ReqUser true "Arrays of delete or add"
// @Success 200 {object} ApiAnswer "Successful"
// @Failure 400 {object} ApiError "Bad request"
// @Router /change_segment/{id} [put]
func (s *UserHandler) changeUserSegment(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if r.Method == "PUT" {
		addReq := new(models.ReqUser)
		if err := json.NewDecoder(r.Body).Decode(&addReq); err != nil {
			return err
		}
		for _, name := range addReq.Add {
			for _, name1 := range addReq.Delete {
				if name == name1 {
					return writeJson(w, http.StatusBadRequest, ApiError{"intersection of arrays"})
				}
			}
		}
		err := s.service.ChangeUserSegment(id, addReq.Add, addReq.Delete)
		if err != nil {
			return writeJson(w, http.StatusBadRequest, ApiError{Error: fmt.Sprintf("%s", err)})
		}
		return writeJson(w, http.StatusOK, ApiAnswer{Ans: "Successful"})
	}
	return writeJson(w, http.StatusBadRequest, ApiError{Error: "Method not allowed"})
}

// @Summary GetUserSegment
// @Description Return all segments of user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User_id"
// @Success 200 {object} models.ResUser "Return User segments"
// @Failure 400 {object} ApiError "Error"
// @Router /user_segment/{id} [get]
func (s *UserHandler) getUserSegment(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		segments, err := s.service.GetUserSegment(id)
		if err != nil {
			return writeJson(w, http.StatusBadRequest, ApiError{Error: fmt.Sprintf("%s", err)})
		}
		userSegment := map[string]interface{}{
			"segments": segments,
		}
		return writeJson(w, http.StatusOK, userSegment)
	}
	return writeJson(w, http.StatusBadRequest, ApiError{Error: "Method not allowed"})
}
