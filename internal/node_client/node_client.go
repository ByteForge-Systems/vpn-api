package node_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/ByteForge-Systems/vpn-api/internal/config"
)

type Client struct {
	ID   string `json:"id"`
	Flow string `json:"flow"`
}

// AddUser отправляет запрос на добавление пользователя в конфиг
func AddUser(nodeURL, newUUID string) (string, error) {
	url := fmt.Sprintf("%s/api/key", nodeURL)

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
func RemoveUser(nodeURL, userID string) error {
	url := fmt.Sprintf("%s/api/key/%s", nodeURL, userID)
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
func GenerateVLESSLink(nodeURL, userID string) (string, error) {
	url := fmt.Sprintf("%s/api/key/%s/link", nodeURL, userID)
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

func GetAllUsers(nodeURL string) ([]Client, error) {
	url := fmt.Sprintf("%s/api/key/all", nodeURL)
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
func RestartXray(nodeURL string) error {
	url := fmt.Sprintf("%s/api/management/restart", nodeURL)
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
func StartXray(nodeURL string) error {
	url := fmt.Sprintf("%s/api/management/start", nodeURL)
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
func StopXray(nodeURL string) error {
	url := fmt.Sprintf("%s/api/management/stop", nodeURL)
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
func GetXrayStatus(nodeURL string) (string, error) {
	url := fmt.Sprintf("%s/api/management/status", nodeURL)
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
