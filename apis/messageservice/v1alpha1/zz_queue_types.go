// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DlqPolicyInitParameters struct {

	// The queue to which dead-letter messages are delivered.
	DeadLetterTargetQueue *string `json:"deadLetterTargetQueue,omitempty" tf:"dead_letter_target_queue,omitempty"`

	// Specifies whether to enable the dead-letter message delivery. Valid values: true, false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The maximum number of retries.
	MaxReceiveCount *float64 `json:"maxReceiveCount,omitempty" tf:"max_receive_count,omitempty"`
}

type DlqPolicyObservation struct {

	// The queue to which dead-letter messages are delivered.
	DeadLetterTargetQueue *string `json:"deadLetterTargetQueue,omitempty" tf:"dead_letter_target_queue,omitempty"`

	// Specifies whether to enable the dead-letter message delivery. Valid values: true, false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The maximum number of retries.
	MaxReceiveCount *float64 `json:"maxReceiveCount,omitempty" tf:"max_receive_count,omitempty"`
}

type DlqPolicyParameters struct {

	// The queue to which dead-letter messages are delivered.
	// +kubebuilder:validation:Optional
	DeadLetterTargetQueue *string `json:"deadLetterTargetQueue,omitempty" tf:"dead_letter_target_queue,omitempty"`

	// Specifies whether to enable the dead-letter message delivery. Valid values: true, false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The maximum number of retries.
	// +kubebuilder:validation:Optional
	MaxReceiveCount *float64 `json:"maxReceiveCount,omitempty" tf:"max_receive_count,omitempty"`
}

type QueueInitParameters struct {

	// The period after which all messages sent to the queue are consumed. Default value: 0. Valid values: 0 to 604800. Unit: seconds.
	DelaySeconds *float64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// The dead-letter queue policy. See dlq_policy below.
	DlqPolicy []DlqPolicyInitParameters `json:"dlqPolicy,omitempty" tf:"dlq_policy,omitempty"`

	// Specifies whether to enable the logging feature. Default value: false. Valid values:
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	MaximumMessageSize *float64 `json:"maximumMessageSize,omitempty" tf:"maximum_message_size,omitempty"`

	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	MessageRetentionPeriod *float64 `json:"messageRetentionPeriod,omitempty" tf:"message_retention_period,omitempty"`

	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0.
	PollingWaitSeconds *float64 `json:"pollingWaitSeconds,omitempty" tf:"polling_wait_seconds,omitempty"`

	// The name of the queue.
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	VisibilityTimeout *float64 `json:"visibilityTimeout,omitempty" tf:"visibility_timeout,omitempty"`
}

type QueueObservation struct {

	// (Available since v1.223.2) The time when the queue was created.
	CreateTime *float64 `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The period after which all messages sent to the queue are consumed. Default value: 0. Valid values: 0 to 604800. Unit: seconds.
	DelaySeconds *float64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// The dead-letter queue policy. See dlq_policy below.
	DlqPolicy []DlqPolicyObservation `json:"dlqPolicy,omitempty" tf:"dlq_policy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether to enable the logging feature. Default value: false. Valid values:
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	MaximumMessageSize *float64 `json:"maximumMessageSize,omitempty" tf:"maximum_message_size,omitempty"`

	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	MessageRetentionPeriod *float64 `json:"messageRetentionPeriod,omitempty" tf:"message_retention_period,omitempty"`

	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0.
	PollingWaitSeconds *float64 `json:"pollingWaitSeconds,omitempty" tf:"polling_wait_seconds,omitempty"`

	// The name of the queue.
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	VisibilityTimeout *float64 `json:"visibilityTimeout,omitempty" tf:"visibility_timeout,omitempty"`
}

type QueueParameters struct {

	// The period after which all messages sent to the queue are consumed. Default value: 0. Valid values: 0 to 604800. Unit: seconds.
	// +kubebuilder:validation:Optional
	DelaySeconds *float64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// The dead-letter queue policy. See dlq_policy below.
	// +kubebuilder:validation:Optional
	DlqPolicy []DlqPolicyParameters `json:"dlqPolicy,omitempty" tf:"dlq_policy,omitempty"`

	// Specifies whether to enable the logging feature. Default value: false. Valid values:
	// +kubebuilder:validation:Optional
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	// +kubebuilder:validation:Optional
	MaximumMessageSize *float64 `json:"maximumMessageSize,omitempty" tf:"maximum_message_size,omitempty"`

	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	// +kubebuilder:validation:Optional
	MessageRetentionPeriod *float64 `json:"messageRetentionPeriod,omitempty" tf:"message_retention_period,omitempty"`

	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0.
	// +kubebuilder:validation:Optional
	PollingWaitSeconds *float64 `json:"pollingWaitSeconds,omitempty" tf:"polling_wait_seconds,omitempty"`

	// The name of the queue.
	// +kubebuilder:validation:Optional
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	// +kubebuilder:validation:Optional
	VisibilityTimeout *float64 `json:"visibilityTimeout,omitempty" tf:"visibility_timeout,omitempty"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider QueueInitParameters `json:"initProvider,omitempty"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Queue is the Schema for the Queues API. Provides a Alicloud Message Service Queue resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,alibabacloud}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.queueName) || (has(self.initProvider) && has(self.initProvider.queueName))",message="spec.forProvider.queueName is a required parameter"
	Spec   QueueSpec   `json:"spec"`
	Status QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
