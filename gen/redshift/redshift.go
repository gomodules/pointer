// Package redshift provides a client for Amazon Redshift.
package redshift

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// RedShift is a client for Amazon Redshift.
type RedShift struct {
	client aws.Client
}

// New returns a new RedShift client.
func New(key, secret, region string, client *http.Client) *RedShift {
	if client == nil {
		client = http.DefaultClient
	}

	return &RedShift{
		client: &aws.QueryClient{
			Client: client,
			Auth: aws.Auth{
				Key:     key,
				Secret:  secret,
				Service: "redshift",
				Region:  region,
			},
			Endpoint:   endpoints.Lookup("redshift", region),
			APIVersion: "2012-12-01",
		},
	}
}

// AuthorizeClusterSecurityGroupIngress adds an inbound (ingress) rule to
// an Amazon Redshift security group. Depending on whether the application
// accessing your cluster is running on the Internet or an EC2 instance,
// you can authorize inbound access to either a Classless Interdomain
// Routing IP address range or an EC2 security group. You can add as many
// as 20 ingress rules to an Amazon Redshift security group. The EC2
// security group must be defined in the AWS region where the cluster
// resides. For an overview of blocks, see the Wikipedia article on
// Classless Inter-Domain Routing . You must also associate the security
// group with a cluster so that clients running on these IP addresses or
// the EC2 instance are authorized to connect to the cluster. For
// information about managing security groups, go to Working with Security
// Groups in the Amazon Redshift Cluster Management Guide
func (c *RedShift) AuthorizeClusterSecurityGroupIngress(req AuthorizeClusterSecurityGroupIngressMessage) (resp *AuthorizeClusterSecurityGroupIngressResult, err error) {
	resp = &AuthorizeClusterSecurityGroupIngressResult{}
	err = c.client.Do("AuthorizeClusterSecurityGroupIngress", "POST", "/", req, resp)
	return
}

// AuthorizeSnapshotAccess authorizes the specified AWS customer account to
// restore the specified snapshot. For more information about working with
// snapshots, go to Amazon Redshift Snapshots in the Amazon Redshift
// Cluster Management Guide .
func (c *RedShift) AuthorizeSnapshotAccess(req AuthorizeSnapshotAccessMessage) (resp *AuthorizeSnapshotAccessResult, err error) {
	resp = &AuthorizeSnapshotAccessResult{}
	err = c.client.Do("AuthorizeSnapshotAccess", "POST", "/", req, resp)
	return
}

// CopyClusterSnapshot copies the specified automated cluster snapshot to a
// new manual cluster snapshot. The source must be an automated snapshot
// and it must be in the available state. When you delete a cluster, Amazon
// Redshift deletes any automated snapshots of the cluster. Also, when the
// retention period of the snapshot expires, Amazon Redshift automatically
// deletes it. If you want to keep an automated snapshot for a longer
// period, you can make a manual copy of the snapshot. Manual snapshots are
// retained until you delete them. For more information about working with
// snapshots, go to Amazon Redshift Snapshots in the Amazon Redshift
// Cluster Management Guide .
func (c *RedShift) CopyClusterSnapshot(req CopyClusterSnapshotMessage) (resp *CopyClusterSnapshotResult, err error) {
	resp = &CopyClusterSnapshotResult{}
	err = c.client.Do("CopyClusterSnapshot", "POST", "/", req, resp)
	return
}

// CreateCluster creates a new cluster. To create the cluster in virtual
// private cloud you must provide cluster subnet group name. If you don't
// provide a cluster subnet group name or the cluster security group
// parameter, Amazon Redshift creates a non-VPC cluster, it associates the
// default cluster security group with the cluster. For more information
// about managing clusters, go to Amazon Redshift Clusters in the Amazon
// Redshift Cluster Management Guide .
func (c *RedShift) CreateCluster(req CreateClusterMessage) (resp *CreateClusterResult, err error) {
	resp = &CreateClusterResult{}
	err = c.client.Do("CreateCluster", "POST", "/", req, resp)
	return
}

// CreateClusterParameterGroup creates an Amazon Redshift parameter group.
// Creating parameter groups is independent of creating clusters. You can
// associate a cluster with a parameter group when you create the cluster.
// You can also associate an existing cluster with a parameter group after
// the cluster is created by using ModifyCluster . Parameters in the
// parameter group define specific behavior that applies to the databases
// you create on the cluster. For more information about managing parameter
// groups, go to Amazon Redshift Parameter Groups in the Amazon Redshift
// Cluster Management Guide .
func (c *RedShift) CreateClusterParameterGroup(req CreateClusterParameterGroupMessage) (resp *CreateClusterParameterGroupResult, err error) {
	resp = &CreateClusterParameterGroupResult{}
	err = c.client.Do("CreateClusterParameterGroup", "POST", "/", req, resp)
	return
}

// CreateClusterSecurityGroup creates a new Amazon Redshift security group.
// You use security groups to control access to non-VPC clusters. For
// information about managing security groups, go to Amazon Redshift
// Cluster Security Groups in the Amazon Redshift Cluster Management Guide
// .
func (c *RedShift) CreateClusterSecurityGroup(req CreateClusterSecurityGroupMessage) (resp *CreateClusterSecurityGroupResult, err error) {
	resp = &CreateClusterSecurityGroupResult{}
	err = c.client.Do("CreateClusterSecurityGroup", "POST", "/", req, resp)
	return
}

// CreateClusterSnapshot creates a manual snapshot of the specified
// cluster. The cluster must be in the available state. For more
// information about working with snapshots, go to Amazon Redshift
// Snapshots in the Amazon Redshift Cluster Management Guide .
func (c *RedShift) CreateClusterSnapshot(req CreateClusterSnapshotMessage) (resp *CreateClusterSnapshotResult, err error) {
	resp = &CreateClusterSnapshotResult{}
	err = c.client.Do("CreateClusterSnapshot", "POST", "/", req, resp)
	return
}

// CreateClusterSubnetGroup creates a new Amazon Redshift subnet group. You
// must provide a list of one or more subnets in your existing Amazon
// Virtual Private Cloud (Amazon when creating Amazon Redshift subnet
// group. For information about subnet groups, go to Amazon Redshift
// Cluster Subnet Groups in the Amazon Redshift Cluster Management Guide .
func (c *RedShift) CreateClusterSubnetGroup(req CreateClusterSubnetGroupMessage) (resp *CreateClusterSubnetGroupResult, err error) {
	resp = &CreateClusterSubnetGroupResult{}
	err = c.client.Do("CreateClusterSubnetGroup", "POST", "/", req, resp)
	return
}

// CreateEventSubscription creates an Amazon Redshift event notification
// subscription. This action requires an ARN (Amazon Resource Name) of an
// Amazon SNS topic created by either the Amazon Redshift console, the
// Amazon SNS console, or the Amazon SNS To obtain an ARN with Amazon you
// must create a topic in Amazon SNS and subscribe to the topic. The ARN is
// displayed in the SNS console. You can specify the source type, and lists
// of Amazon Redshift source IDs, event categories, and event severities.
// Notifications will be sent for all events you want that match those
// criteria. For example, you can specify source type = cluster, source ID
// = my-cluster-1 and mycluster2, event categories = Availability, Backup,
// and severity = The subscription will only send notifications for those
// events in the Availability and Backup categories for the specified
// clusters. If you specify both the source type and source IDs, such as
// source type = cluster and source identifier = my-cluster-1,
// notifications will be sent for all the cluster events for my-cluster-1.
// If you specify a source type but do not specify a source identifier, you
// will receive notice of the events for the objects of that type in your
// AWS account. If you do not specify either the SourceType nor the
// SourceIdentifier, you will be notified of events generated from all
// Amazon Redshift sources belonging to your AWS account. You must specify
// a source type if you specify a source ID.
func (c *RedShift) CreateEventSubscription(req CreateEventSubscriptionMessage) (resp *CreateEventSubscriptionResult, err error) {
	resp = &CreateEventSubscriptionResult{}
	err = c.client.Do("CreateEventSubscription", "POST", "/", req, resp)
	return
}

// CreateHsmClientCertificate creates an HSM client certificate that an
// Amazon Redshift cluster will use to connect to the client's HSM in order
// to store and retrieve the keys used to encrypt the cluster databases.
// The command returns a public key, which you must store in the In
// addition to creating the HSM certificate, you must create an Amazon
// Redshift HSM configuration that provides a cluster the information
// needed to store and use encryption keys in the For more information, go
// to Hardware Security Modules in the Amazon Redshift Cluster Management
// Guide.
func (c *RedShift) CreateHsmClientCertificate(req CreateHsmClientCertificateMessage) (resp *CreateHsmClientCertificateResult, err error) {
	resp = &CreateHsmClientCertificateResult{}
	err = c.client.Do("CreateHsmClientCertificate", "POST", "/", req, resp)
	return
}

// CreateHsmConfiguration creates an HSM configuration that contains the
// information required by an Amazon Redshift cluster to store and use
// database encryption keys in a Hardware Security Module After creating
// the HSM configuration, you can specify it as a parameter when creating a
// cluster. The cluster will then store its encryption keys in the In
// addition to creating an HSM configuration, you must also create an HSM
// client certificate. For more information, go to Hardware Security
// Modules in the Amazon Redshift Cluster Management Guide.
func (c *RedShift) CreateHsmConfiguration(req CreateHsmConfigurationMessage) (resp *CreateHsmConfigurationResult, err error) {
	resp = &CreateHsmConfigurationResult{}
	err = c.client.Do("CreateHsmConfiguration", "POST", "/", req, resp)
	return
}

// CreateTags adds one or more tags to a specified resource. A resource can
// have up to 10 tags. If you try to create more than 10 tags for a
// resource, you will receive an error and the attempt will fail. If you
// specify a key that already exists for the resource, the value for that
// key will be updated with the new value.
func (c *RedShift) CreateTags(req CreateTagsMessage) (err error) {
	// NRE
	err = c.client.Do("CreateTags", "POST", "/", req, nil)
	return
}

// DeleteCluster deletes a previously provisioned cluster. A successful
// response from the web service indicates that the request was received
// correctly. Use DescribeClusters to monitor the status of the deletion.
// The delete operation cannot be canceled or reverted once submitted. For
// more information about managing clusters, go to Amazon Redshift Clusters
// in the Amazon Redshift Cluster Management Guide . If you want to shut
// down the cluster and retain it for future use, set
// SkipFinalClusterSnapshot to false and specify a name for
// FinalClusterSnapshotIdentifier . You can later restore this snapshot to
// resume using the cluster. If a final cluster snapshot is requested, the
// status of the cluster will be "final-snapshot" while the snapshot is
// being taken, then it's "deleting" once Amazon Redshift begins deleting
// the cluster. For more information about managing clusters, go to Amazon
// Redshift Clusters in the Amazon Redshift Cluster Management Guide .
func (c *RedShift) DeleteCluster(req DeleteClusterMessage) (resp *DeleteClusterResult, err error) {
	resp = &DeleteClusterResult{}
	err = c.client.Do("DeleteCluster", "POST", "/", req, resp)
	return
}

// DeleteClusterParameterGroup deletes a specified Amazon Redshift
// parameter group. You cannot delete a parameter group if it is associated
// with a cluster.
func (c *RedShift) DeleteClusterParameterGroup(req DeleteClusterParameterGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteClusterParameterGroup", "POST", "/", req, nil)
	return
}

// DeleteClusterSecurityGroup deletes an Amazon Redshift security group.
// You cannot delete a security group that is associated with any clusters.
// You cannot delete the default security group. For information about
// managing security groups, go to Amazon Redshift Cluster Security Groups
// in the Amazon Redshift Cluster Management Guide .
func (c *RedShift) DeleteClusterSecurityGroup(req DeleteClusterSecurityGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteClusterSecurityGroup", "POST", "/", req, nil)
	return
}

// DeleteClusterSnapshot deletes the specified manual snapshot. The
// snapshot must be in the available state, with no other users authorized
// to access the snapshot. Unlike automated snapshots, manual snapshots are
// retained even after you delete your cluster. Amazon Redshift does not
// delete your manual snapshots. You must delete manual snapshot explicitly
// to avoid getting charged. If other accounts are authorized to access the
// snapshot, you must revoke all of the authorizations before you can
// delete the snapshot.
func (c *RedShift) DeleteClusterSnapshot(req DeleteClusterSnapshotMessage) (resp *DeleteClusterSnapshotResult, err error) {
	resp = &DeleteClusterSnapshotResult{}
	err = c.client.Do("DeleteClusterSnapshot", "POST", "/", req, resp)
	return
}

