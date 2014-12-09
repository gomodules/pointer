// Package elasticcache provides a client for Amazon ElastiCache.
package elasticcache

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// ElasticCache is a client for Amazon ElastiCache.
type ElasticCache struct {
	client aws.Client
}

// New returns a new ElasticCache client.
func New(key, secret, region string, client *http.Client) *ElasticCache {
	if client == nil {
		client = http.DefaultClient
	}

	return &ElasticCache{
		client: &aws.QueryClient{
			Signer: &aws.V4Signer{
				Key:     key,
				Secret:  secret,
				Service: "elasticache",
				Region:  region,
				IncludeXAmzContentSha256: true,
			},
			Client:     client,
			Endpoint:   endpoints.Lookup("elasticache", region),
			APIVersion: "2014-09-30",
		},
	}
}

// AuthorizeCacheSecurityGroupIngress the
// AuthorizeCacheSecurityGroupIngress operation allows network ingress to a
// cache security group. Applications using ElastiCache must be running on
// Amazon EC2, and Amazon EC2 security groups are used as the authorization
// mechanism.
func (c *ElasticCache) AuthorizeCacheSecurityGroupIngress(req AuthorizeCacheSecurityGroupIngressMessage) (resp *AuthorizeCacheSecurityGroupIngressResult, err error) {
	resp = &AuthorizeCacheSecurityGroupIngressResult{}
	err = c.client.Do("AuthorizeCacheSecurityGroupIngress", "POST", "/", req, resp)
	return
}

// CopySnapshot the CopySnapshot operation makes a copy of an existing
// snapshot.
func (c *ElasticCache) CopySnapshot(req CopySnapshotMessage) (resp *CopySnapshotResult, err error) {
	resp = &CopySnapshotResult{}
	err = c.client.Do("CopySnapshot", "POST", "/", req, resp)
	return
}

// CreateCacheCluster the CreateCacheCluster operation creates a cache
// cluster. All nodes in the cache cluster run the same protocol-compliant
// cache engine software, either Memcached or Redis.
func (c *ElasticCache) CreateCacheCluster(req CreateCacheClusterMessage) (resp *CreateCacheClusterResult, err error) {
	resp = &CreateCacheClusterResult{}
	err = c.client.Do("CreateCacheCluster", "POST", "/", req, resp)
	return
}

// CreateCacheParameterGroup the CreateCacheParameterGroup operation
// creates a new cache parameter group. A cache parameter group is a
// collection of parameters that you apply to all of the nodes in a cache
// cluster.
func (c *ElasticCache) CreateCacheParameterGroup(req CreateCacheParameterGroupMessage) (resp *CreateCacheParameterGroupResult, err error) {
	resp = &CreateCacheParameterGroupResult{}
	err = c.client.Do("CreateCacheParameterGroup", "POST", "/", req, resp)
	return
}

// CreateCacheSecurityGroup the CreateCacheSecurityGroup operation creates
// a new cache security group. Use a cache security group to control access
// to one or more cache clusters. Cache security groups are only used when
// you are creating a cache cluster outside of an Amazon Virtual Private
// Cloud If you are creating a cache cluster inside of a use a cache subnet
// group instead. For more information, see CreateCacheSubnetGroup
func (c *ElasticCache) CreateCacheSecurityGroup(req CreateCacheSecurityGroupMessage) (resp *CreateCacheSecurityGroupResult, err error) {
	resp = &CreateCacheSecurityGroupResult{}
	err = c.client.Do("CreateCacheSecurityGroup", "POST", "/", req, resp)
	return
}

// CreateCacheSubnetGroup the CreateCacheSubnetGroup operation creates a
// new cache subnet group. Use this parameter only when you are creating a
// cluster in an Amazon Virtual Private Cloud
func (c *ElasticCache) CreateCacheSubnetGroup(req CreateCacheSubnetGroupMessage) (resp *CreateCacheSubnetGroupResult, err error) {
	resp = &CreateCacheSubnetGroupResult{}
	err = c.client.Do("CreateCacheSubnetGroup", "POST", "/", req, resp)
	return
}

// CreateReplicationGroup the CreateReplicationGroup operation creates a
// replication group. A replication group is a collection of cache
// clusters, where one of the cache clusters is a read/write primary and
// the others are read-only replicas. Writes to the primary are
// automatically propagated to the replicas. When you create a replication
// group, you must specify an existing cache cluster that is in the primary
// role. When the replication group has been successfully created, you can
// add one or more read replica replicas to it, up to a total of five read
// replicas. Note: This action is valid only for Redis.
func (c *ElasticCache) CreateReplicationGroup(req CreateReplicationGroupMessage) (resp *CreateReplicationGroupResult, err error) {
	resp = &CreateReplicationGroupResult{}
	err = c.client.Do("CreateReplicationGroup", "POST", "/", req, resp)
	return
}

// CreateSnapshot the CreateSnapshot operation creates a copy of an entire
// cache cluster at a specific moment in time.
func (c *ElasticCache) CreateSnapshot(req CreateSnapshotMessage) (resp *CreateSnapshotResult, err error) {
	resp = &CreateSnapshotResult{}
	err = c.client.Do("CreateSnapshot", "POST", "/", req, resp)
	return
}

// DeleteCacheCluster the DeleteCacheCluster operation deletes a previously
// provisioned cache cluster. DeleteCacheCluster deletes all associated
// cache nodes, node endpoints and the cache cluster itself. When you
// receive a successful response from this operation, Amazon ElastiCache
// immediately begins deleting the cache cluster; you cannot cancel or
// revert this operation. This API cannot be used to delete a cache cluster
// that is the last read replica of a replication group that has automatic
// failover mode enabled.
func (c *ElasticCache) DeleteCacheCluster(req DeleteCacheClusterMessage) (resp *DeleteCacheClusterResult, err error) {
	resp = &DeleteCacheClusterResult{}
	err = c.client.Do("DeleteCacheCluster", "POST", "/", req, resp)
	return
}

// DeleteCacheParameterGroup the DeleteCacheParameterGroup operation
// deletes the specified cache parameter group. You cannot delete a cache
// parameter group if it is associated with any cache clusters.
func (c *ElasticCache) DeleteCacheParameterGroup(req DeleteCacheParameterGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteCacheParameterGroup", "POST", "/", req, nil)
	return
}

// DeleteCacheSecurityGroup the DeleteCacheSecurityGroup operation deletes
// a cache security group.
func (c *ElasticCache) DeleteCacheSecurityGroup(req DeleteCacheSecurityGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteCacheSecurityGroup", "POST", "/", req, nil)
	return
}

// DeleteCacheSubnetGroup the DeleteCacheSubnetGroup operation deletes a
// cache subnet group.
func (c *ElasticCache) DeleteCacheSubnetGroup(req DeleteCacheSubnetGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteCacheSubnetGroup", "POST", "/", req, nil)
	return
}

// DeleteReplicationGroup the DeleteReplicationGroup operation deletes an
// existing cluster. By default, this operation deletes the entire cluster,
// including the primary node group and all of the read replicas. You can
// optionally delete only the read replicas, while retaining the primary
// node group. When you receive a successful response from this operation,
// Amazon ElastiCache immediately begins deleting the selected resources;
// you cannot cancel or revert this operation.
func (c *ElasticCache) DeleteReplicationGroup(req DeleteReplicationGroupMessage) (resp *DeleteReplicationGroupResult, err error) {
	resp = &DeleteReplicationGroupResult{}
	err = c.client.Do("DeleteReplicationGroup", "POST", "/", req, resp)
	return
}

// DeleteSnapshot the DeleteSnapshot operation deletes an existing
// snapshot. When you receive a successful response from this operation,
// ElastiCache immediately begins deleting the snapshot; you cannot cancel
// or revert this operation.
func (c *ElasticCache) DeleteSnapshot(req DeleteSnapshotMessage) (resp *DeleteSnapshotResult, err error) {
	resp = &DeleteSnapshotResult{}
	err = c.client.Do("DeleteSnapshot", "POST", "/", req, resp)
	return
}

