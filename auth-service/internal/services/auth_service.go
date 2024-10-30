package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/config"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/dto"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/models"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo   *repository.AuthRepository
	config *config.Config
}

func NewAuthService(repo *repository.AuthRepository, config *config.Config) *AuthService {
	return &AuthService{
		repo:   repo,
		config: config,
	}
}

// Register crée un nouveau compte utilisateur
func (s *AuthService) Register(req dto.RegisterRequest) error {
	// Vérifier si l'utilisateur existe déjà
	exists, err := s.repo.EmailExists(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		s.config.PasswordHashCost,
	)
	if err != nil {
		return err
	}

	// Créer l'utilisateur
	user := models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Active:    true,
	}

	return s.repo.CreateUser(user)
}

// Login authentifie un utilisateur et retourne les tokens
func (s *AuthService) Login(req dto.LoginRequest) (*dto.TokenResponse, error) {
	// Récupérer l'utilisateur
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Vérifier le mot de passe
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Générer les tokens
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	// Mettre à jour le refresh token en base
	err = s.repo.UpdateRefreshToken(user.ID, refreshToken)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.config.AccessTokenExpiry.Seconds()),
	}, nil
}

// RefreshToken renouvelle l'access token avec un refresh token valide
func (s *AuthService) RefreshToken(refreshToken string) (*dto.TokenResponse, error) {
	// Valider le refresh token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	// Récupérer l'utilisateur
	userID := uint(claims["user_id"].(float64))
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Vérifier que le refresh token correspond
	if user.RefreshToken != refreshToken {
		return nil, errors.New("refresh token has been revoked")
	}

	// Générer un nouveau access token
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken: accessToken,
		ExpiresIn:   int64(s.config.AccessTokenExpiry.Seconds()),
	}, nil
}

// Logout révoque le refresh token d'un utilisateur
func (s *AuthService) Logout(userID uint) error {
	return s.repo.UpdateRefreshToken(userID, "")
}

// Méthodes privées pour générer les tokens
func (s *AuthService) generateAccessToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(s.config.AccessTokenExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecret))
}

func (s *AuthService) generateRefreshToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(s.config.RefreshTokenExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecret))
}

// GetUserProfile récupère les informations de profil d'un utilisateur
func (s *AuthService) GetUserProfile(userID uint) (dto.UserResponse, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{Email: user.Email, Username: user.FirstName + " " + user.LastName}, nil
}

// UpdateProfile met à jour les informations de profil d'un utilisateur
func (s *AuthService) UpdateProfile(userID uint, req dto.UpdateProfileRequest) (dto.UserResponse, error) {
	user, err := s.repo.UpdateUser(userID, req)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		Email:    user.Email,
		Username: user.FirstName + " " + user.LastName,
	}, nil
}
