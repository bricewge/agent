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

<!-- START GENERATED SECTION: COMPONENTS -->

{{< collapse title="faro" >}}
- [`faro.receiver`]({{< relref "../components/faro.receiver.md" >}})
{{< /collapse >}}
  
{{< collapse title="loki" >}}
- [`loki.process`]({{< relref "../components/loki.process.md" >}})
- [`loki.relabel`]({{< relref "../components/loki.relabel.md" >}})
- [`loki.source.api`]({{< relref "../components/loki.source.api.md" >}})
- [`loki.source.awsfirehose`]({{< relref "../components/loki.source.awsfirehose.md" >}})
- [`loki.source.azure_event_hubs`]({{< relref "../components/loki.source.azure_event_hubs.md" >}})
- [`loki.source.cloudflare`]({{< relref "../components/loki.source.cloudflare.md" >}})
- [`loki.source.docker`]({{< relref "../components/loki.source.docker.md" >}})
- [`loki.source.file`]({{< relref "../components/loki.source.file.md" >}})
- [`loki.source.gcplog`]({{< relref "../components/loki.source.gcplog.md" >}})
- [`loki.source.gelf`]({{< relref "../components/loki.source.gelf.md" >}})
- [`loki.source.heroku`]({{< relref "../components/loki.source.heroku.md" >}})
- [`loki.source.journal`]({{< relref "../components/loki.source.journal.md" >}})
- [`loki.source.kafka`]({{< relref "../components/loki.source.kafka.md" >}})
- [`loki.source.kubernetes`]({{< relref "../components/loki.source.kubernetes.md" >}})
- [`loki.source.kubernetes_events`]({{< relref "../components/loki.source.kubernetes_events.md" >}})
- [`loki.source.podlogs`]({{< relref "../components/loki.source.podlogs.md" >}})
- [`loki.source.syslog`]({{< relref "../components/loki.source.syslog.md" >}})
- [`loki.source.windowsevent`]({{< relref "../components/loki.source.windowsevent.md" >}})
{{< /collapse >}}

{{< collapse title="otelcol" >}}
- [`otelcol.exporter.loki`]({{< relref "../components/otelcol.exporter.loki.md" >}})
{{< /collapse >}}

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


