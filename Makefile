.PHONY: setup-build
setup-build:
	mkdir build
	cp src/config.json build/


.PHONY: dkbuild
dkbuild:
	@sudo docker-compose -f docker-compose.build.yaml up

.PHONY: dkrun
dkrun:
	@sudo docker-compose up -d

.PHONY: dkstop
dkstop:
	@sudo docker-compose stop

.PHONY: dklogs
dklogs:
	@sudo docker logs oraculo --tail 300
