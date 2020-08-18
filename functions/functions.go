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
