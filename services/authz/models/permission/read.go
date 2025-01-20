package permission

import (
	"context"

	e "encore.app/internal/errs"
	"encore.dev/rlog"
)

// GetPermissionsByUserId returns a user's permissions by user id.
func (m *Model) GetPermissionsByUserId(ctx context.Context, id string) ([]Permission, error) {
	rows, err := m.DB.Query(ctx, `
		SELECT DISTINCT permission_.operation_id
			, permission_.resource_id
			, permission_.context_id
			, permission_.description
			, permission_.active
			, permission_.created_at
			, permission_.updated_at
			, permission_.version
		FROM permission_
		INNER JOIN role_permission_xref_
			ON permission_.operation_id = role_permission_xref_.operation_id
			AND permission_.resource_id = role_permission_xref_.resource_id
			AND permission_.context_id = role_permission_xref_.context_id
		INNER JOIN user_role_xref_
			ON role_permission_xref_.role_id = user_role_xref_.role_id
		WHERE user_role_xref_.user_id = $1
			AND permission_.active
	`, id)
	if err != nil {
		rlog.Error("failed to query database", "error", err)
		return nil, e.InternalErrorResponse
	}
	defer rows.Close()

	permissions := []Permission{}
	for rows.Next() {
		var p Permission
		err := rows.Scan(
			&p.OperationId,
			&p.ResourceId,
			&p.ContextId,
			&p.Description,
			&p.Active,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Version,
		)
		if err != nil {
			rlog.Error("failed to scan row", "error", err)
			return nil, e.InternalErrorResponse
		}
		permissions = append(permissions, p)
	}

	return permissions, nil
}
