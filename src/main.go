// Copyright © 2019 Antoine Chiny <antoine.chiny@inria.fr>
// Copyright © 2019 Ryan Ciehanski <ryan@ciehanski.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"github.com/astaxie/beego"
	ui "github.com/webui-dev/go-webui/v2"

	libgen_cli "github.com/ciehanski/libgen-cli/cmd/libgen-cli"
	"github.com/ciehanski/libgen-cli/libgen"
)

func main() {
	w := ui.NewWindow()
	// Bind a Go function.
	ui.Bind(w, "greet", greet)
	// Show frontend.
	w.Show("index.html")
	// Wait until all windows get closed.
	ui.Wait()
	client := http.Client{Timeout: libgen.HTTPClientTimeout, Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	_, err := client.Get("http://clients3.google.com/generate_204")
	if err != nil {
		fmt.Println("\nYou need an internet connection to run libgen-cli.")
		os.Exit(1)
	}

	if err := libgen_cli.Execute(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

func greet(e ui.Event) string {
	name, _ := ui.GetArg[string](e)
	fmt.Printf("%s has reached the backend!\n", name)
	jsResp := fmt.Sprintf("Hello %s 🐇", name)
	return jsResp
}