package account

type NewAggregateParams struct {
	Email        string
	PasswordHash string
	FullName     string
	Gender       Gender
}
