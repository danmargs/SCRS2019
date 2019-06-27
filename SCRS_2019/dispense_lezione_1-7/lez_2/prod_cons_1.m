port A;

Produttore() {					
  while(1) {
    /* produce d */
    send(A,d); // invia d su A
  }
}
 
Consumatore() {
  while(1) {
    receive(A,&d); // riceve d da A
    /* consuma d */
  }
}