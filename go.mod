module knative.dev/eventing-redis

go 1.15

require (
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/gomodule/redigo v1.8.2
	github.com/google/go-cmp v0.5.2
	github.com/kelseyhightower/envconfig v1.4.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/eventing v0.18.1-0.20201001144430-5646fe1b227d
	knative.dev/pkg v0.0.0-20201012163217-54ad6c6d39a7
	knative.dev/reconciler-test v0.0.0-20201009163621-6dcba548c50f
	knative.dev/serving v0.18.0
	knative.dev/test-infra v0.0.0-20201014021030-ae3984a33f82
)

replace (
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/apiserver => k8s.io/apiserver v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/code-generator => k8s.io/code-generator v0.18.8
)
