package card

import "github.com/SonnLarissa/bank/pkg/bank/types"

func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Active && card.Balance > 0 {
		if amount > 0 && amount < card.Balance && amount <= 20_000_00 {
			card.Balance = card.Balance - amount
		}
	}
	return card
}
