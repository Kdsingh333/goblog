package dbclient

import "github.com/callistaenterprise/goblog/accountservice/model"

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountId string) (model.Account, error)
	Seed()
}
