// Copyright (c) 2021-2022 Trim21 <trim21.me@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

//go:build dev

package logger

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// development log config.
func getLogger(level zapcore.Level) *zap.Logger {
	consoleEncoding := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        timeKey,
		NameKey:        nameKey,
		MessageKey:     messageKey,
		CallerKey:      callerKey,
		LevelKey:       levelKey,
		StacktraceKey:  traceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   getCallerEncoder(),
	})

	return zap.New(
		zapcore.NewCore(
			consoleEncoding, zapcore.AddSync(colorable.NewColorableStdout()), level,
		),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.WarnLevel),
		zap.Development(),
	)
}

func FxLogger() fx.Option {
	return fx.Provide() // as nil fx.Option
}
