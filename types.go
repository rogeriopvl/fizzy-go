package fizzy

type Board struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	AllAccess                bool   `json:"all_access"`
	CreatedAt                string `json:"created_at"`
	AutoPostponePeriodInDays int    `json:"auto_postpone_period_in_days,omitempty"`
	URL                      string `json:"url"`
	Creator                  User   `json:"creator"`
	PublicURL                string `json:"public_url,omitempty"`
}

type CreateBoardPayload struct {
	Name                     string `json:"name"`
	AllAccess                *bool  `json:"all_access,omitempty"`
	AutoPostponePeriodInDays int    `json:"auto_postpone_period_in_days,omitempty"`
	PublicDescription        string `json:"public_description,omitempty"`
}

type UpdateBoardPayload struct {
	Name                     string `json:"name,omitempty"`
	AllAccess                *bool  `json:"all_access,omitempty"`
	AutoPostponePeriodInDays *int   `json:"auto_postpone_period_in_days,omitempty"`
	PublicDescription        string `json:"public_description,omitempty"`
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
	HasAttachments  bool     `json:"has_attachments"`
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
	Limit            int // Maximum number of results to return (0 = no limit)
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
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	User                     User   `json:"user"`
	Slug                     string `json:"slug"`
	CardsCount               int    `json:"cards_count,omitempty"`
	CreatedAt                string `json:"created_at"`
	AutoPostponePeriodInDays int    `json:"auto_postpone_period_in_days,omitempty"`
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

type CreateAccessTokenPayload struct {
	Description string `json:"description"`
	Permission  string `json:"permission"`
}

type PersonalAccessToken struct {
	Token       string `json:"token"`
	Description string `json:"description"`
	Permission  string `json:"permission"`
}

type EntropyPayload struct {
	AutoPostponePeriodInDays int `json:"auto_postpone_period_in_days"`
}

type JoinCode struct {
	Code       string `json:"code"`
	UsageCount int    `json:"usage_count"`
	UsageLimit int    `json:"usage_limit"`
	URL        string `json:"url"`
	Active     bool   `json:"active"`
}

type UpdateJoinCodePayload struct {
	UsageLimit int `json:"usage_limit"`
}

type NotificationSettings struct {
	BundleEmailFrequency string `json:"bundle_email_frequency"`
}

type UpdateNotificationSettingsPayload struct {
	BundleEmailFrequency string `json:"bundle_email_frequency"`
}

// BoardAccess represents a single user's access status to a board.
type BoardAccess struct {
	User
	HasAccess   bool   `json:"has_access"`
	Involvement string `json:"involvement,omitempty"`
}

// BoardAccesses is the response from GetBoardAccesses, containing the board
// metadata and the aggregated list of users across all pages.
type BoardAccesses struct {
	BoardID   string        `json:"board_id"`
	AllAccess bool          `json:"all_access"`
	Users     []BoardAccess `json:"users"`
}

type Webhook struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	PayloadURL        string   `json:"payload_url"`
	Active            bool     `json:"active"`
	SigningSecret     string   `json:"signing_secret"`
	SubscribedActions []string `json:"subscribed_actions"`
	CreatedAt         string   `json:"created_at"`
	URL               string   `json:"url"`
	Board             Board    `json:"board"`
}

type Export struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	DownloadURL string `json:"download_url,omitempty"`
}

type WebhookDelivery struct {
	ID        string                   `json:"id"`
	State     string                   `json:"state"`
	CreatedAt string                   `json:"created_at"`
	UpdatedAt string                   `json:"updated_at"`
	Request   *WebhookDeliveryRequest  `json:"request"`
	Response  *WebhookDeliveryResponse `json:"response"`
	Event     WebhookDeliveryEvent     `json:"event"`
}

type WebhookDeliveryRequest struct {
	Headers map[string]string `json:"headers"`
}

type WebhookDeliveryResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

type WebhookDeliveryEvent struct {
	ID         string                         `json:"id"`
	Action     string                         `json:"action"`
	CreatedAt  string                         `json:"created_at"`
	Creator    User                           `json:"creator"`
	Eventable  WebhookDeliveryEventEventable  `json:"eventable"`
}

type WebhookDeliveryEventEventable struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	URL  string `json:"url"`
}

type CreateWebhookPayload struct {
	Name              string   `json:"name"`
	URL               string   `json:"url"`
	SubscribedActions []string `json:"subscribed_actions"`
}

type UpdateWebhookPayload struct {
	Name              string   `json:"name,omitempty"`
	SubscribedActions []string `json:"subscribed_actions,omitempty"`
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
