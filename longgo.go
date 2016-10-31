package echo

//import 	"github.com/go-long/gommon/log"

func (e *Echo) Add(method, path string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	 e.add(method, path, handler, middleware...)
}

//func (e *Echo) SetLogger(l *log.Logger) {
//	  e.logger=l
//}
