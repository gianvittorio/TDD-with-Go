package concurrency

import "sync"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

type Dictionary[T any] interface {
	Get(string) (T, error)
	Add(string, T) error
}

type ConcurrenctDictionary struct {
	dictionary map[string]bool
	lock sync.Mutex
}

type ErrConcurrentDictionary string

const (
	ErrNotFound ErrConcurrentDictionary = "could not find the key you were looking for"
)

func (err ErrConcurrentDictionary) Error() string {
	return string(err)
}

func (dict *ConcurrenctDictionary) Get(key string) (bool, error) {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	switch value, exists := dict.dictionary[key]; {
		case exists:
			return value, nil
		default:
			return false, ErrNotFound
	}
}

func (dict *ConcurrenctDictionary) Add(key string, value bool) error {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	dict.lock.TryLock()

	if _, exists := dict.dictionary[key]; exists {
		return ErrNotFound
	}

	dict.dictionary[key] = value
	return nil
}

func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	result := ConcurrenctDictionary{dictionary: make(map[string]bool)}

	for _, url := range urls {
		go func() {
			result.Add(url, checker(url))
		}()
	}

	return result.dictionary
}
