/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// CTLogApplyConfiguration represents an declarative configuration of the CTLog type for use
// with apply.
type CTLogApplyConfiguration struct {
	IgnoreSCT    *bool   `json:"ignoreSCT,omitempty"`
	CTLogPubKey  *string `json:"pubkey,omitempty"`
	TSACertChain *string `json:"tsaCertChain,omitempty"`
}

// CTLogApplyConfiguration constructs an declarative configuration of the CTLog type for use with
// apply.
func CTLog() *CTLogApplyConfiguration {
	return &CTLogApplyConfiguration{}
}

// WithIgnoreSCT sets the IgnoreSCT field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IgnoreSCT field is set to the value of the last call.
func (b *CTLogApplyConfiguration) WithIgnoreSCT(value bool) *CTLogApplyConfiguration {
	b.IgnoreSCT = &value
	return b
}

// WithCTLogPubKey sets the CTLogPubKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CTLogPubKey field is set to the value of the last call.
func (b *CTLogApplyConfiguration) WithCTLogPubKey(value string) *CTLogApplyConfiguration {
	b.CTLogPubKey = &value
	return b
}

// WithTSACertChain sets the TSACertChain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TSACertChain field is set to the value of the last call.
func (b *CTLogApplyConfiguration) WithTSACertChain(value string) *CTLogApplyConfiguration {
	b.TSACertChain = &value
	return b
}
