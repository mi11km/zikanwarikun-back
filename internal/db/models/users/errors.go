package users

type UnauthenticatedUserAccessError struct{}

func (m *UnauthenticatedUserAccessError) Error() string {
	return "unauthenticated user access denied"
}

type TokenIsNotSetError struct{}

func (t *TokenIsNotSetError) Error() string {
	return "token is not set in Authorization header"
}

type AuthenticateUserCanNotDoThisActionError struct{}

func (t *AuthenticateUserCanNotDoThisActionError) Error() string {
	return "authenticate user can not do this action"
}
