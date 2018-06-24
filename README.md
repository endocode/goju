# Goju

Goju is for _Go JSON UNIT_ tests.

It is intended to test JSON files using other json files

```javascript
{
  "items": [ {
               "spec": {
	       "containers": [ {
                                 "image": "k8s.gcr.io/heapster-amd64:v1.5.0"
			       }
                             ]
               }
             }
           ]
}
```
is checked for regular expression by the JSON file

```javascript
{
  "items": {
    "length" :"1",
    "spec":{
      "containers":{
        "image":{
           "matches":"^(gcr.io/(google[-_]containers|k8s-minikube)|k8s.gcr.io)"
        }
      }
    }
  }
}
```
This means, the _items_ array is checked for _length_ 1 and any images in the _items/spec/containers_ array is checked, if the string matches the regular expression.
More examples are in the data directory.

The idea is to check configurations by other configurations, and implement checks adding additional leaves in JSON or YAML.

The executable checks are invoked by reflection on a _Check_ object and must have the same name as define in the `check.go` file with a leading capital letter.


