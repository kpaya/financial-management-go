package walletRepository

import (
	"database/sql"

	"github.com/kpaya/financial-management-go/src/domain"
)

type WalletRepository struct {
	Db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{Db: db}
}

func (r *WalletRepository) Insert(input *domain.Wallet) error {
	_, err := r.Db.Exec("INSERT INTO wallet (id, user_id, active) VALUES (?, ?, ?)", input.WalletId, input.UserId, input.Active)
	if err != nil {
		return err
	}
	return nil
}

func (r *WalletRepository) GetById(uuid string) (*domain.Wallet, error) {
	var wallet domain.Wallet
	err := r.Db.QueryRow("SELECT id, user_id, active FROM wallet WHERE id = ?", uuid).Scan(&wallet.WalletId, &wallet.UserId, &wallet.Active)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) GetByUserId(userId string) ([]*domain.Wallet, error) {
	rows, err := r.Db.Query("SELECT id, user_id, active FROM wallet WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	wallets := make([]*domain.Wallet, 0)

	for rows.Next() {
		wallet := new(domain.Wallet)
		err := rows.Scan(&wallet.WalletId, &wallet.UserId, &wallet.Active)

		if err != nil {
			continue
		}

		wallets = append(wallets, wallet)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return wallets, nil
}
