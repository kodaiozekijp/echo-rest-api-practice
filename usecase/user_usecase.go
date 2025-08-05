package usecase

import (
	"echo-rest-api-practice/entities"
	"echo-rest-api-practice/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user entities.User) (entities.UserResponse, error)
	Login(user entities.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

// userUseCaseを生成する関数
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

// 引数で受け取ったUserの登録を制御する関数
func (uu *userUsecase) SignUp(user entities.User) (entities.UserResponse, error) {
	// 受け取ったPasswordをハッシュ化する
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return entities.UserResponse{}, err
	}
	// 受け取ったEmailとハッシュ化したPasswordから新しいUserを作成する
	newUser := entities.User{Email: user.Email, Password: string(hash)}
	// 作成したUserを用いてレポジトリ層の登録処理を呼び出す
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return entities.UserResponse{}, err
	}
	// 登録に成功した場合はUserResponseとして返却する
	resUser := entities.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

// ログインを制御する関数
func (uu *userUsecase) Login(user entities.User) (string, error) {
	// 要求されたEmailを用いてレポジトリ層のユーザ取得処理を呼び出す
	storedUser := entities.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// 要求されたPasswordを用いてDBに保存されているハッシュ化されたPasswordと比較する
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// パスワードが一致した場合はJWTの生成を行う
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
