package chat

import (
	"errors"
	"sendtg/internal/config"
)

func ResolveChatID(name string) (int64, error) {
	conf := config.GetConfig()
	id, exists := conf.Chats[name]
	if !exists {
		return 0, errors.New("chat not found")
	}
	return id, nil
}
