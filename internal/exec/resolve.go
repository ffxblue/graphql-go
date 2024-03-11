package exec

import "context"

type contextKey string

const field contextKey = "field"

// Field being resolved for a request.
type Field struct {
	Alias string
	Name  string
}

// FieldFromContext that is currently being resolved.
// May only be called from a resolver function, using the context it was provided.
func FieldFromContext(ctx context.Context) Field {
	return ctx.Value(field).(Field)
}

// WithField being resolved. Used internally for
func WithField(ctx context.Context, f Field) context.Context {
	return context.WithValue(ctx, field, f)
}
