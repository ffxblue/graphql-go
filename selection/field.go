package selection

import (
	"context"

	"github.com/ffxblue/graphql-go/internal/exec"
)

// Field being resolved for a request.
type Field exec.Field

// FieldFromContext that is currently being resolved.
// May only be called from a resolver function, using the context it was provided.
func FieldFromContext(ctx context.Context) Field {
	return Field(exec.FieldFromContext(ctx))
}
