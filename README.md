# Go_world_server

Go_world_server est une application web écrite en Go (aussi appelé Golang), langage créé par Google en 2009.

Cette application web renvoie l'heure et la date courante pour toutes les timezones passées par l'URL.

## Sommaire

1. [Prérequis](#prerequis)

2. [Installer Go_world_server à l'aide du code source sur une machine quelconque](#installation_sur_machine)

3. [Installer Go_world_server avec Docker](#installation_avec_docker)

	* [Installer Go_world_server à l'aide du fichier Dockerfile](#installation_avec_dockerfile)

	* [Installer Go_world_server à l'aide de l'image docker pullé](#installation_avec_pull)

4. [Installer Go_world_server avec Balena](#installation_avec_balena)

5. [Comment obtenir l'heure et la date courante d'une ou de plusieurs timezones ?](#comment_utiliser)

6. [Comment exécuter cette application Web uniquement en local ?](#executer_uniquement_en_local)

	* [Pourquoi exécuter cette application Web uniquement localement ?](#pourquoi_localement_uniquement)

	* [Comment faire à partir du code source ?](#localement_avec_code_source)

	* [Comment faire avec Docker ?](#localement_avec_docker)

7. [Comment obtenir la documentation du projet ?](#comment_documentation)

8. [Conclusion](#conclusion)

<a name="prerequis"></a>
## Prérequis

Pour installer Go_world_server, les prérequis suivants doivent être satisfaits pour la machine:

* Linux dans sa version la plus récente,

* Go dans sa version la plus récente.

<a name="installation_sur_machine"></a>
## Installer Go_world_server à l'aide du code source sur une machine quelconque

Pour installer Go_world_server à l'aide du code source, vérifiez bien que les prérequis listés ci-dessus sont remplis.

Ensuite, clonez le projet Go_world_server à l'aide de cette commande:

```bash
git clone https://gitlab.imerir.com/eric.ghoubiguian/go_world_server
```

Ensuite, placez-vous dans le répertoire créé par la commande précédente à l'aide ce celle-ci:

```bash
cd go_world_server
```

Exécutez après cette commande:

```bash
export GOPATH=${PWD}
```

__petite précision__: Cette commande permet de modifier la valeur de la variable d'environnement `GOPATH` qui spécifie la source du projet go.

Félicitations une fois cela fait, vous pouvez maintenant lancer le projet à l'aide de la commande suivante:

```bash
go run main.go
```
ça y est, le serveur est fonctionnel. Connectez-vous à l'aide du lien URL <a href="http://localhost/">suivant</a>.

<a name="installation_avec_docker"></a>
## Installer Go_world_server avec Docker

Il existe 2 procédés pour installer Go_world_server avec Docker: **à l'aide du fichier Dockerfile** et **à l'aide de l'image docker pullé**.

<a name="installation_avec_dockerfile"></a>
#### Installer Go_world_server à l'aide du fichier Dockerfile

Un fichier nommé `Dockerfile` existe à la racine du projet.

Ce fichier permet à Docker de créer l'image de l'application web et de faire fonctionner ce dernier dans un container Docker.

Tout d'abord, placez-vous dans le projet à l'aide de la commande suivante:

```bash
cd go_world_server
```
Le succès de cette commande dépend du fait que le projet Go_world_server existe sur le serveur d'installation. Si ce n'est pas le cas, tout d'abord clonez-le à l'aide de la commande suivante:

```bash
git clone https://gitlab.imerir.com/eric.ghoubiguian/go_world_server
```
Ensuite, créez l'image Docker de Go_world_server à l'aide de la commande suivante:

```bash
docker image build -t go_world_server:latest .
```

__petite précision__: le caractère `.` précisé à la fin de la commande se réfère au répertoire courant.

Une fois cette commande terminée, l'image `go_world_server` est créée avec `go_world_server` comme nom et `latest` comme étiquette.

Pour vérifier, exécutez la commande suivante pour vérifier que l'image `go_world_server` a bien été créée avec succès:

```bash
docker image ls
```
Vous pouvez de plus consulter la liste de toutes les instructions qui ont été exécutées pour créer l'image Docker de Go_world_server:

```bash
docker image history go_world_server:latest
```
Il est maintenant temps de créer le container Docker et de le faire marcher, pour cela exécutez la commande suivante:

```bash
docker container run -d --name go_world_server -p 80:80 go_world_server:latest
```
__petite précision__: 

Toutes mes félicitations, le container Docker de Go_world_server est créé et marche maintenant.
Pour s'en rendre compte, listez tous les containers Docker à l'aide de la commande suivante:

```bash
docker image ls
```

ou de celle-ci

```bash
docker image ls -a
```
Maintenant qu'il est installé, l'application web fonctionne maintenant.

Pour s'en rendre compte, il vous suffit d'ouvrir votre navigateur web favori puis d'entrer l'URL suivante:

```bash
<adresse_ip_du_container>:80
```
Pour connaître l'adresse IP du container Docker de l'application web Go_world_server, faites la commande suivante:

```bash
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' go_world_server
```
L'adresse IP du container s'affiche ensuite en dessous.

Pour consulter le fichier des logs dans le terminal, tapez la commande suivante:

```bash
docker container exec -it go_world_server cat log.txt
```

<a name="installation_avec_pull"></a>
### Installer Go_world_server à l'aide de l'image docker pullé

Il existe une image Docker de Go world server stockée sur Docker Hub [ici](https://hub.docker.com/r/wicken/go_world_server).

__petite précision__: Docker Hub est un service d'hébergement d'images Docker sur le même principe que GitHub pour les projets.

Obtenez l'image Docker de l'application web Go world server à l'aide de cette commande:

```bash
docker pull wicken/go_world_server
```
Une fois cette commande terminée, l'image `wicken/go_world_server` est importée avec `wicken/go_world_server` comme nom et `latest` comme étiquette.

Pour vérifier, exécutez la commande suivante pour vérifier que l'image `wicken/go_world_server` a bien été créée avec succès:

```bash
docker image ls
```
Il est maintenant temps de créer le container Docker et de le faire marcher, pour cela exécutez la commande suivante:

```bash
docker container run -d --name go_world_server -p 80:80 wicken/go_world_server:latest
```
__petite précision__: 

Toutes mes félicitations, le container Docker de Go_world_server est créé et marche maintenant.
Pour s'en rendre compte, listez tous les containers Docker à l'aide de la commande suivante:

```bash
docker image ls
```

ou de celle-ci

```bash
docker image ls -a
```
Maintenant qu'il est installé, l'application web fonctionne maintenant.

Pour s'en rendre compte, il vous suffit d'ouvrir votre navigateur web favori puis d'entrer l'URL suivante:

```bash
<adresse_ip_du_container>:80
```
Pour connaître l'adresse IP du container Docker de l'application web Go_world_server, faites la commande suivante:

```bash
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' go_world_server
```
L'adresse IP du container s'affiche ensuite en dessous.

Pour consulter le fichier des logs dans le terminal, tapez la commande suivante:

```bash
docker container exec -it go_world_server cat log.txt
```

<a name="installation_avec_balena"></a>
## Installer Go_world_server avec Balena

<a name="comment_utiliser"></a>
## Comment obtenir l'heure et la date courante d'une ou de plusieurs timezones ?

Pour obtenir l'heure et la date courantes pour une ou plusieurs timezones, il faut les renseigner dans l'URL de l'application et ensuite valider tout en remplaçant les `/` composants la timezone par `&`.

__Par exemple__: Partons du principe que vous voulez obtenir l'heure et la date courantes de Paris à France et à Shanghai en République populaire de Chine.

Les timezones correspondantes sont `Europe/Paris` et `Asia/Shanghai`, donc celles-ci devront être obligatoirement renseignées dans l'URL sous la forme suivante: `Europe&Paris` et `Asia&Shanghai`.

En effet:

```bash
http://<adresse_ip_utilisée>:80/Europe&Paris/Asia&Shanghai
```
affichera d'abord l'heure et la date courantes à Paris et ensuite à Shanghai, alors que

```bash
http://<adresse_ip_utilisée>:80/Asia&Shanghai/Europe&Paris
```
affichera d'abord l'heure et la date courantes à Shanghai et ensuite à Paris.

<a name="executer_uniquement_en_local"></a>
## Comment exécuter cette application Web uniquement en local ?

<a name="pourquoi_localement_uniquement"></a>
### Pourquoi exécuter cette application Web uniquement localement ?

<a name="localement_avec_code_source"></a>
### Comment faire à partir du code source ?

<a name="localement_avec_docker"></a>
### Comment faire avec Docker ?

<a name="comment_documentation"></a>
## Comment obtenir la documentation du projet ?

Après avoir installé le projet Go_world_server comme indiqué [ici](#installation_sur_machine), vérifiez bien que vous êtes bien dans le répertoire du projet.

Si ce n'est pas le cas, tapez la commande suivante:

```bash
cd $GOPATH
```

Ensuite pour consulter la documentation du projet exécutez la commande suivante:

```bash
godoc -http=:<port_choisi>
```

La balise `<port_choisi>` correspond au numéro de port que vous voulez.

Une fois la commande lancée, ouvrez une page dans votre navigateur web et dirigez vous à l'adresse suivante `http://localhost:<port_choisi>`.

<a name="conclusion"></a>
## Conclusion
