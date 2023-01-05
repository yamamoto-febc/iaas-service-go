// Copyright 2022-2023 The sacloud/iaas-service-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package archive

import (
	"context"
	"fmt"

	"github.com/sacloud/iaas-api-go"
)

func (s *Service) Update(req *UpdateRequest) (*iaas.Archive, error) {
	return s.UpdateWithContext(context.Background(), req)
}

func (s *Service) UpdateWithContext(ctx context.Context, req *UpdateRequest) (*iaas.Archive, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := iaas.NewArchiveOp(s.caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, fmt.Errorf("reading Archive[%s] failed: %s", req.ID, err)
	}

	params, err := req.ToRequestParameter(current)
	if err != nil {
		return nil, fmt.Errorf("processing request parameter failed: %s", err)
	}

	return client.Update(ctx, req.Zone, req.ID, params)
}
