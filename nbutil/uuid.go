package nbutil

import (
	"github.com/satori/go.uuid"
)

func Uuid() string {
	// 创建UUID
	var err error
	var u1 = uuid.Must(uuid.NewV4(), err).String() //上文介绍的Version 4
	return u1
}
