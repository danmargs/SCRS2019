UTILIZZO DELLO SMART CONTRACT


inserire lo smart contract nel path apposito della rete, in modo tale da farlo installare nei diversi peer (quando viene eseguito script.sh), 
dopo di che entrare nel container CLI ed impersonificare uno dei peer della rete (comandi descritti nel pdf del progetto, ma cambiano a seconda della rete che si utilizza) per poter eseguire i comandi successivi e quindi istanziare, interrogare e fare transazioni con lo smart contract.


### per istanziare il chaincode in un peer

peer chaincode instantiate −o orderer . example .com:7050 −C mychannel −n mycc −v 1.0 −c ’{”Args”:[”init”,”renzo”,”sposato
”,”lucia”,”sposata”,”100”]} ’ −P ”OR (’Org1MSP.member’ , ’Org2MSP. member ’ ) ”



### per aggiungere sul conto in comune un valore di 110

peer chaincode invoke -C mychannel -n mycc -c '{"Args":["aggiungi", "renzo", "lucia", "110"]}'



### per aggiungere sul conto del marito un valore di 100

peer chaincode invoke -C mychannel -n mycc -c '{"Args":["aggiungiMarito", "renzo", "100"]}'



### per aggiungere sul conto del moglie un valore di 100

peer chaincode invoke -C mychannel -n mycc -c '{"Args":["aggiungiMoglie", "lucia", "100"]}'



### se A vuole divorziare da B

peer chaincode query -C mychannel -n mycc -c '{"Args":["divorzia", "A", "B"]}'



### se si vogliono sapere le informazioni Moglie 

peer chaincode query -C mychannel -n mycc -c '{"Args":["queryMoglie", "lucia"]}'



### se si vogliono sapere le informazioni Marito 

peer chaincode query -C mychannel -n mycc -c '{"Args":["queryMarito", "renzo"]}'


