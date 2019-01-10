package email_validator

import "bitbucket.org/EmptyShadow/smtp_use_proxy"

// Валидатор Email адресов
// proxy - объект прокси сервера, через который будет происходить соединение при тестовой отправки, если nil,
// то соединение прямое
//
// Email address validator
// proxy - proxy server object through which the connection will occur during the test send, if nil,
// then the connection is direct
type EmailValidator struct {
	proxy *smtp_use_proxy.Proxy
}

// Создать валидатор
// proxy - объект прокси сервера, через который будет происходить соединение при тестовой отправки, если nil,
// то соединение прямое
//
// Create validator
// proxy - proxy server object through which the connection will occur during the test send, if nil,
//then the connection is direct
func NewEmailValidator(proxy *smtp_use_proxy.Proxy) (*EmailValidator, error) {
	return &EmailValidator{proxy: proxy}, nil
}

// Получить прокси
//
// Get proxy
func (v *EmailValidator) GetProxy() (*smtp_use_proxy.Proxy) {
	return v.proxy
}
