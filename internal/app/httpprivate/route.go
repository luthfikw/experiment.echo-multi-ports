package httpprivate

func (ox *httpPrivate) InitRoute() {
	ox.Instance.GET("/", ox.Handler.Welcome)
	ox.Instance.GET("/v1/:path", ox.Handler.WelcomeA)
	ox.Instance.GET("/v1/foo", ox.Handler.WelcomeB)
}
