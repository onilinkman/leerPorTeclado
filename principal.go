package leerPorTeclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//LeerCadena Ingresa por teclado y te devuelve una cadena sin contar con el espacio
func LeerCadena() (cad string, err error) {
	lector := bufio.NewReader(os.Stdin)
	salida, erro := lector.ReadString('\n')
	if erro != nil {
		fmt.Println(erro)
		return "", erro
	}
	return salida, nil
}

//LeerEntero Te devuelve un error y un entero
func LeerEntero() (ent int, err error) {
	lector := bufio.NewReader(os.Stdin)
	salida, erro := lector.ReadString('\n')
	if erro != nil {
		return 0, erro
	}
	sal := salida[:len(salida)-2]
	entero, err1 := strconv.Atoi(sal)
	if err1 != nil {
		fmt.Println(err1)
		return entero, err1
	}

	return entero, nil
}
