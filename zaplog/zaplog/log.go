package zaplog

import (
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const logTmFmtWithMS = "2006-01-02 15:04:05.000"

var L = NewLogger()

func NewLogger() *zap.SugaredLogger {
	coreList := make([]zapcore.Core, 0)

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format(logTmFmtWithMS) + "]")
	}
	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}

	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + uuid.NewString() + "]")
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	customNameEncoder := func(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + strings.ToUpper(loggerName) + "]")
	}

	// [2023-06-29 15:35:12.313]       [DEBUG] [TEST]  [af378fc1-5d11-4bc4-a454-6e7352969bc8]  [cmd/main.go:15]        this is debug
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = customLevelEncoder
	encoderConfig.EncodeCaller = customCallerEncoder
	encoderConfig.EncodeName = customNameEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	//encoderConfig := zapcore.EncoderConfig{
	//	CallerKey:      "caller_line", // 打印文件名和行数
	//	LevelKey:       "level_name",
	//	MessageKey:     "msg",
	//	TimeKey:        "ts",
	//	NameKey:        "name",
	//	StacktraceKey:  "stacktrace",
	//	LineEnding:     zapcore.DefaultLineEnding,
	//	EncodeTime:     customTimeEncoder,   // 自定义时间格式
	//	EncodeLevel:    customLevelEncoder,  // 小写编码器
	//	EncodeCaller:   customCallerEncoder, // 全路径编码器
	//	EncodeDuration: zapcore.SecondsDurationEncoder,
	//	//EncodeName:     zapcore.FullNameEncoder,
	//	EncodeName: customNameEncoder, // 自定义
	//}

	//// level大写染色编码器
	//if l.enableColor {
	//	encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//}
	//
	//// json 格式化处理
	//if l.jsonFormat {
	//	return zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf),
	//		syncWriter, zap.NewAtomicLevelAt(l.logMinLevel))
	//}

	//encoderConfig := zap.NewProductionEncoderConfig()     //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //指定时间格式
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
