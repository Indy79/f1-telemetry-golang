package main

import (
	"bytes"
	"encoding/binary"
	"errors"
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
	buff := new(bytes.Buffer)
	reader := byteReader(buff)
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
		_, _, udperr := conn.ReadFromUDP(buff.Bytes())
		if udperr != nil {
			fmt.Printf("an error occured")
		}
		err := reader(&header)
		if err != nil {
			fmt.Printf("ooooopss %v", err)
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
			err := reader(&motion)
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
			err := reader(&session)
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
					"SessionUID":     strconv.FormatUint(header.SessionUID, 10),
				},
				sessionMap.Map(),
				time.Now())
			writeAPI.WritePoint(p)
			break
		case 6:
			var telemetry car.PacketCarTelemetryData
			err := reader(&telemetry)
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
					"SessionUID":     strconv.FormatUint(header.SessionUID, 10),
				},
				telemetryMap.Map(),
				time.Now())
			writeAPI.WritePoint(p)
			break
		case 7:
			var status car.PacketCarStatusData
			err := reader(&status)
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
					"SessionUID":     strconv.FormatUint(header.SessionUID, 10),
				},
				statusMap.Map(),
				time.Now())
			writeAPI.WritePoint(p)
			break
		}

	}
	conn.Close()
}

func byteReader(buffer *bytes.Buffer) func(data interface{}) error {
	fmt.Println(buffer)
	return func(data interface{}) error {
		err := binary.Read(buffer, binary.LittleEndian, data)
		if err != nil {
			fmt.Printf("some parsing error : %v", err)
			return errors.New("cannot parse binary")
		}
		return nil
	}
}
