package player

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	DatastoreGroup    = "TheDivisionPlayers"
	DeletionURLFormat = "/r/%s"
)

type Player struct {
	Username     string     `json:"username"`
	Platform     Platform   `json:"platform"`
	Activity     Activity   `json:"activity"`
	Microphone   bool       `json:"microphone"`
	LookingFor   LookingFor `json:"lookingfor"`
	StoryLevel   int        `json:"level"`
	DZLevel      int        `json:"dzlevel"`
	Description  string     `json:"description"`
	Signup       time.Time  `json:"signup"`
	IPAddress    string     `json:"-"`
	ForwardedFor string     `json:"-"`
	DeletionID   string     `json:"-"`
}

var (
	ErrUsernameContainsSpaces = fmt.Errorf("El nombre de usuario contiene espacios.")
	ErrUsernameTooShort       = fmt.Errorf("El nombre de usuario es demasiado corto.")
	ErrUsernameTooLong        = fmt.Errorf("El nombre de usuario es demasiado largo.")
	ErrBadPlatformSelected    = fmt.Errorf("La plataforma elegida es desconocida.")
	ErrBadActivitySelected    = fmt.Errorf("La actividad seleccionada es desconocida.")
	ErrBadLookingForSelected  = fmt.Errorf("No se indicó una intención correcta.")
	ErrBadStoryLevelSelected  = fmt.Errorf("No se seleccionó un tramo de nivel del modo historia correcto.")
	ErrBadDZLevelSelected     = fmt.Errorf("No se seleccionó un tramo de nivel de la Zona Oscura corecto.")
	ErrDescriptionTooShort    = fmt.Errorf("La descripción es demasiado corta.")
	ErrDescriptionTooLong     = fmt.Errorf("La descripción es demasiado larga.")
)

func (p Player) Validate() error {
	// Check if the username contains spaces
	if strings.Contains(p.Username, " ") {
		return ErrUsernameContainsSpaces
	}

	// Check if the username is too short
	if utf8.RuneCountInString(p.Username) < 3 {
		return ErrUsernameTooShort
	}

	// Check if the username is too long
	if utf8.RuneCountInString(p.Username) > 50 {
		return ErrUsernameTooLong
	}

	// Check if Platform was casted
	if p.Platform.String() == "" {
		return ErrBadPlatformSelected
	}

	// Check if Activity was casted
	if p.Activity.String() == "" {
		return ErrBadActivitySelected
	}

	// Check if LookingFor was casted
	if p.LookingFor.String() == "" {
		return ErrBadLookingForSelected
	}

	// Check if StoryLevel was casted
	if p.StoryLevel == 0 || p.StoryLevel > 30 {
		return ErrBadStoryLevelSelected
	}

	// Check if DZLevel was casted
	if p.DZLevel == 0 || p.DZLevel > 99 {
		return ErrBadDZLevelSelected
	}

	// Check if the length of the description is less than optimal
	if utf8.RuneCountInString(p.Description) < 5 {
		return ErrDescriptionTooShort
	}

	// Check if the length of the description is longer than optimal
	if utf8.RuneCountInString(p.Description) > 200 {
		return ErrDescriptionTooLong
	}

	return nil
}

func (p Player) GetHash() string {
	hash := fmt.Sprintf("%v|%v", p.Username, p.Platform.String())
	return strings.ToLower(hash)
}