// DeleteClusterSubnetGroup is undocumented.
func (c *RedShift) DeleteClusterSubnetGroup(req DeleteClusterSubnetGroupMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteClusterSubnetGroup", "POST", "/", req, nil)
	return
}

// DeleteEventSubscription deletes an Amazon Redshift event notification
// subscription.
func (c *RedShift) DeleteEventSubscription(req DeleteEventSubscriptionMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteEventSubscription", "POST", "/", req, nil)
	return
}

// DeleteHsmClientCertificate is undocumented.
func (c *RedShift) DeleteHsmClientCertificate(req DeleteHsmClientCertificateMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteHsmClientCertificate", "POST", "/", req, nil)
	return
}

// DeleteHsmConfiguration deletes the specified Amazon Redshift HSM
// configuration.
func (c *RedShift) DeleteHsmConfiguration(req DeleteHsmConfigurationMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteHsmConfiguration", "POST", "/", req, nil)
	return
}

// DeleteTags deletes a tag or tags from a resource. You must provide the
// ARN of the resource from which you want to delete the tag or tags.
func (c *RedShift) DeleteTags(req DeleteTagsMessage) (err error) {
	// NRE
	err = c.client.Do("DeleteTags", "POST", "/", req, nil)
	return
}

// DescribeClusterParameterGroups returns a list of Amazon Redshift
// parameter groups, including parameter groups you created and the default
// parameter group. For each parameter group, the response includes the
// parameter group name, description, and parameter group family name. You
// can optionally specify a name to retrieve the description of a specific
// parameter group. For more information about managing parameter groups,
// go to Amazon Redshift Parameter Groups in the Amazon Redshift Cluster
// Management Guide . If you specify both tag keys and tag values in the
// same request, Amazon Redshift returns all parameter groups that match
// any combination of the specified keys and values. For example, if you
// have owner and environment for tag keys, and admin and test for tag
// values, all parameter groups that have any combination of those values
// are returned. If both tag keys and values are omitted from the request,
// parameter groups are returned regardless of whether they have tag keys
// or values associated with them.
func (c *RedShift) DescribeClusterParameterGroups(req DescribeClusterParameterGroupsMessage) (resp *DescribeClusterParameterGroupsResult, err error) {
	resp = &DescribeClusterParameterGroupsResult{}
	err = c.client.Do("DescribeClusterParameterGroups", "POST", "/", req, resp)
	return
}

// DescribeClusterParameters returns a detailed list of parameters
// contained within the specified Amazon Redshift parameter group. For each
// parameter the response includes information such as parameter name,
// description, data type, value, whether the parameter value is
// modifiable, and so on. You can specify source filter to retrieve
// parameters of only specific type. For example, to retrieve parameters
// that were modified by a user action such as from
// ModifyClusterParameterGroup , you can specify source equal to user For
// more information about managing parameter groups, go to Amazon Redshift
// Parameter Groups in the Amazon Redshift Cluster Management Guide .
func (c *RedShift) DescribeClusterParameters(req DescribeClusterParametersMessage) (resp *DescribeClusterParametersResult, err error) {
	resp = &DescribeClusterParametersResult{}
	err = c.client.Do("DescribeClusterParameters", "POST", "/", req, resp)
	return
}

// DescribeClusterSecurityGroups returns information about Amazon Redshift
// security groups. If the name of a security group is specified, the
// response will contain only information about only that security group.
// For information about managing security groups, go to Amazon Redshift
// Cluster Security Groups in the Amazon Redshift Cluster Management Guide
// . If you specify both tag keys and tag values in the same request,
// Amazon Redshift returns all security groups that match any combination
// of the specified keys and values. For example, if you have owner and
// environment for tag keys, and admin and test for tag values, all
// security groups that have any combination of those values are returned.
// If both tag keys and values are omitted from the request, security
// groups are returned regardless of whether they have tag keys or values
// associated with them.
func (c *RedShift) DescribeClusterSecurityGroups(req DescribeClusterSecurityGroupsMessage) (resp *DescribeClusterSecurityGroupsResult, err error) {
	resp = &DescribeClusterSecurityGroupsResult{}
	err = c.client.Do("DescribeClusterSecurityGroups", "POST", "/", req, resp)
	return
}

// DescribeClusterSnapshots returns one or more snapshot objects, which
// contain metadata about your cluster snapshots. By default, this
// operation returns information about all snapshots of all clusters that
// are owned by you AWS customer account. No information is returned for
// snapshots owned by inactive AWS customer accounts. If you specify both
// tag keys and tag values in the same request, Amazon Redshift returns all
// snapshots that match any combination of the specified keys and values.
// For example, if you have owner and environment for tag keys, and admin
// and test for tag values, all snapshots that have any combination of
// those values are returned. Only snapshots that you own are returned in
// the response; shared snapshots are not returned with the tag key and tag
// value request parameters. If both tag keys and values are omitted from
// the request, snapshots are returned regardless of whether they have tag
// keys or values associated with them.
func (c *RedShift) DescribeClusterSnapshots(req DescribeClusterSnapshotsMessage) (resp *DescribeClusterSnapshotsResult, err error) {
	resp = &DescribeClusterSnapshotsResult{}
	err = c.client.Do("DescribeClusterSnapshots", "POST", "/", req, resp)
	return
}

// DescribeClusterSubnetGroups returns one or more cluster subnet group
// objects, which contain metadata about your cluster subnet groups. By
// default, this operation returns information about all cluster subnet
// groups that are defined in you AWS account. If you specify both tag keys
// and tag values in the same request, Amazon Redshift returns all subnet
// groups that match any combination of the specified keys and values. For
// example, if you have owner and environment for tag keys, and admin and
// test for tag values, all subnet groups that have any combination of
// those values are returned. If both tag keys and values are omitted from
// the request, subnet groups are returned regardless of whether they have
// tag keys or values associated with them.
func (c *RedShift) DescribeClusterSubnetGroups(req DescribeClusterSubnetGroupsMessage) (resp *DescribeClusterSubnetGroupsResult, err error) {
	resp = &DescribeClusterSubnetGroupsResult{}
	err = c.client.Do("DescribeClusterSubnetGroups", "POST", "/", req, resp)
	return
}

// DescribeClusterVersions returns descriptions of the available Amazon
// Redshift cluster versions. You can call this operation even before
// creating any clusters to learn more about the Amazon Redshift versions.
// For more information about managing clusters, go to Amazon Redshift
// Clusters in the Amazon Redshift Cluster Management Guide
func (c *RedShift) DescribeClusterVersions(req DescribeClusterVersionsMessage) (resp *DescribeClusterVersionsResult, err error) {
	resp = &DescribeClusterVersionsResult{}
	err = c.client.Do("DescribeClusterVersions", "POST", "/", req, resp)
	return
}

// DescribeClusters returns properties of provisioned clusters including
// general cluster properties, cluster database properties, maintenance and
// backup properties, and security and access properties. This operation
// supports pagination. For more information about managing clusters, go to
// Amazon Redshift Clusters in the Amazon Redshift Cluster Management Guide
// . If you specify both tag keys and tag values in the same request,
// Amazon Redshift returns all clusters that match any combination of the
// specified keys and values. For example, if you have owner and
// environment for tag keys, and admin and test for tag values, all
// clusters that have any combination of those values are returned. If both
// tag keys and values are omitted from the request, clusters are returned
// regardless of whether they have tag keys or values associated with them.
func (c *RedShift) DescribeClusters(req DescribeClustersMessage) (resp *DescribeClustersResult, err error) {
	resp = &DescribeClustersResult{}
	err = c.client.Do("DescribeClusters", "POST", "/", req, resp)
	return
}

// DescribeDefaultClusterParameters returns a list of parameter settings
// for the specified parameter group family. For more information about
// managing parameter groups, go to Amazon Redshift Parameter Groups in the
// Amazon Redshift Cluster Management Guide .
func (c *RedShift) DescribeDefaultClusterParameters(req DescribeDefaultClusterParametersMessage) (resp *DescribeDefaultClusterParametersResult, err error) {
	resp = &DescribeDefaultClusterParametersResult{}
	err = c.client.Do("DescribeDefaultClusterParameters", "POST", "/", req, resp)
	return
}

// DescribeEventCategories displays a list of event categories for all
// event source types, or for a specified source type. For a list of the
// event categories and source types, go to Amazon Redshift Event
// Notifications
func (c *RedShift) DescribeEventCategories(req DescribeEventCategoriesMessage) (resp *DescribeEventCategoriesResult, err error) {
	resp = &DescribeEventCategoriesResult{}
	err = c.client.Do("DescribeEventCategories", "POST", "/", req, resp)
	return
}

// DescribeEventSubscriptions lists descriptions of all the Amazon Redshift
// event notifications subscription for a customer account. If you specify
// a subscription name, lists the description for that subscription.
func (c *RedShift) DescribeEventSubscriptions(req DescribeEventSubscriptionsMessage) (resp *DescribeEventSubscriptionsResult, err error) {
	resp = &DescribeEventSubscriptionsResult{}
	err = c.client.Do("DescribeEventSubscriptions", "POST", "/", req, resp)
	return
}

// DescribeEvents returns events related to clusters, security groups,
// snapshots, and parameter groups for the past 14 days. Events specific to
// a particular cluster, security group, snapshot or parameter group can be
// obtained by providing the name as a parameter. By default, the past hour
// of events are returned.
func (c *RedShift) DescribeEvents(req DescribeEventsMessage) (resp *DescribeEventsResult, err error) {
	resp = &DescribeEventsResult{}
	err = c.client.Do("DescribeEvents", "POST", "/", req, resp)
	return
}

// DescribeHsmClientCertificates returns information about the specified
// HSM client certificate. If no certificate ID is specified, returns
// information about all the HSM certificates owned by your AWS customer
// account. If you specify both tag keys and tag values in the same
// request, Amazon Redshift returns all HSM client certificates that match
// any combination of the specified keys and values. For example, if you
// have owner and environment for tag keys, and admin and test for tag
// values, all HSM client certificates that have any combination of those
// values are returned. If both tag keys and values are omitted from the
// request, HSM client certificates are returned regardless of whether they
// have tag keys or values associated with them.
func (c *RedShift) DescribeHsmClientCertificates(req DescribeHsmClientCertificatesMessage) (resp *DescribeHsmClientCertificatesResult, err error) {
	resp = &DescribeHsmClientCertificatesResult{}
	err = c.client.Do("DescribeHsmClientCertificates", "POST", "/", req, resp)
	return
}

// DescribeHsmConfigurations returns information about the specified Amazon
// Redshift HSM configuration. If no configuration ID is specified, returns
// information about all the HSM configurations owned by your AWS customer
// account. If you specify both tag keys and tag values in the same
// request, Amazon Redshift returns all HSM connections that match any
// combination of the specified keys and values. For example, if you have
// owner and environment for tag keys, and admin and test for tag values,
// all HSM connections that have any combination of those values are
// returned. If both tag keys and values are omitted from the request, HSM
// connections are returned regardless of whether they have tag keys or
// values associated with them.
func (c *RedShift) DescribeHsmConfigurations(req DescribeHsmConfigurationsMessage) (resp *DescribeHsmConfigurationsResult, err error) {
	resp = &DescribeHsmConfigurationsResult{}
	err = c.client.Do("DescribeHsmConfigurations", "POST", "/", req, resp)
	return
}

// DescribeLoggingStatus describes whether information, such as queries and
// connection attempts, is being logged for the specified Amazon Redshift
// cluster.
func (c *RedShift) DescribeLoggingStatus(req DescribeLoggingStatusMessage) (resp *DescribeLoggingStatusResult, err error) {
	resp = &DescribeLoggingStatusResult{}
	err = c.client.Do("DescribeLoggingStatus", "POST", "/", req, resp)
	return
}

