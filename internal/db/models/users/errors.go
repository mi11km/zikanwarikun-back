package users

type UnauthenticatedUserAccessError struct{}

func (m *UnauthenticatedUserAccessError) Error() string {
	return "unauthenticated user access denied"
}
