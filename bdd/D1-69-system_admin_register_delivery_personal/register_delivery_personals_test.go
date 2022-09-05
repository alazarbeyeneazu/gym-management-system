package bdd

import "github.com/cucumber/godog"

func iHaveUser(arg1 *godog.Table) error {
	return godog.ErrPending
}

func resultShouldBe(arg1 string) error {
	return godog.ErrPending
}

func systemAdminSendsAnwartuhaGmailcom(arg1, arg2 string, arg3 int, arg4 string, arg5 int, arg6 string, arg7 int) error {
	return godog.ErrPending
}

func systemAdminSendsTuhaAnwartuhaGmailcom(arg1, arg2, arg3 string, arg4 int, arg5 string, arg6 int) error {
	return godog.ErrPending
}

func systemAdminSendsTuhaAutomobile(arg1, arg2, arg3, arg4 string) error {
	return godog.ErrPending
}

func systemAdminSendsTuhaAutomobileParcel_delivery(arg1, arg2, arg3, arg4 string) error {
	return godog.ErrPending
}

func systemAdminSentsTuhaAutomobileParcel_delivery(arg1, arg2, arg3, arg4 string) error {
	return godog.ErrPending
}

func systemAmdinIsLogedIn() error {
	return godog.ErrPending
}

func theResultShouldBe(arg1 string) error {
	return godog.ErrPending
}

func userSentAnd(arg1, arg2 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^i have user$`, iHaveUser)
	ctx.Step(`^result should be "([^"]*)"$`, resultShouldBe)
	ctx.Step(`^system admin sends "([^"]*)" ""([^"]*)"\+(\d+)" "([^"]*)" "(\d+)" "([^"]*)" "anwartuha(\d+)@gmail\.com"$`, systemAdminSendsAnwartuhaGmailcom)
	ctx.Step(`^system admin sends "([^"]*)" "Tuha" "([^"]*)" ""([^"]*)"(\d+)" "([^"]*)" "anwartuha(\d+)@gmail\.com"$`, systemAdminSendsTuhaAnwartuhaGmailcom)
	ctx.Step(`^system admin sends "([^"]*)" "Tuha" "([^"]*)" "automobile" "([^"]*)" "" "([^"]*)"$`, systemAdminSendsTuhaAutomobile)
	ctx.Step(`^system admin sends "([^"]*)" "Tuha" "([^"]*)" "automobile" "([^"]*)" "parcel_delivery" "([^"]*)"$`, systemAdminSendsTuhaAutomobileParcel_delivery)
	ctx.Step(`^system admin sents "([^"]*)" "Tuha" "([^"]*)" "automobile" "([^"]*)" "parcel_delivery" "([^"]*)"$`, systemAdminSentsTuhaAutomobileParcel_delivery)
	ctx.Step(`^system amdin is loged in$`, systemAmdinIsLogedIn)
	ctx.Step(`^the result should be "([^"]*)"$`, theResultShouldBe)
	ctx.Step(`^user sent "([^"]*)" and "([^"]*)"$`, userSentAnd)
}
