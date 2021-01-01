package DBException

const UserNotFoundMessage = "there is no User @"
const PostNotFoundMessage = "there is no post @"
const UserNotAllowedMessage = "permission denied for User @"
const PostEmptyMessage = "cannot set empty post @"
const InternalMessage = "fatal error in database /n"

type UserNotFound struct {
	Message string
}

func (e *UserNotFound) Error() string {
	return e.Message
}

func ThrowUserNotFoundException(userID string) error {
	return &UserNotFound{Message: UserNotFoundMessage + userID}
}

type UserNotAllowed struct {
	Message string
}

func (e *UserNotAllowed) Error() string {
	return e.Message
}

func ThrowUserNotAllowedException(userID string) error {
	return &UserNotAllowed{Message: UserNotAllowedMessage + userID}
}

type PostNotFound struct {
	Message string
}

func (e *PostNotFound) Error() string {
	return e.Message
}

func ThrowPostNotFoundException(postID string) error {
	return &PostNotFound{Message: PostNotFoundMessage + postID}
}

type InternalDBError struct {
	Message string
}

func (e *InternalDBError) Error() string {
	return e.Message
}

func ThrowInternalDBException(mongoErr string) error {
	return &InternalDBError{Message: InternalMessage + mongoErr}
}
