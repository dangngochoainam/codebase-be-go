package main

import (
	"context"
	"example/internal/common/helper/loghelper"
)

func main() {
	loghelper.InitZap("go-example", "dev")

	loghelper.Logger.WithContext(context.Background()).Info(";sldfksd")
	// Logger.DPanic("asdasda")
	// Logger.Debug(";sldfksd")
	// Logger.Error(";sldfksd")
	// Logger.Fatalln(";sldfksd")
	// Logger.Warn(";sldfksd")
}
