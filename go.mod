module go.f110.dev/rules_k8s_controller

go 1.16

require (
	github.com/Masterminds/semver/v3 v3.1.1
	github.com/bazelbuild/buildtools v0.0.0-20210227132407-f2aed9ee205d
	github.com/docker/spdystream v0.0.0-20160310174837-449fdfce4d96 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/google/go-github/v33 v33.0.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.21.0
	k8s.io/apiextensions-apiserver v0.21.0 // indirect
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
	sigs.k8s.io/kind v0.10.0
)
