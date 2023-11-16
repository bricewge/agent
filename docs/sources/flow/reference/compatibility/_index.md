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

Components that _export_ Loki `LogsReceiver`
- [`faro.receiver`]({{< relref "../components/faro.receiver.md" >}})
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
- [`otelcol.exporter.loki`]({{< relref "../components/otelcol.exporter.loki.md" >}})


Components that _consume_ Loki Logs receiver:
- [`loki.echo`]({{< relref "../components/loki.echo.md" >}})
- [`loki.process`]({{< relref "../components/loki.process.md" >}})
- [`loki.relabel`]({{< relref "../components/loki.relabel.md" >}})
- [`loki.write`]({{< relref "../components/loki.write.md" >}})
- [`otelcol.receiver.loki`]({{< relref "../components/otelcol.receiver.loki.md" >}})


## Targets

Targets are a `list(map(string))` - a [list]({{< relref "../../config-language/expressions/types_and_values/#naming-convention" >}}) of [maps]({{< relref "../../config-language/expressions/types_and_values/#naming-convention" >}}) with [string]({{< relref "../../config-language/expressions/types_and_values/#strings" >}}) values. As such, 
they can contain different key-value pairs and can be used with a wide range of 
components. Some components export Targets with key-value pairs specified in
the reference documentation, while other components accept Targets as arguments.
Some components require Targets to contain specific key-value pairs in order
to work correctly. It is recommended to always check component reference for 
details when working with Targets.