// DescribeOrderableClusterOptions returns a list of orderable cluster
// options. Before you create a new cluster you can use this operation to
// find what options are available, such as the EC2 Availability Zones in
// the specific AWS region that you can specify, and the node types you can
// request. The node types differ by available storage, memory, CPU and
// price. With the cost involved you might want to obtain a list of cluster
// options in the specific region and specify values when creating a
// cluster. For more information about managing clusters, go to Amazon
// Redshift Clusters in the Amazon Redshift Cluster Management Guide
func (c *RedShift) DescribeOrderableClusterOptions(req DescribeOrderableClusterOptionsMessage) (resp *DescribeOrderableClusterOptionsResult, err error) {
	resp = &DescribeOrderableClusterOptionsResult{}
	err = c.client.Do("DescribeOrderableClusterOptions", "POST", "/", req, resp)
	return
}

// DescribeReservedNodeOfferings returns a list of the available reserved
// node offerings by Amazon Redshift with their descriptions including the
// node type, the fixed and recurring costs of reserving the node and
// duration the node will be reserved for you. These descriptions help you
// determine which reserve node offering you want to purchase. You then use
// the unique offering ID in you call to PurchaseReservedNodeOffering to
// reserve one or more nodes for your Amazon Redshift cluster. For more
// information about managing parameter groups, go to Purchasing Reserved
// Nodes in the Amazon Redshift Cluster Management Guide .
func (c *RedShift) DescribeReservedNodeOfferings(req DescribeReservedNodeOfferingsMessage) (resp *DescribeReservedNodeOfferingsResult, err error) {
	resp = &DescribeReservedNodeOfferingsResult{}
	err = c.client.Do("DescribeReservedNodeOfferings", "POST", "/", req, resp)
	return
}

// DescribeReservedNodes is undocumented.
func (c *RedShift) DescribeReservedNodes(req DescribeReservedNodesMessage) (resp *DescribeReservedNodesResult, err error) {
	resp = &DescribeReservedNodesResult{}
	err = c.client.Do("DescribeReservedNodes", "POST", "/", req, resp)
	return
}

// DescribeResize returns information about the last resize operation for
// the specified cluster. If no resize operation has ever been initiated
// for the specified cluster, a 404 error is returned. If a resize
// operation was initiated and completed, the status of the resize remains
// as until the next resize. A resize operation can be requested using
// ModifyCluster and specifying a different number or type of nodes for the
// cluster.
func (c *RedShift) DescribeResize(req DescribeResizeMessage) (resp *DescribeResizeResult, err error) {
	resp = &DescribeResizeResult{}
	err = c.client.Do("DescribeResize", "POST", "/", req, resp)
	return
}

// DescribeTags returns a list of tags. You can return tags from a specific
// resource by specifying an or you can return all tags for a given type of
// resource, such as clusters, snapshots, and so on. The following are
// limitations for DescribeTags : You cannot specify an ARN and a
// resource-type value together in the same request. You cannot use the
// MaxRecords and Marker parameters together with the ARN parameter. The
// MaxRecords parameter can be a range from 10 to 50 results to return in a
// request. If you specify both tag keys and tag values in the same
// request, Amazon Redshift returns all resources that match any
// combination of the specified keys and values. For example, if you have
// owner and environment for tag keys, and admin and test for tag values,
// all resources that have any combination of those values are returned. If
// both tag keys and values are omitted from the request, resources are
// returned regardless of whether they have tag keys or values associated
// with them.
func (c *RedShift) DescribeTags(req DescribeTagsMessage) (resp *DescribeTagsResult, err error) {
	resp = &DescribeTagsResult{}
	err = c.client.Do("DescribeTags", "POST", "/", req, resp)
	return
}

// DisableLogging stops logging information, such as queries and connection
// attempts, for the specified Amazon Redshift cluster.
func (c *RedShift) DisableLogging(req DisableLoggingMessage) (resp *DisableLoggingResult, err error) {
	resp = &DisableLoggingResult{}
	err = c.client.Do("DisableLogging", "POST", "/", req, resp)
	return
}

// DisableSnapshotCopy disables the automatic copying of snapshots from one
// region to another region for a specified cluster.
func (c *RedShift) DisableSnapshotCopy(req DisableSnapshotCopyMessage) (resp *DisableSnapshotCopyResult, err error) {
	resp = &DisableSnapshotCopyResult{}
	err = c.client.Do("DisableSnapshotCopy", "POST", "/", req, resp)
	return
}

// EnableLogging starts logging information, such as queries and connection
// attempts, for the specified Amazon Redshift cluster.
func (c *RedShift) EnableLogging(req EnableLoggingMessage) (resp *EnableLoggingResult, err error) {
	resp = &EnableLoggingResult{}
	err = c.client.Do("EnableLogging", "POST", "/", req, resp)
	return
}

// EnableSnapshotCopy enables the automatic copy of snapshots from one
// region to another region for a specified cluster.
func (c *RedShift) EnableSnapshotCopy(req EnableSnapshotCopyMessage) (resp *EnableSnapshotCopyResult, err error) {
	resp = &EnableSnapshotCopyResult{}
	err = c.client.Do("EnableSnapshotCopy", "POST", "/", req, resp)
	return
}

// ModifyCluster modifies the settings for a cluster. For example, you can
// add another security or parameter group, update the preferred
// maintenance window, or change the master user password. Resetting a
// cluster password or modifying the security groups associated with a
// cluster do not need a reboot. However, modifying a parameter group
// requires a reboot for parameters to take effect. For more information
// about managing clusters, go to Amazon Redshift Clusters in the Amazon
// Redshift Cluster Management Guide . You can also change node type and
// the number of nodes to scale up or down the cluster. When resizing a
// cluster, you must specify both the number of nodes and the node type
// even if one of the parameters does not change.
func (c *RedShift) ModifyCluster(req ModifyClusterMessage) (resp *ModifyClusterResult, err error) {
	resp = &ModifyClusterResult{}
	err = c.client.Do("ModifyCluster", "POST", "/", req, resp)
	return
}

// ModifyClusterParameterGroup modifies the parameters of a parameter
// group. For more information about managing parameter groups, go to
// Amazon Redshift Parameter Groups in the Amazon Redshift Cluster
// Management Guide .
func (c *RedShift) ModifyClusterParameterGroup(req ModifyClusterParameterGroupMessage) (resp *ModifyClusterParameterGroupResult, err error) {
	resp = &ModifyClusterParameterGroupResult{}
	err = c.client.Do("ModifyClusterParameterGroup", "POST", "/", req, resp)
	return
}

// ModifyClusterSubnetGroup modifies a cluster subnet group to include the
// specified list of VPC subnets. The operation replaces the existing list
// of subnets with the new list of subnets.
func (c *RedShift) ModifyClusterSubnetGroup(req ModifyClusterSubnetGroupMessage) (resp *ModifyClusterSubnetGroupResult, err error) {
	resp = &ModifyClusterSubnetGroupResult{}
	err = c.client.Do("ModifyClusterSubnetGroup", "POST", "/", req, resp)
	return
}

// ModifyEventSubscription modifies an existing Amazon Redshift event
// notification subscription.
func (c *RedShift) ModifyEventSubscription(req ModifyEventSubscriptionMessage) (resp *ModifyEventSubscriptionResult, err error) {
	resp = &ModifyEventSubscriptionResult{}
	err = c.client.Do("ModifyEventSubscription", "POST", "/", req, resp)
	return
}

// ModifySnapshotCopyRetentionPeriod modifies the number of days to retain
// automated snapshots in the destination region after they are copied from
// the source region.
func (c *RedShift) ModifySnapshotCopyRetentionPeriod(req ModifySnapshotCopyRetentionPeriodMessage) (resp *ModifySnapshotCopyRetentionPeriodResult, err error) {
	resp = &ModifySnapshotCopyRetentionPeriodResult{}
	err = c.client.Do("ModifySnapshotCopyRetentionPeriod", "POST", "/", req, resp)
	return
}

// PurchaseReservedNodeOffering allows you to purchase reserved nodes.
// Amazon Redshift offers a predefined set of reserved node offerings. You
// can purchase one of the offerings. You can call the
// DescribeReservedNodeOfferings API to obtain the available reserved node
// offerings. You can call this API by providing a specific reserved node
// offering and the number of nodes you want to reserve. For more
// information about managing parameter groups, go to Purchasing Reserved
// Nodes in the Amazon Redshift Cluster Management Guide .
func (c *RedShift) PurchaseReservedNodeOffering(req PurchaseReservedNodeOfferingMessage) (resp *PurchaseReservedNodeOfferingResult, err error) {
	resp = &PurchaseReservedNodeOfferingResult{}
	err = c.client.Do("PurchaseReservedNodeOffering", "POST", "/", req, resp)
	return
}

// RebootCluster reboots a cluster. This action is taken as soon as
// possible. It results in a momentary outage to the cluster, during which
// the cluster status is set to rebooting . A cluster event is created when
// the reboot is completed. Any pending cluster modifications (see
// ModifyCluster ) are applied at this reboot. For more information about
// managing clusters, go to Amazon Redshift Clusters in the Amazon Redshift
// Cluster Management Guide
func (c *RedShift) RebootCluster(req RebootClusterMessage) (resp *RebootClusterResult, err error) {
	resp = &RebootClusterResult{}
	err = c.client.Do("RebootCluster", "POST", "/", req, resp)
	return
}

// ResetClusterParameterGroup sets one or more parameters of the specified
// parameter group to their default values and sets the source values of
// the parameters to "engine-default". To reset the entire parameter group
// specify the ResetAllParameters parameter. For parameter changes to take
// effect you must reboot any associated clusters.
func (c *RedShift) ResetClusterParameterGroup(req ResetClusterParameterGroupMessage) (resp *ResetClusterParameterGroupResult, err error) {
	resp = &ResetClusterParameterGroupResult{}
	err = c.client.Do("ResetClusterParameterGroup", "POST", "/", req, resp)
	return
}

// RestoreFromClusterSnapshot creates a new cluster from a snapshot. Amazon
// Redshift creates the resulting cluster with the same configuration as
// the original cluster from which the snapshot was created, except that
// the new cluster is created with the default cluster security and
// parameter group. After Amazon Redshift creates the cluster you can use
// the ModifyCluster API to associate a different security group and
// different parameter group with the restored cluster. If you restore a
// cluster into a you must provide a cluster subnet group where you want
// the cluster restored. For more information about working with snapshots,
// go to Amazon Redshift Snapshots in the Amazon Redshift Cluster
// Management Guide .
func (c *RedShift) RestoreFromClusterSnapshot(req RestoreFromClusterSnapshotMessage) (resp *RestoreFromClusterSnapshotResult, err error) {
	resp = &RestoreFromClusterSnapshotResult{}
	err = c.client.Do("RestoreFromClusterSnapshot", "POST", "/", req, resp)
	return
}

// RevokeClusterSecurityGroupIngress revokes an ingress rule in an Amazon
// Redshift security group for a previously authorized IP range or Amazon
// EC2 security group. To add an ingress rule, see
// AuthorizeClusterSecurityGroupIngress . For information about managing
// security groups, go to Amazon Redshift Cluster Security Groups in the
// Amazon Redshift Cluster Management Guide .
func (c *RedShift) RevokeClusterSecurityGroupIngress(req RevokeClusterSecurityGroupIngressMessage) (resp *RevokeClusterSecurityGroupIngressResult, err error) {
	resp = &RevokeClusterSecurityGroupIngressResult{}
	err = c.client.Do("RevokeClusterSecurityGroupIngress", "POST", "/", req, resp)
	return
}

// RevokeSnapshotAccess removes the ability of the specified AWS customer
// account to restore the specified snapshot. If the account is currently
// restoring the snapshot, the restore will run to completion. For more
// information about working with snapshots, go to Amazon Redshift
// Snapshots in the Amazon Redshift Cluster Management Guide .
func (c *RedShift) RevokeSnapshotAccess(req RevokeSnapshotAccessMessage) (resp *RevokeSnapshotAccessResult, err error) {
	resp = &RevokeSnapshotAccessResult{}
	err = c.client.Do("RevokeSnapshotAccess", "POST", "/", req, resp)
	return
}

// RotateEncryptionKey is undocumented.
func (c *RedShift) RotateEncryptionKey(req RotateEncryptionKeyMessage) (resp *RotateEncryptionKeyResult, err error) {
	resp = &RotateEncryptionKeyResult{}
	err = c.client.Do("RotateEncryptionKey", "POST", "/", req, resp)
	return
}

