package cache

import ( f "redis/file"
	)

func GetRedis() (Cache, error) {
	filePath := "./cacheFile.txt"
	fs, err := f.GetNewLocalFileSystem(filePath)
	return Redis{"myRedis", fs}, err
}

type Redis struct{
	Name string
	FileSystem f.FileSystem
}

func (r Redis) Put(key string, value string) error{
	err := r.FileSystem.Append(key, value)

	return err
}

func (r Redis) Get(key string) (string, error) {
	value, err := r.FileSystem.GetValueForKey(key)

	return value, err
}