package fizzy

type Board struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AllAccess bool   `json:"all_access"`
	CreatedAt string `json:"created_at"`
	URL       string `json:"url"`
	Creator   User   `json:"creator"`
}

type CreateBoardPayload struct {
	Name               string `json:"name"`
	AllAccess          bool   `json:"all_access"`
	AutoPostponePeriod int    `json:"auto_postpone_period"`
	PublicDescription  string `json:"public_description"`
}

type UpdateBoardPayload struct {
	Name               string `json:"name,omitempty"`
	AllAccess          *bool  `json:"all_access,omitempty"`
	AutoPostponePeriod *int   `json:"auto_postpone_period,omitempty"`
	PublicDescription  string `json:"public_description,omitempty"`
}

type Column struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Color     ColorObject `json:"color"`
	CreatedAt string      `json:"created_at"`
}

type ColorObject struct {
	Name  string `json:"name"`
	Value Color  `json:"value"`
}

type CreateColumnPayload struct {
	Name  string `json:"name"`
	Color *Color `json:"color,omitempty"`
}

type UpdateColumnPayload struct {
	Name  string `json:"name,omitempty"`
	Color *Color `json:"color,omitempty"`
}

type Card struct {
	ID              string   `json:"id"`
	Number          int      `json:"number"`
	Title           string   `json:"title"`
	Status          string   `json:"status"`
	Description     string   `json:"description"`
	DescriptionHTML string   `json:"description_html"`
	ImageURL        string   `json:"image_url"`
	Tags            []string `json:"tags"`
	Closed          bool     `json:"closed"`
	Golden          bool     `json:"golden"`
	LastActiveAt    string   `json:"last_active_at"`
	CreatedAt       string   `json:"created_at"`
	URL             string   `json:"url"`
	Board           Board    `json:"board"`
	Column          *Column  `json:"column,omitempty"`
	Creator         User     `json:"creator"`
	CommentsURL     string   `json:"comments_url"`
	Steps           []Step   `json:"steps,omitempty"`
}

type CardFilters struct {
	BoardIDs         []string
	TagIDs           []string
	AssigneeIDs      []string
	CreatorIDs       []string
	CloserIDs        []string
	CardIDs          []string
	IndexedBy        string
	SortedBy         string
	AssignmentStatus string
	CreationStatus   string
	ClosureStatus    string
	Terms            []string
}

type CreateCardPayload struct {
	Title        string   `json:"title"`
	Description  string   `json:"description,omitempty"`
	Status       string   `json:"status,omitempty"`
	ImageURL     string   `json:"image_url,omitempty"`
	TagIDS       []string `json:"tag_ids,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	LastActiveAt string   `json:"last_active_at,omitempty"`
}

type UpdateCardPayload struct {
	Title        string   `json:"title,omitempty"`
	Description  string   `json:"description,omitempty"`
	Status       string   `json:"status,omitempty"`
	TagIDS       []string `json:"tag_ids,omitempty"`
	LastActiveAt string   `json:"last_active_at,omitempty"`
}

type GetMyIdentityResponse struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	User      User   `json:"user"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email_address"`
	Role      string `json:"role"`
	Active    bool   `json:"active"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	URL       string `json:"url"`
}

type Notification struct {
	ID        string        `json:"id"`
	Read      bool          `json:"read"`
	ReadAt    string        `json:"read_at"`
	CreatedAt string        `json:"created_at"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
	Creator   User          `json:"creator"`
	Card      CardReference `json:"card"`
	URL       string        `json:"url"`
}

type CardReference struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type Tag struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	URL       string `json:"url"`
}

type Comment struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Body      struct {
		PlainText string `json:"plain_text"`
		HTML      string `json:"html"`
	} `json:"body"`
	Creator      User          `json:"creator"`
	Card         CardReference `json:"card"`
	ReactionsURL string        `json:"reactions_url"`
	URL          string        `json:"url"`
}

type Reaction struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Reacter User   `json:"reacter"`
	URL     string `json:"url"`
}

// Step represents a checklist item on a card.
type Step struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

// Color represents a CSS color value for columns.
type Color string

const (
	ColorBlue   Color = "var(--color-card-default)"
	ColorGray   Color = "var(--color-card-1)"
	ColorTan    Color = "var(--color-card-2)"
	ColorYellow Color = "var(--color-card-3)"
	ColorLime   Color = "var(--color-card-4)"
	ColorAqua   Color = "var(--color-card-5)"
	ColorViolet Color = "var(--color-card-6)"
	ColorPurple Color = "var(--color-card-7)"
	ColorPink   Color = "var(--color-card-8)"
)

func AllColors() []Color {
	return []Color{
		ColorBlue,
		ColorGray,
		ColorTan,
		ColorYellow,
		ColorLime,
		ColorAqua,
		ColorViolet,
		ColorPurple,
		ColorPink,
	}
}
