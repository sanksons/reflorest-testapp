package formatter

import (
	"github.com/sanksons/reflorest/src/common/logger/message"
)

// FormatInterface interface methods for formatterss
type FormatInterface interface {
	GetFormattedLog(msg *message.LogMsg) interface{}
}
