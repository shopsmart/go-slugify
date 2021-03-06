package slugify_test

import (
	"fmt"

	"github.com/shopsmart/go-slugify"
)

func ExampleSlugify() {
	s := "Nín hǎo. Wǒ shì zhōng guó rén This is a test, Brad's Deals ---"
	fmt.Println(slugify.Slugify(s))
	// Output: nin-hao-wo-shi-zhong-guo-ren-this-is-a-test-brads-deals
}