// AccountWithRestoreAccess is undocumented.
type AccountWithRestoreAccess struct {
	AccountID string `xml:"AccountId"`
}

// AuthorizeClusterSecurityGroupIngressMessage is undocumented.
type AuthorizeClusterSecurityGroupIngressMessage struct {
	CIDRIP                   string `xml:"CIDRIP"`
	ClusterSecurityGroupName string `xml:"ClusterSecurityGroupName"`
	EC2SecurityGroupName     string `xml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerID  string `xml:"EC2SecurityGroupOwnerId"`
}

// AuthorizeClusterSecurityGroupIngressResult is undocumented.
type AuthorizeClusterSecurityGroupIngressResult struct {
	ClusterSecurityGroup ClusterSecurityGroup `xml:"AuthorizeClusterSecurityGroupIngressResult>ClusterSecurityGroup"`
}

// AuthorizeSnapshotAccessMessage is undocumented.
type AuthorizeSnapshotAccessMessage struct {
	AccountWithRestoreAccess  string `xml:"AccountWithRestoreAccess"`
	SnapshotClusterIdentifier string `xml:"SnapshotClusterIdentifier"`
	SnapshotIdentifier        string `xml:"SnapshotIdentifier"`
}

// AuthorizeSnapshotAccessResult is undocumented.
type AuthorizeSnapshotAccessResult struct {
	Snapshot Snapshot `xml:"AuthorizeSnapshotAccessResult>Snapshot"`
}

// AvailabilityZone is undocumented.
type AvailabilityZone struct {
	Name string `xml:"Name"`
}

// Cluster is undocumented.
type Cluster struct {
	AllowVersionUpgrade              bool                             `xml:"AllowVersionUpgrade"`
	AutomatedSnapshotRetentionPeriod int                              `xml:"AutomatedSnapshotRetentionPeriod"`
	AvailabilityZone                 string                           `xml:"AvailabilityZone"`
	ClusterCreateTime                time.Time                        `xml:"ClusterCreateTime"`
	ClusterIdentifier                string                           `xml:"ClusterIdentifier"`
	ClusterNodes                     []ClusterNode                    `xml:"ClusterNodes>member"`
	ClusterParameterGroups           []ClusterParameterGroupStatus    `xml:"ClusterParameterGroups>ClusterParameterGroup"`
	ClusterPublicKey                 string                           `xml:"ClusterPublicKey"`
	ClusterRevisionNumber            string                           `xml:"ClusterRevisionNumber"`
	ClusterSecurityGroups            []ClusterSecurityGroupMembership `xml:"ClusterSecurityGroups>ClusterSecurityGroup"`
	ClusterSnapshotCopyStatus        ClusterSnapshotCopyStatus        `xml:"ClusterSnapshotCopyStatus"`
	ClusterStatus                    string                           `xml:"ClusterStatus"`
	ClusterSubnetGroupName           string                           `xml:"ClusterSubnetGroupName"`
	ClusterVersion                   string                           `xml:"ClusterVersion"`
	DBName                           string                           `xml:"DBName"`
	ElasticIPStatus                  ElasticIPStatus                  `xml:"ElasticIpStatus"`
	Encrypted                        bool                             `xml:"Encrypted"`
	Endpoint                         Endpoint                         `xml:"Endpoint"`
	HsmStatus                        HsmStatus                        `xml:"HsmStatus"`
	KmsKeyID                         string                           `xml:"KmsKeyId"`
	MasterUsername                   string                           `xml:"MasterUsername"`
	ModifyStatus                     string                           `xml:"ModifyStatus"`
	NodeType                         string                           `xml:"NodeType"`
	NumberOfNodes                    int                              `xml:"NumberOfNodes"`
	PendingModifiedValues            PendingModifiedValues            `xml:"PendingModifiedValues"`
	PreferredMaintenanceWindow       string                           `xml:"PreferredMaintenanceWindow"`
	PubliclyAccessible               bool                             `xml:"PubliclyAccessible"`
	RestoreStatus                    RestoreStatus                    `xml:"RestoreStatus"`
	Tags                             []Tag                            `xml:"Tags>Tag"`
	VpcID                            string                           `xml:"VpcId"`
	VpcSecurityGroups                []VpcSecurityGroupMembership     `xml:"VpcSecurityGroups>VpcSecurityGroup"`
}

// ClusterNode is undocumented.
type ClusterNode struct {
	NodeRole         string `xml:"NodeRole"`
	PrivateIPAddress string `xml:"PrivateIPAddress"`
	PublicIPAddress  string `xml:"PublicIPAddress"`
}

// ClusterParameterGroup is undocumented.
type ClusterParameterGroup struct {
	Description          string `xml:"Description"`
	ParameterGroupFamily string `xml:"ParameterGroupFamily"`
	ParameterGroupName   string `xml:"ParameterGroupName"`
	Tags                 []Tag  `xml:"Tags>Tag"`
}

// ClusterParameterGroupDetails is undocumented.
type ClusterParameterGroupDetails struct {
	Marker     string      `xml:"DescribeClusterParametersResult>Marker"`
	Parameters []Parameter `xml:"DescribeClusterParametersResult>Parameters>Parameter"`
}

// ClusterParameterGroupNameMessage is undocumented.
type ClusterParameterGroupNameMessage struct {
	ParameterGroupName   string `xml:"ParameterGroupName"`
	ParameterGroupStatus string `xml:"ParameterGroupStatus"`
}

// ClusterParameterGroupStatus is undocumented.
type ClusterParameterGroupStatus struct {
	ParameterApplyStatus string `xml:"ParameterApplyStatus"`
	ParameterGroupName   string `xml:"ParameterGroupName"`
}

// ClusterParameterGroupsMessage is undocumented.
type ClusterParameterGroupsMessage struct {
	Marker          string                  `xml:"DescribeClusterParameterGroupsResult>Marker"`
	ParameterGroups []ClusterParameterGroup `xml:"DescribeClusterParameterGroupsResult>ParameterGroups>ClusterParameterGroup"`
}

// ClusterSecurityGroup is undocumented.
type ClusterSecurityGroup struct {
	ClusterSecurityGroupName string             `xml:"ClusterSecurityGroupName"`
	Description              string             `xml:"Description"`
	EC2SecurityGroups        []EC2SecurityGroup `xml:"EC2SecurityGroups>EC2SecurityGroup"`
	IPRanges                 []IPRange          `xml:"IPRanges>IPRange"`
	Tags                     []Tag              `xml:"Tags>Tag"`
}

// ClusterSecurityGroupMembership is undocumented.
type ClusterSecurityGroupMembership struct {
	ClusterSecurityGroupName string `xml:"ClusterSecurityGroupName"`
	Status                   string `xml:"Status"`
}

// ClusterSecurityGroupMessage is undocumented.
type ClusterSecurityGroupMessage struct {
	ClusterSecurityGroups []ClusterSecurityGroup `xml:"DescribeClusterSecurityGroupsResult>ClusterSecurityGroups>ClusterSecurityGroup"`
	Marker                string                 `xml:"DescribeClusterSecurityGroupsResult>Marker"`
}

// ClusterSnapshotCopyStatus is undocumented.
type ClusterSnapshotCopyStatus struct {
	DestinationRegion string `xml:"DestinationRegion"`
	RetentionPeriod   int    `xml:"RetentionPeriod"`
}

// ClusterSubnetGroup is undocumented.
type ClusterSubnetGroup struct {
	ClusterSubnetGroupName string   `xml:"ClusterSubnetGroupName"`
	Description            string   `xml:"Description"`
	SubnetGroupStatus      string   `xml:"SubnetGroupStatus"`
	Subnets                []Subnet `xml:"Subnets>Subnet"`
	Tags                   []Tag    `xml:"Tags>Tag"`
	VpcID                  string   `xml:"VpcId"`
}

// ClusterSubnetGroupMessage is undocumented.
type ClusterSubnetGroupMessage struct {
	ClusterSubnetGroups []ClusterSubnetGroup `xml:"DescribeClusterSubnetGroupsResult>ClusterSubnetGroups>ClusterSubnetGroup"`
	Marker              string               `xml:"DescribeClusterSubnetGroupsResult>Marker"`
}

// ClusterVersion is undocumented.
type ClusterVersion struct {
	ClusterParameterGroupFamily string `xml:"ClusterParameterGroupFamily"`
	ClusterVersion              string `xml:"ClusterVersion"`
	Description                 string `xml:"Description"`
}

// ClusterVersionsMessage is undocumented.
type ClusterVersionsMessage struct {
	ClusterVersions []ClusterVersion `xml:"DescribeClusterVersionsResult>ClusterVersions>ClusterVersion"`
	Marker          string           `xml:"DescribeClusterVersionsResult>Marker"`
}

// ClustersMessage is undocumented.
type ClustersMessage struct {
	Clusters []Cluster `xml:"DescribeClustersResult>Clusters>Cluster"`
	Marker   string    `xml:"DescribeClustersResult>Marker"`
}

// CopyClusterSnapshotMessage is undocumented.
type CopyClusterSnapshotMessage struct {
	SourceSnapshotClusterIdentifier string `xml:"SourceSnapshotClusterIdentifier"`
	SourceSnapshotIdentifier        string `xml:"SourceSnapshotIdentifier"`
	TargetSnapshotIdentifier        string `xml:"TargetSnapshotIdentifier"`
}

// CopyClusterSnapshotResult is undocumented.
type CopyClusterSnapshotResult struct {
	Snapshot Snapshot `xml:"CopyClusterSnapshotResult>Snapshot"`
}

// CreateClusterMessage is undocumented.
type CreateClusterMessage struct {
	AllowVersionUpgrade              bool     `xml:"AllowVersionUpgrade"`
	AutomatedSnapshotRetentionPeriod int      `xml:"AutomatedSnapshotRetentionPeriod"`
	AvailabilityZone                 string   `xml:"AvailabilityZone"`
	ClusterIdentifier                string   `xml:"ClusterIdentifier"`
	ClusterParameterGroupName        string   `xml:"ClusterParameterGroupName"`
	ClusterSecurityGroups            []string `xml:"ClusterSecurityGroups>ClusterSecurityGroupName"`
	ClusterSubnetGroupName           string   `xml:"ClusterSubnetGroupName"`
	ClusterType                      string   `xml:"ClusterType"`
	ClusterVersion                   string   `xml:"ClusterVersion"`
	DBName                           string   `xml:"DBName"`
	ElasticIP                        string   `xml:"ElasticIp"`
	Encrypted                        bool     `xml:"Encrypted"`
	HsmClientCertificateIdentifier   string   `xml:"HsmClientCertificateIdentifier"`
	HsmConfigurationIdentifier       string   `xml:"HsmConfigurationIdentifier"`
	KmsKeyID                         string   `xml:"KmsKeyId"`
	MasterUserPassword               string   `xml:"MasterUserPassword"`
	MasterUsername                   string   `xml:"MasterUsername"`
	NodeType                         string   `xml:"NodeType"`
	NumberOfNodes                    int      `xml:"NumberOfNodes"`
	Port                             int      `xml:"Port"`
	PreferredMaintenanceWindow       string   `xml:"PreferredMaintenanceWindow"`
	PubliclyAccessible               bool     `xml:"PubliclyAccessible"`
	Tags                             []Tag    `xml:"Tags>Tag"`
	VpcSecurityGroupIds              []string `xml:"VpcSecurityGroupIds>VpcSecurityGroupId"`
}

// CreateClusterParameterGroupMessage is undocumented.
type CreateClusterParameterGroupMessage struct {
	Description          string `xml:"Description"`
	ParameterGroupFamily string `xml:"ParameterGroupFamily"`
	ParameterGroupName   string `xml:"ParameterGroupName"`
	Tags                 []Tag  `xml:"Tags>Tag"`
}

// CreateClusterParameterGroupResult is undocumented.
type CreateClusterParameterGroupResult struct {
	ClusterParameterGroup ClusterParameterGroup `xml:"CreateClusterParameterGroupResult>ClusterParameterGroup"`
}

// CreateClusterResult is undocumented.
type CreateClusterResult struct {
	Cluster Cluster `xml:"CreateClusterResult>Cluster"`
}

// CreateClusterSecurityGroupMessage is undocumented.
type CreateClusterSecurityGroupMessage struct {
	ClusterSecurityGroupName string `xml:"ClusterSecurityGroupName"`
	Description              string `xml:"Description"`
	Tags                     []Tag  `xml:"Tags>Tag"`
}

// CreateClusterSecurityGroupResult is undocumented.
type CreateClusterSecurityGroupResult struct {
	ClusterSecurityGroup ClusterSecurityGroup `xml:"CreateClusterSecurityGroupResult>ClusterSecurityGroup"`
}

// CreateClusterSnapshotMessage is undocumented.
type CreateClusterSnapshotMessage struct {
	ClusterIdentifier  string `xml:"ClusterIdentifier"`
	SnapshotIdentifier string `xml:"SnapshotIdentifier"`
	Tags               []Tag  `xml:"Tags>Tag"`
}

// CreateClusterSnapshotResult is undocumented.
type CreateClusterSnapshotResult struct {
	Snapshot Snapshot `xml:"CreateClusterSnapshotResult>Snapshot"`
}

// CreateClusterSubnetGroupMessage is undocumented.
type CreateClusterSubnetGroupMessage struct {
	ClusterSubnetGroupName string   `xml:"ClusterSubnetGroupName"`
	Description            string   `xml:"Description"`
	SubnetIds              []string `xml:"SubnetIds>SubnetIdentifier"`
	Tags                   []Tag    `xml:"Tags>Tag"`
}

// CreateClusterSubnetGroupResult is undocumented.
type CreateClusterSubnetGroupResult struct {
	ClusterSubnetGroup ClusterSubnetGroup `xml:"CreateClusterSubnetGroupResult>ClusterSubnetGroup"`
}

// CreateEventSubscriptionMessage is undocumented.
type CreateEventSubscriptionMessage struct {
	Enabled          bool     `xml:"Enabled"`
	EventCategories  []string `xml:"EventCategories>EventCategory"`
	Severity         string   `xml:"Severity"`
	SnsTopicARN      string   `xml:"SnsTopicArn"`
	SourceIds        []string `xml:"SourceIds>SourceId"`
	SourceType       string   `xml:"SourceType"`
	SubscriptionName string   `xml:"SubscriptionName"`
	Tags             []Tag    `xml:"Tags>Tag"`
}

// CreateEventSubscriptionResult is undocumented.
type CreateEventSubscriptionResult struct {
	EventSubscription EventSubscription `xml:"CreateEventSubscriptionResult>EventSubscription"`
}

// CreateHsmClientCertificateMessage is undocumented.
type CreateHsmClientCertificateMessage struct {
	HsmClientCertificateIdentifier string `xml:"HsmClientCertificateIdentifier"`
	Tags                           []Tag  `xml:"Tags>Tag"`
}

// CreateHsmClientCertificateResult is undocumented.
type CreateHsmClientCertificateResult struct {
	HsmClientCertificate HsmClientCertificate `xml:"CreateHsmClientCertificateResult>HsmClientCertificate"`
}

// CreateHsmConfigurationMessage is undocumented.
type CreateHsmConfigurationMessage struct {
	Description                string `xml:"Description"`
	HsmConfigurationIdentifier string `xml:"HsmConfigurationIdentifier"`
	HsmIPAddress               string `xml:"HsmIpAddress"`
	HsmPartitionName           string `xml:"HsmPartitionName"`
	HsmPartitionPassword       string `xml:"HsmPartitionPassword"`
	HsmServerPublicCertificate string `xml:"HsmServerPublicCertificate"`
	Tags                       []Tag  `xml:"Tags>Tag"`
}

// CreateHsmConfigurationResult is undocumented.
type CreateHsmConfigurationResult struct {
	HsmConfiguration HsmConfiguration `xml:"CreateHsmConfigurationResult>HsmConfiguration"`
}

// CreateTagsMessage is undocumented.
type CreateTagsMessage struct {
	ResourceName string `xml:"ResourceName"`
	Tags         []Tag  `xml:"Tags>Tag"`
}

// DefaultClusterParameters is undocumented.
type DefaultClusterParameters struct {
	Marker               string      `xml:"Marker"`
	ParameterGroupFamily string      `xml:"ParameterGroupFamily"`
	Parameters           []Parameter `xml:"Parameters>Parameter"`
}

// DeleteClusterMessage is undocumented.
type DeleteClusterMessage struct {
	ClusterIdentifier              string `xml:"ClusterIdentifier"`
	FinalClusterSnapshotIdentifier string `xml:"FinalClusterSnapshotIdentifier"`
	SkipFinalClusterSnapshot       bool   `xml:"SkipFinalClusterSnapshot"`
}

// DeleteClusterParameterGroupMessage is undocumented.
type DeleteClusterParameterGroupMessage struct {
	ParameterGroupName string `xml:"ParameterGroupName"`
}

// DeleteClusterResult is undocumented.
type DeleteClusterResult struct {
	Cluster Cluster `xml:"DeleteClusterResult>Cluster"`
}

// DeleteClusterSecurityGroupMessage is undocumented.
type DeleteClusterSecurityGroupMessage struct {
	ClusterSecurityGroupName string `xml:"ClusterSecurityGroupName"`
}

// DeleteClusterSnapshotMessage is undocumented.
type DeleteClusterSnapshotMessage struct {
	SnapshotClusterIdentifier string `xml:"SnapshotClusterIdentifier"`
	SnapshotIdentifier        string `xml:"SnapshotIdentifier"`
}

// DeleteClusterSnapshotResult is undocumented.
type DeleteClusterSnapshotResult struct {
	Snapshot Snapshot `xml:"DeleteClusterSnapshotResult>Snapshot"`
}

// DeleteClusterSubnetGroupMessage is undocumented.
type DeleteClusterSubnetGroupMessage struct {
	ClusterSubnetGroupName string `xml:"ClusterSubnetGroupName"`
}

// DeleteEventSubscriptionMessage is undocumented.
type DeleteEventSubscriptionMessage struct {
	SubscriptionName string `xml:"SubscriptionName"`
}

// DeleteHsmClientCertificateMessage is undocumented.
type DeleteHsmClientCertificateMessage struct {
	HsmClientCertificateIdentifier string `xml:"HsmClientCertificateIdentifier"`
}

// DeleteHsmConfigurationMessage is undocumented.
type DeleteHsmConfigurationMessage struct {
	HsmConfigurationIdentifier string `xml:"HsmConfigurationIdentifier"`
}

// DeleteTagsMessage is undocumented.
type DeleteTagsMessage struct {
	ResourceName string   `xml:"ResourceName"`
	TagKeys      []string `xml:"TagKeys>TagKey"`
}

// DescribeClusterParameterGroupsMessage is undocumented.
type DescribeClusterParameterGroupsMessage struct {
	Marker             string   `xml:"Marker"`
	MaxRecords         int      `xml:"MaxRecords"`
	ParameterGroupName string   `xml:"ParameterGroupName"`
	TagKeys            []string `xml:"TagKeys>TagKey"`
	TagValues          []string `xml:"TagValues>TagValue"`
}

// DescribeClusterParametersMessage is undocumented.
type DescribeClusterParametersMessage struct {
	Marker             string `xml:"Marker"`
	MaxRecords         int    `xml:"MaxRecords"`
	ParameterGroupName string `xml:"ParameterGroupName"`
	Source             string `xml:"Source"`
}

// DescribeClusterSecurityGroupsMessage is undocumented.
type DescribeClusterSecurityGroupsMessage struct {
	ClusterSecurityGroupName string   `xml:"ClusterSecurityGroupName"`
	Marker                   string   `xml:"Marker"`
	MaxRecords               int      `xml:"MaxRecords"`
	TagKeys                  []string `xml:"TagKeys>TagKey"`
	TagValues                []string `xml:"TagValues>TagValue"`
}

// DescribeClusterSnapshotsMessage is undocumented.
type DescribeClusterSnapshotsMessage struct {
	ClusterIdentifier  string    `xml:"ClusterIdentifier"`
	EndTime            time.Time `xml:"EndTime"`
	Marker             string    `xml:"Marker"`
	MaxRecords         int       `xml:"MaxRecords"`
	OwnerAccount       string    `xml:"OwnerAccount"`
	SnapshotIdentifier string    `xml:"SnapshotIdentifier"`
	SnapshotType       string    `xml:"SnapshotType"`
	StartTime          time.Time `xml:"StartTime"`
	TagKeys            []string  `xml:"TagKeys>TagKey"`
	TagValues          []string  `xml:"TagValues>TagValue"`
}

// DescribeClusterSubnetGroupsMessage is undocumented.
type DescribeClusterSubnetGroupsMessage struct {
	ClusterSubnetGroupName string   `xml:"ClusterSubnetGroupName"`
	Marker                 string   `xml:"Marker"`
	MaxRecords             int      `xml:"MaxRecords"`
	TagKeys                []string `xml:"TagKeys>TagKey"`
	TagValues              []string `xml:"TagValues>TagValue"`
}

// DescribeClusterVersionsMessage is undocumented.
type DescribeClusterVersionsMessage struct {
	ClusterParameterGroupFamily string `xml:"ClusterParameterGroupFamily"`
	ClusterVersion              string `xml:"ClusterVersion"`
	Marker                      string `xml:"Marker"`
	MaxRecords                  int    `xml:"MaxRecords"`
}

// DescribeClustersMessage is undocumented.
type DescribeClustersMessage struct {
	ClusterIdentifier string   `xml:"ClusterIdentifier"`
	Marker            string   `xml:"Marker"`
	MaxRecords        int      `xml:"MaxRecords"`
	TagKeys           []string `xml:"TagKeys>TagKey"`
	TagValues         []string `xml:"TagValues>TagValue"`
}

// DescribeDefaultClusterParametersMessage is undocumented.
type DescribeDefaultClusterParametersMessage struct {
	Marker               string `xml:"Marker"`
	MaxRecords           int    `xml:"MaxRecords"`
	ParameterGroupFamily string `xml:"ParameterGroupFamily"`
}

// DescribeDefaultClusterParametersResult is undocumented.
type DescribeDefaultClusterParametersResult struct {
	DefaultClusterParameters DefaultClusterParameters `xml:"DescribeDefaultClusterParametersResult>DefaultClusterParameters"`
}

// DescribeEventCategoriesMessage is undocumented.
type DescribeEventCategoriesMessage struct {
	SourceType string `xml:"SourceType"`
}

// DescribeEventSubscriptionsMessage is undocumented.
type DescribeEventSubscriptionsMessage struct {
	Marker           string `xml:"Marker"`
	MaxRecords       int    `xml:"MaxRecords"`
	SubscriptionName string `xml:"SubscriptionName"`
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

// DescribeHsmClientCertificatesMessage is undocumented.
type DescribeHsmClientCertificatesMessage struct {
	HsmClientCertificateIdentifier string   `xml:"HsmClientCertificateIdentifier"`
	Marker                         string   `xml:"Marker"`
	MaxRecords                     int      `xml:"MaxRecords"`
	TagKeys                        []string `xml:"TagKeys>TagKey"`
	TagValues                      []string `xml:"TagValues>TagValue"`
}

// DescribeHsmConfigurationsMessage is undocumented.
type DescribeHsmConfigurationsMessage struct {
	HsmConfigurationIdentifier string   `xml:"HsmConfigurationIdentifier"`
	Marker                     string   `xml:"Marker"`
	MaxRecords                 int      `xml:"MaxRecords"`
	TagKeys                    []string `xml:"TagKeys>TagKey"`
	TagValues                  []string `xml:"TagValues>TagValue"`
}

// DescribeLoggingStatusMessage is undocumented.
type DescribeLoggingStatusMessage struct {
	ClusterIdentifier string `xml:"ClusterIdentifier"`
}

// DescribeOrderableClusterOptionsMessage is undocumented.
type DescribeOrderableClusterOptionsMessage struct {
	ClusterVersion string `xml:"ClusterVersion"`
	Marker         string `xml:"Marker"`
	MaxRecords     int    `xml:"MaxRecords"`
	NodeType       string `xml:"NodeType"`
}

// DescribeReservedNodeOfferingsMessage is undocumented.
type DescribeReservedNodeOfferingsMessage struct {
	Marker                 string `xml:"Marker"`
	MaxRecords             int    `xml:"MaxRecords"`
	ReservedNodeOfferingID string `xml:"ReservedNodeOfferingId"`
}

// DescribeReservedNodesMessage is undocumented.
type DescribeReservedNodesMessage struct {
	Marker         string `xml:"Marker"`
	MaxRecords     int    `xml:"MaxRecords"`
	ReservedNodeID string `xml:"ReservedNodeId"`
}

// DescribeResizeMessage is undocumented.
type DescribeResizeMessage struct {
	ClusterIdentifier string `xml:"ClusterIdentifier"`
}

// DescribeTagsMessage is undocumented.
type DescribeTagsMessage struct {
	Marker       string   `xml:"Marker"`
	MaxRecords   int      `xml:"MaxRecords"`
	ResourceName string   `xml:"ResourceName"`
	ResourceType string   `xml:"ResourceType"`
	TagKeys      []string `xml:"TagKeys>TagKey"`
	TagValues    []string `xml:"TagValues>TagValue"`
}

// DisableLoggingMessage is undocumented.
type DisableLoggingMessage struct {
	ClusterIdentifier string `xml:"ClusterIdentifier"`
}

// DisableSnapshotCopyMessage is undocumented.
type DisableSnapshotCopyMessage struct {
	ClusterIdentifier string `xml:"ClusterIdentifier"`
}

// DisableSnapshotCopyResult is undocumented.
type DisableSnapshotCopyResult struct {
	Cluster Cluster `xml:"DisableSnapshotCopyResult>Cluster"`
}

// EC2SecurityGroup is undocumented.
type EC2SecurityGroup struct {
	EC2SecurityGroupName    string `xml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerID string `xml:"EC2SecurityGroupOwnerId"`
	Status                  string `xml:"Status"`
	Tags                    []Tag  `xml:"Tags>Tag"`
}

