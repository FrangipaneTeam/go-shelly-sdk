package main

import (
	"log"
	"os"
	"strings"

	"github.com/kr/pretty"
	"gopkg.in/yaml.v3"
)

/* TODO: API v1
- [ ] Add legacy http client
- [ ] Finish template
- [ ] Import devices openapi definitions
*/

type devices map[string]struct {
	deviceDetails
	Endpoints map[string]*endpoint
}

type endpoints struct {
	Compatibility []deviceDetails      `yaml:"compatibility"`
	Paths         map[string]*endpoint `yaml:"paths"`
}

type deviceDetails struct {
	Name      string `yaml:"name"`
	ProductID string `yaml:"productId"`
}

type endpoint struct {
	Description string `yaml:"description"`
	Get         struct {
		Tags        []string            `yaml:"tags"`
		OperationID string              `yaml:"operationId"`
		Responses   map[string]response `yaml:"responses"`
	} `yaml:"responses"`
	Parameters []parameter `yaml:"parameters"`
}

type parameter struct {
	Name        string `yaml:"name"`
	In          string `yaml:"in"`
	Description string `yaml:"description"`
	Schema      struct {
		Type string `yaml:"type"`
	} `yaml:"schema"`
}

type response struct {
	Description string `yaml:"description"`
	Content     struct {
		ApplicationJSON struct {
			Schema struct {
				Type       string              `yaml:"type"`
				Properties map[string]property `yaml:"properties"`
			} `yaml:"schema"`
		} `yaml:"application/json"`
	} `yaml:"content"`
}

type property struct {
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
}

func main() {
	var (
		d   = make(devices)
		err error
	)

	// list all yaml files in the template_v1 directory
	files, err := os.ReadDir("template_v1/")
	if err != nil {
		log.Fatal(err)
	}

	// for each yaml file in the template_v1 directory
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		// if file is not yaml file
		if f.Name()[len(f.Name())-5:] != ".yaml" {
			continue
		}

		log.Printf("Reading file %s", f.Name())

		data := &endpoints{
			Compatibility: []deviceDetails{},
			Paths:         map[string]*endpoint{},
		}

		// Read yaml file
		yamlFile, err := os.ReadFile("template_v1/" + f.Name())
		if err != nil {
			log.Fatalf("yamlFile.Get err #%v ", err)
		}

		// Parse yaml
		if err = yaml.Unmarshal(yamlFile, data); err != nil {
			log.Fatalf("error on Unmarshal Yaml file %s: %v", f.Name(), err)
		}

		for _, v := range data.Paths {
			// Remove HTML balises from description fields
			v.Description = removeHTMLBalises(v.Description)

			for i, p := range v.Parameters {
				v.Parameters[i].Description = removeHTMLBalises(p.Description)
			}

			for _, r := range v.Get.Responses {
				r.Description = removeHTMLBalises(r.Description)
				for _, c := range r.Content.ApplicationJSON.Schema.Properties {
					c.Description = removeHTMLBalises(c.Description)
				}
			}
		}

		// Add device in devices list
		for _, v := range data.Compatibility {
			d[v.ProductID] = struct {
				deviceDetails
				Endpoints map[string]*endpoint
			}{
				deviceDetails: v,
				Endpoints:     data.Paths,
			}
		}
	}

	pretty.Print(d)
}

func removeHTMLBalises(s string) string {
	listOfBalises := []string{
		"<p>",
		"</p>",
		"<br>",
		"<br/>",
		"<br />",
		"<strong>",
		"</strong>",
		"<em>",
		"</em>",
		"<code>",
		"</code>",
		"<pre>",
		"</pre>",
		"<ul>",
		"</ul>",
		"<li>",
		"</li>",
		"<ol>",
		"</ol>",
		"<h1>",
		"</h1>",
		"<h2>",
		"</h2>",
		"<h3>",
		"</h3>",
		"<h4>",
		"</h4>",
		"<h5>",
		"</h5>",
		"<h6>",
		"</h6>",
		"<hr>",
		"<hr/>",
		"<hr />",
		"<blockquote>",
		"</blockquote>",
		"<a>",
		"</a>",
		"<table>",
		"</table>",
		"<tr>",
		"</tr>",
		"<th>",
		"</th>",
		"<td>",
		"</td>",
		"<div>",
		"</div>",
		"<span>",
		"</span>",
		"<img>",
		"</img>",
	}

	for _, balise := range listOfBalises {
		s = strings.ReplaceAll(s, balise, "")
	}

	return s
}
