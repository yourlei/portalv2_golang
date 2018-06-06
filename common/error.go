package common

type error interface {
    Error() string
}

// type errors error
// func (e errors) Error(msg string) string {
//     return msg
// }