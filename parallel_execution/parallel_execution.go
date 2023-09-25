package parallel_execution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {

	if n < 1 {
		return errors.New("invalid workers count")
	}

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	// требования по завершению всех горутин
	var wg sync.WaitGroup

	next := make(chan interface{}, n)
	taskChan := make(chan Task)

	// создаём генератор задач
	wg.Add(1)
	go func() {
		defer close(taskChan)
		defer wg.Done()

		// вычитываем слайс задач и передаём в воркеры только если возможна дальнейшая обработка
		for _, task := range tasks {
			if _, ok := <-next; !ok {
				return
			}

			taskChan <- task
		}
	}()

	// запускаем генератор и первые N значений
	for i := 0; i < n; i++ {
		next <- struct{}{}
	}

	// Запускаем N воркеров
	resultChan := make(chan interface{}, n)
	defer close(resultChan)

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(numb int) {
			defer wg.Done()

			for task := range taskChan {
				resultChan <- task()
			}
		}(i)
	}

	errorsCounter := 0
	tasksCounter := 0

	// вычитываем результаты воркеров
	for res := range resultChan {
		tasksCounter += 1

		// считаем ошибки
		if res != nil {
			errorsCounter += 1
		}

		// если все задачи обработаны или количество ошибок достигнуто останавливаем генератор
		// генератор погасит воркеры
		if tasksCounter == len(tasks)-1 || errorsCounter == m {
			close(next)
			break
		}

		// если воркер освободился добавляем новую задачу
		next <- struct{}{}
	}

	wg.Wait()

	if errorsCounter > m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
