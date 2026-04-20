package user

type UserAggregate struct {
	events []DomainEvent // Lưu tạm để bắn đi, không lưu DB
}