// ElasticIPStatus is undocumented.
type ElasticIPStatus struct {
	ElasticIP string `xml:"ElasticIp"`
	Status    string `xml:"Status"`
}

// EnableLoggingMessage is undocumented.
type EnableLoggingMessage struct {
	BucketName        string `xml:"BucketName"`
	ClusterIdentifier string `xml:"ClusterIdentifier"`
	S3KeyPrefix       string `xml:"S3KeyPrefix"`
}

// EnableSnapshotCopyMessage is undocumented.
type EnableSnapshotCopyMessage struct {
	ClusterIdentifier string `xml:"ClusterIdentifier"`
	DestinationRegion string `xml:"DestinationRegion"`
	RetentionPeriod   int    `xml:"RetentionPeriod"`
}

// EnableSnapshotCopyResult is undocumented.
type EnableSnapshotCopyResult struct {
	Cluster Cluster `xml:"EnableSnapshotCopyResult>Cluster"`
}

// Endpoint is undocumented.
type Endpoint struct {
	Address string `xml:"Address"`
	Port    int    `xml:"Port"`
}

// Event is undocumented.
type Event struct {
	Date             time.Time `xml:"Date"`
	EventCategories  []string  `xml:"EventCategories>EventCategory"`
	EventID          string    `xml:"EventId"`
	Message          string    `xml:"Message"`
	Severity         string    `xml:"Severity"`
	SourceIdentifier string    `xml:"SourceIdentifier"`
	SourceType       string    `xml:"SourceType"`
}

