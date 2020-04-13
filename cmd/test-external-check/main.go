package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Comcast/kuberhealthy/v2/pkg/checks/external"
	checkclient "github.com/Comcast/kuberhealthy/v2/pkg/checks/external/checkclient"
)

var reportFailure bool
var reportDelay time.Duration

func init() {

	// enable debug logging on the check client
	checkclient.Debug = true

	var err error

	// parse REPORT_FAILURE environment var
	reportFailure, err = strconv.ParseBool(getEnv("REPORT_FAILURE", "false"))
	if err != nil {
		log.Fatalln("Failed to parse REPORT_FAILURE env var:", err)
	}

	// parse REPORT_DELAY environment var
	reportDelayStr := getEnv("REPORT_DELAY", "10s")
	reportDelay, err = time.ParseDuration(reportDelayStr)
	if err != nil {
		log.Fatalln("Failed to parse REPORT_DELAY env var:", err)
	}
}

// getEnv returns a default value if there is no environment variable set
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {

	log.Println("Using kuberhealthy reporting url", os.Getenv(external.KHReportingURL))
	log.Println("Waiting", reportDelay, "seconds before reporting...")
	time.Sleep(reportDelay)

	var err error
	if reportFailure {
		log.Println("Reporting failure...")
		err = checkclient.ReportFailure([]string{"Test has failed!"})
	} else {
		log.Println("Reporting success...")
		err = checkclient.ReportSuccess()
	}

	if err != nil {
		log.Println("Error reporting to Kuberhealthy servers:", err)
		return
	}
	log.Println("Successfully reported to Kuberhealthy servers")
}
