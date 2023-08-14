package types

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//===========================================================================

type UserModel struct {
	ID          int64
	Nickname    string
	FirstName   sql.NullString
	LastName    sql.NullString
	Password    string
	Location    sql.NullString
	Description sql.NullString
	AvatarLink  sql.NullString
	SearchFor   sql.NullString
	ShowPeople  int
	OldMatches  sql.NullString
	NewMatches  sql.NullString
	LastLogin   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUserModel() *UserModel {
	return &UserModel{
		ShowPeople: 5,
		LastLogin:  time.Now(),
		CreatedAt:  time.Now(),
	}
}

//==========================================================================

func (userModel *UserModel) SetNickName(nick string) *UserModel {
	userModel.Nickname = strings.ToLower(nick)
	return userModel
}

func (userModel *UserModel) SetPassword(password string) (*UserModel, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userModel.Password = string(hashedPassword)
	return userModel, nil
}

func (userModel *UserModel) SetLocation(location string) {
	userModel.Location = sql.NullString{String: location, Valid: true}
}

func (userModel *UserModel) SetNames(first, last string) {
	userModel.FirstName = sql.NullString{String: first, Valid: true}
	userModel.LastName = sql.NullString{String: last, Valid: true}
}

func (userModel *UserModel) SetSearchPref(option string) {
	userModel.SearchFor = sql.NullString{String: option, Valid: true}
}

func (userModel *UserModel) SetDescription(description string) {
	userModel.Description = sql.NullString{String: description, Valid: true}
}

//==========================================================================

func (userModel *UserModel) SetID(id int64) *UserModel {
	userModel.ID = id
	return userModel
}

//==========================================================================

func (userModel *UserModel) Validate(database *Database) error {
	if userModel.Nickname == "" {
		return errors.New("nickname is required")
	}
	//_, err := mail.ParseAddress(userModel.Nickname)
	//if err != nil {
	//	return errors.New("invalid nickname address")
	//}
	if userModel.Password == "" {
		return errors.New("password is required")
	}
	query := `SELECT COUNT(*) FROM "user" WHERE nickname = $1`
	var count int
	err := database.Connection.QueryRow(query, userModel.Nickname).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user already exists")
	}
	return nil
}

//==========================================================================

func (userModel *UserModel) Insert(database *Database) error {
	statement := `INSERT INTO "user" (nickname, password,first_name,last_name,location,description,search_for,show_people,last_login) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	err := database.Connection.QueryRow(statement,
		userModel.Nickname, userModel.Password, userModel.FirstName, userModel.LastName, userModel.Location,
		userModel.Description, userModel.SearchFor, userModel.ShowPeople, time.Now(),
	).Scan(&userModel.ID)
	if err != nil {
		return err
	}
	return nil
}

//==========================================================================

func (userModel *UserModel) Update(database *Database) error {
	statement := `
UPDATE "user"
SET 
   password = $1,
   location = $2,
   description = $3,
   avatar_link = $4,
   search_for = $5,
   show_people = $6,
   old_matches = $7,
   new_matches = $8,
   last_login = $9,
   first_name = $10,
   last_name = $11
WHERE
   id = $12;
	`
	_, err := database.Connection.Exec(statement,
		userModel.Password,
		userModel.Location,
		userModel.Description,
		userModel.AvatarLink,
		userModel.SearchFor,
		userModel.ShowPeople,
		userModel.OldMatches,
		userModel.NewMatches,
		userModel.LastLogin,
		userModel.FirstName,
		userModel.LastName,
		userModel.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

//==========================================================================

func (userModel *UserModel) FindByEmail(database *Database, nickname string) (*UserModel, error) {
	query := `SELECT id, nickname, password FROM "user" WHERE nickname = $1`
	row := database.Connection.QueryRow(query, nickname)
	err := row.Scan(&userModel.ID, &userModel.Nickname, &userModel.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return userModel, nil
}

//==========================================================================

func (userModel *UserModel) DeleteSessionsByUser(database *Database) error {
	query := `DELETE FROM session WHERE user_id = $1`
	_, err := database.Connection.Exec(query, userModel.ID)
	if err != nil {
		return err
	}
	return nil
}

//==========================================================================

func (userModel *UserModel) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}

//==========================================================================

func (userModel *UserModel) Auth(c *gin.Context, database *Database) error {
	sessionToken, err := c.Cookie(os.Getenv("SESSION_TOKEN_KEY"))
	if err != nil {
		return err
	}
	sessionModel := NewSessionModel()
	err = sessionModel.FindByToken(database, sessionToken)
	if err != nil {
		return err
	}
	userModel.SetID(sessionModel.UserID)
	trueUser := NewUserModel()
	err = trueUser.FindById(database, sessionModel.UserID)
	if err != nil {
		return err
	}
	if userModel.ID != trueUser.ID {
		return err
	}
	userModel.SetNickName(trueUser.Nickname)
	_, err = userModel.SetPassword(trueUser.Password)
	if err != nil {
		return fmt.Errorf("cant auth err: %w", err)
	}

	*userModel = *trueUser

	return nil
}

//==========================================================================

func (userModel *UserModel) FindById(database *Database, userId int64) error {
	query := `SELECT id, nickname, password,first_name,last_name,location,
       description,avatar_link	,search_for,show_people,old_matches,new_matches,last_login FROM "user" WHERE id = $1`
	row := database.Connection.QueryRow(query, userId)
	err := row.Scan(&userModel.ID, &userModel.Nickname, &userModel.Password, &userModel.FirstName, &userModel.LastName, &userModel.Location,
		&userModel.Description, &userModel.AvatarLink, &userModel.SearchFor, &userModel.ShowPeople, &userModel.OldMatches, &userModel.NewMatches, &userModel.LastLogin)
	if err != nil {
		return err
	}
	return nil
}

//==========================================================================

// ==========================================================================
func SearchPreferenceLabel(value sql.NullString) string {
	if !value.Valid {
		return ""
	}
	switch value.String {
	case "friend_online":
		return "Friend to meet online"
	case "friend_offline":
		return "Friend to meet offline"
	case "help_project":
		return "Help with my pet project"
	default:
		return ""
	}
}

func LocationLabel(value sql.NullString) string {
	if !value.Valid {
		return ""
	}
	switch value.String {
	case "Georgia":
		return "Georgia"
	case "Serbia":
		return "Serbia"
	//... другие значения
	default:
		return "Someplace else"
	}
}
