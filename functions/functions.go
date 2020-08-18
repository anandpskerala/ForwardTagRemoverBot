/*  Copyright (C) 2020 by Anandpskerala@Github, < https://github.com/Anandpskerala >.
 *
 * This file is part of < https://github.com/Anandpskerala/ForwardTagRemoverBot > project,
 * and is released under the "GNU v3.0 License Agreement".
 * Please see < https://github.com/Anandpskerala/ForwardTagRemoverBot/blob/master/LICENSE >
 *
 * All rights reserved.
 */

package functions

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"go.uber.org/zap"
)

func ForwardText(b ext.Bot, u *gotgbot.Update) error {
	_, err := u.EffectiveMessage.ReplyText(u.EffectiveMessage.Text)
	if err != nil {
		b.Logger.Errorf("Error in sending message", zap.Error(err))
	}
	return err
}

func ForwardVideo(b ext.Bot, u *gotgbot.Update) error {
	_, err := u.EffectiveMessage.ReplyVideo(b.NewFileId(u.EffectiveMessage.Video.FileId))
	if err != nil {
		b.Logger.Errorf("Error in sending video", zap.Error(err))
	}
	return err
}

func ForwardDocument(b ext.Bot, u *gotgbot.Update) error {
	_, err := u.EffectiveMessage.ReplyDocument(b.NewFileId(u.EffectiveMessage.Document.FileId))
	if err != nil {
		b.Logger.Errorf("Error in sending document", zap.Error(err))
	}
	return err
}

func ForwardPhoto(b ext.Bot, u *gotgbot.Update) error {
	_, err := u.EffectiveMessage.ReplyPhoto(b.NewFileId(u.EffectiveMessage.Photo[len(u.EffectiveMessage.Photo)-1].FileId))
	if err != nil {
		b.Logger.Errorf("Error in sending photo", zap.Error(err))
	}
	return err
}

func ForwardSticker(b ext.Bot, u *gotgbot.Update) error {
	_, err := u.EffectiveMessage.ReplySticker(b.NewFileId(u.EffectiveMessage.Sticker.FileId))
	if err != nil {
		b.Logger.Errorf("Error in sending sticker", zap.Error(err))
	}
	return err
}

func ForwardAudio(b ext.Bot, u *gotgbot.Update) error {
	_, err := u.EffectiveMessage.ReplyAudio(b.NewFileId(u.EffectiveMessage.Audio.FileId))
	if err != nil {
		b.Logger.Errorf("Error in sending Audio", zap.Error(err))
	}
	return err
}

func ForwardVoice(b ext.Bot, u *gotgbot.Update) error {
	_, err := u.EffectiveMessage.ReplyVoice(b.NewFileId(u.EffectiveMessage.Voice.FileId))
	if err != nil {
		b.Logger.Errorf("Error in sending voice", zap.Error(err))
	}
	return err
}
