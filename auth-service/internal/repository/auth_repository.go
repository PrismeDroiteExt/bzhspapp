package repository

import (
	"errors"

	"github.com/prismedroiteext/breizhsport/auth-service/internal/dto"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// CreateUser crée un nouvel utilisateur
func (r *AuthRepository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}

// GetUserByEmail récupère un utilisateur par son email
func (r *AuthRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := r.db.Where("email = ? AND active = ?", email, true).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, result.Error
	}
	return user, nil
}

// GetUserByID récupère un utilisateur par son ID
func (r *AuthRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, result.Error
	}
	return user, nil
}

// EmailExists vérifie si un email existe déjà
func (r *AuthRepository) EmailExists(email string) (bool, error) {
	var count int64
	result := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0, result.Error
}

// UpdateRefreshToken met à jour le refresh token d'un utilisateur
func (r *AuthRepository) UpdateRefreshToken(userID uint, token string) error {
	result := r.db.Model(&models.User{}).Where("id = ?", userID).Update("refresh_token", token)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// UpdateLastLogin met à jour la date de dernière connexion
func (r *AuthRepository) UpdateLastLogin(userID uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("last_login", gorm.Expr("NOW()")).Error
}

// DeactivateUser désactive un compte utilisateur
func (r *AuthRepository) DeactivateUser(userID uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("active", false).Error
}

// UpdatePassword met à jour le mot de passe d'un utilisateur
func (r *AuthRepository) UpdatePassword(userID uint, hashedPassword string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("password", hashedPassword).Error
}

// UpdateUser met à jour les informations d'un utilisateur et retourne l'utilisateur mis à jour
func (r *AuthRepository) UpdateUser(userID uint, req dto.UpdateProfileRequest) (models.User, error) {
	var user models.User
	result := r.db.Model(&models.User{}).Where("id = ?", userID).Updates(req).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
