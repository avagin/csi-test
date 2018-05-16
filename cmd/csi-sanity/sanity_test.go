/*
Copyright 2017 Luis Pabón luis@portworx.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sanity

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"encoding/json"

	"github.com/kubernetes-csi/csi-test/pkg/sanity"
)

const (
	prefix string = "csi."
)

var (
	VERSION = "(dev)"
	version bool
	config  sanity.Config
)

func init() {
	flag.StringVar(&config.Address, prefix+"endpoint", "", "CSI endpoint")
	flag.BoolVar(&version, prefix+"version", false, "Version of this program")
	flag.StringVar(&config.TargetPath, prefix+"mountdir", os.TempDir()+"/csi", "Mount point for NodePublish")
	flag.StringVar(&config.StagingPath, prefix+"stagingdir", os.TempDir()+"/csi", "Mount point for NodeStage if staging is supported")
	flag.StringVar(&config.SecretPath, prefix+"secretfile", "", "Path to a file with secrets")
	flag.Parse()

}

func TestSanity(t *testing.T) {
	if config.SecretPath != "" {
		data, err := ioutil.ReadFile(config.SecretPath)
		if err != nil {
			t.Fatalf("Unable to open %s: %s", config.SecretPath, err.Error());
		}
		err = json.Unmarshal(data, &config.Secrets)
		if err != nil {
			t.Fatalf("Unable to parse %s: %s", config.SecretPath, err.Error());
		}
	}
	if version {
		fmt.Printf("Version = %s\n", VERSION)
		return
	}
	if len(config.Address) == 0 {
		t.Fatalf("--%sendpoint must be provided with an CSI endpoint", prefix)
	}
	sanity.Test(t, &config)
}
