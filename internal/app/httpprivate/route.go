package httpprivate

func (ox *httpPrivate) InitRoute() {
	ox.Instance.GET("/", ox.Handler.Welcome)
}
