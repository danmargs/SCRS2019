# SCRS2019
progetto scrs 2019


# Idea Progetto

Il matrimonio prevede degli accordi che riguardano anche i rapporti patrimoniali tra i coniugi. Dal 1975, il regime patrimoniale stabilito dalla legge, durante la celebrazione del matrimonio, `e quello della comunione dei beni. L’idea di questo progetto `e quella di creare uno Smart Contract da eseguire sulla Blockchain, che sia in grado di gestire la comunione dei beni di due coniugi senza l’uso di un intermediario umano come un giudice o un avvocato.


# Premesse per la creazione ed utilizzo dello Smart Contract

La realizzazione dello smart contract di basa su queste premesse:
  • Una coppia sposata con la comunione dei beni deve utilizzare un solo conto in comune;
  • Se una persona chiede il divorzio non deve poter utilizzare (aggiungere o togliere) i soldi sul conto in comune, 
    ma in quel caso pu`o utilizzare il conto separato privato (vale anche per il coniuge che magari non vuole ancora 
    divorziare).
  • Se tutti e due chiedono il divorzio allora i beni presenti nel conto in comune dovranno essere divisi a met`a in ogni 
    caso.


# Passo 1
PER POTER ESEGUIRE IL PROGETTO COME STRUTTURATO NEL PDF È NECESSARIO AVERE 2 MACCHINE VIRTUALI CONNESSE SOTTO LA STESSA RETE CON DOCKER INSTALLATO.

COPIARE LA CARTELLA Build-Multi-Host-Network-Hyperledger E SEGUIRE LE ISTRUZIONI PRESENTI NEL "CAPITOLO 5" DEL PDF PER POTER SETTARE LA RETE E FAR COMUNICARE LE DUE MACCHINE TRAMITE DOCKER.

OPPURE SE NON SI VUOLE PROVARE LO SMART CONTRACT SULLA RETE DEL PROGETTO, MA SU UNA RETE COME LA FIRST NETWORK, BASTA PRENDERE LO SMART CONTRACT PRESENTE NELLA DIRECTORY "⁨Build-Multi-Host-Network-Hyperledger⁩/chaincode⁩/⁨chaincode_example02⁩" E SOSTITUIRLO A QUELLA PRESENTE IN "fabric-samples/chaincode/chaincode_example02/go"


# Passo 2
# UTILIZZO DELLO SMART CONTRACT


inserire lo smart contract nel path apposito della rete, in modo tale da farlo installare nei diversi peer (quando viene eseguito script.sh), 
dopo di che entrare nel container CLI ed impersonificare uno dei peer della rete (comandi descritti nel pdf del progetto, ma cambiano a seconda della rete che si utilizza) per poter eseguire i comandi successivi e quindi istanziare, interrogare e fare transazioni con lo smart contract.


# per istanziare il chaincode in un peer

peer chaincode instantiate −o orderer . example .com:7050 −C mychannel −n mycc −v 1.0 −c ’{”Args”:[”init”,”renzo”,”sposato
”,”lucia”,”sposata”,”100”]} ’ −P ”OR (’Org1MSP.member’ , ’Org2MSP. member ’ ) ”



# per aggiungere sul conto in comune un valore di 110

peer chaincode invoke -C mychannel -n mycc -c '{"Args":["aggiungi", "renzo", "lucia", "110"]}'



# per aggiungere sul conto del marito un valore di 100

peer chaincode invoke -C mychannel -n mycc -c '{"Args":["aggiungiMarito", "renzo", "100"]}'



# per aggiungere sul conto del moglie un valore di 100

peer chaincode invoke -C mychannel -n mycc -c '{"Args":["aggiungiMoglie", "lucia", "100"]}'



# se A vuole divorziare da B

peer chaincode query -C mychannel -n mycc -c '{"Args":["divorzia", "A", "B"]}'



# se si vogliono sapere le informazioni Moglie 

peer chaincode query -C mychannel -n mycc -c '{"Args":["queryMoglie", "lucia"]}'



# se si vogliono sapere le informazioni Marito 

peer chaincode query -C mychannel -n mycc -c '{"Args":["queryMarito", "renzo"]}'


