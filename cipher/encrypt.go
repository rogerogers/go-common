package cipher

type Crypt interface {
	encrypt() string
	decrypt() string
}
