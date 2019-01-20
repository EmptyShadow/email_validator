# Валидатор электронных почт

Модуль позволяет проверять валидность электронной почты тремя этапами:
1. Валидация формата записи.
2. Валидация домена, на наличие mx записей у dns сервера.
3. Пробный запрос к почтовому серверу по адрессу указанному в mx записи для того, чтобы узнать поддерживает ли сервер 
smtp и существует ли пользователь.

```go
func main() {
	email := "example@example.com"
	
	validator, err := NewEmailValidator(nil)
    if err != nil {
    	return false, err
    }
	
	if err := validator.CheckEmail(email); err != nil {
		fmt.println(err)
	} else {
		fmt.println("email valid")
	}
}

```

При выполнении тестового запроса есть возможность использовать прокси сервер.
```go
func main() {
	email := "example@example.com"
	
	proxyURL := "username:password@proxyhost:proxyport"
	proxy := NewProxyByStringURL(proxyURL)
	
	validator, err := NewEmailValidator(proxy)
    if err != nil {
    	return false, err
    }
	
	if err := validator.CheckEmail(email); err != nil {
		fmt.println(err)
	} else {
		fmt.println("email valid")
	}
}

```

Выполнять этапы можно по отдельности.
```go
func main() {
	email := "example@example.com"
	
	validator, err := NewEmailValidator(nil)
    if err != nil {
    	return false, err
    }
	
	if err := validator.CheckFormat(email); err != nil {
		fmt.println(err)
	}
	
	mx, err := validator.CheckDomain(email)
	if err != nil {
		fmt.println(err)
	}
	
	if err := validator.CheckRequest(email, mx[0]); err != nil {
    	fmt.println(err)
	}
}

```