package service

import (
	"context"
	"errors"
	"os"
	"sync"

	localCtx "github.com/glstr/futty_golang/context"

	"google.golang.org/genai"
)

var (
	ErrReqAPIFailed = errors.New("request api failed")
)

const (
	geminiServiceName                = "gemini_service"
	geminiModelNameGemini25Flash     = "gemini-2.5-flash"
	geminiModelNameGemini3ProPreview = "gemini-3-pro-preview"
)

type ChatService interface {
	Chat(logBuffer *localCtx.LogBuffer, message string) (string, error)
}

var defaultGeminiService = new(GeminiService)
var chatServiceMap = map[string]ChatService{
	geminiServiceName: defaultGeminiService,
}

func GetChatService(service string) (ChatService, error) {
	if srv, find := chatServiceMap[service]; find {
		return srv, nil
	}
	return nil, ErrNotFoundService
}

type GeminiService struct {
	client *genai.Client
	once   sync.Once
}

func (s *GeminiService) init(logBuffer *localCtx.LogBuffer) {
	s.once.Do(func() {
		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey == "" {
			apiKey = "AIzaSyD9rPlJdLzpqysjTr_VG-w2BYGsczyY5ns"
		}

		config := &genai.ClientConfig{
			APIKey: apiKey,
		}
		c, err := genai.NewClient(context.Background(), config)
		if err != nil {
			logBuffer.WriteLog("init_client[failed] error_msg[%s]", err.Error())
			panic(err)
		}
		s.client = c
	})
}

func (s *GeminiService) Chat(logBuffer *localCtx.LogBuffer, message string) (string, error) {
	s.init(logBuffer)
	result, err := s.client.Models.GenerateContent(
		context.Background(),
		geminiModelNameGemini25Flash,
		genai.Text(message),
		nil,
	)
	if err != nil {
		logBuffer.WriteLog("request_api[failed] error_msg[%s]", err.Error())
		return "", ErrReqAPIFailed
	}

	return result.Text(), nil
}
