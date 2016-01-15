package routes

// these are functions that are actually greated by the intiated api.

type login_method func(username string, password string) (bool)
var chklogin login_method

type change_password func(username string, password string) (bool)
var chpassword change_password