package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

func All_Update(c *cli.Context) {
	update, err := exec.Command("apt", "update").Output()
	upgrade, err := exec.Command("apt", "upgrade").Output()
	autoremove, err := exec.Command("apt", "autoremove").Output()
	systemtool, err := exec.Command("apt", "install", "wget").Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(update))
	fmt.Println(string(upgrade))
	fmt.Println(string(autoremove))
	fmt.Println(string(systemtool))
	fmt.Println("📦, update and package install")
}

func Setup_apa(c *cli.Context) {
	// Required for PHP, Apache setup
	libxml2, err := exec.Command("apt", "-y", "install", "libxml2-dev", "wget", "pcre", "make", "pcre-devel", "libxml2-dev", "libxml2").Output()
	Set_atr()
	Set_apr_util()
	Set_apache()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(libxml2))

	fmt.Println("📦, install apache")
}

func Set_atr() {
	//APR
	wget_apr, err := exec.Command("wget", "https://www-eu.apache.org/dist//apr/apr-1.7.0.tar.gz").Output()
	tar_apr, err := exec.Command("tar", "zxvf", "apr-1.7.0.tar.gz").Output()
	os.Chdir("apr-1.7.0")
	configure_apr, err := exec.Command("./configure").Output()
	make_apr, err := exec.Command("make").Output()
	makein_apr, err := exec.Command("make", "install").Output()
	fmt.Println(string(wget_apr))
	fmt.Println(string(tar_apr))
	fmt.Println(string(configure_apr))
	fmt.Println(string(make_apr))
	fmt.Println(string(makein_apr))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func Set_apr_util() {
	//APR-util
	wget_apr_util, err := exec.Command("wget", "https://www-eu.apache.org/dist//apr/apr-util-1.6.1.tar.gz").Output()
	tar_apr_util, err := exec.Command("tar", "zxvf", "apr-util-1.6.1.tar.gz").Output()
	os.Chdir("apr-util-1.6.1")
	configure_apr_util, err := exec.Command("./configure").Output()
	make_apr_util, err := exec.Command("make").Output()
	makein_apr_util, err := exec.Command("make", "install").Output()
	fmt.Println(string(wget_apr_util))
	fmt.Println(string(tar_apr_util))
	fmt.Println(string(configure_apr_util))
	fmt.Println(string(make_apr_util))
	fmt.Println(string(makein_apr_util))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func Set_apache() {
	// Apache
	wget_apache, err := exec.Command("wget", "https://www-us.apache.org/dist//httpd/httpd-2.4.39.tar.gz").Output()
	tar_apache, err := exec.Command("tar", "xvzf", "httpd-2.4.39.tar.gz").Output()
	os.Chdir("httpd-2.4.39")
	configure_apache, err := exec.Command("./configure").Output()
	make_apache, err := exec.Command("make").Output()
	makein_apache, err := exec.Command("make", "make install").Output()
	fmt.Println(string(wget_apache))
	fmt.Println(string(tar_apache))
	fmt.Println(string(configure_apache))
	fmt.Println(string(make_apache))
	fmt.Println(string(makein_apache))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func Setup_php(c *cli.Context) {
	wget_php, err := exec.Command("wget", "https://www.php.net/distributions/php-7.3.8.tar.gz").Output()
	tar_php, err := exec.Command("tar", "zxvf", "php-7.3.8.tar.gz").Output()
	os.Chdir("php-7.3.8")
	configure_php, err := exec.Command("./configure").Output()
	make_php, err := exec.Command("make").Output()
	makein_php, err := exec.Command("make", "install").Output()
	fmt.Println(string(wget_php))
	fmt.Println(string(tar_php))
	fmt.Println(string(configure_php))
	fmt.Println(string(make_php))
	fmt.Println(string(makein_php))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
