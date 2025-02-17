package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	nodeAPIBaseURL = "http://185.239.142.164:8083" // вроде бы так будет
)

// отправляет запрос на добавление пользователя в конфиг
func AddUser(userID string) (string, error) {
	url := fmt.Sprintf("%s/api/user", nodeAPIBaseURL)
	payload := map[string]string{"id": userID}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result["id"], nil
}

// отправляет запрос на удаление пользователя из конфига
func RemoveUser(userID string) error {
	url := fmt.Sprintf("%s/api/user/%s", nodeAPIBaseURL, userID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// отправляет запрос на генерацию VLESS-ссылки для пользователя
func GenerateVLESSLink(userID string) (string, error) {
	url := fmt.Sprintf("%s/api/user/%s/link", nodeAPIBaseURL, userID)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result["link"], nil
}

// отправляет запрос на перезапуск Xray
func RestartXray() error {
	url := fmt.Sprintf("%s/api/management/restart", nodeAPIBaseURL)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// отправляет запрос на запуск Xray
func StartXray() error {
    url := fmt.Sprintf("%s/api/management/start", nodeAPIBaseURL)
    resp, err := http.Post(url, "application/json", nil)
    if err != nil {
        return fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }
    return nil
}

// отправляет запрос на остановку Xray
func StopXray() error {
    url := fmt.Sprintf("%s/api/management/stop", nodeAPIBaseURL)
    resp, err := http.Post(url, "application/json", nil)
    if err != nil {
        return fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }
    return nil
}

// отправляет запрос на получение статуса Xray
func GetXrayStatus() (string, error) {
	url := fmt.Sprintf("%s/api/management/status", nodeAPIBaseURL)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result["status"], nil
}