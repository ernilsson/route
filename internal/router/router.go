package router

type Router struct {
	routes []Route
}

func (r *Router) Consume(packet []byte) error {
	for _, route := range r.routes {
		if !route.accept(packet) {
			continue
		}
		if err := route.handle(packet); err != nil {
			return err
		}
	}
	return nil
}

func (r *Router) ForceConsume(packet []byte) {
	for _, route := range r.routes {
		if route.accept(packet) {
			_ = route.handle(packet)
		}
	}
}

func NewRouter(routes ...Route) Router {
	return Router{routes: routes}
}