// EventCategoriesMap is undocumented.
type EventCategoriesMap struct {
	Events     []EventInfoMap `xml:"Events>EventInfoMap"`
	SourceType string         `xml:"SourceType"`
}

// EventCategoriesMessage is undocumented.
type EventCategoriesMessage struct {
	EventCategoriesMapList []EventCategoriesMap `xml:"DescribeEventCategoriesResult>EventCategoriesMapList>EventCategoriesMap"`
}

// EventInfoMap is undocumented.
type EventInfoMap struct {
	EventCategories  []string `xml:"EventCategories>EventCategory"`
	EventDescription string   `xml:"EventDescription"`
	EventID          string   `xml:"EventId"`
	Severity         string   `xml:"Severity"`
}

// EventSubscription is undocumented.
type EventSubscription struct {
	CustSubscriptionID       string    `xml:"CustSubscriptionId"`
	CustomerAwsID            string    `xml:"CustomerAwsId"`
	Enabled                  bool      `xml:"Enabled"`
	EventCategoriesList      []string  `xml:"EventCategoriesList>EventCategory"`
	Severity                 string    `xml:"Severity"`
	SnsTopicARN              string    `xml:"SnsTopicArn"`
	SourceIdsList            []string  `xml:"SourceIdsList>SourceId"`
	SourceType               string    `xml:"SourceType"`
	Status                   string    `xml:"Status"`
	SubscriptionCreationTime time.Time `xml:"SubscriptionCreationTime"`
	Tags                     []Tag     `xml:"Tags>Tag"`
}

// EventSubscriptionsMessage is undocumented.
type EventSubscriptionsMessage struct {
	EventSubscriptionsList []EventSubscription `xml:"DescribeEventSubscriptionsResult>EventSubscriptionsList>EventSubscription"`
	Marker                 string              `xml:"DescribeEventSubscriptionsResult>Marker"`
}

// EventsMessage is undocumented.
type EventsMessage struct {
	Events []Event `xml:"DescribeEventsResult>Events>Event"`
	Marker string  `xml:"DescribeEventsResult>Marker"`
}

// HsmClientCertificate is undocumented.
type HsmClientCertificate struct {
	HsmClientCertificateIdentifier string `xml:"HsmClientCertificateIdentifier"`
	HsmClientCertificatePublicKey  string `xml:"HsmClientCertificatePublicKey"`
	Tags                           []Tag  `xml:"Tags>Tag"`
}

// HsmClientCertificateMessage is undocumented.
type HsmClientCertificateMessage struct {
	HsmClientCertificates []HsmClientCertificate `xml:"DescribeHsmClientCertificatesResult>HsmClientCertificates>HsmClientCertificate"`
	Marker                string                 `xml:"DescribeHsmClientCertificatesResult>Marker"`
}

// HsmConfiguration is undocumented.
type HsmConfiguration struct {
	Description                string `xml:"Description"`
	HsmConfigurationIdentifier string `xml:"HsmConfigurationIdentifier"`
	HsmIPAddress               string `xml:"HsmIpAddress"`
	HsmPartitionName           string `xml:"HsmPartitionName"`
	Tags                       []Tag  `xml:"Tags>Tag"`
}

// HsmConfigurationMessage is undocumented.
type HsmConfigurationMessage struct {
	HsmConfigurations []HsmConfiguration `xml:"DescribeHsmConfigurationsResult>HsmConfigurations>HsmConfiguration"`
	Marker            string             `xml:"DescribeHsmConfigurationsResult>Marker"`
}

// HsmStatus is undocumented.
type HsmStatus struct {
	HsmClientCertificateIdentifier string `xml:"HsmClientCertificateIdentifier"`
	HsmConfigurationIdentifier     string `xml:"HsmConfigurationIdentifier"`
	Status                         string `xml:"Status"`
}

// IPRange is undocumented.
type IPRange struct {
	CIDRIP string `xml:"CIDRIP"`
	Status string `xml:"Status"`
	Tags   []Tag  `xml:"Tags>Tag"`
}

// LoggingStatus is undocumented.
type LoggingStatus struct {
	BucketName                 string    `xml:"BucketName"`
	LastFailureMessage         string    `xml:"LastFailureMessage"`
	LastFailureTime            time.Time `xml:"LastFailureTime"`
	LastSuccessfulDeliveryTime time.Time `xml:"LastSuccessfulDeliveryTime"`
	LoggingEnabled             bool      `xml:"LoggingEnabled"`
	S3KeyPrefix                string    `xml:"S3KeyPrefix"`
}

// ModifyClusterMessage is undocumented.
type ModifyClusterMessage struct {
	AllowVersionUpgrade              bool     `xml:"AllowVersionUpgrade"`
	AutomatedSnapshotRetentionPeriod int      `xml:"AutomatedSnapshotRetentionPeriod"`
	ClusterIdentifier                string   `xml:"ClusterIdentifier"`
	ClusterParameterGroupName        string   `xml:"ClusterParameterGroupName"`
	ClusterSecurityGroups            []string `xml:"ClusterSecurityGroups>ClusterSecurityGroupName"`
	ClusterType                      string   `xml:"ClusterType"`
	ClusterVersion                   string   `xml:"ClusterVersion"`
	HsmClientCertificateIdentifier   string   `xml:"HsmClientCertificateIdentifier"`
	HsmConfigurationIdentifier       string   `xml:"HsmConfigurationIdentifier"`
	MasterUserPassword               string   `xml:"MasterUserPassword"`
	NewClusterIdentifier             string   `xml:"NewClusterIdentifier"`
	NodeType                         string   `xml:"NodeType"`
	NumberOfNodes                    int      `xml:"NumberOfNodes"`
	PreferredMaintenanceWindow       string   `xml:"PreferredMaintenanceWindow"`
	VpcSecurityGroupIds              []string `xml:"VpcSecurityGroupIds>VpcSecurityGroupId"`
}

// ModifyClusterParameterGroupMessage is undocumented.
type ModifyClusterParameterGroupMessage struct {
	ParameterGroupName string      `xml:"ParameterGroupName"`
	Parameters         []Parameter `xml:"Parameters>Parameter"`
}

// ModifyClusterResult is undocumented.
type ModifyClusterResult struct {
	Cluster Cluster `xml:"ModifyClusterResult>Cluster"`
}

// ModifyClusterSubnetGroupMessage is undocumented.
type ModifyClusterSubnetGroupMessage struct {
	ClusterSubnetGroupName string   `xml:"ClusterSubnetGroupName"`
	Description            string   `xml:"Description"`
	SubnetIds              []string `xml:"SubnetIds>SubnetIdentifier"`
}

// ModifyClusterSubnetGroupResult is undocumented.
type ModifyClusterSubnetGroupResult struct {
	ClusterSubnetGroup ClusterSubnetGroup `xml:"ModifyClusterSubnetGroupResult>ClusterSubnetGroup"`
}

// ModifyEventSubscriptionMessage is undocumented.
type ModifyEventSubscriptionMessage struct {
	Enabled          bool     `xml:"Enabled"`
	EventCategories  []string `xml:"EventCategories>EventCategory"`
	Severity         string   `xml:"Severity"`
	SnsTopicARN      string   `xml:"SnsTopicArn"`
	SourceIds        []string `xml:"SourceIds>SourceId"`
	SourceType       string   `xml:"SourceType"`
	SubscriptionName string   `xml:"SubscriptionName"`
}

// ModifyEventSubscriptionResult is undocumented.
type ModifyEventSubscriptionResult struct {
	EventSubscription EventSubscription `xml:"ModifyEventSubscriptionResult>EventSubscription"`
}

// ModifySnapshotCopyRetentionPeriodMessage is undocumented.
type ModifySnapshotCopyRetentionPeriodMessage struct {
	ClusterIdentifier string `xml:"ClusterIdentifier"`
	RetentionPeriod   int    `xml:"RetentionPeriod"`
}

// ModifySnapshotCopyRetentionPeriodResult is undocumented.
type ModifySnapshotCopyRetentionPeriodResult struct {
	Cluster Cluster `xml:"ModifySnapshotCopyRetentionPeriodResult>Cluster"`
}

// OrderableClusterOption is undocumented.
type OrderableClusterOption struct {
	AvailabilityZones []AvailabilityZone `xml:"AvailabilityZones>AvailabilityZone"`
	ClusterType       string             `xml:"ClusterType"`
	ClusterVersion    string             `xml:"ClusterVersion"`
	NodeType          string             `xml:"NodeType"`
}

