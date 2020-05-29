module github.com/vmware-labs/yaml-jsonpointer/yamled

go 1.14

require (
	github.com/vmware-labs/go-yaml-edit v0.1.4
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

replace gopkg.in/yaml.v3 => github.com/atomatt/yaml v0.0.0-20200228174225-55c5cf55e3ee
