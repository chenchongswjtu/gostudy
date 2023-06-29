package zaplog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var L = NewLogger()

func NewLogger() *zap.SugaredLogger {
	coreList := make([]zapcore.Core, 0)

	encoderConfig := zap.NewProductionEncoderConfig()     //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //指定时间格式
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder //显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 输出终端
	//encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))

	// 输出文件（INFO）
	infoJackLogger := &lumberjack.Logger{
		Filename:   "./info.log",
		MaxSize:    50, //MB
		MaxBackups: 30,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(infoJackLogger), zapcore.InfoLevel))

	// 输出文件（ERROR）
	errJackLogger := &lumberjack.Logger{
		Filename:   "./err.log",
		MaxSize:    50, // MB
		MaxBackups: 30, // 最大备份数
		MaxAge:     30, // days
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(errJackLogger), zapcore.ErrorLevel))

	// 将有output汇聚成coreList一次性初始化
	core := zapcore.NewTee(coreList...)

	// 设置是否打印堆栈
	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	zapLogger := zap.New(core, zap.AddCaller() /*zap.AddCallerSkip(1), */, zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

	return zapLogger
}

func MustLogger(name string) *zap.SugaredLogger {
	return NewLogger().Named(name)
}
