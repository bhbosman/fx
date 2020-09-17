package fx_test

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func TestArrayAnnotated(t *testing.T) {
	type a struct {
		name string
	}
	newA := func() *a {
		return &a{name: "foo"}
	}
	newB := func() *a {
		return &a{name: "bar"}
	}
	t.Run("Provided None", func(t *testing.T) {
		app := fxtest.New(t,
			fx.Provide(
				[]fx.Annotated{},
			),
		)
		if !assert.NoError(t, app.Err()) {
			return
		}
		defer app.RequireStart().RequireStop()
	})
	t.Run("Provided 1", func(t *testing.T) {
		var in struct {
			fx.In
			A *a `name:"foo"`
		}
		app := fxtest.New(t,
			fx.Provide(
				[]fx.Annotated{
					{
						Name:   "foo",
						Target: newA,
					},
				},
			),
			fx.Populate(&in),
		)
		if !assert.NoError(t, app.Err()) {
			return
		}
		defer app.RequireStart().RequireStop()
		assert.NotNil(t, in.A, "expected in.A to be injected")
		assert.Equal(t, "foo", in.A.name, "expected to get a type 'a' of name 'foo'")
	})
	t.Run("Provided 2", func(t *testing.T) {
		var in struct {
			fx.In
			A *a `name:"foo"`
			B *a `name:"bar"`
		}
		app := fxtest.New(t,
			fx.Provide(
				[]fx.Annotated{
					{
						Name:   "foo",
						Target: newA,
					},
					{
						Name:   "bar",
						Target: newB,
					},
				},
			),
			fx.Populate(&in),
		)
		if !assert.NoError(t, app.Err()) {
			return
		}
		defer app.RequireStart().RequireStop()
		assert.NotNil(t, in.A, "expected in.A to be injected")
		assert.Equal(t, "foo", in.A.name, "expected to get a type 'a' of name 'foo'")

		assert.NotNil(t, in.B, "expected in.B to be injected")
		assert.Equal(t, "bar", in.B.name, "expected to get a type 'b' of name 'bar'")
	})
}
