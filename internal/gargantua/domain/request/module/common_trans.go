package module

import (
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
)

func TransMapFromListOfString(ofString map[string]*protobuf.ListOfString) map[string][]string {
	result := make(map[string][]string, len(ofString))
	for k, v := range ofString {
		result[k] = v.Strings
	}
	return result
}

func TransMapToListOfString(ofString map[string][]string) map[string]*protobuf.ListOfString {
	result := make(map[string]*protobuf.ListOfString, len(ofString))
	for k, v := range ofString {
		result[k] = &protobuf.ListOfString{Strings: v}
	}
	return result
}

func TransCookieFromPb(cookie []*protobuf.Cookie) []*http.Cookie {
	result := make([]*http.Cookie, len(cookie))
	for i, c := range cookie {
		result[i] = &http.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    c.Expires.AsTime(),
			RawExpires: c.RawExpires,
			MaxAge:     int(c.MaxAge),
			Secure:     c.Secure,
			HttpOnly:   c.HttpOnly,
			SameSite:   http.SameSite(int(c.SameSite)),
			Raw:        c.Raw,
			Unparsed:   c.Unparsed,
		}
	}
	return result
}

func TransCookieToPb(cookie []*http.Cookie) []*protobuf.Cookie {
	result := make([]*protobuf.Cookie, len(cookie))
	for i, c := range cookie {
		result[i] = &protobuf.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    timestamppb.New(c.Expires),
			RawExpires: c.RawExpires,
			MaxAge:     int32(c.MaxAge),
			Secure:     c.Secure,
			HttpOnly:   c.HttpOnly,
			SameSite:   int32(c.SameSite),
			Raw:        c.Raw,
			Unparsed:   c.Unparsed,
		}
	}
	return result
}
