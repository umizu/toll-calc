// simulates an OBU sending geolocation data to the server (the data_receiver)

package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/umizu/tollcalc/types"
)

const (
	sendInterval = time.Second
	wsEndpoint   = "ws://127.0.0.1:30000/ws"
)

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func genLatLon() (float64, float64) {
	return genCoord(), genCoord()
}

func main() {
	obuIDS := generateOBUIDS(3)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for _, v := range obuIDS {
			lat, lon := genLatLon()
			data := types.OBUData{
				OBUId: v,
				Lat:   lat,
				Lon:   lon,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("seeded")
}
