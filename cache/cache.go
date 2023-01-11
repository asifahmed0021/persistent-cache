package cache

type Cache interface{
	Get(string) (string, error)
	Put(string, string) error
}

