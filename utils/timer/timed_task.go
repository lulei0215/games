package timer

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Timer interface {
	// Cron
	FindCronList() map[string]*taskManager
	// Task
	AddTaskByFuncWithSecond(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) // Task Func
	// Task
	AddTaskByJobWithSeconds(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	//
	AddTaskByFunc(cronName string, spec string, task func(), taskName string, option ...cron.Option) (cron.EntryID, error)
	//   Run
	AddTaskByJob(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	// taskNamecron
	FindCron(cronName string) (*taskManager, bool)
	// cron
	StartCron(cronName string)
	// cron
	StopCron(cronName string)
	// crontask
	FindTask(cronName string, taskName string) (*task, bool)
	// idcrontask
	RemoveTask(cronName string, id int)
	// taskNamecrontask
	RemoveTaskByName(cronName string, taskName string)
	// cronName
	Clear(cronName string)
	// cron
	Close()
}

type task struct {
	EntryID  cron.EntryID
	Spec     string
	TaskName string
}

type taskManager struct {
	corn  *cron.Cron
	tasks map[cron.EntryID]*task
}

// timer
type timer struct {
	cronList map[string]*taskManager
	sync.Mutex
}

// AddTaskByFunc
func (t *timer) AddTaskByFunc(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddFunc(spec, fun)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByFuncWithSecond WithSeconds
func (t *timer) AddTaskByFuncWithSecond(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddFunc(spec, fun)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByJob
func (t *timer) AddTaskByJob(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddJob(spec, job)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByJobWithSeconds
func (t *timer) AddTaskByJobWithSeconds(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			corn:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].corn.AddJob(spec, job)
	t.cronList[cronName].corn.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// FindCron cronNamecron
func (t *timer) FindCron(cronName string) (*taskManager, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.cronList[cronName]
	return v, ok
}

// FindTask cronNamecron
func (t *timer) FindTask(cronName string, taskName string) (*task, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.cronList[cronName]
	if !ok {
		return nil, ok
	}
	for _, t2 := range v.tasks {
		if t2.TaskName == taskName {
			return t2, true
		}
	}
	return nil, false
}

// FindCronList
func (t *timer) FindCronList() map[string]*taskManager {
	t.Lock()
	defer t.Unlock()
	return t.cronList
}

// StartCron
func (t *timer) StartCron(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Start()
	}
}

// StopCron
func (t *timer) StopCron(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Stop()
	}
}

// RemoveTask cronName
func (t *timer) RemoveTask(cronName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Remove(cron.EntryID(id))
		delete(v.tasks, cron.EntryID(id))
	}
}

// RemoveTaskByName cronName taskName
func (t *timer) RemoveTaskByName(cronName string, taskName string) {
	fTask, ok := t.FindTask(cronName, taskName)
	if !ok {
		return
	}
	t.RemoveTask(cronName, int(fTask.EntryID))
}

// Clear
func (t *timer) Clear(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.corn.Stop()
		delete(t.cronList, cronName)
	}
}

// Close
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.cronList {
		v.corn.Stop()
	}
}

func NewTimerTask() Timer {
	return &timer{cronList: make(map[string]*taskManager)}
}
