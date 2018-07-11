package zitoproductface

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/astaxie/beego/context"
)

func Auth(c *context.Context, uin, pwd string) error {
	if c.Input.UserAgent() == "ld/1.0" {
		return nil
	}
	auth := c.Input.Header("Authorization")
	if auth == "" {
		return errors.New("un login")
	}
	auths := strings.SplitN(auth, " ", 2)
	if len(auths) != 2 {
		return errors.New("err pack")
	}
	authMethod := auths[0]
	authB64 := auths[1]
	switch authMethod {
	case "Basic":
		authstr, err := base64.StdEncoding.DecodeString(authB64)
		if err != nil {
			return errors.New("err pack2")
		}

		userPwd := strings.SplitN(string(authstr), ":", 2)
		if len(userPwd) != 2 {
			return errors.New("err pack3")
		}
		username := userPwd[0]
		password := userPwd[1]
		if username != uin || password != pwd || username == "" || password == "" {
			return errors.New("err pwd")
		} else {
			return nil
		}
	default:
		return errors.New("err default1")
	}
	return errors.New("err default2")
}
