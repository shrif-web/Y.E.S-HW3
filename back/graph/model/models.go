package model

func (e DuplicateUsernameException) Error() string {
	return *e.Message
}
func (e NoUserFoundException) Error() string {
	return *e.Message
}
func (e UserPassMissMatchException) Error() string {
	return *e.Message
}
func (e InternalServerException) Error() string {
	return *e.Message
}
func (e UserNotAllowedException) Error() string {
	return *e.Message
}
func (e PostNotFoundException) Error() string {
	return *e.Message
}
func (e PostEmptyException) Error() string {
	return *e.Message
}