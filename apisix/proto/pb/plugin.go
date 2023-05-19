/**
 * @Author: lyszhang
 * @Email: ericlyszhang@gmail.com
 * @Description:
 * @File:
 * @Version: 1.0.0
 * @Data:
 */

package pb

type ProxyRewritePlugin struct {
	Plugin struct {
		Scheme   string   `json:"scheme"`
		Uri      string   `json:"uri,omitempty"`
		RegexUri []string `json:"regex_uri,omitempty"`
	} `json:"proxy-rewrite"`
}
