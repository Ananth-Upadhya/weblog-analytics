package main

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const (
	hostRequestCountEndpoint    = "/hostRequestCount"
	httpStatusCodeCountEndpoint = "/httpStatusCodeCount"
	hourWithHighestRequestCountEndpoint = "/hourWithHighestRequestCount"
	pathNameComponentsTop10ResourcesEndpoint = "/pathNameComponents"
)

func main() {
	//create a router and corresponding groups
	router := gin.New()

	//hostRequestCount endpoint
	router.GET(hostRequestCountEndpoint, hostRequestCount)

	//httpStatusCodeCount endpoint
	router.GET(httpStatusCodeCountEndpoint, httpStatusCodeCount)

	//hourWithHighestRequestCount endpoint
	router.GET(hourWithHighestRequestCountEndpoint, hourWithHighestRequestCount)

	//pathNameComponentsTop10Resources endpoint
	router.GET(pathNameComponentsTop10ResourcesEndpoint, pathNameComponentsTop10Resources)

	err := router.Run(":" + "8080")
	if err != nil {
		panic(err)
	}
}

func hostRequestCount() {
	//(host, request-count) tuples for the top-10 frequent hosts
	command := fmt.Sprint("awk '{a[$1]++} END{for(i in a){print i,a[i]}}' Input_file| sort -nrk2.1 | head -10")
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println( "Error:", err )
	} else {
		fmt.Printf( "Top 10 Host IP and Count: %s\n", output )
	}
}

func httpStatusCodeCount() {
	//(HTTP-status-code, count) tuples, sorted by count
	command := fmt.Sprint("awk '{a[$9]++} END{for(i in a) {print i, a[i]}}' Input_file| sort -nrk2.1")
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println( "Error:", err )
	} else {
		fmt.Printf( "HTTP Status Code and Count: %s\n", output )
	}
}

func hourWithHighestRequestCount() {
	//the hour with the highest request count, along with the count
	command := fmt.Sprint("awk -F[:\\ ]  '{count[$5]++}; $12 == 200 { hour[$5]++} END { for (i in hour) print i, count[i] }' Input_file")
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println( "Error:", err )
	} else {
		fmt.Printf("Hour with the highest request count along with count: %s\n", output )
	}
}

func pathNameComponentsTop10Resources() {
	//the path name components of top-10 most frequently accessed resources
	command := fmt.Sprint("awk '{a[$7]++} END{for(i in a){print i,a[i]}}' Input_file| sort -nrk2.1 | head -10")
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println( "Error:", err )
	} else {
		fmt.Printf("Path name components of top-10 most frequently accessed resources: %s\n", output )
	}
}