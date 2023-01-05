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

package mobilegateway

import (
	"context"

	"github.com/sacloud/iaas-api-go"
)

func (s *Service) AddSIM(req *AddSIMRequest) error {
	return s.AddSIMWithContext(context.Background(), req)
}

func (s *Service) AddSIMWithContext(ctx context.Context, req *AddSIMRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	mgwOp := iaas.NewMobileGatewayOp(s.caller)
	simOp := iaas.NewSIMOp(s.caller)

	if err := mgwOp.AddSIM(ctx, req.Zone, req.ID, &iaas.MobileGatewayAddSIMRequest{SIMID: req.SIMID.String()}); err != nil {
		return err
	}
	return simOp.AssignIP(ctx, req.SIMID, &iaas.SIMAssignIPRequest{IP: req.IPAddress})
}
