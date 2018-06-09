package common

type error interface {
    Error() string
}

// type errors error
// func (e errors) Error(msg string) string {
//     return msg
// }
type ErrorCode struct {
    Success int
    Error   int
}