// DescribeCacheClusters the DescribeCacheClusters operation returns
// information about all provisioned cache clusters if no cache cluster
// identifier is specified, or about a specific cache cluster if a cache
// cluster identifier is supplied. By default, abbreviated information
// about the cache clusters(s) will be returned. You can use the optional
// ShowDetails flag to retrieve detailed information about the cache nodes
// associated with the cache clusters. These details include the DNS
// address and port for the cache node endpoint. If the cluster is in the
// state, only cluster level information will be displayed until all of the
// nodes are successfully provisioned. If the cluster is in the state, only
// cluster level information will be displayed. If cache nodes are
// currently being added to the cache cluster, node endpoint information
// and creation time for the additional nodes will not be displayed until
// they are completely provisioned. When the cache cluster state is
// available , the cluster is ready for use. If cache nodes are currently
// being removed from the cache cluster, no endpoint information for the
// removed nodes is displayed.
func (c *ElasticCache) DescribeCacheClusters(req DescribeCacheClustersMessage) (resp *DescribeCacheClustersResult, err error) {
	resp = &DescribeCacheClustersResult{}
	err = c.client.Do("DescribeCacheClusters", "POST", "/", req, resp)
	return
}

// DescribeCacheEngineVersions the DescribeCacheEngineVersions operation
// returns a list of the available cache engines and their versions.
func (c *ElasticCache) DescribeCacheEngineVersions(req DescribeCacheEngineVersionsMessage) (resp *DescribeCacheEngineVersionsResult, err error) {
	resp = &DescribeCacheEngineVersionsResult{}
	err = c.client.Do("DescribeCacheEngineVersions", "POST", "/", req, resp)
	return
}

// DescribeCacheParameterGroups the DescribeCacheParameterGroups operation
// returns a list of cache parameter group descriptions. If a cache
// parameter group name is specified, the list will contain only the
// descriptions for that group.
func (c *ElasticCache) DescribeCacheParameterGroups(req DescribeCacheParameterGroupsMessage) (resp *DescribeCacheParameterGroupsResult, err error) {
	resp = &DescribeCacheParameterGroupsResult{}
	err = c.client.Do("DescribeCacheParameterGroups", "POST", "/", req, resp)
	return
}

// DescribeCacheParameters the DescribeCacheParameters operation returns
// the detailed parameter list for a particular cache parameter group.
func (c *ElasticCache) DescribeCacheParameters(req DescribeCacheParametersMessage) (resp *DescribeCacheParametersResult, err error) {
	resp = &DescribeCacheParametersResult{}
	err = c.client.Do("DescribeCacheParameters", "POST", "/", req, resp)
	return
}

// DescribeCacheSecurityGroups the DescribeCacheSecurityGroups operation
// returns a list of cache security group descriptions. If a cache security
// group name is specified, the list will contain only the description of
// that group.
func (c *ElasticCache) DescribeCacheSecurityGroups(req DescribeCacheSecurityGroupsMessage) (resp *DescribeCacheSecurityGroupsResult, err error) {
	resp = &DescribeCacheSecurityGroupsResult{}
	err = c.client.Do("DescribeCacheSecurityGroups", "POST", "/", req, resp)
	return
}

// DescribeCacheSubnetGroups the DescribeCacheSubnetGroups operation
// returns a list of cache subnet group descriptions. If a subnet group
// name is specified, the list will contain only the description of that
// group.
func (c *ElasticCache) DescribeCacheSubnetGroups(req DescribeCacheSubnetGroupsMessage) (resp *DescribeCacheSubnetGroupsResult, err error) {
	resp = &DescribeCacheSubnetGroupsResult{}
	err = c.client.Do("DescribeCacheSubnetGroups", "POST", "/", req, resp)
	return
}

// DescribeEngineDefaultParameters the DescribeEngineDefaultParameters
// operation returns the default engine and system parameter information
// for the specified cache engine.
func (c *ElasticCache) DescribeEngineDefaultParameters(req DescribeEngineDefaultParametersMessage) (resp *DescribeEngineDefaultParametersResult, err error) {
	resp = &DescribeEngineDefaultParametersResult{}
	err = c.client.Do("DescribeEngineDefaultParameters", "POST", "/", req, resp)
	return
}

// DescribeEvents the DescribeEvents operation returns events related to
// cache clusters, cache security groups, and cache parameter groups. You
// can obtain events specific to a particular cache cluster, cache security
// group, or cache parameter group by providing the name as a parameter. By
// default, only the events occurring within the last hour are returned;
// however, you can retrieve up to 14 days' worth of events if necessary.
func (c *ElasticCache) DescribeEvents(req DescribeEventsMessage) (resp *DescribeEventsResult, err error) {
	resp = &DescribeEventsResult{}
	err = c.client.Do("DescribeEvents", "POST", "/", req, resp)
	return
}

// DescribeReplicationGroups the DescribeReplicationGroups operation
// returns information about a particular replication group. If no
// identifier is specified, DescribeReplicationGroups returns information
// about all replication groups.
func (c *ElasticCache) DescribeReplicationGroups(req DescribeReplicationGroupsMessage) (resp *DescribeReplicationGroupsResult, err error) {
	resp = &DescribeReplicationGroupsResult{}
	err = c.client.Do("DescribeReplicationGroups", "POST", "/", req, resp)
	return
}

// DescribeReservedCacheNodes the DescribeReservedCacheNodes operation
// returns information about reserved cache nodes for this account, or
// about a specified reserved cache node.
func (c *ElasticCache) DescribeReservedCacheNodes(req DescribeReservedCacheNodesMessage) (resp *DescribeReservedCacheNodesResult, err error) {
	resp = &DescribeReservedCacheNodesResult{}
	err = c.client.Do("DescribeReservedCacheNodes", "POST", "/", req, resp)
	return
}

// DescribeReservedCacheNodesOfferings the
// DescribeReservedCacheNodesOfferings operation lists available reserved
// cache node offerings.
func (c *ElasticCache) DescribeReservedCacheNodesOfferings(req DescribeReservedCacheNodesOfferingsMessage) (resp *DescribeReservedCacheNodesOfferingsResult, err error) {
	resp = &DescribeReservedCacheNodesOfferingsResult{}
	err = c.client.Do("DescribeReservedCacheNodesOfferings", "POST", "/", req, resp)
	return
}

// DescribeSnapshots the DescribeSnapshots operation returns information
// about cache cluster snapshots. By default, DescribeSnapshots lists all
// of your snapshots; it can optionally describe a single snapshot, or just
// the snapshots associated with a particular cache cluster.
func (c *ElasticCache) DescribeSnapshots(req DescribeSnapshotsMessage) (resp *DescribeSnapshotsResult, err error) {
	resp = &DescribeSnapshotsResult{}
	err = c.client.Do("DescribeSnapshots", "POST", "/", req, resp)
	return
}

// ModifyCacheCluster the ModifyCacheCluster operation modifies the
// settings for a cache cluster. You can use this operation to change one
// or more cluster configuration parameters by specifying the parameters
// and the new values.
func (c *ElasticCache) ModifyCacheCluster(req ModifyCacheClusterMessage) (resp *ModifyCacheClusterResult, err error) {
	resp = &ModifyCacheClusterResult{}
	err = c.client.Do("ModifyCacheCluster", "POST", "/", req, resp)
	return
}

// ModifyCacheParameterGroup the ModifyCacheParameterGroup operation
// modifies the parameters of a cache parameter group. You can modify up to
// 20 parameters in a single request by submitting a list parameter name
// and value pairs.
func (c *ElasticCache) ModifyCacheParameterGroup(req ModifyCacheParameterGroupMessage) (resp *ModifyCacheParameterGroupResult, err error) {
	resp = &ModifyCacheParameterGroupResult{}
	err = c.client.Do("ModifyCacheParameterGroup", "POST", "/", req, resp)
	return
}

