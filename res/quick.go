package res

import "net/http"

func UnAuth() *Res {
	return ErrorRes(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func PermissionDeny() *Res {
	return ErrorRes(http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

func ServerError() *Res {
	return ErrorRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}
