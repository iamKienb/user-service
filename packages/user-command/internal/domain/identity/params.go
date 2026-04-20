package identity

type NewAggregateParams struct {
	Email        string
	PasswordHash string
	FullName     string
	Gender       string
}

// type UpdateProfileParams struct {
// 	UserID      string
// 	FullName    string
// 	AvatarURL   string
// 	DateOfBirth *time.Time
// }

type UpdateCredentialParams struct {
	UserID            string
	NewHashedPassword string
	NewVersion        int
}

// type UpdateLoginStatsParams struct {
// 	UserID      string
// 	Success     bool
// 	LockedUntil *time.Time
// }
