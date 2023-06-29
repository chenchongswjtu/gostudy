package main

//
//import (
//	"net/http"
//	"os"
//
//	"github.com/natefinch/lumberjack"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//)
//
//var sugarLogger *zap.SugaredLogger
//
//func main() {
//	InitLogger()
//	defer sugarLogger.Sync()
//
//	// 2023-06-28T20:09:30.789+0800    INFO    test    zaplog/main.go:18       ddd     {"channel": "mychannel"}
//	sugarLogger.With("channel", "mychannel").Named("test").Infof("ddd")
//	simpleHttpGet("http://www.baidu.com")
//	simpleHttpGet("http://www.topgoer.com")
//
//}
//
//func InitLogger() {
//	writeSyncer := getLogWriter()
//	encoder := getEncoder()
//	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
//
//	logger := zap.New(core, zap.AddCaller())
//	sugarLogger = logger.Sugar()
//}
//
//func getEncoder() zapcore.Encoder {
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	return zapcore.NewConsoleEncoder(encoderConfig)
//}
//
//func getLogWriter() zapcore.WriteSyncer {
//	lumberJackLogger := &lumberjack.Logger{
//		Filename:   "./test.log",
//		MaxSize:    1,
//		MaxBackups: 5,
//		MaxAge:     30,
//		Compress:   false,
//	}
//	return zapcore.AddSync(lumberJackLogger)
//}
//
//func simpleHttpGet(url string) {
//	sugarLogger.Debugf("Trying to hit GET request for %s", url)
//	resp, err := http.Get(url)
//	if err != nil {
//		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
//	} else {
//		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
//		resp.Body.Close()
//	}
//}

//package main
//
//import (
//	"os"
//
//	"github.com/natefinch/lumberjack"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//)
//
//var log *zap.Logger
//
//func main() {
//	var coreArr []zapcore.Core
//
//	//获取编码器
//	encoderConfig := zap.NewProductionEncoderConfig()            //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //指定时间格式
//	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
//	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径
//	encoder := zapcore.NewConsoleEncoder(encoderConfig)
//
//	//日志级别
//	//highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
//	//	return lev >= zap.ErrorLevel
//	//})
//	//lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
//	//	return lev < zap.ErrorLevel && lev >= zap.DebugLevel
//	//})
//
//	//info文件writeSyncer
//	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
//		Filename:   "./log/info.log", //日志文件存放目录，如果文件夹不存在会自动创建
//		MaxSize:    2,                //文件大小限制,单位MB
//		MaxBackups: 100,              //最大保留日志文件数量
//		MaxAge:     30,               //日志文件保留天数
//		Compress:   false,            //是否压缩处理
//	})
//	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.InfoLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
//	//error文件writeSyncer
//	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
//		Filename:   "./log/error.log", //日志文件存放目录
//		MaxSize:    1,                 //文件大小限制,单位MB
//		MaxBackups: 5,                 //最大保留日志文件数量
//		MaxAge:     30,                //日志文件保留天数
//		Compress:   false,             //是否压缩处理
//	})
//	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.ErrorLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
//
//	coreArr = append(coreArr, infoFileCore)
//	coreArr = append(coreArr, errorFileCore)
//	log = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) //zap.AddCaller()为显示文件名和行号，可省略
//
//	log.Info("hello info")
//	log.Debug("hello debug")
//	log.Error("hello error")
//}
