# autolabel

## background

Automatically add labels according to annotations, mainly for deployment, statefulset and other resources.

## usage 

Add annotations which like `autolabel/<pod name>=<label value>` to the resource, the label named `autolabel` with a value of `<label value>`  will be added to the resource automatically.
