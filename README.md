# mq2p

Read metrics from a Message Queue in Json format and expose them in a Prometheus compatible format.

## Why not Mqtt2prometheus?

[Mqtt2prometheus](https://github.com/hikhvar/mqtt2prometheus) is the inspiration for this project, but falls short on several fronts:

1. Focused on Sensor devices.
2. Works with data coming in from multiple topics - which is a counter-feature.
3. Does not allow generic labelling of metrics.
4. Does not support Histogram and Summary metric types.

