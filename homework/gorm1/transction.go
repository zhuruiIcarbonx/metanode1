package gorm1

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/**
 题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

*/

type Account struct {
	ID      uint    `gorm:"primarykey"`
	Name    string  `gorm:"type:varchar(100);not null"`
	Balance float64 `gorm:"not null"`
}

type Transaction struct {
	ID            uint    `gorm:"primarykey"`
	FromAccountID uint    `gorm:"not null"`
	ToAccountID   uint    `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
}

func Transfer() {
	db := initDb()

	db.AutoMigrate(&Account{}, &Transaction{})

	accountA := &Account{
		ID:      1,
		Name:    "A",
		Balance: 500.0,
	}

	accountB := &Account{
		ID:      2,
		Name:    "B",
		Balance: 300.0,
	}

	db.Debug().Create(accountA)
	db.Debug().Create(accountB)

	// 进行转账操作
	doTransferAuto(db, accountA.ID, accountB.ID, 100.0)
}

// 手动事务
func doTransferManua(db *gorm.DB, fromAccountID, toAccountID uint, amount float64) {

	tx := db.Begin()

	var accountA Account

	if err := tx.Debug().First(&accountA, fromAccountID).Error; err != nil {
		fmt.Println("账户A不存在！")
		tx.Rollback()
		return
	}

	if accountA.Balance < amount {

		fmt.Println("账户A余额不足！")

	}

	var accountB Account
	if err := tx.Debug().First(&accountB, toAccountID).Error; err != nil {
		fmt.Println("账户B不存在！")
		tx.Rollback()
		return
	}

	//减A加B
	tx.Debug().Model(&accountA).Update("balance", accountA.Balance-amount)
	tx.Debug().Model(&accountB).Update("balance", accountB.Balance+amount)

	trans := Transaction{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}

	tx.Debug().Create(&trans)

	tx.Commit()
	fmt.Println("转账成功")

}

// 自动事务
func doTransferAuto(db *gorm.DB, fromAccountID, toAccountID uint, amount float64) {

	db.Transaction(func(tx *gorm.DB) error {

		var accountA Account

		if err := tx.Debug().First(&accountA, fromAccountID).Error; err != nil {
			return errors.New("账户A不存在！")
		}

		if accountA.Balance < amount {
			return errors.New("账户A余额不足！")
		}

		var accountB Account
		if err := tx.Debug().First(&accountB, toAccountID).Error; err != nil {
			return errors.New("账户B不存在！")
		}

		//减A加B
		tx.Debug().Model(&accountA).Update("balance", accountA.Balance-amount)
		tx.Debug().Model(&accountB).Update("balance", accountB.Balance+amount)

		trans := Transaction{
			FromAccountID: fromAccountID,
			ToAccountID:   toAccountID,
			Amount:        amount,
		}

		tx.Debug().Create(&trans)
		return nil

	})

	fmt.Println("转账成功")

}
