Installation Fabric multi host Docker Swarm
===========================================

 

Topologie
---------

 

-   PC 1

    -   Orderer (orderer.exemple.com) port 7050

-   PC 2

    -   CA1 (ca.org1.example.com) -port 7054

    -   Org1 Peer0 (peer0.org1.example.com) -port 7051,8051

    -   CouchDB Peer0

    -   Org1 Peer1 (peer1.org1.example.com) -port 7053,8053

    -   CouchDB Peer1

-   PC 3

    -   CA2 (ca.org2.example.com) -port 7054

    -   Org2 Peer0 (peer0.org2.example.com) -port 7051,8051

    -   CouchDB Peer0

    -   Org2 Peer1(peer1.org2.example.com) -port 7053,8053

    -   CouchDB Peer1

     

Pré-requis
----------

 

Pour pouvoir réaliser l’installation suivante, il est nécessaire d’avoir sur
chacun des hôtes les composants suivant :

-   Fabric — 1.4.0

-   Go

-   Docker - 18.09

-   Docker-Compose — 1.23.2

-   Port nécessaire :

    -   TCP port 2377

    -   TCP et UDP port 7946

    -   UDP port 4789

 

Setup
-----

 

1.  Cloner le Git Fabric-samples dans \$GOPATH/src/github.com/hyperledger

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    git clone https://github.com/hyperledger/fabric-samples.git
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

2.  Allez dans le dossier fabric-samples

3.  Cloner le Git d’installation

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    git clone https://github.com/benoitclerget/Multiple-Network.git
    cd Multiple-Network
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

4.  Une fois dans ce dossiers, générer les fichiers de crypto avec la commandes
    suivante

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    ./bymn.sh generate crypto-config.yaml
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

5.  Modifier les variables d’environnement contenue dans bymn.sh.

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    export ORDERER_HOSTNAME=<host name of PC-1>
    export ORG1_PEER0_HOSTNAME= <host name of PC-2>
    export ORG1_PEER1_HOSTNAME= <host name of PC-2>
    export ORG2_PEER0_HOSTNAME=<host name of PC-3>
    export ORG2_PEER1_HOSTNAME=<host name of PC-3>
    export SWARM_NETWORK=”fabric” 
    export DOCKER_STACK=”fabric”
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

    Une fois les modification appliquées, copier le dossier multiple-network sur
    tous les PC

6.  Créer et configurer Docker Swarm. Pour cela, sur le PC1 :

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    docker swarm init --advertise-addr <PC-1 IP address>
    docker swarm join-token manager
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

    Après ces commandes, copier/coller la sortie de la dernière commande sur les
    PC 2 et 3. La sortie ressemble à la commande suivante :

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    docker swarm join --token SWMTKN-1–3anjn4oxwcn278hie3413zaakr4higjdqr2x89r5605p1dosui-a4u407pt6c5ta2ont7pqdnm 137.116.147.36:2377 --advertise-addr <PC-2 Ip Address>

    docker swarm join --token SWMTKN-1–3anjn4oxwcn278hie3413zaakr4higjdqr2x89r5605p1dosui-a4u407pt6c5ta2ont7pqdnm 137.116.147.36:2377 --advertise-addr <PC-3 Ip Address>
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

7.  Créer un réseau overlay

    Sur le PC1 :

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    docker network create --attachable --driver overlay fabric
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

8.  Démarrer l'Orderer

    Sur le PC1 :

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    ./bymn.sh up -f docker-compose-orderer.yaml
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

9.  Démarrer l’organisation 1 :

    Sur le PC2 :

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    ./bymn.sh up -f docker-compose-org1.yaml
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

10. Démarrer l’organisation 2 :

    Sur le PC3 :

    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    ./bymn.sh up -f docker-compose-org2.yaml
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

 

Test du réseau
--------------

 

Pour tester le réseaux, il est possible d’utiliser le script de test présent
dans les conteneur CLI. Pour ce faire, se connecter à un des docker CLI (fabric tool) présent
sur le PC2 ou PC3.

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
docker exec -it <IdDockerCLI> bash
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Une fois connecté lancé le script :

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
./script/script.sh
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Si tout est correct, vous devriez avoir :

![](https://cdn-images-1.medium.com/max/1600/1*TTgzN9CB5Spfkye8yEDdNg.png)
