package handlers

import (
	"avito/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type SegmentService interface {
	CreateSegment(seg *models.Segment) error
	DeleteSegment(name string) error
}

type SegmentHandler struct {
	service SegmentService
}

func newSegmentHandler(router *mux.Router, service SegmentService) {
	s := &SegmentHandler{service: service}
	router.HandleFunc("/create_segment", makeHTTPHandleFunc(s.createSegment))
	router.HandleFunc("/delete_segment", makeHTTPHandleFunc(s.deleteSegment))
}

// @Summary CreateSegment
// @Description Create a new segment
// @Tags Segments
// @Accept json
// @Produce json
// @Param segment body models.Segment true "Segment name"
// @Success 200 {object} ApiAnswer "Segment created"
// @Failure 400 {object} ApiError "Bad request"
// @Router /create_segment [post]
func (s *SegmentHandler) createSegment(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		createSegmentReq := new(models.Segment)
		if err := json.NewDecoder(r.Body).Decode(createSegmentReq); err != nil {
			return err
		}
		segment := models.NewSegment(createSegmentReq.Name)
		if err := s.service.CreateSegment(segment); err != nil {
			return writeJson(w, http.StatusBadRequest, ApiError{Error: fmt.Sprintf("%v", err)})
		}
		return writeJson(w, http.StatusOK, ApiAnswer{Ans: "Segment created"})
	}
	return writeJson(w, http.StatusOK, ApiError{Error: fmt.Sprintf("Method not allowed")})

}

// @Summary DeleteSegment
// @Description Create a new segment
// @Tags Segments
// @Accept json
// @Produce json
// @Param segment body models.Segment true "Segment name"
// @Success 200 {object} ApiAnswer "Segment deleted"
// @Failure 400 {object} ApiError "Bad request"
// @Router /delete_segment [delete]
func (s *SegmentHandler) deleteSegment(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "DELETE" {
		deletedReq := new(models.Segment)
		if err := json.NewDecoder(r.Body).Decode(&deletedReq); err != nil {
			return err
		}
		err := s.service.DeleteSegment(deletedReq.Name)
		if err != nil {
			return writeJson(w, http.StatusBadRequest, ApiError{Error: fmt.Sprintf("%v", err)})
		}
		return writeJson(w, http.StatusOK, ApiAnswer{Ans: fmt.Sprintf("Segment deleted")})
	}
	return writeJson(w, http.StatusBadRequest, ApiError{Error: fmt.Sprintf("Method not allowed")})
}
