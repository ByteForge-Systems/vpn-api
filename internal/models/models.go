package models

// SuccessResponse представляет успешный ответ API
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse представляет ошибку API
type ErrorResponse struct {
	Error string `json:"error"`
}

// User представляет данные пользователя
type User struct {
	ID   string `json:"id"`
	Flow string `json:"flow"`
}

type Node struct {
	ID          int    `json:"id"`
	IP          string `json:"ip"`
	Port        int    `json:"port"`
	PublicKey   string `json:"public_key,omitempty"`
	PrivateKey  string `json:"private_key,omitempty"`
	Country     string `json:"country"`
	Comment     string `json:"comment"`
	IsOnline    bool   `json:"is_online"`
	SSHPort     int    `json:"ssh_port"`
	SSHUsername string `json:"ssh_username"`
	SSHPassword string `json:"ssh_password,omitempty"`
}

// VLESSLink представляет VLESS-ссылку
type VLESSLink struct {
	Link string `json:"link"`
}

// XrayStatus представляет статус Xray
type XrayStatus struct {
	Status string `json:"status"`
}
