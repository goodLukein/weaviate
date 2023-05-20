//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package clients

import (
	"fmt"

	"github.com/weaviate/weaviate/modules/text2vec-huggingface/ent"
)

type huggingFaceUrlBuilder struct {
	origin   string
	pathMask string
}

func newHuggingFaceUrlBuilder(config ent.VectorizationConfig) *huggingFaceUrlBuilder {
	return &huggingFaceUrlBuilder{
		origin:   config.Origin,
		pathMask: config.PathMask,
	}
}

func (o *huggingFaceUrlBuilder) url(model string) string {
	return fmt.Sprintf("%s%s", o.origin, o.getPath(model))
}

func (o *huggingFaceUrlBuilder) getPath(model string) string {
	return fmt.Sprintf(o.pathMask, model)
}
