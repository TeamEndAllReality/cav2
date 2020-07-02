package cav2

//Curseforge api

//Addon of curseforge
type Addon struct {
	Attachments              []AddonAttachments   `json:"attachments"`
	Authors                  []AddonAuthors       `json:"authors"`
	AvatarURL                interface{}          `json:"avatarUrl"`
	Categories               []AddonCategories    `json:"categories"`
	CategoryList             string               `json:"categoryList"`
	CategorySection          AddonCategorySection `json:"categorySection"`
	ClientURL                string               `json:"clientUrl"`
	CommentCount             int                  `json:"commentCount"`
	DateCreated              string               `json:"dateCreated"`
	DateModified             string               `json:"dateModified"`
	DateReleased             string               `json:"dateReleased"`
	DefaultFileID            int                  `json:"defaultFileId"`
	DonationURL              interface{}          `json:"donationUrl"`
	DownloadCount            float64              `json:"downloadCount"`
	ExternalURL              interface{}          `json:"externalUrl"`
	FullDescription          string               `json:"fullDescription"`
	GameID                   int                  `json:"gameId"`
	GameName                 string               `json:"gameName"`
	GamePopularityRank       int                  `json:"gamePopularityRank"`
	GameVersionLatestFiles   []AddonGameVersion   `json:"gameVersionLatestFiles"`
	ID                       int                  `json:"id"`
	InstallCount             int                  `json:"installCount"`
	IsAvailable              bool                 `json:"isAvailable"`
	IsFeatured               bool                 `json:"isFeatured"`
	LatestFiles              []AddonLatestFile    `json:"latestFiles"`
	Likes                    int                  `json:"likes"`
	Name                     string               `json:"name"`
	PackageType              int                  `json:"packageType"`
	PopularityScore          float64              `json:"popularityScore"`
	PortalName               string               `json:"portalName"`
	PrimaryAuthorName        string               `json:"primaryAuthorName"`
	PrimaryCategoryAvatarURL string               `json:"primaryCategoryAvatarUrl"`
	PrimaryCategoryName      string               `json:"primaryCategoryName"`
	Rating                   int                  `json:"rating"`
	SectionName              string               `json:"sectionName"`
	Slug                     string               `json:"slug"`
	Stage                    int                  `json:"stage"`
	Status                   int                  `json:"status"`
	Summary                  string               `json:"summary"`
	WebsiteURL               string               `json:"websiteUrl"`
}

