// +build linux darwin

/*
Copyright 2022 CodeNotary, Inc. All rights reserved.

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

package config

// ConfigImmudb ...
var ConfigImmudb = []byte(`dir = "/var/lib/immudb"
network = "tcp"
address = "0.0.0.0"
port = 3322
dbname = "data"
pidfile = "/var/lib/immudb/immudb.pid"
logfile = "/var/log/immudb/immudb.log"
mtls = false
detached = false
auth = true
pkey = ""
certificate = ""
clientcas = ""
devmode = true
admin-password = "immudb"`)