// ModifyCacheSubnetGroup the ModifyCacheSubnetGroup operation modifies an
// existing cache subnet group.
func (c *ElasticCache) ModifyCacheSubnetGroup(req ModifyCacheSubnetGroupMessage) (resp *ModifyCacheSubnetGroupResult, err error) {
	resp = &ModifyCacheSubnetGroupResult{}
	err = c.client.Do("ModifyCacheSubnetGroup", "POST", "/", req, resp)
	return
}

// ModifyReplicationGroup the ModifyReplicationGroup operation modifies the
// settings for a replication group.
func (c *ElasticCache) ModifyReplicationGroup(req ModifyReplicationGroupMessage) (resp *ModifyReplicationGroupResult, err error) {
	resp = &ModifyReplicationGroupResult{}
	err = c.client.Do("ModifyReplicationGroup", "POST", "/", req, resp)
	return
}

// PurchaseReservedCacheNodesOffering the
// PurchaseReservedCacheNodesOffering operation allows you to purchase a
// reserved cache node offering.
func (c *ElasticCache) PurchaseReservedCacheNodesOffering(req PurchaseReservedCacheNodesOfferingMessage) (resp *PurchaseReservedCacheNodesOfferingResult, err error) {
	resp = &PurchaseReservedCacheNodesOfferingResult{}
	err = c.client.Do("PurchaseReservedCacheNodesOffering", "POST", "/", req, resp)
	return
}

// RebootCacheCluster the RebootCacheCluster operation reboots some, or
// all, of the cache nodes within a provisioned cache cluster. This API
// will apply any modified cache parameter groups to the cache cluster. The
// reboot action takes place as soon as possible, and results in a
// momentary outage to the cache cluster. During the reboot, the cache
// cluster status is set to The reboot causes the contents of the cache
// (for each cache node being rebooted) to be lost. When the reboot is
// complete, a cache cluster event is created.
func (c *ElasticCache) RebootCacheCluster(req RebootCacheClusterMessage) (resp *RebootCacheClusterResult, err error) {
	resp = &RebootCacheClusterResult{}
	err = c.client.Do("RebootCacheCluster", "POST", "/", req, resp)
	return
}

// ResetCacheParameterGroup the ResetCacheParameterGroup operation modifies
// the parameters of a cache parameter group to the engine or system
// default value. You can reset specific parameters by submitting a list of
// parameter names. To reset the entire cache parameter group, specify the
// ResetAllParameters and CacheParameterGroupName parameters.
func (c *ElasticCache) ResetCacheParameterGroup(req ResetCacheParameterGroupMessage) (resp *ResetCacheParameterGroupResult, err error) {
	resp = &ResetCacheParameterGroupResult{}
	err = c.client.Do("ResetCacheParameterGroup", "POST", "/", req, resp)
	return
}

// RevokeCacheSecurityGroupIngress the RevokeCacheSecurityGroupIngress
// operation revokes ingress from a cache security group. Use this
// operation to disallow access from an Amazon EC2 security group that had
// been previously authorized.
func (c *ElasticCache) RevokeCacheSecurityGroupIngress(req RevokeCacheSecurityGroupIngressMessage) (resp *RevokeCacheSecurityGroupIngressResult, err error) {
	resp = &RevokeCacheSecurityGroupIngressResult{}
	err = c.client.Do("RevokeCacheSecurityGroupIngress", "POST", "/", req, resp)
	return
}

