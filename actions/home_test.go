package actions

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}

func (as *ActionSuite) Test_RoutesHandler() {
	res := as.HTML("/routes").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
