package keymanager

import (
	"bufio"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"os"

	crypto "github.com/libp2p/go-libp2p-crypto"
)

// estructura de una clave
type Claves struct {
	Private        crypto.PrivKey
	Private_bytes  []byte
	Private_string string
	Public         crypto.PubKey
	Public_bytes   []byte
	Public_string  string
}

// devuelve una estructura a partir de las claves
func NuevaClave(private crypto.PrivKey) Claves {
	public := private.GetPublic()
	priv_bytes, _ := crypto.MarshalPrivateKey(private)
	pub_bytes, _ := crypto.MarshalPublicKey(public)

	return Claves{
		Private:        private,
		Private_bytes:  priv_bytes,
		Private_string: crypto.ConfigEncodeKey(priv_bytes),
		Public:         public,
		Public_bytes:   pub_bytes,
		Public_string:  crypto.ConfigEncodeKey(pub_bytes),
	}
}

// crea una clave privada
func Newkey() Claves {
	// generamos nueva key
	private, _, err := crypto.GenerateRSAKeyPair(2048, rand.Reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(" [i] Nueva clave privada creada.")
	return NuevaClave(private)
}

// TODO: permitir personalizar el nombre del fichero al exportar
func (s *Claves) Export() {
	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: s.Private_bytes,
	}
	// preparamos fichero
	pemPrivateFile, err := os.Create("private_key.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// guardamos datos en fichero
	err = pem.Encode(pemPrivateFile, pemPrivateBlock)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemPrivateFile.Close()
	fmt.Println(" [i] Clave privada exportada al fichero 'private_key.pem'")
}

// TODO: permitir personalizar el nombre del fichero al importar
// importa clave privada desde un fichero
func Import_key() Claves {
	// cargamos fichero
	privateKeyFile, err := os.Open("private_key.pem")
	if err != nil {
		fmt.Println(" [!] No hay ninguna clave creada, usa la opcion -newkey para generar una.")
		os.Exit(1)
	}
	// cargamos bytes del fichero
	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	// decodificamos pem y lo convertimos a privkey
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()
	private, _ := crypto.UnmarshalPrivateKey(data.Bytes)
	fmt.Println(" [i] Clave privada importada del fichero 'private_key.pem'")
	// devolvemos la estructura de esta clave
	return NuevaClave(private)
}