// OrderableClusterOptionsMessage is undocumented.
type OrderableClusterOptionsMessage struct {
	Marker                  string                   `xml:"DescribeOrderableClusterOptionsResult>Marker"`
	OrderableClusterOptions []OrderableClusterOption `xml:"DescribeOrderableClusterOptionsResult>OrderableClusterOptions>OrderableClusterOption"`
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

// PendingModifiedValues is undocumented.
type PendingModifiedValues struct {
	AutomatedSnapshotRetentionPeriod int    `xml:"AutomatedSnapshotRetentionPeriod"`
	ClusterIdentifier                string `xml:"ClusterIdentifier"`
	ClusterType                      string `xml:"ClusterType"`
	ClusterVersion                   string `xml:"ClusterVersion"`
	MasterUserPassword               string `xml:"MasterUserPassword"`
	NodeType                         string `xml:"NodeType"`
	NumberOfNodes                    int    `xml:"NumberOfNodes"`
}

// PurchaseReservedNodeOfferingMessage is undocumented.
type PurchaseReservedNodeOfferingMessage struct {
	NodeCount              int    `xml:"NodeCount"`
	ReservedNodeOfferingID string `xml:"ReservedNodeOfferingId"`
}

// PurchaseReservedNodeOfferingResult is undocumented.
type PurchaseReservedNodeOfferingResult struct {
	ReservedNode ReservedNode `xml:"PurchaseReservedNodeOfferingResult>ReservedNode"`
}

// RebootClusterMessage is undocumented.
type RebootClusterMessage struct {
	ClusterIdentifier string `xml:"ClusterIdentifier"`
}

// RebootClusterResult is undocumented.
type RebootClusterResult struct {
	Cluster Cluster `xml:"RebootClusterResult>Cluster"`
}

// RecurringCharge is undocumented.
type RecurringCharge struct {
	RecurringChargeAmount    float64 `xml:"RecurringChargeAmount"`
	RecurringChargeFrequency string  `xml:"RecurringChargeFrequency"`
}

// ReservedNode is undocumented.
type ReservedNode struct {
	CurrencyCode           string            `xml:"CurrencyCode"`
	Duration               int               `xml:"Duration"`
	FixedPrice             float64           `xml:"FixedPrice"`
	NodeCount              int               `xml:"NodeCount"`
	NodeType               string            `xml:"NodeType"`
	OfferingType           string            `xml:"OfferingType"`
	RecurringCharges       []RecurringCharge `xml:"RecurringCharges>RecurringCharge"`
	ReservedNodeID         string            `xml:"ReservedNodeId"`
	ReservedNodeOfferingID string            `xml:"ReservedNodeOfferingId"`
	StartTime              time.Time         `xml:"StartTime"`
	State                  string            `xml:"State"`
	UsagePrice             float64           `xml:"UsagePrice"`
}

// ReservedNodeOffering is undocumented.
type ReservedNodeOffering struct {
	CurrencyCode           string            `xml:"CurrencyCode"`
	Duration               int               `xml:"Duration"`
	FixedPrice             float64           `xml:"FixedPrice"`
	NodeType               string            `xml:"NodeType"`
	OfferingType           string            `xml:"OfferingType"`
	RecurringCharges       []RecurringCharge `xml:"RecurringCharges>RecurringCharge"`
	ReservedNodeOfferingID string            `xml:"ReservedNodeOfferingId"`
	UsagePrice             float64           `xml:"UsagePrice"`
}

// ReservedNodeOfferingsMessage is undocumented.
type ReservedNodeOfferingsMessage struct {
	Marker                string                 `xml:"DescribeReservedNodeOfferingsResult>Marker"`
	ReservedNodeOfferings []ReservedNodeOffering `xml:"DescribeReservedNodeOfferingsResult>ReservedNodeOfferings>ReservedNodeOffering"`
}

// ReservedNodesMessage is undocumented.
type ReservedNodesMessage struct {
	Marker        string         `xml:"DescribeReservedNodesResult>Marker"`
	ReservedNodes []ReservedNode `xml:"DescribeReservedNodesResult>ReservedNodes>ReservedNode"`
}

// ResetClusterParameterGroupMessage is undocumented.
type ResetClusterParameterGroupMessage struct {
	ParameterGroupName string      `xml:"ParameterGroupName"`
	Parameters         []Parameter `xml:"Parameters>Parameter"`
	ResetAllParameters bool        `xml:"ResetAllParameters"`
}

// ResizeProgressMessage is undocumented.
type ResizeProgressMessage struct {
	AvgResizeRateInMegaBytesPerSecond  float64  `xml:"DescribeResizeResult>AvgResizeRateInMegaBytesPerSecond"`
	ElapsedTimeInSeconds               int      `xml:"DescribeResizeResult>ElapsedTimeInSeconds"`
	EstimatedTimeToCompletionInSeconds int      `xml:"DescribeResizeResult>EstimatedTimeToCompletionInSeconds"`
	ImportTablesCompleted              []string `xml:"DescribeResizeResult>ImportTablesCompleted>member"`
	ImportTablesInProgress             []string `xml:"DescribeResizeResult>ImportTablesInProgress>member"`
	ImportTablesNotStarted             []string `xml:"DescribeResizeResult>ImportTablesNotStarted>member"`
	ProgressInMegaBytes                int      `xml:"DescribeResizeResult>ProgressInMegaBytes"`
	Status                             string   `xml:"DescribeResizeResult>Status"`
	TargetClusterType                  string   `xml:"DescribeResizeResult>TargetClusterType"`
	TargetNodeType                     string   `xml:"DescribeResizeResult>TargetNodeType"`
	TargetNumberOfNodes                int      `xml:"DescribeResizeResult>TargetNumberOfNodes"`
	TotalResizeDataInMegaBytes         int      `xml:"DescribeResizeResult>TotalResizeDataInMegaBytes"`
}

// RestoreFromClusterSnapshotMessage is undocumented.
type RestoreFromClusterSnapshotMessage struct {
	AllowVersionUpgrade              bool     `xml:"AllowVersionUpgrade"`
	AutomatedSnapshotRetentionPeriod int      `xml:"AutomatedSnapshotRetentionPeriod"`
	AvailabilityZone                 string   `xml:"AvailabilityZone"`
	ClusterIdentifier                string   `xml:"ClusterIdentifier"`
	ClusterParameterGroupName        string   `xml:"ClusterParameterGroupName"`
	ClusterSecurityGroups            []string `xml:"ClusterSecurityGroups>ClusterSecurityGroupName"`
	ClusterSubnetGroupName           string   `xml:"ClusterSubnetGroupName"`
	ElasticIP                        string   `xml:"ElasticIp"`
	HsmClientCertificateIdentifier   string   `xml:"HsmClientCertificateIdentifier"`
	HsmConfigurationIdentifier       string   `xml:"HsmConfigurationIdentifier"`
	KmsKeyID                         string   `xml:"KmsKeyId"`
	OwnerAccount                     string   `xml:"OwnerAccount"`
	Port                             int      `xml:"Port"`
	PreferredMaintenanceWindow       string   `xml:"PreferredMaintenanceWindow"`
	PubliclyAccessible               bool     `xml:"PubliclyAccessible"`
	SnapshotClusterIdentifier        string   `xml:"SnapshotClusterIdentifier"`
	SnapshotIdentifier               string   `xml:"SnapshotIdentifier"`
	VpcSecurityGroupIds              []string `xml:"VpcSecurityGroupIds>VpcSecurityGroupId"`
}

// RestoreFromClusterSnapshotResult is undocumented.
type RestoreFromClusterSnapshotResult struct {
	Cluster Cluster `xml:"RestoreFromClusterSnapshotResult>Cluster"`
}

// RestoreStatus is undocumented.
type RestoreStatus struct {
	CurrentRestoreRateInMegaBytesPerSecond float64 `xml:"CurrentRestoreRateInMegaBytesPerSecond"`
	ElapsedTimeInSeconds                   int     `xml:"ElapsedTimeInSeconds"`
	EstimatedTimeToCompletionInSeconds     int     `xml:"EstimatedTimeToCompletionInSeconds"`
	ProgressInMegaBytes                    int     `xml:"ProgressInMegaBytes"`
	SnapshotSizeInMegaBytes                int     `xml:"SnapshotSizeInMegaBytes"`
	Status                                 string  `xml:"Status"`
}

// RevokeClusterSecurityGroupIngressMessage is undocumented.
type RevokeClusterSecurityGroupIngressMessage struct {
	CIDRIP                   string `xml:"CIDRIP"`
	ClusterSecurityGroupName string `xml:"ClusterSecurityGroupName"`
	EC2SecurityGroupName     string `xml:"EC2SecurityGroupName"`
	EC2SecurityGroupOwnerID  string `xml:"EC2SecurityGroupOwnerId"`
}

// RevokeClusterSecurityGroupIngressResult is undocumented.
type RevokeClusterSecurityGroupIngressResult struct {
	ClusterSecurityGroup ClusterSecurityGroup `xml:"RevokeClusterSecurityGroupIngressResult>ClusterSecurityGroup"`
}

// RevokeSnapshotAccessMessage is undocumented.
type RevokeSnapshotAccessMessage struct {
	AccountWithRestoreAccess  string `xml:"AccountWithRestoreAccess"`
	SnapshotClusterIdentifier string `xml:"SnapshotClusterIdentifier"`
	SnapshotIdentifier        string `xml:"SnapshotIdentifier"`
}

// RevokeSnapshotAccessResult is undocumented.
type RevokeSnapshotAccessResult struct {
	Snapshot Snapshot `xml:"RevokeSnapshotAccessResult>Snapshot"`
}

// RotateEncryptionKeyMessage is undocumented.
type RotateEncryptionKeyMessage struct {
	ClusterIdentifier string `xml:"ClusterIdentifier"`
}

// RotateEncryptionKeyResult is undocumented.
type RotateEncryptionKeyResult struct {
	Cluster Cluster `xml:"RotateEncryptionKeyResult>Cluster"`
}

// Snapshot is undocumented.
type Snapshot struct {
	AccountsWithRestoreAccess              []AccountWithRestoreAccess `xml:"AccountsWithRestoreAccess>AccountWithRestoreAccess"`
	ActualIncrementalBackupSizeInMegaBytes float64                    `xml:"ActualIncrementalBackupSizeInMegaBytes"`
	AvailabilityZone                       string                     `xml:"AvailabilityZone"`
	BackupProgressInMegaBytes              float64                    `xml:"BackupProgressInMegaBytes"`
	ClusterCreateTime                      time.Time                  `xml:"ClusterCreateTime"`
	ClusterIdentifier                      string                     `xml:"ClusterIdentifier"`
	ClusterVersion                         string                     `xml:"ClusterVersion"`
	CurrentBackupRateInMegaBytesPerSecond  float64                    `xml:"CurrentBackupRateInMegaBytesPerSecond"`
	DBName                                 string                     `xml:"DBName"`
	ElapsedTimeInSeconds                   int                        `xml:"ElapsedTimeInSeconds"`
	Encrypted                              bool                       `xml:"Encrypted"`
	EncryptedWithHSM                       bool                       `xml:"EncryptedWithHSM"`
	EstimatedSecondsToCompletion           int                        `xml:"EstimatedSecondsToCompletion"`
	KmsKeyID                               string                     `xml:"KmsKeyId"`
	MasterUsername                         string                     `xml:"MasterUsername"`
	NodeType                               string                     `xml:"NodeType"`
	NumberOfNodes                          int                        `xml:"NumberOfNodes"`
	OwnerAccount                           string                     `xml:"OwnerAccount"`
	Port                                   int                        `xml:"Port"`
	SnapshotCreateTime                     time.Time                  `xml:"SnapshotCreateTime"`
	SnapshotIdentifier                     string                     `xml:"SnapshotIdentifier"`
	SnapshotType                           string                     `xml:"SnapshotType"`
	SourceRegion                           string                     `xml:"SourceRegion"`
	Status                                 string                     `xml:"Status"`
	Tags                                   []Tag                      `xml:"Tags>Tag"`
	TotalBackupSizeInMegaBytes             float64                    `xml:"TotalBackupSizeInMegaBytes"`
	VpcID                                  string                     `xml:"VpcId"`
}

// SnapshotMessage is undocumented.
type SnapshotMessage struct {
	Marker    string     `xml:"DescribeClusterSnapshotsResult>Marker"`
	Snapshots []Snapshot `xml:"DescribeClusterSnapshotsResult>Snapshots>Snapshot"`
}

// Subnet is undocumented.
type Subnet struct {
	SubnetAvailabilityZone AvailabilityZone `xml:"SubnetAvailabilityZone"`
	SubnetIdentifier       string           `xml:"SubnetIdentifier"`
	SubnetStatus           string           `xml:"SubnetStatus"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// TaggedResource is undocumented.
type TaggedResource struct {
	ResourceName string `xml:"ResourceName"`
	ResourceType string `xml:"ResourceType"`
	Tag          Tag    `xml:"Tag"`
}

// TaggedResourceListMessage is undocumented.
type TaggedResourceListMessage struct {
	Marker          string           `xml:"DescribeTagsResult>Marker"`
	TaggedResources []TaggedResource `xml:"DescribeTagsResult>TaggedResources>TaggedResource"`
}

// VpcSecurityGroupMembership is undocumented.
type VpcSecurityGroupMembership struct {
	Status             string `xml:"Status"`
	VpcSecurityGroupID string `xml:"VpcSecurityGroupId"`
}

// DescribeClusterParameterGroupsResult is a wrapper for ClusterParameterGroupsMessage.
type DescribeClusterParameterGroupsResult struct {
	XMLName xml.Name `xml:"DescribeClusterParameterGroupsResponse"`

	Marker          string                  `xml:"DescribeClusterParameterGroupsResult>Marker"`
	ParameterGroups []ClusterParameterGroup `xml:"DescribeClusterParameterGroupsResult>ParameterGroups>ClusterParameterGroup"`
}

// DescribeClusterParametersResult is a wrapper for ClusterParameterGroupDetails.
type DescribeClusterParametersResult struct {
	XMLName xml.Name `xml:"DescribeClusterParametersResponse"`

	Marker     string      `xml:"DescribeClusterParametersResult>Marker"`
	Parameters []Parameter `xml:"DescribeClusterParametersResult>Parameters>Parameter"`
}

// DescribeClusterSecurityGroupsResult is a wrapper for ClusterSecurityGroupMessage.
type DescribeClusterSecurityGroupsResult struct {
	XMLName xml.Name `xml:"DescribeClusterSecurityGroupsResponse"`

	ClusterSecurityGroups []ClusterSecurityGroup `xml:"DescribeClusterSecurityGroupsResult>ClusterSecurityGroups>ClusterSecurityGroup"`
	Marker                string                 `xml:"DescribeClusterSecurityGroupsResult>Marker"`
}

// DescribeClusterSnapshotsResult is a wrapper for SnapshotMessage.
type DescribeClusterSnapshotsResult struct {
	XMLName xml.Name `xml:"DescribeClusterSnapshotsResponse"`

	Marker    string     `xml:"DescribeClusterSnapshotsResult>Marker"`
	Snapshots []Snapshot `xml:"DescribeClusterSnapshotsResult>Snapshots>Snapshot"`
}

// DescribeClusterSubnetGroupsResult is a wrapper for ClusterSubnetGroupMessage.
type DescribeClusterSubnetGroupsResult struct {
	XMLName xml.Name `xml:"DescribeClusterSubnetGroupsResponse"`

	ClusterSubnetGroups []ClusterSubnetGroup `xml:"DescribeClusterSubnetGroupsResult>ClusterSubnetGroups>ClusterSubnetGroup"`
	Marker              string               `xml:"DescribeClusterSubnetGroupsResult>Marker"`
}

// DescribeClusterVersionsResult is a wrapper for ClusterVersionsMessage.
type DescribeClusterVersionsResult struct {
	XMLName xml.Name `xml:"DescribeClusterVersionsResponse"`

	ClusterVersions []ClusterVersion `xml:"DescribeClusterVersionsResult>ClusterVersions>ClusterVersion"`
	Marker          string           `xml:"DescribeClusterVersionsResult>Marker"`
}

// DescribeClustersResult is a wrapper for ClustersMessage.
type DescribeClustersResult struct {
	XMLName xml.Name `xml:"DescribeClustersResponse"`

	Clusters []Cluster `xml:"DescribeClustersResult>Clusters>Cluster"`
	Marker   string    `xml:"DescribeClustersResult>Marker"`
}

// DescribeEventCategoriesResult is a wrapper for EventCategoriesMessage.
type DescribeEventCategoriesResult struct {
	XMLName xml.Name `xml:"DescribeEventCategoriesResponse"`

	EventCategoriesMapList []EventCategoriesMap `xml:"DescribeEventCategoriesResult>EventCategoriesMapList>EventCategoriesMap"`
}

// DescribeEventSubscriptionsResult is a wrapper for EventSubscriptionsMessage.
type DescribeEventSubscriptionsResult struct {
	XMLName xml.Name `xml:"DescribeEventSubscriptionsResponse"`

	EventSubscriptionsList []EventSubscription `xml:"DescribeEventSubscriptionsResult>EventSubscriptionsList>EventSubscription"`
	Marker                 string              `xml:"DescribeEventSubscriptionsResult>Marker"`
}

// DescribeEventsResult is a wrapper for EventsMessage.
type DescribeEventsResult struct {
	XMLName xml.Name `xml:"DescribeEventsResponse"`

	Events []Event `xml:"DescribeEventsResult>Events>Event"`
	Marker string  `xml:"DescribeEventsResult>Marker"`
}

// DescribeHsmClientCertificatesResult is a wrapper for HsmClientCertificateMessage.
type DescribeHsmClientCertificatesResult struct {
	XMLName xml.Name `xml:"DescribeHsmClientCertificatesResponse"`

	HsmClientCertificates []HsmClientCertificate `xml:"DescribeHsmClientCertificatesResult>HsmClientCertificates>HsmClientCertificate"`
	Marker                string                 `xml:"DescribeHsmClientCertificatesResult>Marker"`
}

// DescribeHsmConfigurationsResult is a wrapper for HsmConfigurationMessage.
type DescribeHsmConfigurationsResult struct {
	XMLName xml.Name `xml:"DescribeHsmConfigurationsResponse"`

	HsmConfigurations []HsmConfiguration `xml:"DescribeHsmConfigurationsResult>HsmConfigurations>HsmConfiguration"`
	Marker            string             `xml:"DescribeHsmConfigurationsResult>Marker"`
}

// DescribeLoggingStatusResult is a wrapper for LoggingStatus.
type DescribeLoggingStatusResult struct {
	XMLName xml.Name `xml:"Response"`

	BucketName                 string    `xml:"DescribeLoggingStatusResult>BucketName"`
	LastFailureMessage         string    `xml:"DescribeLoggingStatusResult>LastFailureMessage"`
	LastFailureTime            time.Time `xml:"DescribeLoggingStatusResult>LastFailureTime"`
	LastSuccessfulDeliveryTime time.Time `xml:"DescribeLoggingStatusResult>LastSuccessfulDeliveryTime"`
	LoggingEnabled             bool      `xml:"DescribeLoggingStatusResult>LoggingEnabled"`
	S3KeyPrefix                string    `xml:"DescribeLoggingStatusResult>S3KeyPrefix"`
}

// DescribeOrderableClusterOptionsResult is a wrapper for OrderableClusterOptionsMessage.
type DescribeOrderableClusterOptionsResult struct {
	XMLName xml.Name `xml:"DescribeOrderableClusterOptionsResponse"`

	Marker                  string                   `xml:"DescribeOrderableClusterOptionsResult>Marker"`
	OrderableClusterOptions []OrderableClusterOption `xml:"DescribeOrderableClusterOptionsResult>OrderableClusterOptions>OrderableClusterOption"`
}

// DescribeReservedNodeOfferingsResult is a wrapper for ReservedNodeOfferingsMessage.
type DescribeReservedNodeOfferingsResult struct {
	XMLName xml.Name `xml:"DescribeReservedNodeOfferingsResponse"`

	Marker                string                 `xml:"DescribeReservedNodeOfferingsResult>Marker"`
	ReservedNodeOfferings []ReservedNodeOffering `xml:"DescribeReservedNodeOfferingsResult>ReservedNodeOfferings>ReservedNodeOffering"`
}

// DescribeReservedNodesResult is a wrapper for ReservedNodesMessage.
type DescribeReservedNodesResult struct {
	XMLName xml.Name `xml:"DescribeReservedNodesResponse"`

	Marker        string         `xml:"DescribeReservedNodesResult>Marker"`
	ReservedNodes []ReservedNode `xml:"DescribeReservedNodesResult>ReservedNodes>ReservedNode"`
}

// DescribeResizeResult is a wrapper for ResizeProgressMessage.
type DescribeResizeResult struct {
	XMLName xml.Name `xml:"DescribeResizeResponse"`

	AvgResizeRateInMegaBytesPerSecond  float64  `xml:"DescribeResizeResult>AvgResizeRateInMegaBytesPerSecond"`
	ElapsedTimeInSeconds               int      `xml:"DescribeResizeResult>ElapsedTimeInSeconds"`
	EstimatedTimeToCompletionInSeconds int      `xml:"DescribeResizeResult>EstimatedTimeToCompletionInSeconds"`
	ImportTablesCompleted              []string `xml:"DescribeResizeResult>ImportTablesCompleted>member"`
	ImportTablesInProgress             []string `xml:"DescribeResizeResult>ImportTablesInProgress>member"`
	ImportTablesNotStarted             []string `xml:"DescribeResizeResult>ImportTablesNotStarted>member"`
	ProgressInMegaBytes                int      `xml:"DescribeResizeResult>ProgressInMegaBytes"`
	Status                             string   `xml:"DescribeResizeResult>Status"`
	TargetClusterType                  string   `xml:"DescribeResizeResult>TargetClusterType"`
	TargetNodeType                     string   `xml:"DescribeResizeResult>TargetNodeType"`
	TargetNumberOfNodes                int      `xml:"DescribeResizeResult>TargetNumberOfNodes"`
	TotalResizeDataInMegaBytes         int      `xml:"DescribeResizeResult>TotalResizeDataInMegaBytes"`
}

// DescribeTagsResult is a wrapper for TaggedResourceListMessage.
type DescribeTagsResult struct {
	XMLName xml.Name `xml:"DescribeTagsResponse"`

	Marker          string           `xml:"DescribeTagsResult>Marker"`
	TaggedResources []TaggedResource `xml:"DescribeTagsResult>TaggedResources>TaggedResource"`
}

// DisableLoggingResult is a wrapper for LoggingStatus.
type DisableLoggingResult struct {
	XMLName xml.Name `xml:"Response"`

	BucketName                 string    `xml:"DisableLoggingResult>BucketName"`
	LastFailureMessage         string    `xml:"DisableLoggingResult>LastFailureMessage"`
	LastFailureTime            time.Time `xml:"DisableLoggingResult>LastFailureTime"`
	LastSuccessfulDeliveryTime time.Time `xml:"DisableLoggingResult>LastSuccessfulDeliveryTime"`
	LoggingEnabled             bool      `xml:"DisableLoggingResult>LoggingEnabled"`
	S3KeyPrefix                string    `xml:"DisableLoggingResult>S3KeyPrefix"`
}

// EnableLoggingResult is a wrapper for LoggingStatus.
type EnableLoggingResult struct {
	XMLName xml.Name `xml:"Response"`

	BucketName                 string    `xml:"EnableLoggingResult>BucketName"`
	LastFailureMessage         string    `xml:"EnableLoggingResult>LastFailureMessage"`
	LastFailureTime            time.Time `xml:"EnableLoggingResult>LastFailureTime"`
	LastSuccessfulDeliveryTime time.Time `xml:"EnableLoggingResult>LastSuccessfulDeliveryTime"`
	LoggingEnabled             bool      `xml:"EnableLoggingResult>LoggingEnabled"`
	S3KeyPrefix                string    `xml:"EnableLoggingResult>S3KeyPrefix"`
}

// ModifyClusterParameterGroupResult is a wrapper for ClusterParameterGroupNameMessage.
type ModifyClusterParameterGroupResult struct {
	XMLName xml.Name `xml:"Response"`

	ParameterGroupName   string `xml:"ModifyClusterParameterGroupResult>ParameterGroupName"`
	ParameterGroupStatus string `xml:"ModifyClusterParameterGroupResult>ParameterGroupStatus"`
}

// ResetClusterParameterGroupResult is a wrapper for ClusterParameterGroupNameMessage.
type ResetClusterParameterGroupResult struct {
	XMLName xml.Name `xml:"Response"`

	ParameterGroupName   string `xml:"ResetClusterParameterGroupResult>ParameterGroupName"`
	ParameterGroupStatus string `xml:"ResetClusterParameterGroupResult>ParameterGroupStatus"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
