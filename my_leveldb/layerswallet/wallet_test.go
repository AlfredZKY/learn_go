package layerswallet

import (
	"myleveldb/myleveldb"
	"testing"
)

func TestWriteWalletDB(t *testing.T) {
	ky := WalletInit("walletDb")
	key := "hello"
	value := "world"
	ky.WriteData(key,value)
}

func TestUpdateWalletDb(t *testing.T) {
	ky := WalletInit("walletDb")
	key := "1234"
	value := "1234"
	ky.UpdateData(key,value)
}

func TestSearchWalletDB(t *testing.T) {
	ky := WalletInit("walletDb")
	key := "hello"
	ky.SearchData(key)
}

func TestDeleteWalletDb(t *testing.T) {
	ky := WalletInit("walletDb")
	key := "hello"
	ky.DeleteData(key)
	if 2 > 3{
		myleveldb.RecursiveDeleteDic(ky.DbPath)
	}
	myleveldb.RecursiveDeleteDic(ky.DbPath)
}

