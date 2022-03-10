package main

import (
	"encoding/json"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	config "github.com/IBM/mq2prom/pkg/config"
	manager "github.com/IBM/mq2prom/pkg/metric-manager"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	conf, err := config.ReadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	log.Info("Successfully read config file.")

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

	log.Info("Connected to MQTT Bus.")

	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		data := msg.Payload()

		var mqp manager.MQPayload
		err := json.Unmarshal(data, &mqp)
		if err != nil {
			log.Errorf("Unrecognized payload: %v\n", err)
		}

		sm.Update(mqp)
	}

	if token := c.Subscribe(conf.MqConfig.Topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Error("Error in subscribing to topic: ", token.Error())
		os.Exit(1)
	}
	log.Info("Subscribed to topic: ", conf.MqConfig.Topic)

	http.Handle("/metrics", promhttp.Handler())

	log.Info("Starting up metric endpoint at :9641")
	log.Fatal(http.ListenAndServe(":9641", nil))
}
