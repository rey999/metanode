package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Accounts struct {
	Id      int64 `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Balance int64 `gorm:"column:balance" json:"balance"`
}

type Transactions struct {
	Id           int64 `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Amount       int64 `gorm:"column:amount" json:"amount"`
	AccountsId   int64 `gorm:"column:accounts_id" json:"accounts_id"`
	ToAccountsId int64 `gorm:"column:to_accounts_id" json:"to_accounts_id"`
}

func transferAmount100(A *Accounts, B *Accounts, db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if A.Balance < 100 {
			return errors.New("余额不足")
		}
		A.Balance -= 100
		if err := tx.Save(A).Error; err != nil {
			return err
		}
		B.Balance += 100
		if err := tx.Save(B).Error; err != nil {
			return err

		}
		if err := tx.Create(&Transactions{Amount: 100, AccountsId: A.Id, ToAccountsId: B.Id}).Error; err != nil {
			return err
		}
		if err := tx.Create(&Transactions{Amount: -100, AccountsId: B.Id, ToAccountsId: A.Id}).Error; err != nil {
			return err
		}
		return nil
	})

}

func PrintlnAcounts(db *gorm.DB) {
	var accounts []Accounts
	db.Find(&accounts)
	for _, v := range accounts {
		fmt.Println(v)
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	db.AutoMigrate(&Accounts{}, &Transactions{})
	db.Where("1 = 1").Delete(&Accounts{})
	db.Where("1 = 1").Delete(&Transactions{})
	A := Accounts{Balance: 200}
	db.Create(&A)

	B := Accounts{Balance: 0}
	db.Create(&B)

	err1 := transferAmount100(&A, &B, db)
	if err1 != nil {
		fmt.Println("err1:", err1)
	}
	PrintlnAcounts(db)

	err2 := transferAmount100(&A, &B, db)
	if err2 != nil {
		fmt.Println("err2:", err2)
	}
	PrintlnAcounts(db)

	err3 := transferAmount100(&A, &B, db)
	if err3 != nil {
		fmt.Println("err3:", err3)
	}
	PrintlnAcounts(db)

}
