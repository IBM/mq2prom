# mq2prom

Read metrics from a Message Queue in Json format and expose them in a Prometheus compatible format.

## Installing

```
make build
```
builds the `mq2prom` binary which can be run directly. By default metrics are exposed on port 9641.

## Testing

Make sure that a MQTT compatible MQ is running at the port specified in the config file `config.yaml`. Run this binary.

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
