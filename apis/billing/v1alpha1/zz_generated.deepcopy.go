// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamBinding) DeepCopyInto(out *AccountIamBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamBinding.
func (in *AccountIamBinding) DeepCopy() *AccountIamBinding {
	if in == nil {
		return nil
	}
	out := new(AccountIamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountIamBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamBindingList) DeepCopyInto(out *AccountIamBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountIamBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamBindingList.
func (in *AccountIamBindingList) DeepCopy() *AccountIamBindingList {
	if in == nil {
		return nil
	}
	out := new(AccountIamBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountIamBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamBindingObservation) DeepCopyInto(out *AccountIamBindingObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamBindingObservation.
func (in *AccountIamBindingObservation) DeepCopy() *AccountIamBindingObservation {
	if in == nil {
		return nil
	}
	out := new(AccountIamBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamBindingParameters) DeepCopyInto(out *AccountIamBindingParameters) {
	*out = *in
	if in.BillingAccountID != nil {
		in, out := &in.BillingAccountID, &out.BillingAccountID
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]ConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamBindingParameters.
func (in *AccountIamBindingParameters) DeepCopy() *AccountIamBindingParameters {
	if in == nil {
		return nil
	}
	out := new(AccountIamBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamBindingSpec) DeepCopyInto(out *AccountIamBindingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamBindingSpec.
func (in *AccountIamBindingSpec) DeepCopy() *AccountIamBindingSpec {
	if in == nil {
		return nil
	}
	out := new(AccountIamBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamBindingStatus) DeepCopyInto(out *AccountIamBindingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamBindingStatus.
func (in *AccountIamBindingStatus) DeepCopy() *AccountIamBindingStatus {
	if in == nil {
		return nil
	}
	out := new(AccountIamBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamMember) DeepCopyInto(out *AccountIamMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamMember.
func (in *AccountIamMember) DeepCopy() *AccountIamMember {
	if in == nil {
		return nil
	}
	out := new(AccountIamMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountIamMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamMemberConditionObservation) DeepCopyInto(out *AccountIamMemberConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamMemberConditionObservation.
func (in *AccountIamMemberConditionObservation) DeepCopy() *AccountIamMemberConditionObservation {
	if in == nil {
		return nil
	}
	out := new(AccountIamMemberConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamMemberConditionParameters) DeepCopyInto(out *AccountIamMemberConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamMemberConditionParameters.
func (in *AccountIamMemberConditionParameters) DeepCopy() *AccountIamMemberConditionParameters {
	if in == nil {
		return nil
	}
	out := new(AccountIamMemberConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamMemberList) DeepCopyInto(out *AccountIamMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountIamMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamMemberList.
func (in *AccountIamMemberList) DeepCopy() *AccountIamMemberList {
	if in == nil {
		return nil
	}
	out := new(AccountIamMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountIamMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamMemberObservation) DeepCopyInto(out *AccountIamMemberObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamMemberObservation.
func (in *AccountIamMemberObservation) DeepCopy() *AccountIamMemberObservation {
	if in == nil {
		return nil
	}
	out := new(AccountIamMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamMemberParameters) DeepCopyInto(out *AccountIamMemberParameters) {
	*out = *in
	if in.BillingAccountID != nil {
		in, out := &in.BillingAccountID, &out.BillingAccountID
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]AccountIamMemberConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamMemberParameters.
func (in *AccountIamMemberParameters) DeepCopy() *AccountIamMemberParameters {
	if in == nil {
		return nil
	}
	out := new(AccountIamMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamMemberSpec) DeepCopyInto(out *AccountIamMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamMemberSpec.
func (in *AccountIamMemberSpec) DeepCopy() *AccountIamMemberSpec {
	if in == nil {
		return nil
	}
	out := new(AccountIamMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamMemberStatus) DeepCopyInto(out *AccountIamMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamMemberStatus.
func (in *AccountIamMemberStatus) DeepCopy() *AccountIamMemberStatus {
	if in == nil {
		return nil
	}
	out := new(AccountIamMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamPolicy) DeepCopyInto(out *AccountIamPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamPolicy.
func (in *AccountIamPolicy) DeepCopy() *AccountIamPolicy {
	if in == nil {
		return nil
	}
	out := new(AccountIamPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountIamPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamPolicyList) DeepCopyInto(out *AccountIamPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountIamPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamPolicyList.
func (in *AccountIamPolicyList) DeepCopy() *AccountIamPolicyList {
	if in == nil {
		return nil
	}
	out := new(AccountIamPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountIamPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamPolicyObservation) DeepCopyInto(out *AccountIamPolicyObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamPolicyObservation.
func (in *AccountIamPolicyObservation) DeepCopy() *AccountIamPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(AccountIamPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamPolicyParameters) DeepCopyInto(out *AccountIamPolicyParameters) {
	*out = *in
	if in.BillingAccountID != nil {
		in, out := &in.BillingAccountID, &out.BillingAccountID
		*out = new(string)
		**out = **in
	}
	if in.PolicyData != nil {
		in, out := &in.PolicyData, &out.PolicyData
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamPolicyParameters.
func (in *AccountIamPolicyParameters) DeepCopy() *AccountIamPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(AccountIamPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamPolicySpec) DeepCopyInto(out *AccountIamPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamPolicySpec.
func (in *AccountIamPolicySpec) DeepCopy() *AccountIamPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AccountIamPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountIamPolicyStatus) DeepCopyInto(out *AccountIamPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountIamPolicyStatus.
func (in *AccountIamPolicyStatus) DeepCopy() *AccountIamPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AccountIamPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllUpdatesRuleObservation) DeepCopyInto(out *AllUpdatesRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllUpdatesRuleObservation.
func (in *AllUpdatesRuleObservation) DeepCopy() *AllUpdatesRuleObservation {
	if in == nil {
		return nil
	}
	out := new(AllUpdatesRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllUpdatesRuleParameters) DeepCopyInto(out *AllUpdatesRuleParameters) {
	*out = *in
	if in.DisableDefaultIamRecipients != nil {
		in, out := &in.DisableDefaultIamRecipients, &out.DisableDefaultIamRecipients
		*out = new(bool)
		**out = **in
	}
	if in.MonitoringNotificationChannels != nil {
		in, out := &in.MonitoringNotificationChannels, &out.MonitoringNotificationChannels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PubsubTopic != nil {
		in, out := &in.PubsubTopic, &out.PubsubTopic
		*out = new(string)
		**out = **in
	}
	if in.SchemaVersion != nil {
		in, out := &in.SchemaVersion, &out.SchemaVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllUpdatesRuleParameters.
func (in *AllUpdatesRuleParameters) DeepCopy() *AllUpdatesRuleParameters {
	if in == nil {
		return nil
	}
	out := new(AllUpdatesRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmountObservation) DeepCopyInto(out *AmountObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmountObservation.
func (in *AmountObservation) DeepCopy() *AmountObservation {
	if in == nil {
		return nil
	}
	out := new(AmountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmountParameters) DeepCopyInto(out *AmountParameters) {
	*out = *in
	if in.LastPeriodAmount != nil {
		in, out := &in.LastPeriodAmount, &out.LastPeriodAmount
		*out = new(bool)
		**out = **in
	}
	if in.SpecifiedAmount != nil {
		in, out := &in.SpecifiedAmount, &out.SpecifiedAmount
		*out = make([]SpecifiedAmountParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmountParameters.
func (in *AmountParameters) DeepCopy() *AmountParameters {
	if in == nil {
		return nil
	}
	out := new(AmountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Budget) DeepCopyInto(out *Budget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Budget.
func (in *Budget) DeepCopy() *Budget {
	if in == nil {
		return nil
	}
	out := new(Budget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Budget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetFilterObservation) DeepCopyInto(out *BudgetFilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetFilterObservation.
func (in *BudgetFilterObservation) DeepCopy() *BudgetFilterObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetFilterParameters) DeepCopyInto(out *BudgetFilterParameters) {
	*out = *in
	if in.CreditTypes != nil {
		in, out := &in.CreditTypes, &out.CreditTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CreditTypesTreatment != nil {
		in, out := &in.CreditTypesTreatment, &out.CreditTypesTreatment
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Projects != nil {
		in, out := &in.Projects, &out.Projects
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Subaccounts != nil {
		in, out := &in.Subaccounts, &out.Subaccounts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetFilterParameters.
func (in *BudgetFilterParameters) DeepCopy() *BudgetFilterParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetList) DeepCopyInto(out *BudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Budget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetList.
func (in *BudgetList) DeepCopy() *BudgetList {
	if in == nil {
		return nil
	}
	out := new(BudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetObservation) DeepCopyInto(out *BudgetObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetObservation.
func (in *BudgetObservation) DeepCopy() *BudgetObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetParameters) DeepCopyInto(out *BudgetParameters) {
	*out = *in
	if in.AllUpdatesRule != nil {
		in, out := &in.AllUpdatesRule, &out.AllUpdatesRule
		*out = make([]AllUpdatesRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = make([]AmountParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BillingAccount != nil {
		in, out := &in.BillingAccount, &out.BillingAccount
		*out = new(string)
		**out = **in
	}
	if in.BudgetFilter != nil {
		in, out := &in.BudgetFilter, &out.BudgetFilter
		*out = make([]BudgetFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ThresholdRules != nil {
		in, out := &in.ThresholdRules, &out.ThresholdRules
		*out = make([]ThresholdRulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetParameters.
func (in *BudgetParameters) DeepCopy() *BudgetParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSpec) DeepCopyInto(out *BudgetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSpec.
func (in *BudgetSpec) DeepCopy() *BudgetSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetStatus) DeepCopyInto(out *BudgetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetStatus.
func (in *BudgetStatus) DeepCopy() *BudgetStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionObservation) DeepCopyInto(out *ConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionObservation.
func (in *ConditionObservation) DeepCopy() *ConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionParameters) DeepCopyInto(out *ConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionParameters.
func (in *ConditionParameters) DeepCopy() *ConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecifiedAmountObservation) DeepCopyInto(out *SpecifiedAmountObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecifiedAmountObservation.
func (in *SpecifiedAmountObservation) DeepCopy() *SpecifiedAmountObservation {
	if in == nil {
		return nil
	}
	out := new(SpecifiedAmountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpecifiedAmountParameters) DeepCopyInto(out *SpecifiedAmountParameters) {
	*out = *in
	if in.CurrencyCode != nil {
		in, out := &in.CurrencyCode, &out.CurrencyCode
		*out = new(string)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(int64)
		**out = **in
	}
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpecifiedAmountParameters.
func (in *SpecifiedAmountParameters) DeepCopy() *SpecifiedAmountParameters {
	if in == nil {
		return nil
	}
	out := new(SpecifiedAmountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThresholdRulesObservation) DeepCopyInto(out *ThresholdRulesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThresholdRulesObservation.
func (in *ThresholdRulesObservation) DeepCopy() *ThresholdRulesObservation {
	if in == nil {
		return nil
	}
	out := new(ThresholdRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThresholdRulesParameters) DeepCopyInto(out *ThresholdRulesParameters) {
	*out = *in
	if in.SpendBasis != nil {
		in, out := &in.SpendBasis, &out.SpendBasis
		*out = new(string)
		**out = **in
	}
	if in.ThresholdPercent != nil {
		in, out := &in.ThresholdPercent, &out.ThresholdPercent
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThresholdRulesParameters.
func (in *ThresholdRulesParameters) DeepCopy() *ThresholdRulesParameters {
	if in == nil {
		return nil
	}
	out := new(ThresholdRulesParameters)
	in.DeepCopyInto(out)
	return out
}
