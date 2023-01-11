package main
import (
		"fmt"
		c "redis/cache"
		"os"
)

func main() {
	redis, err := c.GetRedis()
	if(err!=nil){
		fmt.Println("failed to initialize redis!")
		os.Exit(0)
	}

	var flag int
	fmt.Println("Press 1 to put and 2 to get")
	fmt.Scanln(&flag)

	if(flag==1){
		fmt.Println("its a put")
		var key string
		var value string
		fmt.Println("Enter key")
		fmt.Scanln(&key)
		fmt.Println("Enter value")
		fmt.Scanln(&value)

		err := redis.Put(key, value)
		if(err!=nil){
			fmt.Print("Error: ")
			fmt.Println(err)
		} else {
			fmt.Println("Success: key value pair inserted")
		}

	} else {
		fmt.Println("its a get")
		var key string
		fmt.Println("Enter key")
		fmt.Scanln(&key)

		value, err := redis.Get(key)
		if(err!=nil){
			fmt.Print("Error: ")
			fmt.Println(err)
		} else{
		fmt.Println("Success: " + value)
		}
	}

}