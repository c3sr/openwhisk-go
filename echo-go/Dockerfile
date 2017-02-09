# Dockerfile for example whisk docker action
FROM c3sr/openwhisk-go:go1.7.5
 
ADD example.go /action/example.go

RUN cd /action && go build && mv action exec

CMD ["/bin/bash", "-c", "cd /actionProxy && python -u actionproxy.py"]