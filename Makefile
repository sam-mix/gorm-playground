.PHONY:sql
sql:
	@cd proto && \
	protoc --sql_out=. --go_out=. person.proto && \
	protoc -I. --include_imports --include_source_info --descriptor_set_out=person.pb person.proto && \
	protoc-go-inject-tag -input=./person.pb.go && \
	gofmt -s -w ./person.pb.go

.PHONY:clean
clean: 
	@rm proto/person.pb.go  && rm proto/person.sql.go && rm proto/person.pb

.PHONY:test
test:
	@cd tests && go test -timeout 30s -run ^TestSometging$ playground/tests


.PHONY:rtx
rtx:
	@echo ========================== v2 ==========================
	@cd cases/tx/v2 && go run main.go
	@echo ========================== v1 ==========================
	@cd cases/tx/v1 && go run main.go

.PHONY:tx
tx:
	@echo ========================== v1 ==========================
	@cd cases/tx/v1 && go run main.go
	@echo ========================== v2 ==========================
	@cd cases/tx/v2 && go run main.go

.PHONY:db
db:
	@cd cases/db && go run main.go

.PHONY:up
up:
	@docker compose up -d

.PHONY:down
down:
	@docker compose down

.PHONY:reboot
reboot:
	@docker compose down
	@docker compose up -d

.PHONY:ctx
ctx:
	@cd cases/ctx-tx/v2 && go run main.go

.PHONY:e
e:
	@cd cases/embed/v2 && go run main.go

.PHONY:err001
err001:
	@cd cases/err001 && go run main.go

.PHONY:xx
xx:
	@cd cases/xx && go run main.go

.PHONY:x1
x1:
	@cd cases/x1 && go run main.go

.PHONY:install_sql
install_sql:
	@go install code.guanmai.cn/public_code/protoc-gen-sql@v0.3.0

.PHONY:t
t:
	@cd cases/tow_table/v2 && go run main.go

.PHONY:umt
umt:
	@cd cases/tx-demos/umt && go run main.go

.PHONY:umt2
umt2:
	@cd cases/tx-demos/umt2 && go run main.go

.PHONY:umt3
umt3:
	@cd cases/tx-demos/umt3 && go run main.go

.PHONY:umt4
umt4:
	@cd cases/tx-demos/umt4 && go run main.go

.PHONY:umt5
umt5:
	@cd cases/tx-demos/umt5/v1 && go run main.go

.PHONY:umt6
umt6:
	@echo ========================== v1 ==========================
	@cd cases/tx-demos/umt6/v1 && go run main.go
	@echo ========================== v2 ==========================
	@cd cases/tx-demos/umt6/v2 && go run main.go
	@echo ========================== v2 2 ==========================
	@cd cases/tx-demos/umt6/v2_2 && go run main.go

.PHONY:drop_all_table
drop_all_table:
	@cd cases/drop_all_table && go run main.go


.PHONY:mutex1
mutex1:
	@echo ========================== v1 ==========================
	@cd cases/tx-demos/mutex1/v1 && go run main.go
	@echo ========================== v2 ==========================
	@cd cases/tx-demos/mutex1/v2 && go run main.go



.PHONY:decimal001
decimal001:
	@cd playground/decimal/001 && go run main.go

.PHONY:save
save:
	@cd cases/save && go run main.go

.PHONY:save-v1
save-v1:
	@cd cases/save/v1 && go run main.go

.PHONY:save-ok
save-ok:
	@cd cases/save/ok && go run main.go

.PHONY:chainable-api
chainable-api:
	@cd cases/chainable-api/table/failed && go run main.go

.PHONY:chainable-api-ok
chainable-api-ok:
	@cd cases/chainable-api/table/success && go run main.go


.PHONY:in
in:
	@cd cases/in && go run main.go


.PHONY:pk
pk:
	@echo ========================== v1 ==========================
	@cd cases/pk/v1 && go run main.go
	@echo ========================== v2 ==========================
	@cd cases/pk/v2 && go run main.go

.PHONY:tl
tl:
	@echo ========================== auto ========================== 
	@cd cases/tx-lock/auto && go run main.go
	@echo ========================== man ==========================
	@cd cases/tx-lock/man && go run main.go


.PHONY:raw_sql
raw_sql:
	@cd cases/raw_sql && go run main.go

.PHONY:loop-save
loop-save:
	@cd cases/loop-save && go run main.go


.PHONY:pna-save
pna-save:
	@echo ========================== v1 ========================== 
	@cd cases/pna-save/v1 && go run main.go
	@echo ========================== v2 ==========================
	@cd cases/pna-save/v2 && go run main.go
	