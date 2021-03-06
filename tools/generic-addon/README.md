  # The generic-addon
  
  ## Introduction
  The generic-addon is simply used to generate different manifest for running an addon with the generic controller.
  You pass in the kind, group, and channel for the addon and it gives you all the necessary manifest required.
  
  ## List of Manifest
  The list of manifests it generates are:
  
  1. A CustomResourceDefinition for the addon
  2. A CustomResource of the addon
  3. RBAC(Clusterrole and clusterrolebinding) for creating, deleting, updating the custom resource
  4. A Generic resource for the addon
  5. The rbac for the addon manifest(The generic controller requires this to apply the manifest)
  
  ## Usage
  
  1. Get the code
    
  ```shell script
  go get https://github.com/kubernetes-sigs/cluster-addons
  ```
     
  2. Build
  
  ```shell script
  cd $GOPATH/kubernetes-sigs/cluster-addons/tools/generic-addon
  go install
  ```
 
  3. Run the go program

```shell script
generic-addon <KIND> <GROUP> <CHANNEL>
```

Example:
```shell script
generic-addon Dashboard addons.x-k8s.io ../channels
```
This command generates two files - `dashboard_crd_rbac.yaml` and `dashboard_sample.yaml`

You can apply this to your cluster. Remember to apply the `crd` file as that creates the CRD.
This also assumes you have already applied the CRD for the `Generic` resource.

Have fun! :tada: 

