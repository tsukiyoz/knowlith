package cmd

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultHomeDir    = ".knowlith"
	defaultConfigName = "config.yaml"
)

func onInitialize() {
	initLog()
	initConfig()
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		for _, dir := range searchDirs() {
			viper.AddConfigPath(dir)
		}

		viper.SetConfigType("yaml")

		viper.SetConfigName(defaultConfigName)
	}

	setupEnvironmentVariables()
	_ = viper.ReadInConfig()
	slog.Info("Config initialized success.")
}

func setupEnvironmentVariables() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("KNOWLITH")
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

func searchDirs() []string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultHomeDir), "."}
}

func filePath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, defaultHomeDir, defaultConfigName)
}

func initLog() {
	// 获取日志配置
	format := viper.GetString("log.format") // 日志格式，支持：json、text
	level := viper.GetString("log.level")   // 日志级别，支持：debug, info, warn, error
	output := viper.GetString("log.output") // 日志输出路径，支持：标准输出stdout和文件

	// 转换日志级别
	var slevel slog.Level
	switch level {
	case "debug":
		slevel = slog.LevelDebug
	case "info":
		slevel = slog.LevelInfo
	case "warn":
		slevel = slog.LevelWarn
	case "error":
		slevel = slog.LevelError
	default:
		slevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: slevel}

	var w io.Writer
	var err error
	// 转换日志输出路径
	switch output {
	case "":
		w = os.Stdout
	case "stdout":
		w = os.Stdout
	default:
		w, err = os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
		if err != nil {
			panic(err)
		}
	}

	// 转换日志格式
	if err != nil {
		return
	}
	var handler slog.Handler
	switch format {
	case "json":
		handler = slog.NewJSONHandler(w, opts)
	case "text":
		handler = slog.NewTextHandler(w, opts)
	default:
		handler = slog.NewJSONHandler(w, opts)

	}

	// 设置全局的日志实例为自定义的日志实例
	slog.SetDefault(slog.New(handler))
	slog.Info("Log initialized success.")
}
