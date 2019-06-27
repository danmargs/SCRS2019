canale C(int);

Processo P1 {
   int n;
   ...
   send C(n);
   ... 
}

Processo P2 {
   int m;
   ...
   receive C(m);
   ...
}