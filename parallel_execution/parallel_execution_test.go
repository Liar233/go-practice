package parallel_execution

import (
	"errors"
	"testing"
)

func TestRunNoError(t *testing.T) {

	tasks := make([]Task, 10)

	for i := 0; i < len(tasks); i++ {
		tasks[i] = func() error {
			return nil
		}
	}

	err := Run(tasks, 3, 3)

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
}

func TestRunFirstThreeErrors(t *testing.T) {

	tasks := make([]Task, 10)

	for i := 0; i < len(tasks); i++ {

		var fn Task

		if i < 3 {
			fn = func() error {
				return errors.New("error")
			}
		} else {
			fn = func() error {
				return nil
			}
		}

		tasks[i] = fn
	}

	err := Run(tasks, 3, 3)

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
}

func TestRunMiddleThreeErrors(t *testing.T) {

	tasks := make([]Task, 10)

	for i := 0; i < len(tasks); i++ {

		var fn Task

		if i > 4 || i < 7 {
			fn = func() error {
				return errors.New("error")
			}
		} else {
			fn = func() error {
				return nil
			}
		}

		tasks[i] = fn
	}

	err := Run(tasks, 3, 3)

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
}

func TestRunTasksLessThenWorkers(t *testing.T) {

	tasks := make([]Task, 5)

	for i := 0; i < len(tasks); i++ {
		tasks[i] = func() error {
			return nil
		}
	}

	err := Run(tasks, 7, 3)

	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
}
