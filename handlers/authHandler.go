package handlers

import (
	"jwt_auth_golang/modules/auth"
)

// ProfileHandler struct
type profileHandler struct {
	Rd auth.AuthInterface
	Tk auth.TokenInterface
}

func NewProfile(rd auth.AuthInterface, tk auth.TokenInterface) *profileHandler {
	return &profileHandler{rd, tk}
}
