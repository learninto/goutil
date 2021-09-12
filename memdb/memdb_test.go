package memdb

import (
	"context"
	"testing"

	"github.com/learninto/goutil/conf"
)

func TestMemDb(t *testing.T) {
	conf.Set("MEMDB_DSN_foo", "redis://localhost:6379/")

	ctx := context.Background()
	ctx, db := Get(ctx, "foo")

	s := db.Set(ctx, "a", "123", 0)
	if err := s.Err(); err != nil {
		t.Fatal(err)
	}

	sc := db.Get(ctx, "a")
	if v, err := sc.Result(); err != nil {
		t.Fatal(err)
	} else if v != "123" {
		t.Fatal("invalid string: " + v)
	}
}
