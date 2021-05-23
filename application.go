package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"fr.serli.f1/application/packets"
	"fr.serli.f1/application/packets/car"
	"github.com/fatih/structs"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port, err := strconv.Atoi(os.Getenv("F1-SERVER-PORT"))
	if err != nil {
		fmt.Printf("Cannot parse port : %v", err)
	}
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(""),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	for {
		var header packets.PacketHeader
		_, _, udperr := conn.ReadFromUDP(p)
		if udperr != nil {
			fmt.Printf("an error occured")
		}
		buffer := bytes.NewReader(p)
		err := binary.Read(buffer, binary.LittleEndian, &header)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		// Create a new client using an InfluxDB server base URL and an authentication token
		// and set batch size to 20

		token := os.Getenv("INFLUX_TOKEN")
		influxDbUrl := os.Getenv("INFLUX_URI")
		organisation := os.Getenv("INFLUX_ORGA")
		bucket := os.Getenv("INFLUX_BUCKET")

		client := influxdb2.NewClientWithOptions(influxDbUrl, token, influxdb2.DefaultOptions().SetBatchSize(20))
		// Get non-blocking write client
		writeAPI := client.WriteAPI(organisation, bucket)

		switch header.PacketId {
		case 0:
			var motion packets.PacketMotionData
			err := binary.Read(buffer, binary.LittleEndian, &motion)
			if err != nil {
				fmt.Printf("Some error  %v", err)
				continue
			}
			motionMap := structs.New(motion)
			p := influxdb2.NewPoint(
				"motion",
				map[string]string{
					"PacketId":       string(header.PacketId),
					"PlayerCarIndex": string(header.PlayerCarIndex),
					"SessionUID":     strconv.FormatUint(header.SessionUID, 10),
				},
				motionMap.Map(),
				time.Now())
			writeAPI.WritePoint(p)
			break
		case 1:
			var session packets.PacketSessionData
			err := binary.Read(buffer, binary.LittleEndian, &session)
			if err != nil {
				fmt.Printf("Some error  %v", err)
				continue
			}
			sessionMap := structs.New(session)
			p := influxdb2.NewPoint(
				"session",
				map[string]string{
					"PacketId":       string(header.PacketId),
					"PlayerCarIndex": string(header.PlayerCarIndex),
					"SessionUID":     string(header.SessionUID),
				},
				sessionMap.Map(),
				time.Now())
			writeAPI.WritePoint(p)
			break
		case 6:
			var telemetry car.PacketCarTelemetryData
			err := binary.Read(buffer, binary.LittleEndian, &telemetry)
			if err != nil {
				fmt.Printf("Some error  %v", err)
				continue
			}
			telemetryMap := structs.New(telemetry.M_carTelemetryData[header.PlayerCarIndex])
			//fmt.Printf("%v", telemetryMap)
			p := influxdb2.NewPoint(
				"telemetry",
				map[string]string{
					"PacketId":       string(header.PacketId),
					"PlayerCarIndex": string(header.PlayerCarIndex),
					"SessionUID":     string(header.SessionUID),
				},
				telemetryMap.Map(),
				time.Now())
			writeAPI.WritePoint(p)
			break
		case 7:
			var status car.PacketCarStatusData
			err := binary.Read(buffer, binary.LittleEndian, &status)
			if err != nil {
				fmt.Printf("Some error  %v", err)
				continue
			}
			statusMap := structs.New(status.M_carStatusData[header.PlayerCarIndex])
			//fmt.Printf("%v", statusMap)
			p := influxdb2.NewPoint(
				"status",
				map[string]string{
					"PacketId":       string(header.PacketId),
					"PlayerCarIndex": string(header.PlayerCarIndex),
					"SessionUID":     string(header.SessionUID),
				},
				statusMap.Map(),
				time.Now())
			writeAPI.WritePoint(p)
			break
		}

	}
	conn.Close()
}
