package mtrf

/*
Примеры реальных сообщений:
Mode=RX,Ch=44,Cmd=15,Fmt=1,D0=3
bind start
{173,1,0,20,44,15,1,3,0,0,0,0,0,0,0,1,174}

Mode=RX,Ch=44,Cmd=15
bind end
{173,1,0,21,44,15,0,0,0,0,0,0,0,0,0,254,174}

Датчик света включил свет
{173,1,0,26,44,2,0,0,0,0,0,0,0,0,0,246,174}


{173,1,0,30,41,21,7,207,32,44,255,0,0,0,0,43,174}
[173,1,0,6,42,21,7,215,32,45,255,0,0,0,0,29,174]

*/
