define build
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -tags=jsoniter,poll_opt -gcflags "all=-N -l" -o build/$(1) cmd/$(1)/main.go
endef


.PHONY: ip
ip:
	$(call build,pigupal-ip)

.PHONY: slave
slave:
	$(call build,pigupal-slave)

.PHONY: master
master:
	$(call build,pigupal)
	ssh -o StrictHostKeyChecking=no root@10.10.10.20 'systemctl stop pigupal'
	scp build/pigupal root@10.10.10.20:/root/palworld
	ssh -o StrictHostKeyChecking=no root@10.10.10.20 'systemctl start pigupal'