package rollback

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRollbacker_Add(t *testing.T) {
	var (
		giveCtx         = context.TODO()
		giveDescription = "test"
		giveRbFunc      = func(ctx context.Context) error { return nil }
	)
	rbr := &Rollbacker{}
	rbr.Add(giveCtx, giveDescription, giveRbFunc)
	assert.Len(t, rbr.rollbacks, 1)
	assert.Exactly(t, giveDescription, rbr.rollbacks[0].description)
	assert.Exactly(
		t,
		reflect.ValueOf(giveRbFunc).Pointer(),
		reflect.ValueOf(rbr.rollbacks[0].funk).Pointer(),
	)
}

func TestRollbacker_Rollback_ShouldBeIntSlice321_WhenCallThatIsAppendedAddIntFunc(t *testing.T) {
	var (
		added = []int{}
		want  = []int{3, 2, 1}
	)

	rbr := &Rollbacker{
		rollbacks: []*rollback{
			{
				description: "add 1",
				funk: func(c context.Context) error {
					added = append(added, 1)
					return nil
				},
			},
			{
				description: "add 2",
				funk: func(c context.Context) error {
					added = append(added, 2)
					return nil
				},
			},
			{
				description: "add 3",
				funk: func(c context.Context) error {
					added = append(added, 3)
					return nil
				},
			},
		},
	}

	err := rbr.Rollback(context.TODO())
	assert.NoError(t, err)
	assert.Exactly(t, want, added)
}
func TestRollbacker_Rollback_ShouldReturnErrAndBeIntSlice31_WhenHappendErrAtSecondRollback(t *testing.T) {
	var (
		happendErr = errors.New("second err")
		added      = []int{}
		want       = []int{3, 1}
	)

	rbr := &Rollbacker{
		rollbacks: []*rollback{
			{
				description: "add 1",
				funk: func(c context.Context) error {
					added = append(added, 1)
					return nil
				},
			},
			{
				description: "add 2",
				funk: func(c context.Context) error {
					return happendErr
				},
			},
			{
				description: "add 3",
				funk: func(c context.Context) error {
					added = append(added, 3)
					return nil
				},
			},
		},
	}

	err := rbr.Rollback(context.TODO())
	assert.ErrorIs(t, err, happendErr)
	assert.Exactly(t, want, added)
}

func TestRollbacker_RollbackForPanic_ShouldCallRollback_WhenHappendPanic(t *testing.T) {
	var (
		addTarget = []int{}
		want      = []int{1}
	)
	rbr := &Rollbacker{
		rollbacks: []*rollback{
			{
				description: "add 1",
				funk: func(c context.Context) error {
					addTarget = append(addTarget, 1)
					return nil
				},
			},
		},
	}

	go func() {
		defer rbr.RollbackForPanic(context.TODO())
		panic("happend panic error!!!")
	}()

	time.Sleep(time.Second)

	assert.Exactly(t, addTarget, want)
}
func TestRollbacker_RollbackForPanic_ShouldNotCallRollback_WhenNoPanic(t *testing.T) {
	var (
		addTarget = []int{}
		want      = []int{}
	)
	rbr := &Rollbacker{
		rollbacks: []*rollback{
			{
				description: "add 1",
				funk: func(c context.Context) error {
					addTarget = append(addTarget, 1)
					return nil
				},
			},
		},
	}

	go func() {
		defer rbr.RollbackForPanic(context.TODO())
	}()

	time.Sleep(time.Second)

	assert.Exactly(t, addTarget, want)
}
