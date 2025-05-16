COUCHBASE_USER ?= jaba_admin
COUCHBASE_PASSWORD ?= jaba_pwd

.PHONY: init-cluster add-nodes rebalance

init-cluster:
	@echo "Initializing Couchbase cluster..."
	docker exec -it db1.lan bash -c "couchbase-cli cluster-init -c db1.lan:8091 \
		--cluster-username $(COUCHBASE_USER) \
		--cluster-password $(COUCHBASE_PASSWORD) \
		--services data,index,query \
		--cluster-name medium\
		--cluster-ramsize 3072 \
		--cluster-index-ramsize 512 \
		--cluster-analytics-ramsize 1024 \
		--cluster-eventing-ramsize 256"\

add-nodes:
	@echo "Adding db2.lan..."
	docker exec -it db1.lan bash -c "couchbase-cli server-add -c db1.lan:8091 \
		-u $(COUCHBASE_USER) -p $(COUCHBASE_PASSWORD) \
		--server-add=db2.lan \
		--server-add-username=$(COUCHBASE_USER) \
		--server-add-password=$(COUCHBASE_PASSWORD) \
		--services=data,index,query"
	@echo "Adding db3.lan..."
	docker exec -it db1.lan bash -c "couchbase-cli server-add -c db1.lan:8091 \
		-u $(COUCHBASE_USER) -p $(COUCHBASE_PASSWORD) \
		--server-add=db3.lan \
		--server-add-username=$(COUCHBASE_USER) \
		--server-add-password=$(COUCHBASE_PASSWORD) \
		--services=data,index,query"

rename-bucket:
	curl -u jaba_admin:jaba_pwd \
		 -X POST http://localhost:8091/node/controller/rename \
		 -d hostname=db1.lan


rebalance:
	@echo "Rebalancing cluster..."
	docker exec -it db1.lan bash -c "couchbase-cli rebalance -c db1.lan:8091 \
		-u $(COUCHBASE_USER) \
		-p $(COUCHBASE_PASSWORD)"
create-bucket:
	docker exec -it db1.lan bash -c "couchbase-cli bucket-create -c db1.lan:8091 \
  -u $(COUCHBASE_USER) -p $(COUCHBASE_PASSWORD) \
  --bucket=doom-scrolling \
  --bucket-ramsize=512\
   --bucket-type=couchbase"

cluster-up: init-cluster add-nodes rebalance create-bucket
luster-up: init-cluster create-bucket