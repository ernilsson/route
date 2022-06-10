package packet

type Router struct {
	routes []Route
}

func (r *Router) Consume(ctx Context) error {
	for _, route := range r.routes {
		if !route.accept(ctx.Payload) {
			continue
		}
		if err := route.handle(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (r *Router) ForceConsume(ctx Context) {
	for _, route := range r.routes {
		if route.accept(ctx.Payload) {
			_ = route.handle(ctx)
		}
	}
}

func NewRouter(routes ...Route) Router {
	return Router{routes: routes}
}