// AuthorizeCacheSecurityGroupIngressMessage is undocumented.
type AuthorizeCacheSecurityGroupIngressMessage struct {
	CacheSecurityGroupName  string `xml:"CacheSecurityGroupName"`
	EC2SecurityGroupName    string `xml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerID string `xml:"EC2SecurityGroupOwnerId"`
}

// AuthorizeCacheSecurityGroupIngressResult is undocumented.
type AuthorizeCacheSecurityGroupIngressResult struct {
	CacheSecurityGroup CacheSecurityGroup `xml:"AuthorizeCacheSecurityGroupIngressResult>CacheSecurityGroup"`
}

// AvailabilityZone is undocumented.
type AvailabilityZone struct {
	Name string `xml:"Name"`
}

// CacheCluster is undocumented.
type CacheCluster struct {
	AutoMinorVersionUpgrade    bool                           `xml:"AutoMinorVersionUpgrade"`
	CacheClusterCreateTime     time.Time                      `xml:"CacheClusterCreateTime"`
	CacheClusterID             string                         `xml:"CacheClusterId"`
	CacheClusterStatus         string                         `xml:"CacheClusterStatus"`
	CacheNodeType              string                         `xml:"CacheNodeType"`
	CacheNodes                 []CacheNode                    `xml:"CacheNodes>CacheNode"`
	CacheParameterGroup        CacheParameterGroupStatus      `xml:"CacheParameterGroup"`
	CacheSecurityGroups        []CacheSecurityGroupMembership `xml:"CacheSecurityGroups>CacheSecurityGroup"`
	CacheSubnetGroupName       string                         `xml:"CacheSubnetGroupName"`
	ClientDownloadLandingPage  string                         `xml:"ClientDownloadLandingPage"`
	ConfigurationEndpoint      Endpoint                       `xml:"ConfigurationEndpoint"`
	Engine                     string                         `xml:"Engine"`
	EngineVersion              string                         `xml:"EngineVersion"`
	NotificationConfiguration  NotificationConfiguration      `xml:"NotificationConfiguration"`
	NumCacheNodes              int                            `xml:"NumCacheNodes"`
	PendingModifiedValues      PendingModifiedValues          `xml:"PendingModifiedValues"`
	PreferredAvailabilityZone  string                         `xml:"PreferredAvailabilityZone"`
	PreferredMaintenanceWindow string                         `xml:"PreferredMaintenanceWindow"`
	ReplicationGroupID         string                         `xml:"ReplicationGroupId"`
	SecurityGroups             []SecurityGroupMembership      `xml:"SecurityGroups>member"`
	SnapshotRetentionLimit     int                            `xml:"SnapshotRetentionLimit"`
	SnapshotWindow             string                         `xml:"SnapshotWindow"`
}

// CacheClusterMessage is undocumented.
type CacheClusterMessage struct {
	CacheClusters []CacheCluster `xml:"DescribeCacheClustersResult>CacheClusters>CacheCluster"`
	Marker        string         `xml:"DescribeCacheClustersResult>Marker"`
}

// CacheEngineVersion is undocumented.
type CacheEngineVersion struct {
	CacheEngineDescription        string `xml:"CacheEngineDescription"`
	CacheEngineVersionDescription string `xml:"CacheEngineVersionDescription"`
	CacheParameterGroupFamily     string `xml:"CacheParameterGroupFamily"`
	Engine                        string `xml:"Engine"`
	EngineVersion                 string `xml:"EngineVersion"`
}

// CacheEngineVersionMessage is undocumented.
type CacheEngineVersionMessage struct {
	CacheEngineVersions []CacheEngineVersion `xml:"DescribeCacheEngineVersionsResult>CacheEngineVersions>CacheEngineVersion"`
	Marker              string               `xml:"DescribeCacheEngineVersionsResult>Marker"`
}

// CacheNode is undocumented.
type CacheNode struct {
	CacheNodeCreateTime      time.Time `xml:"CacheNodeCreateTime"`
	CacheNodeID              string    `xml:"CacheNodeId"`
	CacheNodeStatus          string    `xml:"CacheNodeStatus"`
	CustomerAvailabilityZone string    `xml:"CustomerAvailabilityZone"`
	Endpoint                 Endpoint  `xml:"Endpoint"`
	ParameterGroupStatus     string    `xml:"ParameterGroupStatus"`
	SourceCacheNodeID        string    `xml:"SourceCacheNodeId"`
}

// CacheNodeTypeSpecificParameter is undocumented.
type CacheNodeTypeSpecificParameter struct {
	AllowedValues               string                       `xml:"AllowedValues"`
	CacheNodeTypeSpecificValues []CacheNodeTypeSpecificValue `xml:"CacheNodeTypeSpecificValues>CacheNodeTypeSpecificValue"`
	DataType                    string                       `xml:"DataType"`
	Description                 string                       `xml:"Description"`
	IsModifiable                bool                         `xml:"IsModifiable"`
	MinimumEngineVersion        string                       `xml:"MinimumEngineVersion"`
	ParameterName               string                       `xml:"ParameterName"`
	Source                      string                       `xml:"Source"`
}

// CacheNodeTypeSpecificValue is undocumented.
type CacheNodeTypeSpecificValue struct {
	CacheNodeType string `xml:"CacheNodeType"`
	Value         string `xml:"Value"`
}

// CacheParameterGroup is undocumented.
type CacheParameterGroup struct {
	CacheParameterGroupFamily string `xml:"CacheParameterGroupFamily"`
	CacheParameterGroupName   string `xml:"CacheParameterGroupName"`
	Description               string `xml:"Description"`
}

// CacheParameterGroupDetails is undocumented.
type CacheParameterGroupDetails struct {
	CacheNodeTypeSpecificParameters []CacheNodeTypeSpecificParameter `xml:"DescribeCacheParametersResult>CacheNodeTypeSpecificParameters>CacheNodeTypeSpecificParameter"`
	Marker                          string                           `xml:"DescribeCacheParametersResult>Marker"`
	Parameters                      []Parameter                      `xml:"DescribeCacheParametersResult>Parameters>Parameter"`
}

// CacheParameterGroupNameMessage is undocumented.
type CacheParameterGroupNameMessage struct {
	CacheParameterGroupName string `xml:"CacheParameterGroupName"`
}

// CacheParameterGroupStatus is undocumented.
type CacheParameterGroupStatus struct {
	CacheNodeIdsToReboot    []string `xml:"CacheNodeIdsToReboot>CacheNodeId"`
	CacheParameterGroupName string   `xml:"CacheParameterGroupName"`
	ParameterApplyStatus    string   `xml:"ParameterApplyStatus"`
}

// CacheParameterGroupsMessage is undocumented.
type CacheParameterGroupsMessage struct {
	CacheParameterGroups []CacheParameterGroup `xml:"DescribeCacheParameterGroupsResult>CacheParameterGroups>CacheParameterGroup"`
	Marker               string                `xml:"DescribeCacheParameterGroupsResult>Marker"`
}

// CacheSecurityGroup is undocumented.
type CacheSecurityGroup struct {
	CacheSecurityGroupName string             `xml:"CacheSecurityGroupName"`
	Description            string             `xml:"Description"`
	EC2SecurityGroups      []EC2SecurityGroup `xml:"EC2SecurityGroups>EC2SecurityGroup"`
	OwnerID                string             `xml:"OwnerId"`
}

// CacheSecurityGroupMembership is undocumented.
type CacheSecurityGroupMembership struct {
	CacheSecurityGroupName string `xml:"CacheSecurityGroupName"`
	Status                 string `xml:"Status"`
}

// CacheSecurityGroupMessage is undocumented.
type CacheSecurityGroupMessage struct {
	CacheSecurityGroups []CacheSecurityGroup `xml:"DescribeCacheSecurityGroupsResult>CacheSecurityGroups>CacheSecurityGroup"`
	Marker              string               `xml:"DescribeCacheSecurityGroupsResult>Marker"`
}

// CacheSubnetGroup is undocumented.
type CacheSubnetGroup struct {
	CacheSubnetGroupDescription string   `xml:"CacheSubnetGroupDescription"`
	CacheSubnetGroupName        string   `xml:"CacheSubnetGroupName"`
	Subnets                     []Subnet `xml:"Subnets>Subnet"`
	VpcID                       string   `xml:"VpcId"`
}

// CacheSubnetGroupMessage is undocumented.
type CacheSubnetGroupMessage struct {
	CacheSubnetGroups []CacheSubnetGroup `xml:"DescribeCacheSubnetGroupsResult>CacheSubnetGroups>CacheSubnetGroup"`
	Marker            string             `xml:"DescribeCacheSubnetGroupsResult>Marker"`
}

// CopySnapshotMessage is undocumented.
type CopySnapshotMessage struct {
	SourceSnapshotName string `xml:"SourceSnapshotName"`
	TargetSnapshotName string `xml:"TargetSnapshotName"`
}

// CopySnapshotResult is undocumented.
type CopySnapshotResult struct {
	Snapshot Snapshot `xml:"CopySnapshotResult>Snapshot"`
}

// CreateCacheClusterMessage is undocumented.
type CreateCacheClusterMessage struct {
	AZMode                     string   `xml:"AZMode"`
	AutoMinorVersionUpgrade    bool     `xml:"AutoMinorVersionUpgrade"`
	CacheClusterID             string   `xml:"CacheClusterId"`
	CacheNodeType              string   `xml:"CacheNodeType"`
	CacheParameterGroupName    string   `xml:"CacheParameterGroupName"`
	CacheSecurityGroupNames    []string `xml:"CacheSecurityGroupNames>CacheSecurityGroupName"`
	CacheSubnetGroupName       string   `xml:"CacheSubnetGroupName"`
	Engine                     string   `xml:"Engine"`
	EngineVersion              string   `xml:"EngineVersion"`
	NotificationTopicARN       string   `xml:"NotificationTopicArn"`
	NumCacheNodes              int      `xml:"NumCacheNodes"`
	Port                       int      `xml:"Port"`
	PreferredAvailabilityZone  string   `xml:"PreferredAvailabilityZone"`
	PreferredAvailabilityZones []string `xml:"PreferredAvailabilityZones>PreferredAvailabilityZone"`
	PreferredMaintenanceWindow string   `xml:"PreferredMaintenanceWindow"`
	ReplicationGroupID         string   `xml:"ReplicationGroupId"`
	SecurityGroupIds           []string `xml:"SecurityGroupIds>SecurityGroupId"`
	SnapshotARNs               []string `xml:"SnapshotArns>SnapshotArn"`
	SnapshotName               string   `xml:"SnapshotName"`
	SnapshotRetentionLimit     int      `xml:"SnapshotRetentionLimit"`
	SnapshotWindow             string   `xml:"SnapshotWindow"`
}

// CreateCacheClusterResult is undocumented.
type CreateCacheClusterResult struct {
	CacheCluster CacheCluster `xml:"CreateCacheClusterResult>CacheCluster"`
}

// CreateCacheParameterGroupMessage is undocumented.
type CreateCacheParameterGroupMessage struct {
	CacheParameterGroupFamily string `xml:"CacheParameterGroupFamily"`
	CacheParameterGroupName   string `xml:"CacheParameterGroupName"`
	Description               string `xml:"Description"`
}

// CreateCacheParameterGroupResult is undocumented.
type CreateCacheParameterGroupResult struct {
	CacheParameterGroup CacheParameterGroup `xml:"CreateCacheParameterGroupResult>CacheParameterGroup"`
}

// CreateCacheSecurityGroupMessage is undocumented.
type CreateCacheSecurityGroupMessage struct {
	CacheSecurityGroupName string `xml:"CacheSecurityGroupName"`
	Description            string `xml:"Description"`
}

// CreateCacheSecurityGroupResult is undocumented.
type CreateCacheSecurityGroupResult struct {
	CacheSecurityGroup CacheSecurityGroup `xml:"CreateCacheSecurityGroupResult>CacheSecurityGroup"`
}

// CreateCacheSubnetGroupMessage is undocumented.
type CreateCacheSubnetGroupMessage struct {
	CacheSubnetGroupDescription string   `xml:"CacheSubnetGroupDescription"`
	CacheSubnetGroupName        string   `xml:"CacheSubnetGroupName"`
	SubnetIds                   []string `xml:"SubnetIds>SubnetIdentifier"`
}

// CreateCacheSubnetGroupResult is undocumented.
type CreateCacheSubnetGroupResult struct {
	CacheSubnetGroup CacheSubnetGroup `xml:"CreateCacheSubnetGroupResult>CacheSubnetGroup"`
}

// CreateReplicationGroupMessage is undocumented.
type CreateReplicationGroupMessage struct {
	AutoMinorVersionUpgrade     bool     `xml:"AutoMinorVersionUpgrade"`
	AutomaticFailoverEnabled    bool     `xml:"AutomaticFailoverEnabled"`
	CacheNodeType               string   `xml:"CacheNodeType"`
	CacheParameterGroupName     string   `xml:"CacheParameterGroupName"`
	CacheSecurityGroupNames     []string `xml:"CacheSecurityGroupNames>CacheSecurityGroupName"`
	CacheSubnetGroupName        string   `xml:"CacheSubnetGroupName"`
	Engine                      string   `xml:"Engine"`
	EngineVersion               string   `xml:"EngineVersion"`
	NotificationTopicARN        string   `xml:"NotificationTopicArn"`
	NumCacheClusters            int      `xml:"NumCacheClusters"`
	Port                        int      `xml:"Port"`
	PreferredCacheClusterAZs    []string `xml:"PreferredCacheClusterAZs>AvailabilityZone"`
	PreferredMaintenanceWindow  string   `xml:"PreferredMaintenanceWindow"`
	PrimaryClusterID            string   `xml:"PrimaryClusterId"`
	ReplicationGroupDescription string   `xml:"ReplicationGroupDescription"`
	ReplicationGroupID          string   `xml:"ReplicationGroupId"`
	SecurityGroupIds            []string `xml:"SecurityGroupIds>SecurityGroupId"`
	SnapshotARNs                []string `xml:"SnapshotArns>SnapshotArn"`
	SnapshotName                string   `xml:"SnapshotName"`
	SnapshotRetentionLimit      int      `xml:"SnapshotRetentionLimit"`
	SnapshotWindow              string   `xml:"SnapshotWindow"`
}

// CreateReplicationGroupResult is undocumented.
type CreateReplicationGroupResult struct {
	ReplicationGroup ReplicationGroup `xml:"CreateReplicationGroupResult>ReplicationGroup"`
}

// CreateSnapshotMessage is undocumented.
type CreateSnapshotMessage struct {
	CacheClusterID string `xml:"CacheClusterId"`
	SnapshotName   string `xml:"SnapshotName"`
}

// CreateSnapshotResult is undocumented.
type CreateSnapshotResult struct {
	Snapshot Snapshot `xml:"CreateSnapshotResult>Snapshot"`
}

// DeleteCacheClusterMessage is undocumented.
type DeleteCacheClusterMessage struct {
	CacheClusterID          string `xml:"CacheClusterId"`
	FinalSnapshotIdentifier string `xml:"FinalSnapshotIdentifier"`
}

// DeleteCacheClusterResult is undocumented.
type DeleteCacheClusterResult struct {
	CacheCluster CacheCluster `xml:"DeleteCacheClusterResult>CacheCluster"`
}

// DeleteCacheParameterGroupMessage is undocumented.
type DeleteCacheParameterGroupMessage struct {
	CacheParameterGroupName string `xml:"CacheParameterGroupName"`
}

// DeleteCacheSecurityGroupMessage is undocumented.
type DeleteCacheSecurityGroupMessage struct {
	CacheSecurityGroupName string `xml:"CacheSecurityGroupName"`
}

// DeleteCacheSubnetGroupMessage is undocumented.
type DeleteCacheSubnetGroupMessage struct {
	CacheSubnetGroupName string `xml:"CacheSubnetGroupName"`
}

// DeleteReplicationGroupMessage is undocumented.
type DeleteReplicationGroupMessage struct {
	FinalSnapshotIdentifier string `xml:"FinalSnapshotIdentifier"`
	ReplicationGroupID      string `xml:"ReplicationGroupId"`
	RetainPrimaryCluster    bool   `xml:"RetainPrimaryCluster"`
}

// DeleteReplicationGroupResult is undocumented.
type DeleteReplicationGroupResult struct {
	ReplicationGroup ReplicationGroup `xml:"DeleteReplicationGroupResult>ReplicationGroup"`
}

// DeleteSnapshotMessage is undocumented.
type DeleteSnapshotMessage struct {
	SnapshotName string `xml:"SnapshotName"`
}

// DeleteSnapshotResult is undocumented.
type DeleteSnapshotResult struct {
	Snapshot Snapshot `xml:"DeleteSnapshotResult>Snapshot"`
}

// DescribeCacheClustersMessage is undocumented.
type DescribeCacheClustersMessage struct {
	CacheClusterID    string `xml:"CacheClusterId"`
	Marker            string `xml:"Marker"`
	MaxRecords        int    `xml:"MaxRecords"`
	ShowCacheNodeInfo bool   `xml:"ShowCacheNodeInfo"`
}

// DescribeCacheEngineVersionsMessage is undocumented.
type DescribeCacheEngineVersionsMessage struct {
	CacheParameterGroupFamily string `xml:"CacheParameterGroupFamily"`
	DefaultOnly               bool   `xml:"DefaultOnly"`
	Engine                    string `xml:"Engine"`
	EngineVersion             string `xml:"EngineVersion"`
	Marker                    string `xml:"Marker"`
	MaxRecords                int    `xml:"MaxRecords"`
}

// DescribeCacheParameterGroupsMessage is undocumented.
type DescribeCacheParameterGroupsMessage struct {
	CacheParameterGroupName string `xml:"CacheParameterGroupName"`
	Marker                  string `xml:"Marker"`
	MaxRecords              int    `xml:"MaxRecords"`
}

// DescribeCacheParametersMessage is undocumented.
type DescribeCacheParametersMessage struct {
	CacheParameterGroupName string `xml:"CacheParameterGroupName"`
	Marker                  string `xml:"Marker"`
	MaxRecords              int    `xml:"MaxRecords"`
	Source                  string `xml:"Source"`
}

// DescribeCacheSecurityGroupsMessage is undocumented.
type DescribeCacheSecurityGroupsMessage struct {
	CacheSecurityGroupName string `xml:"CacheSecurityGroupName"`
	Marker                 string `xml:"Marker"`
	MaxRecords             int    `xml:"MaxRecords"`
}

// DescribeCacheSubnetGroupsMessage is undocumented.
type DescribeCacheSubnetGroupsMessage struct {
	CacheSubnetGroupName string `xml:"CacheSubnetGroupName"`
	Marker               string `xml:"Marker"`
	MaxRecords           int    `xml:"MaxRecords"`
}

// DescribeEngineDefaultParametersMessage is undocumented.
type DescribeEngineDefaultParametersMessage struct {
	CacheParameterGroupFamily string `xml:"CacheParameterGroupFamily"`
	Marker                    string `xml:"Marker"`
	MaxRecords                int    `xml:"MaxRecords"`
}

// DescribeEngineDefaultParametersResult is undocumented.
type DescribeEngineDefaultParametersResult struct {
	EngineDefaults EngineDefaults `xml:"DescribeEngineDefaultParametersResult>EngineDefaults"`
}

// DescribeEventsMessage is undocumented.
type DescribeEventsMessage struct {
	Duration         int       `xml:"Duration"`
	EndTime          time.Time `xml:"EndTime"`
	Marker           string    `xml:"Marker"`
	MaxRecords       int       `xml:"MaxRecords"`
	SourceIdentifier string    `xml:"SourceIdentifier"`
	SourceType       string    `xml:"SourceType"`
	StartTime        time.Time `xml:"StartTime"`
}

// DescribeReplicationGroupsMessage is undocumented.
type DescribeReplicationGroupsMessage struct {
	Marker             string `xml:"Marker"`
	MaxRecords         int    `xml:"MaxRecords"`
	ReplicationGroupID string `xml:"ReplicationGroupId"`
}

// DescribeReservedCacheNodesMessage is undocumented.
type DescribeReservedCacheNodesMessage struct {
	CacheNodeType                string `xml:"CacheNodeType"`
	Duration                     string `xml:"Duration"`
	Marker                       string `xml:"Marker"`
	MaxRecords                   int    `xml:"MaxRecords"`
	OfferingType                 string `xml:"OfferingType"`
	ProductDescription           string `xml:"ProductDescription"`
	ReservedCacheNodeID          string `xml:"ReservedCacheNodeId"`
	ReservedCacheNodesOfferingID string `xml:"ReservedCacheNodesOfferingId"`
}

// DescribeReservedCacheNodesOfferingsMessage is undocumented.
type DescribeReservedCacheNodesOfferingsMessage struct {
	CacheNodeType                string `xml:"CacheNodeType"`
	Duration                     string `xml:"Duration"`
	Marker                       string `xml:"Marker"`
	MaxRecords                   int    `xml:"MaxRecords"`
	OfferingType                 string `xml:"OfferingType"`
	ProductDescription           string `xml:"ProductDescription"`
	ReservedCacheNodesOfferingID string `xml:"ReservedCacheNodesOfferingId"`
}

// DescribeSnapshotsListMessage is undocumented.
type DescribeSnapshotsListMessage struct {
	Marker    string     `xml:"DescribeSnapshotsResult>Marker"`
	Snapshots []Snapshot `xml:"DescribeSnapshotsResult>Snapshots>Snapshot"`
}

// DescribeSnapshotsMessage is undocumented.
type DescribeSnapshotsMessage struct {
	CacheClusterID string `xml:"CacheClusterId"`
	Marker         string `xml:"Marker"`
	MaxRecords     int    `xml:"MaxRecords"`
	SnapshotName   string `xml:"SnapshotName"`
	SnapshotSource string `xml:"SnapshotSource"`
}

// EC2SecurityGroup is undocumented.
type EC2SecurityGroup struct {
	EC2SecurityGroupName    string `xml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerID string `xml:"EC2SecurityGroupOwnerId"`
	Status                  string `xml:"Status"`
}

