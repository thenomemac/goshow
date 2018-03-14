package actions

func (as *ActionSuite) Test_ListingsResource_List() {
	res := as.HTML("/routes").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ListingsResource_Show() {
	// TODO: as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ListingsResource_New() {
	// TODO: as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ListingsResource_Create() {
	// TODO: as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ListingsResource_Edit() {
	// TODO: as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ListingsResource_Update() {
	// TODO: as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ListingsResource_Destroy() {
	// TODO: as.Fail("Not Implemented!")
}
