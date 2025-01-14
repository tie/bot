package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func tokHash(token string) string {
	h := md5.Sum([]byte(token + "gotd-token-salt")) // #nosec
	return fmt.Sprintf("%x", h[:5])
}

func sessionFileName(token string) string {
	return fmt.Sprintf("bot.%s.session.json", tokHash(token))
}

func ignoreClose(closer io.Closer) {
	_ = closer.Close()
}