Components that _output_ Targets:
- [`discovery.azure`]({{< relref "../components/discovery.azure.md" >}})
- [`discovery.consul`]({{< relref "../components/discovery.consul.md" >}})
- [`discovery.consulagent`]({{< relref "../components/discovery.consulagent.md" >}})
- [`discovery.digitalocean`]({{< relref "../components/discovery.digitalocean.md" >}})
- [`discovery.dns`]({{< relref "../components/discovery.dns.md" >}})
- [`discovery.docker`]({{< relref "../components/discovery.docker.md" >}})
- [`discovery.dockerswarm`]({{< relref "../components/discovery.dockerswarm.md" >}})
- [`discovery.ec2`]({{< relref "../components/discovery.ec2.md" >}})
- [`discovery.eureka`]({{< relref "../components/discovery.eureka.md" >}})
- [`discovery.file`]({{< relref "../components/discovery.file.md" >}})
- [`discovery.gce`]({{< relref "../components/discovery.gce.md" >}})
- [`discovery.hetzner`]({{< relref "../components/discovery.hetzner.md" >}})
- [`discovery.http`]({{< relref "../components/discovery.http.md" >}})
- [`discovery.ionos`]({{< relref "../components/discovery.ionos.md" >}})
- [`discovery.kubelet`]({{< relref "../components/discovery.kubelet.md" >}})
- [`discovery.kubernetes`]({{< relref "../components/discovery.kubernetes.md" >}})
- [`discovery.kuma`]({{< relref "../components/discovery.kuma.md" >}})
- [`discovery.lightsail`]({{< relref "../components/discovery.lightsail.md" >}})
- [`discovery.linode`]({{< relref "../components/discovery.linode.md" >}})
- [`discovery.marathon`]({{< relref "../components/discovery.marathon.md" >}})
- [`discovery.nerve`]({{< relref "../components/discovery.nerve.md" >}})
- [`discovery.nomad`]({{< relref "../components/discovery.nomad.md" >}})
- [`discovery.openstack`]({{< relref "../components/discovery.openstack.md" >}})
- [`discovery.puppetdb`]({{< relref "../components/discovery.puppetdb.md" >}})
- [`discovery.relabel`]({{< relref "../components/discovery.relabel.md" >}})
- [`discovery.scaleway`]({{< relref "../components/discovery.scaleway.md" >}})
- [`discovery.serverset`]({{< relref "../components/discovery.serverset.md" >}})
- [`discovery.triton`]({{< relref "../components/discovery.triton.md" >}})
- [`discovery.uyuni`]({{< relref "../components/discovery.uyuni.md" >}})
- [`local.file_match`]({{< relref "../components/local.file_match.md" >}})
- [`prometheus.exporter.agent`]({{< relref "../components/prometheus.exporter.agent.md" >}})
- [`prometheus.exporter.apache`]({{< relref "../components/prometheus.exporter.apache.md" >}})
- [`prometheus.exporter.azure`]({{< relref "../components/prometheus.exporter.azure.md" >}})
- [`prometheus.exporter.blackbox`]({{< relref "../components/prometheus.exporter.blackbox.md" >}})
- [`prometheus.exporter.cadvisor`]({{< relref "../components/prometheus.exporter.cadvisor.md" >}})
- [`prometheus.exporter.cloudwatch`]({{< relref "../components/prometheus.exporter.cloudwatch.md" >}})
- [`prometheus.exporter.consul`]({{< relref "../components/prometheus.exporter.consul.md" >}})
- [`prometheus.exporter.dnsmasq`]({{< relref "../components/prometheus.exporter.dnsmasq.md" >}})
- [`prometheus.exporter.elasticsearch`]({{< relref "../components/prometheus.exporter.elasticsearch.md" >}})
- [`prometheus.exporter.gcp`]({{< relref "../components/prometheus.exporter.gcp.md" >}})
- [`prometheus.exporter.github`]({{< relref "../components/prometheus.exporter.github.md" >}})
- [`prometheus.exporter.kafka`]({{< relref "../components/prometheus.exporter.kafka.md" >}})
- [`prometheus.exporter.memcached`]({{< relref "../components/prometheus.exporter.memcached.md" >}})
- [`prometheus.exporter.mongodb`]({{< relref "../components/prometheus.exporter.mongodb.md" >}})
- [`prometheus.exporter.mssql`]({{< relref "../components/prometheus.exporter.mssql.md" >}})
- [`prometheus.exporter.mysql`]({{< relref "../components/prometheus.exporter.mysql.md" >}})
- [`prometheus.exporter.oracledb`]({{< relref "../components/prometheus.exporter.oracledb.md" >}})
- [`prometheus.exporter.postgres`]({{< relref "../components/prometheus.exporter.postgres.md" >}})
- [`prometheus.exporter.process`]({{< relref "../components/prometheus.exporter.process.md" >}})
- [`prometheus.exporter.redis`]({{< relref "../components/prometheus.exporter.redis.md" >}})
- [`prometheus.exporter.snmp`]({{< relref "../components/prometheus.exporter.snmp.md" >}})
- [`prometheus.exporter.snowflake`]({{< relref "../components/prometheus.exporter.snowflake.md" >}})
- [`prometheus.exporter.squid`]({{< relref "../components/prometheus.exporter.squid.md" >}})
- [`prometheus.exporter.statsd`]({{< relref "../components/prometheus.exporter.statsd.md" >}})
- [`prometheus.exporter.unix`]({{< relref "../components/prometheus.exporter.unix.md" >}})
- [`prometheus.exporter.vsphere`]({{< relref "../components/prometheus.exporter.vsphere.md" >}})
- [`prometheus.exporter.windows`]({{< relref "../components/prometheus.exporter.windows.md" >}})


Components that _accept_ Targets:
- [`discovery.relabel`]({{< relref "../components/discovery.relabel.md" >}})
- [`local.file_match`]({{< relref "../components/local.file_match.md" >}})
- [`loki.source.docker`]({{< relref "../components/loki.source.docker.md" >}})
- [`loki.source.file`]({{< relref "../components/loki.source.file.md" >}})
- [`loki.source.kubernetes`]({{< relref "../components/loki.source.kubernetes.md" >}})
- [`otelcol.processor.discovery`]({{< relref "../components/otelcol.processor.discovery.md" >}})
- [`prometheus.scrape`]({{< relref "../components/prometheus.scrape.md" >}})
- [`pyroscope.scrape`]({{< relref "../components/pyroscope.scrape.md" >}})

