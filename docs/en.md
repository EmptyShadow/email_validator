# Email Validator

The module allows you to check the validity of e-mail in three stages:
1. Validation of the recording format.
2. Validation of the domain for the presence of mx records in the dns server.
3. A test request to the mail server at the address specified in the mx entry in order to find out if the server supports
smtp and whether the user exists.

```go
func main () {
    email: = "example@example.com"

    validator, err: = NewEmailValidator (nil)
    if err! = nil {
        return false, err
    }

    if err: = validator.CheckEmail (email); err! = nil {
        fmt.println (err)
    } else {
        fmt.println ("email valid")
    }
}
```

When performing a test request, it is possible to use a proxy server.
```go
func main () {
    email: = "example@example.com"

    proxyURL: = "username: password @ proxyhost: proxyport"
    proxy: = NewProxyByStringURL (proxyURL)

    validator, err: = NewEmailValidator (proxy)
    if err! = nil {
        return false, err
    }

    if err: = validator.CheckEmail (email); err! = nil {
        fmt.println (err)
    } else {
        fmt.println ("email valid")
    }
}
```

You can perform the steps separately.
```go
func main () {
    email: = "example@example.com"

    validator, err: = NewEmailValidator (nil)
    if err! = nil {
        return false, err
    }

    if err: = validator.CheckFormat (email); err! = nil {
        fmt.println (err)
    }

    mx, err: = validator.CheckDomain (email)
    if err! = nil {
        fmt.println (err)
    }

    if err: = validator.CheckRequest (email, mx [0]); err! = nil {
        fmt.println (err)
    }
}

```