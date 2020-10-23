package debug

import clog "github.com/fdurand/coredns/plugin/pkg/log"

func init() { clog.Discard() }