// Endpoint is undocumented.
type Endpoint struct {
	Address string `xml:"Address"`
	Port    int    `xml:"Port"`
}

// EngineDefaults is undocumented.
type EngineDefaults struct {
	CacheNodeTypeSpecificParameters []CacheNodeTypeSpecificParameter `xml:"CacheNodeTypeSpecificParameters>CacheNodeTypeSpecificParameter"`
	CacheParameterGroupFamily       string                           `xml:"CacheParameterGroupFamily"`
	Marker                          string                           `xml:"Marker"`
	Parameters                      []Parameter                      `xml:"Parameters>Parameter"`
}

// Event is undocumented.
type Event struct {
	Date             time.Time `xml:"Date"`
	Message          string    `xml:"Message"`
	SourceIdentifier string    `xml:"SourceIdentifier"`
	SourceType       string    `xml:"SourceType"`
}

// EventsMessage is undocumented.
type EventsMessage struct {
	Events []Event `xml:"DescribeEventsResult>Events>Event"`
	Marker string  `xml:"DescribeEventsResult>Marker"`
}

// ModifyCacheClusterMessage is undocumented.
type ModifyCacheClusterMessage struct {
	AZMode                     string   `xml:"AZMode"`
	ApplyImmediately           bool     `xml:"ApplyImmediately"`
	AutoMinorVersionUpgrade    bool     `xml:"AutoMinorVersionUpgrade"`
	CacheClusterID             string   `xml:"CacheClusterId"`
	CacheNodeIdsToRemove       []string `xml:"CacheNodeIdsToRemove>CacheNodeId"`
	CacheParameterGroupName    string   `xml:"CacheParameterGroupName"`
	CacheSecurityGroupNames    []string `xml:"CacheSecurityGroupNames>CacheSecurityGroupName"`
	EngineVersion              string   `xml:"EngineVersion"`
	NewAvailabilityZones       []string `xml:"NewAvailabilityZones>PreferredAvailabilityZone"`
	NotificationTopicARN       string   `xml:"NotificationTopicArn"`
	NotificationTopicStatus    string   `xml:"NotificationTopicStatus"`
	NumCacheNodes              int      `xml:"NumCacheNodes"`
	PreferredMaintenanceWindow string   `xml:"PreferredMaintenanceWindow"`
	SecurityGroupIds           []string `xml:"SecurityGroupIds>SecurityGroupId"`
	SnapshotRetentionLimit     int      `xml:"SnapshotRetentionLimit"`
	SnapshotWindow             string   `xml:"SnapshotWindow"`
}

