.SILENT:

dcrun: buildclient
	docker run -dp 3000:3000 --name clientf client

dbrun: buildbackend
	docker run -dp 8080:8080 --name backend back

buildbackend:
	docker build -t back ./newBackend

buildclient:
	docker build -t client ./client

dstop: dbstop
	docker stop backend

dbstop:
	docker stop clientf

drmi: dbdelete
	docker rmi client

dbdelete:
	docker rmi back

dclear:
	docker system prune -a