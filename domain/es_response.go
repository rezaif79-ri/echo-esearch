package domain

type EsSearchResponse[sourceT any] struct {
	Took     int                    `json:"took"`
	TimedOut bool                   `json:"timed_out"`
	Shards   map[string]interface{} `json:"_shards"`
	Hits     struct {
		Total    map[string]interface{} `json:"total"`
		MaxScore float32                `json:"max_score"`
		Hits     []struct {
			Source sourceT `json:"_source"`
		} `json:"hits"`
	}
}

// map[string]interface{}{
// 	"took": 2,
// 	"timed_out": false,
// 	"_shards": map[string]interface{}{
// 		"total": 1,
// 		"successful": 1,
// 		"skipped": 0,
// 		"failed": 0,
// 	},
// 	"hits": map[string]interface{}{
// 		"total": map[string]interface{}{
// 			"value": 1,
// 			"relation": "eq",
// 		},
// 		"max_score": 1,
// 		"hits": []interface{}{
// 			map[string]interface{}{
// 				"_index": "echo_books",
// 				"_type": "_doc",
// 				"_id": "01907e5d-5600-758c-aa1f-e597e207ef53",
// 				"_score": 1,
// 				"_source": map[string]interface{}{
// 					"book_id": "01907e5d-5600-758c-aa1f-e597e207ef53",
// 					"title": "How to be a good manager part 3",
// 					"pages": 1120,
// 					"author": "Norman",
// 				},
// 			},
// 		},
// 	},
// }
