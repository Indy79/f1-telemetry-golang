package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "fr.serli.f1/application/packets"
    "net"
)

func main() {
    p :=  make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 20777,
        IP: net.ParseIP(""),
    }
    conn, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    for {
        var header packets.PacketHeader
        _, _,udperr := conn.ReadFromUDP(p)
        if udperr != nil {
            fmt.Printf("an error occured")
        }
        buffer := bytes.NewReader(p)
        err := binary.Read(buffer, binary.LittleEndian, &header)
        if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }
        switch header.PacketId {
            case 0:
                var motion packets.PacketMotionData
                err := binary.Read(buffer, binary.LittleEndian, &motion)
                if err !=  nil {
                    fmt.Printf("Some error  %v", err)
                    continue
                }
                fmt.Printf("%+v\n", motion)
                break
            case 1:
                var session packets.PacketSessionData
                err := binary.Read(buffer, binary.LittleEndian, &session)
                if err !=  nil {
                    fmt.Printf("Some error  %v", err)
                    continue
                }
                fmt.Printf("%+v\n", session)
                break
        }

    }
    conn.Close()
}