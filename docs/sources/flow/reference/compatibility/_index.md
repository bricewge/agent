---
aliases:
- /docs/grafana-cloud/agent/flow/reference/compatible-components/
- /docs/grafana-cloud/monitor-infrastructure/agent/flow/reference/compatible-components/
- /docs/grafana-cloud/monitor-infrastructure/integrations/agent/flow/reference/compatible-components/
- /docs/grafana-cloud/send-data/agent/flow/reference/compatible-components/
canonical: https://grafana.com/docs/agent/latest/flow/reference/compatible-components/
description: Learn about which components are compatible with each other in Grafana Agent Flow
title: Compatible components
weight: 400
---

# Compatible components

This section provides an overview of _some_ of the possible connections between 
compatible components in Grafana Agent Flow. 

For each common telemetry data type, we provide a list of compatible components
that can export or consume it.

> Note that connecting some components may require further configuration to make
> the connection work correctly. Please refer to the linked documentation for more
> details.

## Loki `LogsReceiver`

`LogsReceiver` is a [capsule]({{< relref "../../config-language/expressions/types_and_values/#capsules" >}})
that is exported by components that can receive Loki logs. Components that
consume `LogsReceiver` as an argument typically send logs to it. Use the
following components to build your Loki logs pipeline:

### Exporters
Below are components that _export_ Loki `LogsReceiver` grouped by namespace. Click
on the namespace to expand and see more detail.

<!-- START GENERATED SECTION: EXPORTERS OF Loki `LogsReceiver` -->dummy
dummy
dummy
dummy
dummy
<!-- END GENERATED SECTION: EXPORTERS OF Loki `LogsReceiver` -->

### Consumers
Below are components that _consume_ Loki `LogsReceiver` grouped by namespace. Click
on the namespace to expand and see more detail.

{{< collapse title="loki" >}}
- [`loki.echo`]({{< relref "../components/loki.echo.md" >}})
- [`loki.process`]({{< relref "../components/loki.process.md" >}})
- [`loki.relabel`]({{< relref "../components/loki.relabel.md" >}})
- [`loki.write`]({{< relref "../components/loki.write.md" >}})
{{< /collapse >}}

{{< collapse title="otelcol" >}}
- [`otelcol.receiver.loki`]({{< relref "../components/otelcol.receiver.loki.md" >}})
{{< /collapse >}}



<!-- START GENERATED SECTION: EXPORTERS OF Targets -->dummy
dummy
dummy
dummy
dummy
<!-- END GENERATED SECTION: EXPORTERS OF Targets -->