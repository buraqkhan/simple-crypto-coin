package blockchain

type transaction struct {
	Sender   string
	Reciever string
	Amount   string
}

func (t *transaction) CreateTransaction(sender string, reciever string, amount string) *transaction {
	return &transaction{
		Sender:   sender,
		Reciever: reciever,
		Amount:   amount,
	}
}
