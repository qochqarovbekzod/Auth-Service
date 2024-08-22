package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"users/api/token"
	pb "users/generated/users"
	"users/model"

	"github.com/gin-gonic/gin"
)

// SignUpHandler handles the creation of a new User item.
// @Summary Create User
// @Description Create a new User item
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body users.SignUpRequest true "Create User"
// @Success 200 {object} users.SignUpResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/auth/register [post]
func (h Handler) SignUpHandler(ctx *gin.Context) {
	var req pb.SignUpRequest

	if err := ctx.ShouldBind(&req); err != nil {

		h.Log.Error("SignUpHandler errors ")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.UserRepo.SignUp(&req)
	if err != nil {
		h.Log.Error("SignUpHandler gkjfg error")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// LogInHandler handles user login.
// @Summary Login User
// @Description Login a user
// @Tags User
// @Accept json
// @Produce json
// @Param Create body users.LogInRequest true "Login"
// @Success 200 {object} users.LogInResponse "Login Successful"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/auth/login [post]
func (h Handler) LogInHandler(ctx *gin.Context) {
	var req pb.LogInRequest

	if err := ctx.ShouldBind(&req); err != nil {
		h.Log.Error("LogInHandler error")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.UserRepo.LogIn(&req)
	fmt.Println(resp)
	if err != nil {
		h.Log.Error("LogInHandler error")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := token.GenerateAccessJWT(resp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	refreshToken, err := token.GenerateAccessJWT(resp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, pb.LogInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// ViewProfileHandler handles fetching a user by ID.
// @Summary Get User by ID
// @Description Get a user by their ID
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} users.ViewProfileResponse "Get User Successful"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/auth/profile [get]
func (h Handler) ViewProfileHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println("mdfg", id)
	resp, err := h.UserRepo.ViewProfile(id)

	if err != nil {
		h.Log.Error("ViewProfileHandler error")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// EditProfile handles updating a user.
// @Summary Update User
// @Description Update an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Param Update body users.EditProfileRequeste true "Update"
// @Success 200 {object} users.EditProfileResponse "Update Successful"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/auth/profile [put]
func (h Handler) EditProfile(ctx *gin.Context) {
	var req pb.EditProfileRequeste

	if err := ctx.ShouldBind(&req); err != nil {
		h.Log.Error("EditProfile error")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.UserRepo.EditProfile(&req)

	if err != nil {
		h.Log.Error("EditProfile error")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// ChangeUserTypeHandler handles updating a user.
// @Summary Update User Type
// @Description Update user type
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Param Update body users.ChangeUserTypeRequeste true "Update"
// @Success 200 {object} users.ChangeUserTypeResponse "Update Successful"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/auth/type [put]
func (h Handler) ChangeUserTypeHandler(ctx *gin.Context) {
	var req pb.ChangeUserTypeRequeste

	if err := ctx.ShouldBind(&req); err != nil {
		h.Log.Error("ChangeUserTypeHandler error")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.UserRepo.ChangeUserType(&req)

	if err != nil {
		h.Log.Error("ChangeUserTypeHandler error")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetAllUsersHandler handles fetching all users with optional filters.
// @Summary Get all users
// @Description Get all users with optional filtering
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param user_name query string false "User Name"
// @Param password query string false "Password"
// @Param email query string false "Email"
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} users.GetAllUsersResponse "Get All Users Successful"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/auth/ [get]
func (h Handler) GetAllUsersHandler(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	var limit2, offset2 int
	var err error

	if limit != "" {
		limit2, err = strconv.Atoi(limit)
		if err != nil {
			h.Log.Error("GetAllUsersHandler error")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		limit2 = 5
	}

	if offset != "" {
		offset2, err = strconv.Atoi(offset)
		if err != nil {
			h.Log.Error("GetAllUsersHandler error")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	resp, err := h.UserRepo.GetAllUsers(&pb.GetAllUsersRequest{
		Offset: int32(offset2),
		Limit:  int32(limit2)})

	if err != nil {
		h.Log.Error("GetAllUsersHandler error")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// DeleteUserHandler handles deleting a user.
// @Summary Delete User
// @Description Delete an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} users.DeleteUserResponse "Delete Successful"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/auth/{id} [delete]
func (h Handler) DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.UserRepo.DeleteUser(id)

	if err != nil {
		h.Log.Error("DeleteUserHandler error")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// PasswordResetHandler handles password reset for a user.
// @Summary Reset User Password
// @Description Reset an existing user's password
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Param Update body users.PasswordResetRequest true "Reset Password"
// @Success 200 {object} users.PasswordResetResponse "Password Reset Successful"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/auth/reset-password [post]
func (h Handler) PasswordResetHandler(ctx *gin.Context) {
	var req pb.PasswordResetRequest

	if err := ctx.ShouldBind(&req); err != nil {
		h.Log.Error("PasswordResetHandler error")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.UserRepo.PasswordReset(&req)

	if err != nil {
		h.Log.Error("PasswordResetHandler error")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// TokenGenerationHandler handles token generation for a user.
// @Summary Generate Token
// @Description Generate a new token for an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Authorization header users.TokenGenerationRequest true "Generate Token"
// @Success 200 {object} users.TokenGenerationResponse "Token Generation Successful"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/auth/refresh [post]
func (h Handler) RefreshToken(c *gin.Context) {

	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		h.Log.Error("ExtractClaims error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
		return
	}

	claims, err := token.ExtractClaims(refreshToken)
	if err != nil {
		h.Log.Error("ExtractClaims error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		h.Log.Error("ExtractClaims error")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		return
	}

	newAccessToken, err := token.GenerateAccessJWT(&model.LoginResponse{
		Id:       claims.UserId,
		Username: claims.Username,
		Email:    claims.Email,
	})
	if err != nil {
		h.Log.Error("ExtractClaims error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
		return
	}
	h.Log.Info("ishladi")
	c.JSON(http.StatusOK, pb.TokenGenerationResponse{
		AccessToken:  newAccessToken,
		RefreshToken: refreshToken,
	})
}
