# mq2prom

![tests](https://github.com/IBM/mq2prom/actions/workflows/unit-test.yml/badge.svg)

Read metrics from a Message Queue in Json format and expose them in a Prometheus compatible format.

**Currently only works for MQTT compatible MQs.**

## Installing

```
make build
```
builds the `mq2prom` binary which can be run directly. By default metrics are exposed on port 9641.

Alternatively, the provided Docker image can be used directly.

## Message Formats

Consult the provided `payload.schema.json` for the JSON Schema specification of the expected metric payload message on the MQ. This had been built taking into account the [Prometheus constraints on metric and label names](https://prometheus.io/docs/concepts/data_model/). Validate your json payloads against this schema using tools specified [here](https://json-schema.org/implementations.html). Currently, the code **does not use the schema for validation**.

Note: all relevant metrics have to be configured in `config.yaml` to be exported. You are also recommended to following the [Prometheus naming guidelines](https://prometheus.io/docs/practices/naming/).

## Testing

Run unit tests:
```
go test ./...
```

For an end-to-end test, make sure that a MQTT compatible MQ is running at the port specified in the config file `config.yaml`. Run this binary.

Submit a message of the form `'[{"name": "simple_metric", "value": 10, "labels": {"label1": "value1"}}]'` onto the MQ.

Now, check that the metric is showing up in the metrics:
```
curl localhost:9641/metrics | grep "simple_metric"
```

## Why not Mqtt2prometheus?

[Mqtt2prometheus](https://github.com/hikhvar/mqtt2prometheus) is the inspiration for this project, but falls short on several fronts:

1. Focused on Sensor devices.
2. Works with data coming in from multiple topics - which is a counter-feature.
3. Does not allow generic labelling of metrics.
4. Does not support Histogram and Summary metric types.


## License

Licensed under Apache License 2.0.
