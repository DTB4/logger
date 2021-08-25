package logger

import "fmt"

func Info(message string, data interface{}) {
	fmt.Println("Log info message:", message, "Log info:", data)
}
func Error(message string, data interface{}) {
	fmt.Println("Log Error message:", message, "Log error:", data)
}
func Warning(message string, data interface{}) {
	fmt.Println("Log Warning message:", message, "Log warning:", data)
}
func Fatal (message string, data interface{}) {
	fmt.Println("Log Fatal message:", message, "Log Fatal:", data)
}