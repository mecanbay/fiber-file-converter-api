package metrics

import "github.com/prometheus/client_golang/prometheus/collectors"

func Init() {
	Registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
}
