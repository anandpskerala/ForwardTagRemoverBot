/*  Copyright (C) 2020 by Anandpskerala@Github, < https://github.com/Anandpskerala >.
 *
 * This file is part of < https://github.com/Anandpskerala/ForwardTagRemoverBot > project,
 * and is released under the "GNU v3.0 License Agreement".
 * Please see < https://github.com/Anandpskerala/ForwardTagRemoverBot/blob/master/LICENSE >
 *
 * All rights reserved.
 */

package main

import (
	"fmt"
	"os"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Anandpskerala/ForwardTagRemoverBot/captions"
	"github.com/Anandpskerala/ForwardTagRemoverBot/commands"
	"github.com/Anandpskerala/ForwardTagRemoverBot/functions"
)

func main() {
	cfg := zap.NewProductionEncoderConfig()

	cfg.EncodeLevel = zapcore.CapitalLevelEncoder

	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zap.InfoLevel))

	defer logger.Sync()

	l := logger.Sugar()

	token := "" //os.Getenv("BOT_TOKEN")
	u, err := gotgbot.NewUpdater(logger, token)
	if err != nil {
		l.Fatalw("Updater failed starting", zap.Error(err))
		return
	}

	bot := u.Bot.FirstName
	fmt.Printf("Successfully logged as %s", bot)

	u.Dispatcher.AddHandler(handlers.NewCommand("start", commands.Start))
	u.Dispatcher.AddHandler(handlers.NewCommand("help", commands.Help))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Text, captions.SetCaption))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Document, functions.ForwardDocument))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Video, functions.ForwardVideo))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Photo, functions.ForwardPhoto))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Voice, functions.ForwardVoice))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Audio, functions.ForwardAudio))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Sticker, functions.ForwardSticker))

	err = u.StartPolling()
	if err != nil {
		l.Fatalw("Polling failed", zap.Error(err))
		return
	}
	u.Idle()

}
