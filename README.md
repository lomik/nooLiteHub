# nooLiteHub
MQTT gateway for MTRF-64 

## Примеры

Везде в примерах используется число 15 в качестве номера канала, с которым производятся действия. В качестве консольного клиента MQTT используется mosquitto_pub

### Cиловой блок без обратной связи

### Силовой блок Noolite-F

**Привязать**
1. Перевести силовой блок в режим привязки
2. Послать сообщение
```
mosquitto_pub -t nooLiteHub/write/tx-f/15/bind -m ""
```

**Включить**
```
mosquitto_pub -t nooLiteHub/write/tx-f/15/on -m ""
```

**Выключить**
```
mosquitto_pub -t nooLiteHub/write/tx-f/15/off -m ""
```

**Переключить**
```
mosquitto_pub -t nooLiteHub/write/tx-f/15/switch -m ""
```

