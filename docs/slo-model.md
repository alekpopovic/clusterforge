# SLO model

An SLO expresses a measured reliability objective over a defined window; it is
not the same as a one-time health check. ClusterForge stores availability,
latency and error-rate targets as metadata so owners can document expectations.
It does not calculate error budgets or claim that a target is achieved.

Define the measurement source, window, eligible traffic, exclusions, alerting,
owner and response before treating a value as an operational SLO. Production
validation still requires monitoring evidence, user-journey tests and human
review; config values alone prove nothing.
