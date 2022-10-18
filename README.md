# Teste Pessoa DevOps Amo Promo V1.0

# ![Amopromo](https://www.amopromo.com/images/logo-amo-promo.svg)

#### Nossas Marcas
|   Passagens Promo   |  Seguros Promo  |    Parceiros Promo    |
| :---:         |     :---:      |          :---: |
| ![PassagensPromo](https://www.amopromo.com/images/logo-passagens-promo.svg)   | ![SegurosPromo](https://www.amopromo.com/images/logo-seguros-promo.svg) |  ![ParceirosPromo](https://www.amopromo.com/images/logo-parceiros-promo.svg) |

### O que queremos:
Estamos interessados em entender sua lógica de programação e resolução objetiva de problemas, portanto não faz parte do desafio a criação de artefatos de planejamento/gestão de projetos, como: cronogramas, diagramas, burndown, quadro de demandas e etc. Nós amamos testes, apesar de não ser obrigatório, seria um grande diferencial qualquer tipo de teste (unitários, integração e etc).

Sua capacidade de interpretação também está sendo avaliada, mas lembre-se que não existe uma solução única para o teste! Então, siga sua interpretação e use a criatividade para completar as tarefas especificadas. Havendo tempo e disponibilidade, incrementar com features/detalhes que julgar interessantes. Boa sorte!


## Introdução
Na Amo Promo, trabalhamos com diversos produtos, cada um com suas particularidades, backlog próprio e processos únicos.

### O que queremos:
1. Infraestrutura como Código (Terraform, Ansible, CloudFormation, Chef, Puppet)
2. Pipeline CI/CD (Github Actions, CircleCI, TravisCI, GitlabCI, Jenkins, Bamboo)
3. Docker/Container
4. Monitoramento (CloudWatch, Graylog, ElasticSearch, Kibana)
5. Cloud (GCP, AWS, Azure, Oracle)
6. Git (Github, Gitlab, BitBucket)

### O desafio:
Seu desafio é criar um pipeline para implementação da aplicação de forma automatizada e disponibilizar acesso via web.

> Recomendamos a criação de uma conta de nível gratuito (free tier) para não incorrer em custos.

> As instruções para rodar a aplicação em Go estão descritas em um README no diretório `app`.

Estamos testando sua capacidade de implementar infraestrutura automatizada moderna, bem como conhecimentos gerais sobre administração de sistemas. Na sua solução, você deve enfatizar a legibilidade, a manutenção e as metodologias de DevOps.

0. Criar um fork público deste repositório para as implementações a seguir.
1. Criar um Dockerfile para rodar a aplicação em Go disponibilizada no diretório `app`.
   - Atentar-se às variáveis de ambiente
2. Criar um Workflow (Github Actions) com os passos abaixo(Build/Test/Deploy).
   - Build: Checkout do código e build da aplicação
   - Test: Testes bem sucedidos
   - Deploy: Utilizar de recursos disponíveis no Github Actions ou Criar um script para implementar o deploy de forma automatizada. O Deploy deve ser feito no local de sua escolha. Recomendamos AWS (EC2, ECS, Fargate, Beanstalk, etc) ou Google (Compute Engine, Kubernetes Engine, App Engine ou Cloud Run).
3. Criar a Infraestrurura como Código (IaaS). Recomendamos Terraform ou Ansible.
4. Disponibilizar uma URL/IP da aplicação
   - Testaremos a rota `/ping`
   - A rota deve retornar o seguinte json `{"message":"pong! Company Name: Amo Promo"}`
5. Criar um script Shell para gerar uma **build local** do app e que permita escolher entre dois tipos de plataformas: `linux/amd64` e `darwin/arm64`
   - A plataforma/arquitetura pode ser informada via argumento ao chamar o script *Ex: ./gera_build.sh linux amd64* ou da forma que achar mais prática e eficiente.

## Entrega do desafio
1. Faça um fork público deste repositório e complemente com as implementações solicitadas: Dockerfile, Worflow do Github Actions e o arquivo(s) de IaaS.
2. Envie o link do repositório ao responsável pelo seu teste

## Diferenciais
- Trazer alguma solução de Monitoramento
- Utilizar a plataforma Cloud do Google (GCP)
