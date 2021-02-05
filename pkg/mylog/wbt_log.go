package mylog

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Logger struct {
	Entity *logrus.Entry
	//transactionId string
}

const (
	TransactionIdKey = "transaction_id"
)

var logrusLogger *logrus.Logger

func init() {
	logrusLogger = getLogrusLogger()
	logrusLogger.AddHook(&DefaultTransactionIdHook{})
}

func NewLoggerWithTransactionId(ctx context.Context) *Logger {

	//如果ctx.Value没有TransactionIdKey，则生成一个新的TransactionIdKey
	if ctx.Value(TransactionIdKey) == nil {
		transId := fmt.Sprintf("%d", time.Now().UnixNano())
		ctx = context.WithValue(ctx, "transaction_id", transId)
	}

	entity := logrusLogger.WithContext(ctx)

	logger := &Logger{Entity: entity}

	return logger
}

func (l *Logger) Info(args ...interface{}) {
	l.Entity.Info(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Entity.Error(args...)
}

func getLogrusLogger() *logrus.Logger {
	logger := logrus.New()
	//logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logFile := "e:/logrus-wbt.log"
	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger.SetOutput(file)
	return logger
}
