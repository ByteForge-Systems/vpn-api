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
    ID    string `json:"id"`
    Flow  string `json:"flow"`
}

// VLESSLink представляет VLESS-ссылку
type VLESSLink struct {
    Link string `json:"link"`
}

// XrayStatus представляет статус Xray
type XrayStatus struct {
    Status string `json:"status"`
}