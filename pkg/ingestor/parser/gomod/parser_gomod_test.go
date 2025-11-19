//
// Copyright 2024 The GUAC Authors.
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

package gomod

import (
	"context"
	"testing"

	"github.com/guacsec/guac/internal/testing/mockverifier"
	"github.com/guacsec/guac/internal/testing/testdata"
	"github.com/guacsec/guac/pkg/assembler"
	"github.com/guacsec/guac/pkg/handler/processor"
	"github.com/guacsec/guac/pkg/ingestor/parser/common"
	"github.com/guacsec/guac/pkg/ingestor/verifier"
	"github.com/guacsec/guac/pkg/logging"
)

func Test_GoModParser(t *testing.T) {
	ctx := logging.WithLogger(context.Background())
	err := verifier.RegisterVerifier(mockverifier.NewMockSigstoreVerifier(), "sigstore")
	if err != nil {
		t.Errorf("verifier.RegisterVerifier() failed with error: %v", err)
	}
	tests := []struct {
		name           string
		doc            *processor.Document
		wantPredicates *assembler.IngestPredicates
		wantIdentities []common.TrustInformation
		wantErr        bool
	}{{
		name:           "testing",
		doc:            &testdata.GoModDoc,
		wantPredicates: nil,
		wantIdentities: nil,
		wantErr:        false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewGoModParser()
			err := d.Parse(ctx, tt.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("slsa.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
		})
	}
}
