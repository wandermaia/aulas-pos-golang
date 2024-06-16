# Anotçãoes Aulas

Inicializando o módulo

```bash

go mod init github.com/wandermaia/aulas-pos-golang/tree/main/19-deploy-k8s

```

Para remover todos os containers

```bash

docker rm -f $(docker ps -a -q)

```

DWARF - Debugging with arbitrary record format. Flags de debugging.
Normalamente são removidos em produção.


Gerando a imagem do container e executando (comandos aula)

```bash

docker build -t wandermaia/19-deploy-k8s:latest -f Dockerfile.prod .

docker images | grep 19-deploy-k8s

docker run --rm -p 8080:8080 wandermaia/19-deploy-k8s:latest

```

Multi-stage builds

O go importa bibliotecas do C (CGO). Isso faz que o programa tenha dependências do C e venha com a função GCO habilitada por padrão.

```bash

docker push wandermaia/19-deploy-k8s:latest

```

## Kubernetes (Kind)

Criar o cluster 

```bash

kind create cluster --name=goexpert

kubectl cluster-info --context kind-goexpert


```
Code sniper no VScode para deployments.



## Referências

Multi-stage builds

https://docs.docker.com/build/building/multi-stage/

Kind

https://kind.sigs.k8s.io/

Install and Set Up kubectl on Linux

https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/