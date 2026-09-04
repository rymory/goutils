// Copyright (c) 2017-2026 Onur Yaşar
// Licensed under AGPL v3 + Commercial Exception
// See LICENSE.txt

// https://github.com/rymory/goutils/blob/main/LICENSE
// https://github.com/rymory/goutils/blob/main/README.md
// rymory.org
// onuryasar.org
// onxorg@proton.me 

package api

import (
	"math/rand"
	"time"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func GeneratePassword(length int, useLetters bool, useSpecial bool, useNum bool) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		if useLetters {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else if useSpecial {
			b[i] = specialBytes[rand.Intn(len(specialBytes))]
		} else if useNum {
			b[i] = numBytes[rand.Intn(len(numBytes))]
		}
	}
	return string(b)
}
