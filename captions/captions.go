package captions

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"go.uber.org/zap"
)

func SetCaption(b ext.Bot, u *gotgbot.Update) error {
	var replyMessage ext.Message
	var fileId string

	chatId := u.EffectiveChat.Id
	message := u.EffectiveMessage
	captionText := message.Text

	if message.ReplyToMessage != nil {
		replyMessage = *message.ReplyToMessage
	}

	if replyMessage.Audio != nil {
		fileId = replyMessage.Audio.FileId
		msg := b.NewSendableAudio(chatId, captionText)
		msg.Audio = b.NewFileId(fileId)
		_, err := msg.Send()
		if err != nil {
			b.Logger.Errorf("Error in sending", zap.Error(err))
		}
		return err
	} else if replyMessage.Document != nil {
		fileId = replyMessage.Document.FileId
		msg := b.NewSendableDocument(chatId, captionText)
		msg.Document = b.NewFileId(fileId)
		_, err := msg.Send()
		if err != nil {
			b.Logger.Errorf("Error in sending", zap.Error(err))
		}
		return err
	} else if replyMessage.Photo != nil {
		fileId = replyMessage.Photo[len(replyMessage.Photo)-1].FileId
		msg := b.NewSendablePhoto(chatId, captionText)
		msg.Photo = b.NewFileId(fileId)
		_, err := msg.Send()
		if err != nil {
			b.Logger.Errorf("Error in sending", zap.Error(err))
		}
		return err
	} else if replyMessage.Video != nil {
		fileId = replyMessage.Video.FileId
		msg := b.NewSendableVideo(chatId, captionText)
		msg.Video = b.NewFileId(fileId)
		_, err := msg.Send()
		if err != nil {
			b.Logger.Errorf("Error in sending", zap.Error(err))
		}
		return err
	} else if replyMessage.Voice != nil {
		fileId = replyMessage.Voice.FileId
		msg := b.NewSendableVoice(chatId, captionText)
		msg.Voice = b.NewFileId(fileId)
		_, err := msg.Send()
		if err != nil {
			b.Logger.Errorf("Error in sending", zap.Error(err))
		}
		return err
	} else {
		msg := b.NewSendableMessage(chatId, captionText)
		_, err := msg.Send()
		if err != nil {
			b.Logger.Errorf("Error in sending", zap.Error(err))
		}
		return err
	}
}
