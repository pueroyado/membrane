package controllers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"errors"
	"net/http"
)

type HandlerUser struct {
	userRepo models.UserRepository
}

type RegRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewHandlerUser(userRepo models.UserRepository) *HandlerUser {
	return &HandlerUser{
		userRepo: userRepo,
	}
}

func (req *RegRequest) Validate() error {
	if req.Email == "" {
		return errors.New("not set email")
	}
	if req.Password == "" {
		return errors.New("not set password")
	}
	return nil
}
func (a *AuthRequest) Validate() error {
	if a.Email == "" {
		return errors.New("not set email")
	}
	if a.Password == "" {
		return errors.New("not set password")
	}
	return nil
}

// Reg @BasePath /user
// @Tags User
// @Produce json
// @Accept json
// @Summary User registration
// @Param email body string true "john@mail.ru"
// @Param password body string true "secret"
// @Success 200 {array} models.JwtPayload "Successful operation"
// @Router /user/reg [post]
func (h *HandlerUser) Reg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		regRequest := RegRequest{}

		if err := json.NewDecoder(r.Body).Decode(&regRequest); err != nil {
			utils.ErrorMessage(w, 400, "Error parse request, detail: "+err.Error())
			return
		}
		if err := regRequest.Validate(); err != nil {
			utils.ErrorMessage(w, 400, "Error validate request: detail: "+err.Error())
			return
		}

		jwt, errReg := h.userRepo.Reg(regRequest.Email, regRequest.Password)
		if errReg != nil {
			utils.ErrorMessage(w, 400, "Error reg, : "+errReg.Error())
			return
		}
		byteData, _ := json.Marshal(jwt)

		w.Header().Set("Content-Type", "application/json")
		w.Write(byteData)
	}
}

// Auth @BasePath /user
// @Tags User
// @Produce json
// @Accept json
// @Summary User auth
// @Param email body string true "john@mail.ru"
// @Param password body string true "secret"
// @Success 200 {array} models.JwtPayload "Successful operation"
// @Router /user/auth [post]
func (h *HandlerUser) Auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authRequest := AuthRequest{}
		if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
			utils.ErrorMessage(w, 400, "Error parse request, detail: "+err.Error())
			return
		}
		if err := authRequest.Validate(); err != nil {
			utils.ErrorMessage(w, 400, "Error validate request: detail: "+err.Error())
			return
		}

		jwt, errAuth := h.userRepo.Auth(authRequest.Email, authRequest.Password)
		if errAuth != nil {
			utils.ErrorMessage(w, 401, errAuth.Error())
			return
		}

		byteData, _ := json.Marshal(jwt)
		w.Header().Set("Content-Type", "application/json")
		w.Write(byteData)
	}
}
