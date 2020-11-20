package domain

import "context"

// CommandHandler receives commands and does the desired operation
// @Port
type CommandHandler interface {
	Invoke(ctx context.Context, command interface{}) error
}
