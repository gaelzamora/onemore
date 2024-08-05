package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"onemore/models"
	"strings"
	"time"
)

func ValidToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) < 2 {
		return false, fmt.Errorf("Invalid token"), ""
	}

	userInfo, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, fmt.Errorf("Failed to decode token: " + err.Error()), ""
	}

	var tkj models.TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		return false, fmt.Errorf("Failed to unmarshal JSON: "+err.Error()), ""
	}

	now := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(now) {
		return false, fmt.Errorf("Token has expired"), ""
	}

	return true, nil, tkj.Username
}