package main

import (
    "log"
    "time"
    "math/rand"
    
    "desukit/asciiArt"
    "desukit/listener"
    "desukit/httpui"
    "desukit/sqldb"
    "desukit/sqldb/schema"

    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        log.Println("Error in config file: ", err)
    }

    // Configuration for listener 
    
    CONN_HOST := viper.GetString("listener.host")
    CONN_PORT := viper.GetString("listener.port")
    min := 1
    max := 4 
    rand.Seed(time.Now().UnixNano())
    n := min + rand.Intn(max-min+1)
    asciiArt.PrintArt(n)

    // ----------- START DB -----------------------
    // database functionality. sqldb makes the database a global variable
    // ConnectDB() only needs to be ran once from main and the db is 
    // accessible anywhere by importing sqldb package like within httpui for example
    sqldb.ConnectDB()
    //schema.InitializeDB() will create the tables if they haven't been created yet
    schema.InitializeDB()
    
    // Start Listener
    go listener.Start(CONN_HOST, CONN_PORT)

    // Start webserver for web interface
    go httpui.Start()
    
    // "Event loop"
    for {

    }
}