// ModifyCacheClusterResult is undocumented.
type ModifyCacheClusterResult struct {
	CacheCluster CacheCluster `xml:"ModifyCacheClusterResult>CacheCluster"`
}

// ModifyCacheParameterGroupMessage is undocumented.
type ModifyCacheParameterGroupMessage struct {
	CacheParameterGroupName string               `xml:"CacheParameterGroupName"`
	ParameterNameValues     []ParameterNameValue `xml:"ParameterNameValues>ParameterNameValue"`
}

// ModifyCacheSubnetGroupMessage is undocumented.
type ModifyCacheSubnetGroupMessage struct {
	CacheSubnetGroupDescription string   `xml:"CacheSubnetGroupDescription"`
	CacheSubnetGroupName        string   `xml:"CacheSubnetGroupName"`
	SubnetIds                   []string `xml:"SubnetIds>SubnetIdentifier"`
}

// ModifyCacheSubnetGroupResult is undocumented.
type ModifyCacheSubnetGroupResult struct {
	CacheSubnetGroup CacheSubnetGroup `xml:"ModifyCacheSubnetGroupResult>CacheSubnetGroup"`
}

// ModifyReplicationGroupMessage is undocumented.
type ModifyReplicationGroupMessage struct {
	ApplyImmediately            bool     `xml:"ApplyImmediately"`
	AutoMinorVersionUpgrade     bool     `xml:"AutoMinorVersionUpgrade"`
	AutomaticFailoverEnabled    bool     `xml:"AutomaticFailoverEnabled"`
	CacheParameterGroupName     string   `xml:"CacheParameterGroupName"`
	CacheSecurityGroupNames     []string `xml:"CacheSecurityGroupNames>CacheSecurityGroupName"`
	EngineVersion               string   `xml:"EngineVersion"`
	NotificationTopicARN        string   `xml:"NotificationTopicArn"`
	NotificationTopicStatus     string   `xml:"NotificationTopicStatus"`
	PreferredMaintenanceWindow  string   `xml:"PreferredMaintenanceWindow"`
	PrimaryClusterID            string   `xml:"PrimaryClusterId"`
	ReplicationGroupDescription string   `xml:"ReplicationGroupDescription"`
	ReplicationGroupID          string   `xml:"ReplicationGroupId"`
	SecurityGroupIds            []string `xml:"SecurityGroupIds>SecurityGroupId"`
	SnapshotRetentionLimit      int      `xml:"SnapshotRetentionLimit"`
	SnapshotWindow              string   `xml:"SnapshotWindow"`
	SnapshottingClusterID       string   `xml:"SnapshottingClusterId"`
}

// ModifyReplicationGroupResult is undocumented.
type ModifyReplicationGroupResult struct {
	ReplicationGroup ReplicationGroup `xml:"ModifyReplicationGroupResult>ReplicationGroup"`
}

// NodeGroup is undocumented.
type NodeGroup struct {
	NodeGroupID      string            `xml:"NodeGroupId"`
	NodeGroupMembers []NodeGroupMember `xml:"NodeGroupMembers>NodeGroupMember"`
	PrimaryEndpoint  Endpoint          `xml:"PrimaryEndpoint"`
	Status           string            `xml:"Status"`
}

