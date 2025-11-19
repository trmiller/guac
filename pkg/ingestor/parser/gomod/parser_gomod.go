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
	"fmt"
	"time"

	"golang.org/x/mod/modfile"

	"github.com/guacsec/guac/pkg/assembler"
	model "github.com/guacsec/guac/pkg/assembler/clients/generated"
	"github.com/guacsec/guac/pkg/handler/processor"
	"github.com/guacsec/guac/pkg/ingestor/parser/common"
)

type goModParser struct {
	doc *processor.Document

	topLevelPackages map[string][]*model.PackageQualifierInputSpec
	timeScanned      time.Time
}

// NewDSSEParser initializes the dsseParser
func NewGoModParser() common.DocumentParser {
	return &goModParser{}
}

// Parse breaks out the document into the graph components
func (g *goModParser) Parse(ctx context.Context, doc *processor.Document) error {
	g.doc = doc

	gomod_doc, err := modfile.Parse("", doc.Blob, nil)
	if err != nil {
		return fmt.Errorf("failed to parse gomod document: %v", err)
	}

	fmt.Printf("Module Path: %s\n", gomod_doc.Module.Mod.Path)
	fmt.Println("Dependencies:")
	for _, require := range gomod_doc.Require {
		fmt.Printf("- %s %s\n", require.Mod.Path, require.Mod.Version)
	}

	fmt.Println("Replace:")
	for _, replace := range gomod_doc.Replace {
		fmt.Printf("- %s %s\n", replace.Old, replace.New)
	}

	return nil
}

func (g *goModParser) GetPredicates(ctx context.Context) *assembler.IngestPredicates {
	return nil
}

func (g *goModParser) GetIdentities(ctx context.Context) []common.TrustInformation {
	return nil
}

func (g *goModParser) GetIdentifiers(ctx context.Context) (*common.IdentifierStrings, error) {
	return nil, nil
}
