package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	manager "github.ibm.com/chandergovind/mq2p/pkg/metric-manager"
)

func main() {
	sm := manager.NewSimpleManager()

	// Setup a single metric - Gauge
	sm.Init(manager.MetricSpec{
		Metrics: []manager.MetricConfig{
			{MQName: "simple_metric", PromName: "simple_metric_prom", Help: "dummy metric value", Type: manager.Gauge, Labels: []string{"label1"}},
		},
	})

	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://0.0.0.0:1883").SetCleanSession(true)

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

	if token := c.Subscribe("topic", 0, messageHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9146", nil))
}
