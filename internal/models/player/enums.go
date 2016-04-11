package player

type Platform int
type Activity int
type LookingFor int
type StoryLevel int
type DZLevel int

const (
	PC_PLATFORM Platform = iota + 1
	PS4_PLATFORM
	XBOX_PLATFORM
)

const (
	DAILY_CHALLENGING Activity = iota + 1
	DAILY_HARD
	DARKZONE_FARMING
	DARKZONE_ROGUE
	STORY_NORMAL
	STORY_HARD
	STORY_CHALLENGE
	FREE_FARMING
	FREE_FREE
	FREE_SIDEMISSIONS
	TRADING
)

const (
	LOOKING_GROUP LookingFor = iota + 1
	LOOKING_FRIENDS
)

const (
	SL_FIRST_TIER StoryLevel = iota + 1
	SL_SECOND_TIER
	SL_THIRD_TIER
)

const (
	DZ_FIRST_TIER DZLevel = iota + 1
	DZ_SECOND_TIER
	DZ_THIRD_TIER
)

func (t Platform) String() string {
	switch t {
	case XBOX_PLATFORM:
		return "Xbox"
	case PS4_PLATFORM:
		return "PlayStation 4"
	case PC_PLATFORM:
		return "PC"
	default:
		return ""
	}
}

func (t Activity) String() string {
	switch t {
	case DAILY_CHALLENGING:
		return "Misión Diaria — Desafiante"
	case DAILY_HARD:
		return "Misión Diaria — Difícil"
	case DARKZONE_FARMING:
		return "Zona Oscura — Farmeo / Leveleo"
	case DARKZONE_ROGUE:
		return "Zona Oscura — Modo Renegado"
	case STORY_NORMAL:
		return "Modo Historia — Niveles dificultad Normal"
	case STORY_HARD:
		return "Modo Historia — Niveles dificultad Difícil"
	case STORY_CHALLENGE:
		return "Modo Historia — Niveles dificultad Desafiante"
	case FREE_FARMING:
		return "Modo Libre — Farmeo / Leveleo"
	case FREE_FREE:
		return "Modo Libre"
	case FREE_SIDEMISSIONS:
		return "Modo Libre — Misiones Secundarias"
	case TRADING:
		return "Intercambio / Trading"
	default:
		return ""
	}
}

func (t LookingFor) String() string {
	switch t {
	case LOOKING_GROUP:
		return "Buscando Equipo"
	case LOOKING_FRIENDS:
		return "Buscando Amistad"
	default:
		return ""
	}
}

func (t StoryLevel) String() string {
	switch t {
	case SL_FIRST_TIER:
		return "Nivel entre 1 y 10"
	case SL_SECOND_TIER:
		return "Nivel entre 11 y 20"
	case SL_THIRD_TIER:
		return "Nivel entre 21 y 30"
	default:
		return ""
	}
}

func (t DZLevel) String() string {
	switch t {
	case DZ_FIRST_TIER:
		return "Nivel menor a 30"
	case DZ_SECOND_TIER:
		return "Nivel entre 30 y 50"
	case DZ_THIRD_TIER:
		return "Nivel mayor a 50"
	default:
		return ""
	}
}
