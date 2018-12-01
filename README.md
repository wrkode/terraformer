#Terraforming(need other name)

CLI tool for generate tf + tfstate files from current configuration(reverse terraform)

###Install
1. git clone project OR binary
2. Run `GO111MODULE=on go mod vendor`
3. go build -v
4. Copy your terraform provider plugin to ~/.terraform.d/plugins/{darwin,linux}_amd64 => https://www.terraform.io/docs/configuration/providers.html


#####Usage:
#####For GCP:
GOOGLE_CLOUD_PROJECT=YOUR_PROJECT ./terraforming google YOUR_SERVICE YOUR_ZONE
Examples: 

````
./terraform google firewalls europe-west1-c
````
````
./terraform google gcs europe-west1-c
````

````
./terraform google addresses europe-west1-c
````

List of support GCP services:
````
addresses
autoscalers
backendBuckets 
backendServices
disks
firewalls
forwardingRules
globalAddresses
globalForwardingRules
healthChecks
httpHealthChecks
httpsHealthChecks
images
instanceGroupManagers
instanceGroups
instanceTemplates - formatting HCL bug
instances - Values must match the following regular expression: '[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?|[1-9][0-9]{0,19}', invalidParameter(zone?)
networks
regionAutoscalers
regionBackendServices
regionDisks
regionInstanceGroupManagers
routers
routes
securityPolicies 
sslPolicies
subnetworks
targetHttpProxies - bug with proxy_id uint64 issue
targetHttpsProxies
targetSslProxies
targetTcpProxies
urlMaps
vpnTunnels
gcs
````



Your tf and tfstate file generate to `generated/gcp/zone/service`

#####For AWS:
./terraforming aws YOUR_SERVICE YOUR_REGION


````
./terraform aws sg eu-west1
````
````
./terraform aws s3 eu-west1
````
````
./terraform aws subnet eu-west1
````
List of support AWS services:
````
elb
iam - bug and not finish
igw
nacl
s3
sg
subnet
vpc
vpn_connection
vpn_gateway
````

###Developing:
Process for generate tf + tfstate files
1. Call GCP/AWS/other api and get list of resources
2. Iterate on resources and take only ID(not need mapping fields!!!)
3. Remove Attributes Reference.
4. Call to infrastructure and take tf + tfstate.


#####Infrastructure
1. Create empty(only IDs) tfstate file.
2. Run terraform refresh with this tfstate file.
3. Convert tfstate files to Go struct.
4. Generate HCL file - tf files.

All mapping of resource make by providers and terraform. Any upgrade need only for providers.
 
#####GCP compute resources
For GCP compute resources use generate code from `gcp_terraforming/compute_resources/gcp_compute_code_generator`

Regenerate code 
````
go run gcp_terraforming/compute_resources/gcp_compute_code_generator/*.go
````