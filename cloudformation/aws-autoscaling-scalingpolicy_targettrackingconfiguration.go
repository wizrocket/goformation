package cloudformation

// AWSAutoScalingScalingPolicy_TargetTrackingConfiguration AWS CloudFormation Resource (AWS::AutoScaling::ScalingPolicy.TargetTrackingConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html
type AWSAutoScalingScalingPolicy_TargetTrackingConfiguration struct {

	// CustomizedMetricSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-customizedmetricspecification
	CustomizedMetricSpecification *AWSAutoScalingScalingPolicy_CustomizedMetricSpecification `json:"CustomizedMetricSpecification,omitempty"`

	// DisableScaleIn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-disablescalein
	DisableScaleIn bool `json:"DisableScaleIn,omitempty"`

	// PredefinedMetricSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-predefinedmetricspecification
	PredefinedMetricSpecification *AWSAutoScalingScalingPolicy_PredefinedMetricSpecification `json:"PredefinedMetricSpecification,omitempty"`

	// TargetValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-targetvalue
	TargetValue float64 `json:"TargetValue,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingScalingPolicy_TargetTrackingConfiguration) AWSCloudFormationType() string {
	return "AWS::AutoScaling::ScalingPolicy.TargetTrackingConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSAutoScalingScalingPolicy_TargetTrackingConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.5.0"
}