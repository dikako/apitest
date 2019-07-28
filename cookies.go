package apitest

import (
	"fmt"
	"net/http"
	"time"
)

// Cookie used to represent an http cookie
type Cookie struct {
	name     *string
	value    *string
	path     *string
	domain   *string
	expires  *time.Time
	maxAge   *int
	secure   *bool
	httpOnly *bool
}

// NewCookie creates a new Cookie with the provided name
func NewCookie(name string) *Cookie {
	return &Cookie{
		name: &name,
	}
}

// Value sets the value of the Cookie
func (cookie *Cookie) Value(value string) *Cookie {
	cookie.value = &value
	return cookie
}

// Path sets the path of the Cookie
func (cookie *Cookie) Path(path string) *Cookie {
	cookie.path = &path
	return cookie
}

// Domain sets the domain of the Cookie
func (cookie *Cookie) Domain(domain string) *Cookie {
	cookie.domain = &domain
	return cookie
}

// Expires sets the expires time of the Cookie
func (cookie *Cookie) Expires(expires time.Time) *Cookie {
	cookie.expires = &expires
	return cookie
}

// MaxAge sets the maxage of the Cookie
func (cookie *Cookie) MaxAge(maxAge int) *Cookie {
	cookie.maxAge = &maxAge
	return cookie
}

// Secure sets the secure bool of the Cookie
func (cookie *Cookie) Secure(secure bool) *Cookie {
	cookie.secure = &secure
	return cookie
}

// HttpOnly sets the httpOnly bool of the Cookie
func (cookie *Cookie) HttpOnly(httpOnly bool) *Cookie {
	cookie.httpOnly = &httpOnly
	return cookie
}

// ToHttpCookie transforms the Cookie to an http cookie
func (cookie *Cookie) ToHttpCookie() *http.Cookie {
	httpCookie := http.Cookie{}

	if cookie.name != nil {
		httpCookie.Name = *cookie.name
	}

	if cookie.value != nil {
		httpCookie.Value = *cookie.value
	}

	if cookie.path != nil {
		httpCookie.Path = *cookie.path
	}

	if cookie.domain != nil {
		httpCookie.Domain = *cookie.domain
	}

	if cookie.expires != nil {
		httpCookie.Expires = *cookie.expires
	}

	if cookie.maxAge != nil {
		httpCookie.MaxAge = *cookie.maxAge
	}

	if cookie.secure != nil {
		httpCookie.Secure = *cookie.secure
	}

	if cookie.httpOnly != nil {
		httpCookie.HttpOnly = *cookie.httpOnly
	}

	return &httpCookie
}

// Compares cookies based on only the provided fields from Cookie.
// Supported fields are Name, Value, Domain, Path, Expires, MaxAge, Secure and HttpOnly
func compareCookies(expectedCookie *Cookie, actualCookie *http.Cookie) (bool, []string) {
	cookieFound := *expectedCookie.name == actualCookie.Name
	compareErrors := make([]string, 0)

	if cookieFound {

		formatError := func(name string, expectedValue, actualValue interface{}) string {
			return fmt.Sprintf("Mismatched field %s. Expected %v but received %v",
				name,
				expectedValue,
				actualValue)
		}

		if expectedCookie.value != nil && *expectedCookie.value != actualCookie.Value {
			compareErrors = append(compareErrors, formatError("Value", *expectedCookie.value, actualCookie.Value))
		}

		if expectedCookie.domain != nil && *expectedCookie.domain != actualCookie.Domain {
			compareErrors = append(compareErrors, formatError("Domain", *expectedCookie.domain, actualCookie.Domain))
		}

		if expectedCookie.path != nil && *expectedCookie.path != actualCookie.Path {
			compareErrors = append(compareErrors, formatError("Path", *expectedCookie.path, actualCookie.Path))
		}

		if expectedCookie.expires != nil && !(*expectedCookie.expires).Equal(actualCookie.Expires) {
			compareErrors = append(compareErrors, formatError("Expires", *expectedCookie.expires, actualCookie.Expires))
		}

		if expectedCookie.maxAge != nil && *expectedCookie.maxAge != actualCookie.MaxAge {
			compareErrors = append(compareErrors, formatError("MaxAge", *expectedCookie.maxAge, actualCookie.MaxAge))
		}

		if expectedCookie.secure != nil && *expectedCookie.secure != actualCookie.Secure {
			compareErrors = append(compareErrors, formatError("Secure", *expectedCookie.secure, actualCookie.Secure))
		}

		if expectedCookie.httpOnly != nil && *expectedCookie.httpOnly != actualCookie.HttpOnly {
			compareErrors = append(compareErrors, formatError("HttpOnly", *expectedCookie.httpOnly, actualCookie.HttpOnly))
		}
	}

	return cookieFound, compareErrors
}
