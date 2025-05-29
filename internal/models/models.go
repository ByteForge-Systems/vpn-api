package models

import (
	"github.com/google/uuid"
	"time"
)

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
	ID         int64     `json:"id" db:"id"`
	UUID       uuid.UUID `json:"uuid" db:"uuid"`
	MacAddress string    `json:"mac_address" db:"mac_address"`
	NodeID     int       `json:"node_id" db:"node_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Status     string    `json:"status" db:"status"`
}

type UserRequest struct {
	MacAddress string `json:"mac_address" binding:"required"`
}

type Node struct {
	ID          int    `json:"id"`
	IP          string `json:"ip"`
	Port        int    `json:"port"`
	SSHPort     int    `json:"ssh_port"`
	Country     string `json:"country"`
	Comment     string `json:"comment"`
	IsOnline    bool   `json:"is_online" db:"is_online"`
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
