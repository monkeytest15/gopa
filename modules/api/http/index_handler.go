/*
Copyright 2016 Medcl (m AT medcl.net)

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

package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/medcl/gopa/core/env"
	"net/http"
)

func (this Handler) IndexAction(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.URL.Path != "/" {
		this.error404(w)
		return
	}

	data := map[string]interface{}{}
	data["cluster_name"] = this.Env.SystemConfig.ClusterName
	data["name"] = this.Env.SystemConfig.NodeName
	version := map[string]interface{}{}
	version["number"] = env.VERSION
	version["build_commit"] = env.LastCommitLog
	version["build_date"] = env.BuildDate
	data["version"] = version
	data["tagline"] = "You Know, for Web"

	this.WriteJson(w, &data, http.StatusOK)
}