// NodeGroupMember is undocumented.
type NodeGroupMember struct {
	CacheClusterID            string   `xml:"CacheClusterId"`
	CacheNodeID               string   `xml:"CacheNodeId"`
	CurrentRole               string   `xml:"CurrentRole"`
	PreferredAvailabilityZone string   `xml:"PreferredAvailabilityZone"`
	ReadEndpoint              Endpoint `xml:"ReadEndpoint"`
}

// NodeSnapshot is undocumented.
type NodeSnapshot struct {
	CacheNodeCreateTime time.Time `xml:"CacheNodeCreateTime"`
	CacheNodeID         string    `xml:"CacheNodeId"`
	CacheSize           string    `xml:"CacheSize"`
	SnapshotCreateTime  time.Time `xml:"SnapshotCreateTime"`
}

// NotificationConfiguration is undocumented.
type NotificationConfiguration struct {
	TopicARN    string `xml:"TopicArn"`
	TopicStatus string `xml:"TopicStatus"`
}

// Parameter is undocumented.
type Parameter struct {
	AllowedValues        string `xml:"AllowedValues"`
	DataType             string `xml:"DataType"`
	Description          string `xml:"Description"`
	IsModifiable         bool   `xml:"IsModifiable"`
	MinimumEngineVersion string `xml:"MinimumEngineVersion"`
	ParameterName        string `xml:"ParameterName"`
	ParameterValue       string `xml:"ParameterValue"`
	Source               string `xml:"Source"`
}

// ParameterNameValue is undocumented.
type ParameterNameValue struct {
	ParameterName  string `xml:"ParameterName"`
	ParameterValue string `xml:"ParameterValue"`
}

// PendingModifiedValues is undocumented.
type PendingModifiedValues struct {
	CacheNodeIdsToRemove []string `xml:"CacheNodeIdsToRemove>CacheNodeId"`
	EngineVersion        string   `xml:"EngineVersion"`
	NumCacheNodes        int      `xml:"NumCacheNodes"`
}

// PurchaseReservedCacheNodesOfferingMessage is undocumented.
type PurchaseReservedCacheNodesOfferingMessage struct {
	CacheNodeCount               int    `xml:"CacheNodeCount"`
	ReservedCacheNodeID          string `xml:"ReservedCacheNodeId"`
	ReservedCacheNodesOfferingID string `xml:"ReservedCacheNodesOfferingId"`
}

// PurchaseReservedCacheNodesOfferingResult is undocumented.
type PurchaseReservedCacheNodesOfferingResult struct {
	ReservedCacheNode ReservedCacheNode `xml:"PurchaseReservedCacheNodesOfferingResult>ReservedCacheNode"`
}

// RebootCacheClusterMessage is undocumented.
type RebootCacheClusterMessage struct {
	CacheClusterID       string   `xml:"CacheClusterId"`
	CacheNodeIdsToReboot []string `xml:"CacheNodeIdsToReboot>CacheNodeId"`
}

// RebootCacheClusterResult is undocumented.
type RebootCacheClusterResult struct {
	CacheCluster CacheCluster `xml:"RebootCacheClusterResult>CacheCluster"`
}

// RecurringCharge is undocumented.
type RecurringCharge struct {
	RecurringChargeAmount    float64 `xml:"RecurringChargeAmount"`
	RecurringChargeFrequency string  `xml:"RecurringChargeFrequency"`
}

// ReplicationGroup is undocumented.
type ReplicationGroup struct {
	AutomaticFailover     string                                `xml:"AutomaticFailover"`
	Description           string                                `xml:"Description"`
	MemberClusters        []string                              `xml:"MemberClusters>ClusterId"`
	NodeGroups            []NodeGroup                           `xml:"NodeGroups>NodeGroup"`
	PendingModifiedValues ReplicationGroupPendingModifiedValues `xml:"PendingModifiedValues"`
	ReplicationGroupID    string                                `xml:"ReplicationGroupId"`
	SnapshottingClusterID string                                `xml:"SnapshottingClusterId"`
	Status                string                                `xml:"Status"`
}

// ReplicationGroupMessage is undocumented.
type ReplicationGroupMessage struct {
	Marker            string             `xml:"DescribeReplicationGroupsResult>Marker"`
	ReplicationGroups []ReplicationGroup `xml:"DescribeReplicationGroupsResult>ReplicationGroups>ReplicationGroup"`
}

// ReplicationGroupPendingModifiedValues is undocumented.
type ReplicationGroupPendingModifiedValues struct {
	AutomaticFailoverStatus string `xml:"AutomaticFailoverStatus"`
	PrimaryClusterID        string `xml:"PrimaryClusterId"`
}

// ReservedCacheNode is undocumented.
type ReservedCacheNode struct {
	CacheNodeCount               int               `xml:"CacheNodeCount"`
	CacheNodeType                string            `xml:"CacheNodeType"`
	Duration                     int               `xml:"Duration"`
	FixedPrice                   float64           `xml:"FixedPrice"`
	OfferingType                 string            `xml:"OfferingType"`
	ProductDescription           string            `xml:"ProductDescription"`
	RecurringCharges             []RecurringCharge `xml:"RecurringCharges>RecurringCharge"`
	ReservedCacheNodeID          string            `xml:"ReservedCacheNodeId"`
	ReservedCacheNodesOfferingID string            `xml:"ReservedCacheNodesOfferingId"`
	StartTime                    time.Time         `xml:"StartTime"`
	State                        string            `xml:"State"`
	UsagePrice                   float64           `xml:"UsagePrice"`
}

// ReservedCacheNodeMessage is undocumented.
type ReservedCacheNodeMessage struct {
	Marker             string              `xml:"DescribeReservedCacheNodesResult>Marker"`
	ReservedCacheNodes []ReservedCacheNode `xml:"DescribeReservedCacheNodesResult>ReservedCacheNodes>ReservedCacheNode"`
}

// ReservedCacheNodesOffering is undocumented.
type ReservedCacheNodesOffering struct {
	CacheNodeType                string            `xml:"CacheNodeType"`
	Duration                     int               `xml:"Duration"`
	FixedPrice                   float64           `xml:"FixedPrice"`
	OfferingType                 string            `xml:"OfferingType"`
	ProductDescription           string            `xml:"ProductDescription"`
	RecurringCharges             []RecurringCharge `xml:"RecurringCharges>RecurringCharge"`
	ReservedCacheNodesOfferingID string            `xml:"ReservedCacheNodesOfferingId"`
	UsagePrice                   float64           `xml:"UsagePrice"`
}

// ReservedCacheNodesOfferingMessage is undocumented.
type ReservedCacheNodesOfferingMessage struct {
	Marker                      string                       `xml:"DescribeReservedCacheNodesOfferingsResult>Marker"`
	ReservedCacheNodesOfferings []ReservedCacheNodesOffering `xml:"DescribeReservedCacheNodesOfferingsResult>ReservedCacheNodesOfferings>ReservedCacheNodesOffering"`
}

// ResetCacheParameterGroupMessage is undocumented.
type ResetCacheParameterGroupMessage struct {
	CacheParameterGroupName string               `xml:"CacheParameterGroupName"`
	ParameterNameValues     []ParameterNameValue `xml:"ParameterNameValues>ParameterNameValue"`
	ResetAllParameters      bool                 `xml:"ResetAllParameters"`
}

