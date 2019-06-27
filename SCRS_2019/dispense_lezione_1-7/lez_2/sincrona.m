canale C1(int);
canale C2(int);

Processo P1 {
   int n;
   ...
   synchSend C2(n);
   receive C1(n);
   ... 
}

Processo P2 {
   int m;
   ...
   receive C2(m);
   synchSend C1(m+1);
   ...
}