package mylog

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type DefaultTransactionIdHook struct {
}

func (hook *DefaultTransactionIdHook) Fire(entry *logrus.Entry) error {
	//entry.Data["appName"] = "MyAppName"
	entry.Data[TransactionIdKey] = entry.Context.Value(TransactionIdKey)
	fmt.Println("DEBUG DefaultTransactionIdHook.Fire")
	return nil
}

func (hook *DefaultTransactionIdHook) Levels() []logrus.Level {
	fmt.Println("DEBUG DefaultTransactionIdHook.Levels")
	return logrus.AllLevels
}