//AddonAttachments embedded attachments on the curseforge page (images,etc)
type AddonAttachments struct {
	Description  string `json:"description"`
	ID           int    `json:"id"`
	IsDefault    bool   `json:"isDefault"`
	ProjectID    int    `json:"projectID"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}

//AddonAuthors Represents an author of an addon
type AddonAuthors struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

//AddonCategories represents the category or an addon
type AddonCategories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

//AddonCategorySection Section of addon
type AddonCategorySection struct {
	ID                      int         `json:"Id"`
	ExtraIncludePattern     interface{} `json:"extraIncludePattern"`
	GameID                  int         `json:"gameId"`
	InitialInclusionPattern string      `json:"initialInclusionPattern"`
	Name                    string      `json:"name"`
	PackageType             int         `json:"packageType"`
	Path                    string      `json:"path"`
}

//AddonGameVersion AddonGameVersion
type AddonGameVersion struct {
	FileType        int    `json:"fileType"`
	GameVersion     string `json:"gameVersion"`
	ProjectFileID   int    `json:"projectFileId"`
	ProjectFileName string `json:"projectFileName"`
}

//AddonLatestFile Latest file uploads
type AddonLatestFile struct {
	AlternateFileID    int               `json:"alternateFileId"`
	Dependencies       []AddonDependency `json:"dependencies"`
	DownloadURL        string            `json:"downloadUrl"`
	FileDate           string            `json:"fileDate"`
	FileName           string            `json:"fileName"`
	FileNameOnDisk     string            `json:"fileNameOnDisk"`
	FileStatus         int               `json:"fileStatus"`
	GameVersion        []string          `json:"gameVersion"`
	ID                 int               `json:"id"`
	IsAlternate        bool              `json:"isAlternate"`
	IsAvailable        bool              `json:"isAvailable"`
	Modules            []AddonModules    `json:"modules"`
	PackageFingerprint int               `json:"packageFingerprint"`
	ReleaseType        int               `json:"releaseType"`
}

//AddonDependency a dependency of an addon
type AddonDependency struct {
	AddonID int `json:"addonId"`
	Type    int `json:"type"`
}

//AddonModules Folders in addon File
type AddonModules struct {
	Fingerprint int    `json:"fingerprint"`
	FolderName  string `json:"folderName"`
}

//File a single file
type File struct {
	ID                 int               `json:"id"`
	FileName           string            `json:"fileName"`
	FileNameOnDisk     string            `json:"fileNameOnDisk"`
	FileDate           string            `json:"fileDate"`
	ReleaseType        int               `json:"releaseType"`
	FileStatus         int               `json:"fileStatus"`
	DownloadURL        string            `json:"downloadUrl"`
	IsAlternate        bool              `json:"isAlternate"`
	AlternateFileID    int               `json:"alternateFileId"`
	Dependencies       []AddonDependency `json:"dependencies"`
	IsAvailable        bool              `json:"isAvailable"`
	Modules            []AddonModules    `json:"modules"`
	PackageFingerprint int64             `json:"packageFingerprint"`
	GameVersion        []string          `json:"gameVersion"`
}

//FingerprintList List of fingerprint matches
type FingerprintList struct {
	IsCacheBuilt          bool               `json:"isCacheBuilt"`
	ExactMatches          []FingerprintMatch `json:"exactMatches"`
	ExactFingerprints     []int              `json:"exactFingerprints"`
	InstalledFingerprints []int              `json:"installedFingerprints"`
}

//FingerprintMatch match of addon fingerprint
type FingerprintMatch struct {
	AddonID int  `json:"id"`
	File    File `json:"file"`
}

// Auth

//Register user/pass object
type Register struct {
	Username   string `json:"Username"`
	Password   string `json:"Password"`
	Email      string `json:"Email"`
	Newsletter bool   `json:"Newsletter"`
}

//AuthenticationResponse Response of auth
type AuthenticationResponse struct {
	Status        int         `json:"Status"`
	StatusMessage string      `json:"StatusMessage"`
	Session       AuthSession `json:"Session"`
}

//AuthSession Session of auth
type AuthSession struct {
	UserID                 int    `json:"UserID"`
	Username               string `json:"Username"`
	DisplayName            string `json:"DisplayName"`
	SessionID              string `json:"SessionID"`
	Token                  string `json:"Token"`
	EmailAddress           string `json:"EmailAddress"`
	EffectivePremiumStatus bool   `json:"EffectivePremiumStatus"`
	ActualPremiumStatus    bool   `json:"ActualPremiumStatus"`
	SubscriptionToken      int    `json:"SubscriptionToken"`
	Expires                int    `json:"Expires"`
	RenewAfter             int    `json:"RenewAfter"`
	IsTemporaryAccount     bool   `json:"IsTemporaryAccount"`
	IsMerged               bool   `json:"IsMerged"`
	Bans                   int    `json:"Bans"`
}

//Login user/pass
type Login struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//TwitchOAuthRequest https://logins-v1.curseapp.net/Help/Api/POST-login-twitch-oauth
type TwitchOAuthRequest struct {
	ClientID    string `json:"ClientID"`
	Code        string `json:"Code"`
	RedirectURI string `json:"RedirectUri"`
	State       string `json:"State"`
}

//TwitchOAuthResponse Response from above
type TwitchOAuthResponse struct {
	Status            int         `json:"Status"`
	Session           AuthSession `json:"Session"`
	TwitchUsername    string      `json:"TwitchUsername"`
	TwitchDisplayName string      `json:"TwitchDisplayName"`
	TwitchAvatar      string      `json:"TwitchAvatar"`
	TwitchUserID      string      `json:"TwitchUserID"`
}

//RenewTokenResponseContract OAuth stuffs
type RenewTokenResponseContract struct {
	Token      string `json:"Token"`
	Expires    int    `json:"Expires"`
	RenewAfter int    `json:"RenewAfter"`
}
