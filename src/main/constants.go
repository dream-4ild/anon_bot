package main

import (
	"os"
	"strconv"
)

var (
	moderChatID, _    = strconv.ParseInt(os.Getenv("MODER_CHAT_ID"), 10, 64)
	targetChatID, _   = strconv.ParseInt(os.Getenv("TARGET_CHAT_ID"), 10, 64)
	targetThreadID, _ = strconv.Atoi(os.Getenv("TARGET_THREAD_ID"))
)
