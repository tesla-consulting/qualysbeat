module github.com/fs015/qualysbeat

go 1.13

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.2.0+incompatible
	github.com/Shopify/sarama => github.com/elastic/sarama v0.0.0-20191122160421-355d120d0970
	github.com/docker/docker => github.com/docker/engine v0.0.0-20191113042239-ea84732a7725
	github.com/docker/go-plugins-helpers => github.com/elastic/go-plugins-helpers v0.0.0-20200207104224-bdf17607b79f
	github.com/dop251/goja => github.com/andrewkroh/goja v0.0.0-20190128172624-dd2ac4456e20
	github.com/fsnotify/fsevents => github.com/elastic/fsevents v0.0.0-20181029231046-e1d381a4d270
	github.com/fsnotify/fsnotify => github.com/adriansr/fsnotify v0.0.0-20180417234312-c9bbe1f46f1d
	github.com/google/gopacket => github.com/adriansr/gopacket v1.1.18-0.20200327165309-dd62abfa8a41
	github.com/insomniacslk/dhcp => github.com/elastic/dhcp v0.0.0-20200227161230-57ec251c7eb3 // indirect
	github.com/tonistiigi/fifo => github.com/containerd/fifo v0.0.0-20190816180239-bda0ff6ed73c
)

require (
	github.com/basgys/goxml2json v1.1.0
	github.com/dlclark/regexp2 v1.2.0 // indirect
	github.com/dop251/goja v0.0.0-20200326102500-6438c8ddc517 // indirect
	github.com/dop251/goja_nodejs v0.0.0-20200128125109-2d688c7e0ac4 // indirect
	github.com/elastic/beats/v7 v7.0.0-alpha2.0.20200407202534-97750c8544b1
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/josephspurrier/goversioninfo v0.0.0-20200309025242-14b0ab84c6ca // indirect
	github.com/magefile/mage v1.10.0
	github.com/mitchellh/gox v1.0.1
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/pierrre/gotestcover v0.0.0-20160113212533-7b94f124d338
	github.com/prometheus/procfs v0.0.11 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/reviewdog/reviewdog v0.9.17
	github.com/tidwall/gjson v1.6.0
	github.com/tsg/go-daemon v0.0.0-20200207173439-e704b93fd89b
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f
	golang.org/x/sys v0.0.0-20200408040146-ea54a3c99b9b // indirect
	golang.org/x/tools v0.0.0-20200408032209-46bd65c8538f
	howett.net/plist v0.0.0-20200225050739-77e249a2e2ba // indirect
)
