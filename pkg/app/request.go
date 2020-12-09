package app

import (
	"github.com/astaxie/beego/validation"
	"log"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Printf("err.key: %s,err.message:%s", err.Key, err.Message)
	}
}