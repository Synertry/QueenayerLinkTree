/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package Logger

import (
	"bytes"
	"io"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/Synertry/QueenayerLinkTree/Build"
)

var (
	LogBuffer bytes.Buffer
	Buf, Out  *slog.Logger
)

func init() {
	// Create handler options that mimic the old logFlags behavior
	handlerOpts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true, // Equivalent to log.Lshortfile
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Customize time format to include microseconds
			if a.Key == "time" {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format("2006-01-02 15:04:05.000000"))
				}
			}
			return a
		},
	}

	Buf = slog.New(slog.NewTextHandler(&LogBuffer, handlerOpts))

	// MultiWriter for slog to stdout and LogBuffer
	multiOutput := io.MultiWriter(os.Stdout, &LogBuffer)
	Out = slog.New(slog.NewTextHandler(multiOutput, handlerOpts))
}

// LogAndExit logs to stderr, writes to error log, SyncPics.log here and exits
func LogAndExit(errDesc string, errLog error) {
	Out.Error(errLog.Error())

	err := os.WriteFile(Build.App+".log", LogBuffer.Bytes(), 0666)
	if err != nil {
		log.Panic("Error creating main log ", Build.App, ".log: ", err.Error())
	}
	os.Exit(-1)
}

// TimeTrack prints elapsed times
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	//Out.Printf("%s took %s\n", name, elapsed)
	Out.Info("TimeTrack", slog.String("name", name), slog.String("elapsed", elapsed.String()))
}
