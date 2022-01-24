package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	config "github.ibm.com/chandergovind/mq2p/pkg/config"
	manager "github.ibm.com/chandergovind/mq2p/pkg/metric-manager"
)

func main() {
	conf, err := config.ReadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	sm := manager.NewSimpleManager()

	sm.Init(manager.MetricSpec{
		Metrics: conf.Metrics,
	})

	opts := mqtt.NewClientOptions()
	opts.AddBroker(conf.MqConfig.Server).SetCleanSession(true)

	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		data := msg.Payload()

		var mqp manager.MQPayload
		err := json.Unmarshal(data, &mqp)
		if err != nil {
			fmt.Printf("Unrecognized payload. Error: %v\n", err)
		}

		sm.Update(mqp)
	}

	if token := c.Subscribe(conf.MqConfig.Topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9641", nil))
}
