package timer

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var job = mockJob{}

type mockJob struct{}

func (job mockJob) Run() {
	mockFunc()
}

func mockFunc() {
	time.Sleep(time.Second)
	fmt.Println("1s...")
}

func TestNewTimerTask(t *testing.T) {
	tm := NewTimerTask()
	_tm := tm.(*timer)

	{
		_, err := tm.AddTaskByFunc("func", "@every 1s", mockFunc, "mockfunc")
		assert.Nil(t, err)
		_, ok := _tm.cronList["func"]
		if !ok {
			t.Error("no find func")
		}
	}

	{
		_, err := tm.AddTaskByJob("job", "@every 1s", job, "job mockfunc")
		assert.Nil(t, err)
		_, ok := _tm.cronList["job"]
		if !ok {
			t.Error("no find job")
		}
	}

	{
		_, ok := tm.FindCron("func")
		if !ok {
			t.Error("no find func")
		}
		_, ok = tm.FindCron("job")
		if !ok {
			t.Error("no find job")
		}
		_, ok = tm.FindCron("none")
		if ok {
			t.Error("find none")
		}
	}
	{
		tm.Clear("func")
		_, ok := tm.FindCron("func")
		if ok {
			t.Error("find func")
		}
	}
	{
		a := tm.FindCronList()
		b, c := tm.FindCron("job")
		fmt.Println(a, b, c)
	}
}
