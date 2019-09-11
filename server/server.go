/**
* Created by GoLand.
* User: link1st
* Date: 2019/9/11
* Time: 09:55
 */

package server

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func Info() {
	fmt.Println("server into")

	var (
		requestId string
	)
	UUID := uuid.NewV4()
	requestId = UUID.String()

	fmt.Println("uuid", requestId)

}
