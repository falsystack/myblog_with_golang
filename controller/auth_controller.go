package controller

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
	"time"
	"toyproject_recruiting_community/usecases"
	"toyproject_recruiting_community/usecases/input"
)

var googleOauthConfig oauth2.Config

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type GoogleOAuth struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

type AuthController interface {
	GoogleLogin(ctx *gin.Context)
	GoogleAuthCallback(ctx *gin.Context)
	Logout(ctx *gin.Context)
	// TODO: Create
}

func NewAuthController(au usecases.AuthUsecase) AuthController {
	log.Println("init auth controller", os.Getenv("GOOGLE_CLIENT_ID"))
	googleOauthConfig = oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	return &authController{au: au}
}

type authController struct {
	au usecases.AuthUsecase
}

func (a *authController) GoogleAuthCallback(ctx *gin.Context) {
	oauthstate, err := ctx.Cookie("oauthstate")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	state := ctx.Request.URL.Query().Get("state")
	if state != oauthstate {
		log.Println("invalid oauth state", state)
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := ctx.Request.URL.Query().Get("code")
	token, err := googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var goauth GoogleOAuth
	err = json.NewDecoder(resp.Body).Decode(&goauth)
	if err != nil {
		log.Println(err)
	}

	// TODO: UserがDBにないと登録する
	_, err = a.au.FindByID(goauth.ID)
	if err != nil {
		err = a.au.Create(&input.CreateUser{
			ID:    goauth.ID,
			Email: goauth.Email,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// jwtを生成
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   goauth.ID,
		"email": goauth.Email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	signedJwt, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": &signedJwt})
}

func (a *authController) GoogleLogin(ctx *gin.Context) {
	uid := uuid.New().String()
	state := base64.URLEncoding.EncodeToString([]byte(uid))

	cookie := &http.Cookie{
		Name:   "oauthstate",
		Value:  state,
		MaxAge: int(time.Hour.Milliseconds() * 2),
	}
	http.SetCookie(ctx.Writer, cookie)
	permissionRequiredURL := googleOauthConfig.AuthCodeURL(state)
	ctx.Redirect(http.StatusTemporaryRedirect, permissionRequiredURL)
}

func (a *authController) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
