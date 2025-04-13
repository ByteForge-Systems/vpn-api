package node_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ByteForge-Systems/vpn-api/internal/config"
)

var nodeAPIBaseURL string

func init() {
	config.LoadEnv()
	nodeAPIBaseURL = config.GetEnv("NODE_API_BASE_URL")

}

type Client struct {
	ID   string `json:"id"`
	Flow string `json:"flow"`
}

// AddUser отправляет запрос на добавление пользователя в конфиг
func AddUser(newUUID string) (string, error) {
	url := fmt.Sprintf("%s/api/key", nodeAPIBaseURL)

	requestBody := map[string]string{
		"uuid": newUUID,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result["id"], nil
}

// RemoveUser отправляет запрос на удаление пользователя из конфига
func RemoveUser(userID string) error {
	url := fmt.Sprintf("%s/api/key/%s", nodeAPIBaseURL, userID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// GenerateVLESSLink отправляет запрос на генерацию VLESS-ссылки для пользователя
func GenerateVLESSLink(userID string) (string, error) {
	url := fmt.Sprintf("%s/api/key/%s/link", nodeAPIBaseURL, userID)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result["link"], nil
}

func GetAllUsers() ([]Client, error) {
	url := fmt.Sprintf("%s/api/key/all", nodeAPIBaseURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Декодируем в структуру с полем "users"
	var response struct {
		Users []Client `json:"users"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return response.Users, nil
}

// RestartXray отправляет запрос на перезапуск Xray
func RestartXray() error {
	url := fmt.Sprintf("%s/api/management/restart", nodeAPIBaseURL)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// StartXray отправляет запрос на запуск Xray
func StartXray() error {
	url := fmt.Sprintf("%s/api/management/start", nodeAPIBaseURL)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// StopXray отправляет запрос на остановку Xray
func StopXray() error {
	url := fmt.Sprintf("%s/api/management/stop", nodeAPIBaseURL)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// GetXrayStatus отправляет запрос на получение статуса Xray
func GetXrayStatus() (string, error) {
	url := fmt.Sprintf("%s/api/management/status", nodeAPIBaseURL)
	fmt.Println("Requesting URL:", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result["status"], nil
}
