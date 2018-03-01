package main

import (
	"github.com/spf13/viper"
	"fmt"
	//"github.com/fsnotify/fsnotify"
	//"time"
)

func main(){
	viper.SetConfigName("config")
	viper.AddConfigPath("./examples/viperExample")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	host:=viper.GetString("host.port")
	fmt.Println(host)
	//for i:=0;i<=1;i++ {
	//	time.Sleep(10*time.Second)
	//	viper.WatchConfig()
	//	viper.OnConfigChange(func(e fsnotify.Event) {
	//		fmt.Println("Config file changed:", e.Name)
	//		host:=viper.GetString("host.port")
	//		fmt.Println(host)
//
	//	})
	//}
}