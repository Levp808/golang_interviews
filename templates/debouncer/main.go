package debouncer

/*
Шаблон Debounce (Антидребезг) ограничивает частоту вызова функции,
вследствие чего выполняется только первый или последний вызов в группе.
*/

import (
	"context"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()
		defer func() {
			threshold = time.Now().Add(d)
			m.Unlock()
		}()
		//Если событие произошло до заданного промежутка, то возвращаем закэшированный результат
		if time.Now().Before(threshold) {
			return result, err
		}
		result, err = circuit(ctx)
		return result, err
	}
}

func DebounceLast(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time = time.Now()
	var ticker *time.Ticker
	var result string
	var err error
	var once sync.Once
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()
		defer m.Unlock()
		threshold = time.Now().Add(d)
		once.Do(func() {
			ticker = time.NewTicker(time.Millisecond * 100)
			go func() {
				defer func() {
					m.Lock()
					ticker.Stop()
					once = sync.Once{}
					m.Unlock()
				}()
				for {
					select {
					case <-ticker.C:
						m.Lock()
						if time.Now().After(threshold) {
							result, err = circuit(ctx)
							m.Unlock()
							return
						}
						m.Unlock()
					case <-ctx.Done():
						m.Lock()
						result, err = "", ctx.Err()
						m.Unlock()
						return
					}
				}
			}()
		})
		return result, err
	}
}
