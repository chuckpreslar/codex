package aux

import (
	"reflect"
)

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes/base"
)

var _intr = reflect.TypeOf((*interfaces.Acceptor)(nil)).Elem()

// CreateAcceptorFrom ...
func CreateAcceptorFrom(item interface{}) interfaces.Acceptor {
	if nil == item {
		return base.InitializeNil()
	}

	var typ = reflect.TypeOf(item)

	if typ.Implements(_intr) {
		return item.(interfaces.Acceptor)
	}

	switch item.(type) {
	case int:
		item = base.InitializeInteger(item.(int))
	case string:
		item = base.InitializeString(item.(string))
	case bool:
		item = base.InitializeBoolean(item.(bool))
	}

	return item.(interfaces.Acceptor)
}
