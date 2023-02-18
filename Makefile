generate-server:
	rm -rf $(genPath)/oapi/$(name)dto && rm -rf $(genPath)/oapi/$(name)srv
	mkdir $(genPath)/oapi/$(name)dto && mkdir $(genPath)/oapi/$(name)srv
	sed 's/#name/$(name)dto/g' $(specPath)/server.cfg.yaml >> $(specPath)/server-$(name).cfg.yaml
	(oapi-codegen -generate types -o $(genPath)/oapi/$(name)dto/$(name)dto.go -package $(name)dto $(specPath)/$(name).yaml || rm -rf $(specPath)/server-$(name).cfg.yaml ) && (oapi-codegen --config $(specPath)/server-$(name).cfg.yaml -package $(name)srv -o $(genPath)/oapi/$(name)srv/$(name)srv.go $(specPath)/$(name).yaml || rm -rf $(specPath)/server-$(name).cfg.yaml)
	rm -rf $(specPath)/server-$(name).cfg.yaml
