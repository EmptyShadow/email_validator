package email_validator

import (
	"bitbucket.org/EmptyShadow/smtp_use_proxy"
	"errors"
	"net"
	"regexp"
	"strings"
)

// ErrorFormat - ошибка формата записи
var (
	ErrorFormat = errors.New("format is not valid")
	ErrorDomain = errors.New("domain is not valid")
)

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

// Проверка формата записи email
// email - адрес, который требуется проверить
//
// Check email entry format
// email - the address you want to check
func (v *EmailValidator) CheckFormat(email string) (error) {
	emailRegexp, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		return err
	}

	if !emailRegexp.MatchString(email) {
		return ErrorFormat
	}

	return nil
}

// Проверка валидности домена
// email - адрес, домен которого требуется проверить
//
// Check domain
// email - address whose domain want to check
func (v *EmailValidator) CheckDomain(email string) ([]*net.MX, error) {
	_, host := splitUserAndHost(email)

	mx, err := net.LookupMX(host)
	if err != nil {
		return nil, err
	}

	if mx == nil || len(mx) == 0 {
		return nil, ErrorDomain
	}

	return mx, nil
}


// Разбиение email адреса на аккаунт и хост
//
// Splitting email address into account and host
func splitUserAndHost(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i + 1:]
	return
}
