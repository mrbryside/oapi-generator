generate-server:
	rm -rf internal/generated/oapi/$(name)dto && rm -rf internal/generated/oapi/$(name)srv
	mkdir internal/generated/oapi/$(name)dto && mkdir internal/generated/oapi/$(name)srv
	sed 's/#name/$(name)dto/g' internal/rest/v1/spec/server.cfg.yaml >> internal/rest/v1/spec/server-$(name).cfg.yaml
	(oapi-codegen -generate types -o internal/generated/oapi/$(name)dto/$(name)dto.go -package $(name)dto internal/rest/v1/spec/$(name).yaml || rm -rf internal/rest/v1/spec/server-$(name).cfg.yaml ) && (oapi-codegen --config internal/rest/v1/spec/server-$(name).cfg.yaml -package $(name)srv -o internal/generated/oapi/$(name)srv/$(name)srv.go internal/rest/v1/spec/$(name).yaml || rm -rf internal/rest/v1/spec/server-$(name).cfg.yaml)
	rm -rf internal/rest/v1/spec/server-$(name).cfg.yaml
