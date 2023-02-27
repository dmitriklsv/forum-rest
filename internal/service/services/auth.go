package services

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"forum/internal/entity"
	"forum/internal/service"

	"github.com/gofrs/uuid"
)

const salt = "9j0oigjio3t8kllmcvblj31"

type authService struct {
	userRepo    service.UserRepo
	sessionRepo service.SessionRepo
}

func NewAuthService(uRepo service.UserRepo, sRepo service.SessionRepo) *authService {
	log.Println("| | authentication service is done!")
	return &authService{
		userRepo:    uRepo,
		sessionRepo: sRepo,
	}
}

func (a *authService) GetUser(ctx context.Context, id uint64) (entity.User, error) {
	return a.userRepo.FindByID(ctx, id)
}

func (a *authService) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	foundUser, _ := a.userRepo.FindOne(ctx, entity.User{
		Email:    user.Email,
		Username: user.Username,
		Password: generatePasswordHash(user.Password),
	})

	if foundUser.Email == user.Email {
		return -1, fmt.Errorf("email alredy exists")
	}
	if foundUser.Username == user.Username {
		return -1, fmt.Errorf("username already exists")
	}

	user.Password = generatePasswordHash(user.Password)
	return a.userRepo.CreateUser(ctx, user)
}

func (a *authService) UpdateSession(ctx context.Context, session entity.Session) (entity.Session, error) {
	s, err := generateSession(session.UserID)
	if err != nil {
		return entity.Session{}, err
	}
	s.ID = session.ID

	return a.sessionRepo.UpdateSession(ctx, s)
}

func (a *authService) SetSession(ctx context.Context, user entity.User) (entity.Session, error) {
	user, err := a.userRepo.FindOne(ctx, entity.User{
		Email:    user.Email,
		Username: user.Username,
		Password: generatePasswordHash(user.Password),
	})
	if err != nil {
		return entity.Session{}, fmt.Errorf("incorrect username or password")
	}
	// fmt.Println(u)

	if err := a.sessionRepo.DeleteSession(ctx, user.ID); err != nil {
		return entity.Session{}, err
	}

	session, err := generateSession(user.ID)
	if err != nil {
		return entity.Session{}, err
	}

	if err := a.sessionRepo.CreateSession(ctx, session); err != nil {
		return entity.Session{}, err
	}

	return session, nil
}

func (a *authService) GetSession(ctx context.Context, sessionToken string) (entity.Session, error) {
	session, err := a.sessionRepo.GetSession(ctx, sessionToken)
	if err != nil {
		return entity.Session{}, err
	}

	if isExpired(session) {
		// fmt.Println("expried")
		if err := a.sessionRepo.DeleteSession(ctx, session.UserID); err != nil {
			return entity.Session{}, err
		}
		return entity.Session{}, fmt.Errorf("session expired")
	}

	return session, nil
}

func isExpired(session entity.Session) bool {
	return time.Now().After(session.ExpireTime)
}

func generateSession(userID uint64) (entity.Session, error) {
	sessionToken, err := uuid.NewV4()
	if err != nil {
		return entity.Session{}, err
	}

	return entity.Session{
		UserID:     userID,
		Token:      sessionToken.String(),
		ExpireTime: time.Now().Add(12 * time.Hour),
	}, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
