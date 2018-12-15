# go-p2p-keymanager
Libreria para crear, exportar e importar las claves privadas usadas con 'go-libp2p'

# install
go get -u github.com/katmai1/go-p2p-keymanager/...

# Usage
    
  import "github.com/katmai1/go-p2p-keymanager/keymanager"

- To generate a new pairs:
    clave := keymanager.Newkey()

- Export private key to pem file:
    clave.Export()

- Import private key from pem file:
    clave := keymanager.Import()