// RevokeCacheSecurityGroupIngressMessage is undocumented.
type RevokeCacheSecurityGroupIngressMessage struct {
	CacheSecurityGroupName  string `xml:"CacheSecurityGroupName"`
	EC2SecurityGroupName    string `xml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerID string `xml:"EC2SecurityGroupOwnerId"`
}

// RevokeCacheSecurityGroupIngressResult is undocumented.
type RevokeCacheSecurityGroupIngressResult struct {
	CacheSecurityGroup CacheSecurityGroup `xml:"RevokeCacheSecurityGroupIngressResult>CacheSecurityGroup"`
}

// SecurityGroupMembership is undocumented.
type SecurityGroupMembership struct {
	SecurityGroupID string `xml:"SecurityGroupId"`
	Status          string `xml:"Status"`
}

// Snapshot is undocumented.
type Snapshot struct {
	AutoMinorVersionUpgrade    bool           `xml:"AutoMinorVersionUpgrade"`
	CacheClusterCreateTime     time.Time      `xml:"CacheClusterCreateTime"`
	CacheClusterID             string         `xml:"CacheClusterId"`
	CacheNodeType              string         `xml:"CacheNodeType"`
	CacheParameterGroupName    string         `xml:"CacheParameterGroupName"`
	CacheSubnetGroupName       string         `xml:"CacheSubnetGroupName"`
	Engine                     string         `xml:"Engine"`
	EngineVersion              string         `xml:"EngineVersion"`
	NodeSnapshots              []NodeSnapshot `xml:"NodeSnapshots>NodeSnapshot"`
	NumCacheNodes              int            `xml:"NumCacheNodes"`
	Port                       int            `xml:"Port"`
	PreferredAvailabilityZone  string         `xml:"PreferredAvailabilityZone"`
	PreferredMaintenanceWindow string         `xml:"PreferredMaintenanceWindow"`
	SnapshotName               string         `xml:"SnapshotName"`
	SnapshotRetentionLimit     int            `xml:"SnapshotRetentionLimit"`
	SnapshotSource             string         `xml:"SnapshotSource"`
	SnapshotStatus             string         `xml:"SnapshotStatus"`
	SnapshotWindow             string         `xml:"SnapshotWindow"`
	TopicARN                   string         `xml:"TopicArn"`
	VpcID                      string         `xml:"VpcId"`
}

// Subnet is undocumented.
type Subnet struct {
	SubnetAvailabilityZone AvailabilityZone `xml:"SubnetAvailabilityZone"`
	SubnetIdentifier       string           `xml:"SubnetIdentifier"`
}

// DescribeCacheClustersResult is a wrapper for CacheClusterMessage.
type DescribeCacheClustersResult struct {
	XMLName xml.Name `xml:"DescribeCacheClustersResponse"`

	CacheClusters []CacheCluster `xml:"DescribeCacheClustersResult>CacheClusters>CacheCluster"`
	Marker        string         `xml:"DescribeCacheClustersResult>Marker"`
}

// DescribeCacheEngineVersionsResult is a wrapper for CacheEngineVersionMessage.
type DescribeCacheEngineVersionsResult struct {
	XMLName xml.Name `xml:"DescribeCacheEngineVersionsResponse"`

	CacheEngineVersions []CacheEngineVersion `xml:"DescribeCacheEngineVersionsResult>CacheEngineVersions>CacheEngineVersion"`
	Marker              string               `xml:"DescribeCacheEngineVersionsResult>Marker"`
}

// DescribeCacheParameterGroupsResult is a wrapper for CacheParameterGroupsMessage.
type DescribeCacheParameterGroupsResult struct {
	XMLName xml.Name `xml:"DescribeCacheParameterGroupsResponse"`

	CacheParameterGroups []CacheParameterGroup `xml:"DescribeCacheParameterGroupsResult>CacheParameterGroups>CacheParameterGroup"`
	Marker               string                `xml:"DescribeCacheParameterGroupsResult>Marker"`
}

// DescribeCacheParametersResult is a wrapper for CacheParameterGroupDetails.
type DescribeCacheParametersResult struct {
	XMLName xml.Name `xml:"DescribeCacheParametersResponse"`

	CacheNodeTypeSpecificParameters []CacheNodeTypeSpecificParameter `xml:"DescribeCacheParametersResult>CacheNodeTypeSpecificParameters>CacheNodeTypeSpecificParameter"`
	Marker                          string                           `xml:"DescribeCacheParametersResult>Marker"`
	Parameters                      []Parameter                      `xml:"DescribeCacheParametersResult>Parameters>Parameter"`
}

// DescribeCacheSecurityGroupsResult is a wrapper for CacheSecurityGroupMessage.
type DescribeCacheSecurityGroupsResult struct {
	XMLName xml.Name `xml:"DescribeCacheSecurityGroupsResponse"`

	CacheSecurityGroups []CacheSecurityGroup `xml:"DescribeCacheSecurityGroupsResult>CacheSecurityGroups>CacheSecurityGroup"`
	Marker              string               `xml:"DescribeCacheSecurityGroupsResult>Marker"`
}

// DescribeCacheSubnetGroupsResult is a wrapper for CacheSubnetGroupMessage.
type DescribeCacheSubnetGroupsResult struct {
	XMLName xml.Name `xml:"DescribeCacheSubnetGroupsResponse"`

	CacheSubnetGroups []CacheSubnetGroup `xml:"DescribeCacheSubnetGroupsResult>CacheSubnetGroups>CacheSubnetGroup"`
	Marker            string             `xml:"DescribeCacheSubnetGroupsResult>Marker"`
}

// DescribeEventsResult is a wrapper for EventsMessage.
type DescribeEventsResult struct {
	XMLName xml.Name `xml:"DescribeEventsResponse"`

	Events []Event `xml:"DescribeEventsResult>Events>Event"`
	Marker string  `xml:"DescribeEventsResult>Marker"`
}

// DescribeReplicationGroupsResult is a wrapper for ReplicationGroupMessage.
type DescribeReplicationGroupsResult struct {
	XMLName xml.Name `xml:"DescribeReplicationGroupsResponse"`

	Marker            string             `xml:"DescribeReplicationGroupsResult>Marker"`
	ReplicationGroups []ReplicationGroup `xml:"DescribeReplicationGroupsResult>ReplicationGroups>ReplicationGroup"`
}

// DescribeReservedCacheNodesOfferingsResult is a wrapper for ReservedCacheNodesOfferingMessage.
type DescribeReservedCacheNodesOfferingsResult struct {
	XMLName xml.Name `xml:"DescribeReservedCacheNodesOfferingsResponse"`

	Marker                      string                       `xml:"DescribeReservedCacheNodesOfferingsResult>Marker"`
	ReservedCacheNodesOfferings []ReservedCacheNodesOffering `xml:"DescribeReservedCacheNodesOfferingsResult>ReservedCacheNodesOfferings>ReservedCacheNodesOffering"`
}

// DescribeReservedCacheNodesResult is a wrapper for ReservedCacheNodeMessage.
type DescribeReservedCacheNodesResult struct {
	XMLName xml.Name `xml:"DescribeReservedCacheNodesResponse"`

	Marker             string              `xml:"DescribeReservedCacheNodesResult>Marker"`
	ReservedCacheNodes []ReservedCacheNode `xml:"DescribeReservedCacheNodesResult>ReservedCacheNodes>ReservedCacheNode"`
}

// DescribeSnapshotsResult is a wrapper for DescribeSnapshotsListMessage.
type DescribeSnapshotsResult struct {
	XMLName xml.Name `xml:"DescribeSnapshotsResponse"`

	Marker    string     `xml:"DescribeSnapshotsResult>Marker"`
	Snapshots []Snapshot `xml:"DescribeSnapshotsResult>Snapshots>Snapshot"`
}

// ModifyCacheParameterGroupResult is a wrapper for CacheParameterGroupNameMessage.
type ModifyCacheParameterGroupResult struct {
	XMLName xml.Name `xml:"Response"`

	CacheParameterGroupName string `xml:"ModifyCacheParameterGroupResult>CacheParameterGroupName"`
}

// ResetCacheParameterGroupResult is a wrapper for CacheParameterGroupNameMessage.
type ResetCacheParameterGroupResult struct {
	XMLName xml.Name `xml:"Response"`

	CacheParameterGroupName string `xml:"ResetCacheParameterGroupResult>CacheParameterGroupName"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
