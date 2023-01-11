package file

type FileSystem interface{
	Append(string,string) error
	GetValueForKey(key string) (string,error)
}