package httppublic

func (ox *httpPublic) InitRoute() {
	ox.Instance.GET("/", ox.Handler.Welcome)
}
