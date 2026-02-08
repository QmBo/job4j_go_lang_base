package trackerstore

import (
	"fmt"
)

func ErrRPoolExec(err error) error {
	return fmt.Errorf("r.pool.Exec: %w", err)
}
func ErrRPoolQuery(err error) error {
	return fmt.Errorf("r.pool.Query: %w", err)
}
func ErrRPoolQueryRow(err error) error {
	return fmt.Errorf("r.pool.QueryRow: %w", err)
}
func ErrRRowScan(err error) error {
	return fmt.Errorf("r.row.Scan: %w", err)
}
func ErrRRowError(err error) error {
	return fmt.Errorf("r.row.Err: %w", err)
}
func ErrCreate(err error) error {
	return fmt.Errorf("failed to create items: %w", err)
}
func ErrGet(err error) error {
	return fmt.Errorf("failed to get items: %w", err)
}
func ErrUpdate(err error) error {
	return fmt.Errorf("failed to update items: %w", err)
}
func ErrDelete(err error) error {
	return fmt.Errorf("failed to delete items: %w", err)
}
func ErrReorderAfterDelete(err error) error {
	return fmt.Errorf("failed to reorder after delete items: %w", err)
}
func ErrFind(err error) error {
	return fmt.Errorf("failed to find items: %w", err)
}
func ErrGetLastPosition(err error) error {
	return fmt.Errorf("failed to find last position : %w", err)
}
func ErrUi(err error) error {
	return fmt.Errorf("ui problem: %w", err)
}
