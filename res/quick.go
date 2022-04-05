package res

import "net/http"

func UnAuth() *res {
	return NewRes().ErrorRes(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func PermissionDeny() *res {
	return NewRes().ErrorRes(http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

func ServerError() *res {
	return NewRes().ErrorRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}
