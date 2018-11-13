package decimal

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSQLScan(t *testing.T) {
	t.Run("Read Decimal from SQL", func(t *testing.T) {
		db, err := sql.Open("sqlite3", ":memory:")
		require.NoError(t, err)
		defer db.Close()
		rows, err := db.Query(`SELECT 1`)
		require.NoError(t, err)
		defer rows.Close()
		decimals := make([]Decimal, 0)
		for rows.Next() {
			var d Decimal
			require.NoError(t, rows.Scan(&d))
			decimals = append(decimals, d)
		}
		require.NoError(t, rows.Err())
		require.Len(t, decimals, 1)
		require.Equal(t, "1", decimals[0].String())
	})

	t.Run("Read Invalid from SQL", func(t *testing.T) {
		db, err := sql.Open("sqlite3", ":memory:")
		require.NoError(t, err)
		defer db.Close()
		rows, err := db.Query(`SELECT "AB"`)
		require.NoError(t, err)
		defer rows.Close()
		for rows.Next() {
			var d Decimal
			require.Error(t, rows.Scan(&d))
		}
	})

	t.Run("Write to SQL", func(t *testing.T) {
		db, err := sql.Open("sqlite3", ":memory:")
		require.NoError(t, err)
		defer db.Close()
		rows, err := db.Query(`SELECT 1 as Foo WHERE Foo = $1`, NewFromInt(1))
		require.NoError(t, err)
		defer rows.Close()
		for rows.Next() {
			var d Decimal
			require.NoError(t, rows.Scan(&d))
			require.Equal(t, "1", d.String())
		}
	})
}
