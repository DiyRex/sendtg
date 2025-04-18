package bot

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"sendtg/internal/config"
	"strings"
	"path/filepath"
	"io"
)

func ResolveBotToken(name string) (string, error) {
	conf := config.GetConfig()
	if name == "" && len(conf.Bots) > 0 {
		for _, token := range conf.Bots {
			return token, nil // First bot
		}
	}
	token, exists := conf.Bots[name]
	if !exists {
		return "", errors.New("bot not found")
	}
	return token, nil
}

func SendFile(token string, chatID int64, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &strings.Builder{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("document", filepath.Base(filePath))
	if err != nil {
		return err
	}
	_, _ = io.Copy(part, file)
	writer.WriteField("chat_id", fmt.Sprint(chatID))
	writer.Close()

	req, err := http.NewRequest("POST",
		fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", token),
		strings.NewReader(body.String()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Telegram API error")
	}
	return nil
}
