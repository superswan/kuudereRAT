package model

type Client struct {
	Uuid string
	Ipaddr string
	Ts_first string
	Ts_last string
}

type Task struct {
	Uuid string
	Task_ID int
	Task_queued int
	Queue int
	Ts_last string
}

type Host struct {
	Uuid string
	Os_name string
	Location string
	User string
}