module knative.dev/eventing-redis

go 1.15

require (
	github.com/alecthomas/units v0.0.0-20201120081800-1786d5ef83d4 // indirect
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/go-redis/redis/v8 v8.4.2
	github.com/gomodule/redigo v1.8.3
	github.com/google/go-cmp v0.5.4
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.6.1
	go.uber.org/zap v1.16.0
	k8s.io/api v0.18.12
	k8s.io/apimachinery v0.18.12
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/eventing v0.19.1-0.20210112102830-a414aee50a2b
	knative.dev/hack v0.0.0-20201214230143-4ed1ecb8db24
	knative.dev/pkg v0.0.0-20210107022335-51c72e24c179
	knative.dev/serving v0.19.1-0.20210111190837-560a4f678dd7
)

replace (
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/apiserver => k8s.io/apiserver v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/code-generator => k8s.io/code-generator v0.18.8
)
