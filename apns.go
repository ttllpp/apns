package main

import (
	"crypto/tls"
	"fmt"
	"github.com/sideshow/apns2"
	"log"
)

func main() {
	cert, err := tls.LoadX509KeyPair("certificate.pem", "key.unencrypted.pem")

	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "token"
	notification.Topic = "包名"
	notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`)

	client := apns2.NewClient(cert).Development()
	res, err := client.Push(notification)

	if res.Sent() {
		log.Println("Sent:", res.ApnsID)
	} else {
		fmt.Printf("Not Sent: %v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
	}
}
