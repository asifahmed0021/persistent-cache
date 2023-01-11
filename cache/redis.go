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