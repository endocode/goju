# Goju

Goju is for _Go JSON UNIT_ tests.

It is intended to test JSON files using other json files

```javascript
{
  "items": [ 
			       {
				       "spec": {
							 "containers": [
							                 {
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

Examples are in the data directory


