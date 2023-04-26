package a

/**
 * adding _ within import block to supress import but not used error
 */
import (
	"fmt"
	_ "github.com/tjcchen/go-practice/init/b"
)

func init() {
	fmt.Println("init from a")
}