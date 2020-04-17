package layerswallet

import (
	"my_leveldb/myleveldb"
)

type WalletDbOperator interface {
	WalletInit(args string) myleveldb.KeyValue
	WriteWalletDB(key string,value string)error
	SearchWalletDB(key string)error
	UpdateWalletDb(key string,value string)error
	DeleteWalletDb(key string)error
}

func WalletInit(args string) myleveldb.KeyValue{
	ky :=myleveldb.KeyValue{DbPath: myleveldb.Init(args)}
	return ky
}

func WriteWalletDB(key string,value string)error{
	ky := WalletInit("walletDb")
	err := ky.WriteData(key,value)
	if err!= nil{
		return err
	}
	return nil
}

func SearchWalletDB(key string)error{
	ky := WalletInit("walletDb")
	err := ky.SearchData(key)
	if err != nil{
		return err
	}
	return nil
}

func UpdateWalletDb(key string,value string)error{
	ky := WalletInit("walletDb")
	err := ky.UpdateData(key,value)
	if err != nil{
		return err
	}
	return nil
}

func DeleteWalletDb(key string)error{
	ky := WalletInit("walletDb")
	err := ky.DeleteData(key)
	if err != nil{
		return err
	}
	return nil
}





